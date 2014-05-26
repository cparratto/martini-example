package main

import (
    "database/sql"
    "log"
    "os"

    "github.com/go-martini/martini"
    "github.com/martini-contrib/binding"
    "github.com/martini-contrib/render"

    _ "github.com/lib/pq"
    "github.com/coopernurse/gorp"


    "github.com/joiggama/martini-example/app/models"
    "github.com/joiggama/martini-example/app/controllers"
)

func establishDbConnection() *gorp.DbMap {
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil { panic(err) }
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

    app.Group("/products", func(router martini.Router){
        router.Post("", controllers.ProductsCreate)
        router.Delete("/:id", controllers.ProductsDelete)
        router.Get("", controllers.ProductsIndex)
        router.Get("/:id", binding.Json(models.Product{}), controllers.ProductsShow)
        router.Put("/:id", binding.Json(models.Product{}), controllers.ProductsUpdate)

        router.Post("/bulk", binding.Json(models.Products{}), controllers.ProductsBulkCreate)
    })

  app.Run()
}
