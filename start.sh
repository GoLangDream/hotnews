#!/usr/bin/env sh

cat application.pid  | xargs kill -9
rm application.log
git pull
go build
export ICEBERG_ENV=production
nohup ./hot_news > application.log 2>&1 &