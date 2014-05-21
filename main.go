package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"

    "github.com/martini-contrib/binding"

    _ "github.com/lib/pq"
    "github.com/go-xorm/xorm"

    "fmt"
    "net/http"
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
  Id   int64
  Code string `json:"code" binding:"required"`
  Name string `json:"name" binding:"required"`
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
    router.Get("", func(params martini.Params, render render.Render, request *http.Request) {
      query     := request.URL.Query()
      limit, _  := strconv.Atoi(query.Get("limit"))
      offset, _ := strconv.Atoi(query.Get("offset"))

      var products []Product

      err := engine.Limit(limit, offset).Find(&products)

      panicIf(err)
      render.JSON(200, products)
    })

    // show
    router.Get("/:id", func(params martini.Params, render render.Render){
      id, err := strconv.Atoi(params["id"])
      var product = Product{Id: int64(id)}
      found, err := engine.Get(&product)
      panicIf(err)
      if found {
        render.JSON(200, product)
      } else {
        render.JSON(404, map[string]string{ "error": "Not Found" })
      }
    })

    // create
    router.Post("", binding.Json(Product{}), func(p Product, r render.Render ){
      // binding should work... but doesn't
      success, err := engine.Insert(&p)
      fmt.Println(success)
      if err == nil {
        r.JSON(201, p)
      } else {
        r.JSON(422, map[string]string{ "error" : "Unprocessable Entity"})
      }
    })
  })

  app.Run()
}
