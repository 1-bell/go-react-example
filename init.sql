CREATE DATABASE IF NOT EXISTS app_db;
USE app_db;

CREATE TABLE IF NOT EXISTS Persons (
     Name varchar(255),
     Age int,
     Balance double,
     Email varchar(255) NOT NULL,
     Address varchar(255),
     PRIMARY KEY (Email)
);
