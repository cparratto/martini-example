package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"

    _ "github.com/lib/pq"
    "github.com/go-xorm/xorm"
)

func panicIf(err error) {
  if err != nil {
    panic(err)
  }
}

func establishDbConnection() *xorm.Engine {
  engine, err := xorm.NewEngine("postgres", "dbname=example_app_dev sslmode=disable")
  panicIf(err)
  return engine
}

type Products struct {
  Id   int32  `json:"id"`
  Code string `json:"code"`
  Name string `json:"name"`
}

func main() {
  app := martini.Classic()

  app.Map(establishDbConnection())
  app.Use(render.Renderer())

  // Routes
  app.Get("/products", func(db *xorm.Engine, r render.Render) {
    var products []Products
    err := db.Find(&products)
    panicIf(err)
    r.JSON(200, products)
  })

  app.Run()
}
