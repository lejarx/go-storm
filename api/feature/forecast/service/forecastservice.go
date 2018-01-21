package service

import (
	"github.com/andream16/go-storm/model/request"
	"database/sql"
	"fmt"
	"errors"
)

func GetForecasts(itemId string, db *sql.DB) (request.Forecast, error) {
	emptyResponse := request.Forecast{Item: itemId, Name: "", Forecast: []request.ForecastEntry{}}
	stmt, err := db.Prepare(`SELECT price,date,score FROM forecast WHERE item=$1 ORDER BY date ASC`); if err != nil {
		defer stmt.Close()
		return request.Forecast{}, err
	}
	defer stmt.Close()
	rows, queryError := stmt.Query(itemId); if queryError != nil {
		return emptyResponse, nil
	}
	defer rows.Close()
	forecastNameQuery, forecastNameQueryError := db.Prepare(`SELECT name FROM forecast WHERE item=$1 LIMIT 1`)
	if forecastNameQueryError != nil {
		return emptyResponse, nil
	}
	defer forecastNameQuery.Close()
	var forecastName string
	nameError := forecastNameQuery.QueryRow(itemId).Scan(&forecastName); if nameError != nil || len(forecastName) == 0 {
		return emptyResponse, nil
	}
	var forecasts request.Forecast
	forecasts.Name = forecastName
	forecasts.Item = itemId
	scoreHasBeenSet := false
	for rows.Next() {
		var forecast request.ForecastEntry
		rowError := rows.Scan(&forecast.Price, &forecast.Date, &forecast.Score); if rowError != nil {
			return request.Forecast{}, errors.New(fmt.Sprintf("Unable to unmarshal forecasts for item %s. Error: %s", itemId, rowError.Error()))
		}
		if !scoreHasBeenSet {
			forecasts.Score = forecast.Score
			scoreHasBeenSet = true
		}
		forecasts.Forecast = append(forecasts.Forecast, request.ForecastEntry{
			Price: forecast.Price,
			Date: forecast.Date,
		})
	}
	iterationError := rows.Err(); if iterationError != nil {
		return request.Forecast{}, errors.New(fmt.Sprintf("No forecasts found for item %s. Error: %s", itemId, iterationError.Error()))
	}
	return forecasts, nil
}

func AddForecasts(forecasts request.Forecast, db *sql.DB) error {
	itemId := forecasts.Item
	forecastsEntries, _ := GetForecasts(itemId, db); if len(forecastsEntries.Forecast) > 0 {
		deleteForecastsError := DeleteForecast(itemId, db); if deleteForecastsError != nil {
			return deleteForecastsError
		}
	}
	stmt, err := db.Prepare(`INSERT INTO forecast(name,item,price,date,score) VALUES ($1,$2,$3,$4,$5)`); if err != nil {
		defer stmt.Close()
		return err
	}
	defer stmt.Close()
	forecastType := forecasts.Name
	for _, forecast := range forecasts.Forecast {
		_, insertError := stmt.Exec(forecastType, itemId, forecast.Price, forecast.Date, forecast.Score)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func EditForecast(forecasts request.Forecast, db *sql.DB) error {
	stmtInsert, err := db.Prepare(`INSERT INTO forecast(name,item,price,date,score) VALUES ($1,$2,$3,$4,$5)`); if err != nil {
		defer stmtInsert.Close()
		return err
	}
	defer stmtInsert.Close()
	forecastType := forecasts.Name
	deleteError := DeleteForecast(forecasts.Item, db); if deleteError != nil {
		return deleteError
	}
	for _, forecast := range forecasts.Forecast {
		_, insertError := stmtInsert.Exec(forecastType, forecasts.Item, forecast.Price, forecast.Date, forecast.Score)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func DeleteForecast(itemId string, db *sql.DB) error {
	stmtDelete, err := db.Prepare(`DELETE FROM forecast WHERE item=$1`); if err != nil {
		defer stmtDelete.Close()
		return err
	}
	defer stmtDelete.Close()
	_, deleteError := stmtDelete.Exec(itemId); if deleteError != nil {
		return deleteError
	}
	return nil
}
