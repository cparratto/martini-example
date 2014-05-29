package main

import(
    "github.com/go-martini/martini"

    "github.com/martini-contrib/binding"
    "github.com/martini-contrib/render"

    "github.com/codegangsta/envy/lib"

    "github.com/joiggama/martini-example/app/controllers"
    "github.com/joiggama/martini-example/app/models"
    "github.com/joiggama/martini-example/config"
)


func main() {
    envy.Bootstrap()

    app := martini.Classic()

    app.Map(config.DB())

    app.Use(render.Renderer())

    app.Group("/products", func(router martini.Router) {
        router.Post("", binding.Json(models.Product{}), controllers.ProductsCreate)
        router.Delete("/:id", controllers.ProductsDelete)
        router.Get("", controllers.ProductsIndex)
        router.Get("/:id", controllers.ProductsShow)
        router.Put("/:id", binding.Json(models.Product{}), controllers.ProductsUpdate)
        router.Post("/bulk", binding.Json(models.Products{}), controllers.ProductsBulkCreate)
    }, controllers.ApiAuth())

    app.Run()
}
