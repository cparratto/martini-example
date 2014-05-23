package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"

    "github.com/martini-contrib/binding"

    "database/sql"
    _ "github.com/lib/pq"
    "github.com/coopernurse/gorp"

    "os"
    "time"
    "net/http"
    "strconv"
)

func panicIf(err error) {
  if err != nil {
    panic(err)
  }
}

func establishDbConnection() *gorp.DbMap {
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  panicIf(err)
  dbmap   := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
  dbmap.AddTableWithName(Product{}, "products").SetKeys(true, "Id")
  return dbmap
}

type Product struct {
  Id   int64          `json:"id"`
  Code string         `json:"code" binding:"required"`
  Name string         `json:"name" binding:"required"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
  UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (p Product) TableName() string {
  return "products"
}

func main() {
  app := martini.Classic()

  dbmap := establishDbConnection()

  app.Map(dbmap)
  app.Use(render.Renderer())

  app.Group("/products", func(router martini.Router) {

    // index
    router.Get("", func(params martini.Params, render render.Render, request *http.Request) {
      query  := request.URL.Query()

      var limit, offset string

      if query.Get("limit") != "" { limit = " LIMIT " + query.Get("limit") }
      if query.Get("offset") != "" { offset = " OFFSET " + query.Get("offset") }

      var products []Product
      _ , err := dbmap.Select(&products, "SELECT * FROM products" + limit + offset)

      panicIf(err)
      render.JSON(200, products)
    })

    // show
    router.Get("/:id", func(params martini.Params, render render.Render){
      id, err := strconv.Atoi(params["id"])
      panicIf(err)
      res, err := dbmap.Get(Product{}, id)
      product := res.(*Product)
      if err == nil {
        render.JSON(200, product)
      } else {
        render.JSON(404, map[string]string{ "error": "Not Found" })
      }
    })

    // create
    router.Post("", binding.Json(Product{}), func(p Product, r render.Render ){
      err := dbmap.Insert(&p)
      if err == nil {
        r.JSON(201, p)
      } else {
        r.JSON(422, map[string]string{ "error" : "Unprocessable Entity"})
      }
    })
  })

  app.Run()
}
