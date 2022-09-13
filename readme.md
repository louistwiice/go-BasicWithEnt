# Documentation

## Step 1: Set env file
|Env name| Description|Example|
|---|---|---|
|SERVER_PORT|Application running port|:9000|
|MYSQL_ROOT_PASSWORD| Mysql root passwod||
|MYSQL_DATABASE| Mysql Database||
|MYSQL_USER| Mysql User||
|MYSQL_PASSWORD| Mysql password||
|MYSQL_HOST| Server IP host |localhost|

Other environment variable will go there in this file


## Step 2: Start mysql container

``` text
docker-compose up -d
```

## Step 3: Create tables
``` text
go generate ./ent
```

## Step 4: Start application
```text
go run api/main.go
```
