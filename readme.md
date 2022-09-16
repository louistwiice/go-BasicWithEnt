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
|TOKEN_HOUR_LIFESPAN| Duration of authentication token you in Login in Hour |1|
|API_SECRET| Secret key that allow you to generate each login token |secret_1246@@@@!!/shghj_---QaZerftQWWWfz|
|TOKEN_PREFIX| authorization token prefix used |Bearer|

Other environment variable will go there in this file


## Step 2: Start mysql container

``` text
docker-compose up -d

# Or
make start-db
```

## Step 3: Create tables
``` text
go generate ./ent

# Or
make generate-schema
```

## Step 4: Start application
```text
go run api/main.go

# or
make start-server
```
