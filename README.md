# Martini-Example


## System dependencies

* Install Go `brew install go`
* Install Postgres `brew install postgres`

## Configuration

* Create the database `createdb example_app_dev`

* Populate it using the example `example_app_dev.sql` file:
  `psql -d example_app_dev < example_app_dev.sql`

* Create `.env` file:
  ```
  DATABASE_URL="dbaname=example_app_dev sslmode=disable"
  API_USERNAME=<username>
  API_PASSWORD=<password>
  ```

* Run `go get` to install app dependencies

* Run app `go run main.go`

## WIP

* How to run the test suite
* Deployment instructions
