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
|ACCESS_TOKEN_HOUR_LIFESPAN| Duration of authentication access token you in Login in Hour |1|
|REFRESH_TOKEN_HOUR_LIFESPAN| Duration of authentication refresh token you in Login in Hour |1|
|ACCESS_TOKEN_SECRET| Access Secret key that allow you to generate each login token |secret_1246@@@@!!/shghj_---QaZerftQWWWfz|
|REFRESH_TOKEN_SECRET| Refresh Secret key that allow you to generate each login token |secret_1246@@@@!!/shghj_---QaZerftQWWWfz|
|TOKEN_PREFIX| authorization token prefix used |Bearer|

Other environment variable will go there in this file


## Step 2: Start mysql container

``` text
make start-db
```

## Step 3: Start application
```text
make start-server
```
