## GO Restful
Test are not fully covered.

## Installation Steps
* `sudo docker-compose up -d`
* use `docker inspect go-restful-api_db_1` to use postgres ip for config `app/config/local.yaml` and put it on
* now you can use endpoint on localhost:5000
* `make migrate` to migrate
* `make test` to run test

## Things left to do
* Model relation has not been maintained
* Full test coverage left
* Model has not been validated due to time constraint

## Structure
All codes are placed inside internal with domain style -: 
* api.go handles routing part and makes a service selection
* service.go handles the bridge connection between router and repository
* repository handles the persistence and business logic
