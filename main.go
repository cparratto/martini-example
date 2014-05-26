package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/binding"
    "github.com/martini-contrib/render"

    _ "github.com/lib/pq"
    "github.com/coopernurse/gorp"

    "database/sql"
    "log"
    "net/http"
    "os"
    "strconv"

    "github.com/joiggama/martini-example/app/models"
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
  dbmap.AddTableWithName(models.Product{}, "products").SetKeys(true, "Id")
  dbmap.TraceOn("[gorp]", log.New(os.Stdout, "", log.Lmicroseconds))
  return dbmap
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

      // There are no ORM chain methods for querying
      if query.Get("limit") != "" { limit = " LIMIT " + query.Get("limit") }
      if query.Get("offset") != "" { offset = " OFFSET " + query.Get("offset") }

      var products []models.Product
      _ , err := dbmap.Select(&products, "SELECT * FROM products" + limit + offset)

      panicIf(err)
      render.JSON(200, products)
    })

    // show
    router.Get("/:id", func(params martini.Params, render render.Render){
      id, err := strconv.Atoi(params["id"])
      panicIf(err)
      product, err := dbmap.Get(models.Product{}, id)
      if err == nil {
        render.JSON(200, product)
      } else {
        render.JSON(404, map[string]string{ "error": err.Error() })
      }
    })

    // create
    router.Post("", binding.Json(models.Product{}), func(product models.Product, render render.Render) {
      err := dbmap.Insert(&product)
      if err == nil {
        render.JSON(201, product)
      } else {
        render.JSON(422, map[string]string{ "error" : err.Error() })
      }
    })

    // update
    // Currently there's no way to only update specific columns, update maps struct fields to table :(
    // https://github.com/coopernurse/gorp/issues/92
    router.Put("/:id", binding.Json(models.Product{}), func(product models.Product, render render.Render, params martini.Params){
      id, _ := strconv.ParseInt(params["id"], 0, 64)
      product.Id = id
      _, err := dbmap.Update(&product)
      if err == nil {
        render.JSON(200, product)
      } else {
        render.JSON(422, err.Error())
      }
    })

    // destroy
    router.Delete("/:id", func(params martini.Params, render render.Render) {
      id, _ := strconv.ParseInt(params["id"], 0, 64)
      _, err := dbmap.Delete(&models.Product{Id: id})
      if err == nil {
        render.JSON(204, "No content")
      } else {
        render.JSON(404, "Not found") // Let's just say that's the reason
      }
    })
  })

  // Bulk Operations
  app.Group("/bulk/products", func(router martini.Router){

    // create
    router.Post("", binding.Json(models.Products{}), func(products models.Products, render render.Render){
      for _, product := range products.Collection {
        err := dbmap.Insert(&product)
        panicIf(err)
      }

      render.JSON(201, products)
    })

    // Bulk modify won't be useful, since gorp overwrites every struct field mapped to the database
    // Bulk delete, what for?
  })

  app.Run()
}
