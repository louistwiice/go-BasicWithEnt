# Documentation

## Step 1: Set env file
|Env name| Description|Example|
|---|---|---|
|SERVER_PORT|Application running port|:9000|
|DB_ROOT_PASSWORD| Database root passwod||
|DB_NAME| Database name||
|DB_USER| Database User||
|DB_PASSWORD| Database password||
|DB_HOST| Database IP host |localhost|

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
