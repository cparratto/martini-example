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

## Tests

None :cry:

## Deployment

Deployment instructions for heroku.

### Setup

* Make sure to have the [heroku toolbelt](https://toolbelt.heroku.com/)
  installed.

* Then run:
  ```bash
  heroku create <app-name> -b https://github.com/kr/heroku-buildpack-go.git
  ```

* Add the heroku git repository to your git remotes

* Install `Godep` to manage app dependencies by running:
  `go get github.com/tools/godep`

### Before pushing

* Every time that you add a new a dependency, make sure run `godep save`
  and add changed files inside `Godeps` directory.

### Release

* `git push <heroku-remote>` or `git push <heroku-remote> <feature-branch>:master`


