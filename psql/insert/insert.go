package insert

var INSERTS = []string {
	"INSERT INTO manufacturer(name) VALUES('no_manufacturer')",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test1','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test2','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test3','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test4','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test5','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test6','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test7','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test8','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test9','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test10','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO price(item,price,date) VALUES('test1',100.00,'10-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',100.10,'11-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',101.00,'12-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',103.99,'13-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',100.00,'14-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',110.00,'15-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',140.00,'16-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',180.89,'17-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',77.00,'18-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',1.00,'19-10-2017')",
	"INSERT INTO category(category) VALUES('cat1')",
	"INSERT INTO category(category) VALUES('cat2')",
	"INSERT INTO category(category) VALUES('cat3')",
	"INSERT INTO category(category) VALUES('cat4')",
	"INSERT INTO category(category) VALUES('cat5')",
	"INSERT INTO category(category) VALUES('cat6')",
	"INSERT INTO category(category) VALUES('cat7')",
	"INSERT INTO category(category) VALUES('cat8')",
	"INSERT INTO category(category) VALUES('cat9')",
	"INSERT INTO category(category) VALUES('cat10')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat1')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat2')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat3')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat4')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat5')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat6')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat1')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat2')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat3')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat4')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat5')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat6')",
}