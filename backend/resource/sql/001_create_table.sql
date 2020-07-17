
DROP DATABASE IF EXISTS ACL;
CREATE DATABASE ACL;
use ACL;


DROP TABLE  IF EXISTS users;
DROP TABLE  IF EXISTS userQuery;
DROP TABLE  IF EXISTS groups1;
DROP TABLE  IF EXISTS userGroup;
DROP TABLE  IF EXISTS Files;


CREATE TABLE users (
	id 	INT AUTO_INCREMENT	PRIMARY KEY,
	first_name	CHAR(25) 	NOT NULL,
	last_name	CHAR(25) 	NOT NULL,
	batch       INT(4)      NOT NULL,     
	email   	CHAR(50)	NOT NULL,
	contact_number CHAR(10)    NOT NULL,
	address     CHAR(100)   NOT NULL,
	password 	VARBINARY(128)	NOT NULL,
	creation_date	DATETIME	DEFAULT CURRENT_TIMESTAMP
)ENGINE = INNODB CHARACTER SET=utf8;

CREATE TABLE userQuery (
	id INT AUTO_INCREMENT PRIMARY KEY,
	user_id INT,
	query CHAR(200) NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE userHealth (
	id INT AUTO_INCREMENT PRIMARY KEY,
	user_id INT,
	age INT,
	temp INT,
	sour_throat BOOLEAN,
	medical_history TEXT NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
);
# groups table
CREATE TABLE groups1 (
	id 	INT     PRIMARY KEY,
	group_name	CHAR(25),
	group_info	CHAR(50),
	group_creation_date		DATETIME 	DEFAULT CURRENT_TIMESTAMP 
)ENGINE = INNODB CHARACTER SET=utf8;

# group-user table 
CREATE TABLE userGroup (
	#id 	INT 	AUTO_INCREMENT UNIQUE KEY,
	user_id 	INT,
	group_id 	INT,
	PRIMARY KEY (user_id, group_id),
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (group_id) REFERENCES groups1(id)
)ENGINE = INNODB CHARACTER SET=utf8;

# file and folder is considered as single entity

CREATE TABLE Files (
    fid   int NOT NULL AUTO_INCREMENT,
    FileName varchar(10) NOT NULL,
    FilePath varchar(50),
    PRIMARY KEY (fid)
)ENGINE = INNODB CHARACTER SET=utf8;


