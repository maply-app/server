#!/bin/bash

while ! nc -z postgresql 5432
do
  echo "Failure connected to PostgreSQL"
  sleep 3
done

while ! nc -z redis 6379
do
  echo "Failure connected to Redis"
  sleep 3
done

go build -ldflags="-s -w" -o ./bin/app .
./bin/app