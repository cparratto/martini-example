package main

import "github.com/go-martini/martini"

//import "net/http"
import "database/sql"
import _ "github.com/lib/pq"
import "github.com/martini-contrib/render"

func connect() *sql.DB {
  db, err := sql.Open("postgres", "dbname=example_app_dev sslmode=disable")

  if err != nil {
    panic(err)
  }

  return db
}

func main() {
  m := martini.Classic()

  m.Map(connect())
  m.Use(render.Renderer())

  m.Get("/products", func(db *sql.DB, rnd render.Render) string {
    products, _ := db.Query("SELECT * FROM products")
    rnd.JSON(200, products)
    return ""
  })

  m.Run()
}
