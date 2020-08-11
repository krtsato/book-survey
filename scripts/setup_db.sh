#!/bin/bash

until mysql -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD -e 'quit'; do
  echo >&2 'MySQL is sleeping now.'
  sleep 1
done

echo >&2 'MySQL is up.'

mysql -h $MYSQL_HOST -u root -p$MYSQL_ROOT_PASSWORD -e "grant all on *.* to 'mysql_user'@'%'"

echo 'MYSQL_USER get ALL PRIVILEGES.'

mysql -h $MYSQL_HOST -u root -p$MYSQL_ROOT_PASSWORD -e "create database if not exists survey_db;"

mysql -h $MYSQL_HOST -u root -p$MYSQL_ROOT_PASSWORD -e "create table if not exists survey_db.genres (
  id int primary key AUTO_INCREMENT,
  genre_code varchar(40) not null,
  name varchar(16) not null
);"

mysql -h $MYSQL_HOST -u root -p$MYSQL_ROOT_PASSWORD -e "create table if not exists survey_db.reviews (
  id int primary key AUTO_INCREMENT,
  review_cnt int not null,
  review_avg float not null
);"


# 本来ならば foreign key(genre_id) references genres(id), とすべきだが
# internal/db/insert.go に記載した "注" のため, ひとまず制約を持たせていない
mysql -h $MYSQL_HOST -u root -p$MYSQL_ROOT_PASSWORD -e "create table if not exists survey_db.items (
  id int primary key AUTO_INCREMENT,
  genre_id int not null,
  review_id int not null,
  title varchar(100) not null,
  price int not null,
  sales_date date not null,
  url varchar(2083) not null,
  foreign key(review_id) references reviews(id)
);"

echo 'MySQL has survey_db and its tables'
