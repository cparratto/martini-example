package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"

    _ "github.com/lib/pq"
    "github.com/go-xorm/xorm"

    "strconv"
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

type Product struct {
  Id   int32  `json:"id"`
  Code string `json:"code"`
  Name string `json:"name"`

}

func (p Product) TableName() string {
  return "products"
}

func main() {
  app := martini.Classic()

  engine := establishDbConnection()
  engine.Sync(new(Product))

  app.Map(engine)
  app.Use(render.Renderer())

  app.Group("/products", func(router martini.Router) {

    // index
    router.Get("", func(params martini.Params, render render.Render) {
      var products []Product
      err := engine.Find(&products)
      panicIf(err)
      render.JSON(200, products)
    })

    // show
    router.Get("/:id", func(params martini.Params, render render.Render){
      id, err := strconv.Atoi(params["id"])
      var product = Product{Id: int32(id)}
      found, err := engine.Get(&product)
      panicIf(err)
      if found {
        render.JSON(200, product)
      } else {
        render.JSON(404, map[string]string{ "error": "Not Found" })
      }
    })

  })

  app.Run()
}
