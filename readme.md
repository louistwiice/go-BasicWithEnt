# Documentation

## Step 1: Set env file
|Env name| Description|Example|
|---|---|---|
|SERVER_PORT|Application running port|:9000|
|DB_ROOT_PASSWORD| Mysql root passwod||
|DB_NAME| Mysql Database||
|DB_USER| Mysql User||
|DB_PASSWORD| Mysql password||
|DB_HOST| Server IP host |localhost|

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
