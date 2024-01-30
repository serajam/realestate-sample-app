# Real Estate sample app

Example of api built with golang

- app design example
- image storage example - AWS S3 compatible storage **minio**
- **domain design** architecture example
- event approach usage example for email sending using bus **nats**
- user session handling with **redis**
- database **postgresql** usage with a help of query builder
- postgres **postgis** usage for geo searching properties within an area or radius 
- **e2e** tests
- swagger for api docs

## Requirements

- make
- docker + docker compose plugin

## Running the app using docker-compose

### If make is installed 

You can run with clean data or without just by specifying the env variables `GENERATE_TEST_DATA` and `CLEANUP_DB` respectively before make cmd.
Clean run, will take some time to generate test data: 
```
GENERATE_TEST_DATA=true CLEANUP_DB=true make upd
```

Run in the background: 

```
make upd
```

## Swagger 

- UI http://localhost:8080/swagger/index.html
- JSON description docs/swagger/swagger.json
``