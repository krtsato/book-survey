create database if not exists survey_db;

use survey_db;

create table if not exists items (
  id int primary key AUTO_INCREMENT,
  genre_id int foreign key,
  title varchar(100) not null,
  price int not null,
  sales_date date not null,
  review_cnt int not null,
  review_avg int not null,
  url varchar(2083) not null
);

create table if not exists genres (
  id int primary key AUTO_INCREMENT,
  genre_code varchar(40) not null,
  name varchar(16) not null
);


