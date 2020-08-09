#!/bin/bash

mysqldump -u $MYSQL_USER -p$MYSQL_PASSWORD -h $MYSQL_HOST $MYSQL_NAME > /var/dump/dump.sql