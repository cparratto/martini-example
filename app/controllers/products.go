package controllers

import (
	"net/http"
	"strconv"

	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/joiggama/martini-example/app/models"
)

func ProductsIndex(db *gorp.DbMap, params martini.Params, render render.Render, request *http.Request) {
	query := request.URL.Query()

	var limit, offset string

	//// There are no ORM chain methods for querying
	if query.Get("limit") != "" {
		limit = " LIMIT " + query.Get("limit")
	}
	if query.Get("offset") != "" {
		offset = " OFFSET " + query.Get("offset")
	}

	var products []models.Product
	_, err := db.Select(&products, "SELECT * FROM products"+limit+offset)

	if err == nil {
		render.JSON(200, products)

	} else {
		render.JSON(404, "")
	}

}

func ProductsShow(db *gorp.DbMap, params martini.Params, render render.Render) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	product, err := db.Get(models.Product{}, id)

	if err == nil {
		render.JSON(200, product)
	} else {
		render.JSON(404, map[string]string{"error": err.Error()})
	}
}

func ProductsCreate(product models.Product, db *gorp.DbMap, render render.Render) {
	err := db.Insert(&product)
	if err == nil {
		render.JSON(201, &product)
	} else {
		render.JSON(422, map[string]string{"error": err.Error()})
	}
}

// Currently there's no way to only update specific columns, update maps struct fields to table :(
// https://github.com/coopernurse/gorp/issues/92
func ProductsUpdate(product models.Product, db *gorp.DbMap, params martini.Params, render render.Render) {
	id, _ := strconv.ParseInt(params["id"], 0, 64)
	product.Id = id
	_, err := db.Update(&product)

	if err == nil {
		render.JSON(200, product)
	} else {
		render.JSON(422, err.Error())
	}
}

func ProductsDelete(db *gorp.DbMap, params martini.Params, render render.Render) {
	id, _ := strconv.ParseInt(params["id"], 0, 64)
	_, err := db.Delete(&models.Product{Id: id})

	if err == nil {
		render.JSON(204, "No content")
	} else {
		render.JSON(404, "Not found") // Let's just say that's the reason
	}
}

func ProductsBulkCreate(products models.Products, db *gorp.DbMap, render render.Render) {
	for _, product := range products.Collection {
		err := db.Insert(&product)
		if err != nil {
			panic(err)
		}
	}

	render.JSON(201, products)
}
