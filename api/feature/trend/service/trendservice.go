package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
	"github.com/go-errors/errors"
	"fmt"
)

func GetTrendByManufacturer(manufacturer string, db *sql.DB) (request.Trend, error) {
	rows, getTrendError := db.Query(`SELECT * FROM trend WHERE manufacturer=$1`, manufacturer); if getTrendError != nil {
		return request.Trend{}, errors.New(fmt.Sprintf("Unable to find trend rows for manufacturer %s", manufacturer))
	}
	defer rows.Close()
	var trend request.Trend; trend.Manufacturer = manufacturer
	for rows.Next() {
		var trendEntry request.TrendEntry
		rowError := rows.Scan(&trendEntry.Date, &trendEntry.Value); if rowError != nil {
			return request.Trend{}, errors.New(fmt.Sprintf("Unable to unmarshal trend entries for manufacturer %s. Error: %s", manufacturer, rowError.Error()))
		}
		trend.Trend = append(trend.Trend, trendEntry)
	}; iterationError := rows.Err(); if iterationError != nil {
		return request.Trend{}, errors.New(fmt.Sprintf("No trend entries found for manufacturer %s. Error: %s", manufacturer, iterationError.Error()))
	}; if len(trend.Trend) == 0 {
		return request.Trend{}, errors.New(fmt.Sprintf("No trend entries found for manufacturer %s", manufacturer))
	}
	return trend, nil
}

func AddTrendByManufacturer(trend request.Trend, db *sql.DB) error {
	return addTrend(trend, db)
}

func EditTrendByManufacturer(trend request.Trend, db *sql.DB) error {
	deleteError := deleteTrend(trend, db); if deleteError != nil {
		return deleteError
	}
	return  addTrend(trend, db)
}

func DeleteTrendByManufacturer(trend request.Trend, db *sql.DB) error {
	return deleteTrend(trend, db)
}

func addTrend(trend request.Trend, db *sql.DB) error {
	for _, trendEntry := range trend.Trend {
		_, insertError := db.Query(`INSERT INTO trend(manufacturer,value,date) VALUES ($1,$2,$3)`, trend.Manufacturer, trendEntry.Value, trendEntry.Date)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func deleteTrend(trend request.Trend, db *sql.DB) error {
	_, deleteError := db.Query(`DELETE FROM trend WHERE manufacturer=$1`, trend.Manufacturer); if deleteError != nil {
		return deleteError
	}
	return nil
}
