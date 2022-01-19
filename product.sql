CREATE TABLE derekshop.product (
	product_no INT PRIMARY KEY AUTO_INCREMENT,
	product_name VARCHAR(20),
	product_price INT,
	product_desc VARCHAR(100),
	img_url VARCHAR(30),
	create_time DATETIME,
	product_status INT
);