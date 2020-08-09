#!/bin/bash

until mysql -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD -e 'quit'; do
  echo >&2 'MySQL is sleeping now.'
  sleep 3
done

echo >&2 'MySQL is up.'

mysql -h $MYSQL_HOST -u root -p$MYSQL_ROOT_PASSWORD -e "GRANT ALL ON *.* TO 'mysql_user'@'%'"

echo 'MYSQL_USER get ALL PRIVILEGES.'