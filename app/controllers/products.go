package controllers

import (
	"net/http"
	"strconv"

	"github.com/joiggama/hood"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/joiggama/martini-example/app/models"
)

func ProductsIndex(db *hood.Hood, params martini.Params, render render.Render, request *http.Request) {
	query := request.URL.Query()

	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	var products []models.Products
	err := db.Limit(limit).Offset(offset).Find(&products)

	if err == nil {
		render.JSON(200, products)

	} else {
		render.JSON(404, err.Error())
	}

}

func ProductsShow(db *hood.Hood, params martini.Params, render render.Render) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	var products []models.Products
	err = db.Where("id", "=", id).Find(&products)

	if err == nil {
		render.JSON(200, products[0])
	} else {
		render.JSON(404, map[string]string{"error": err.Error()})
	}
}

func ProductsCreate(product models.Products, db *hood.Hood, render render.Render) {
	_, err := db.Save(&product)
	if err == nil {
		render.JSON(201, &product)
	} else {
		render.JSON(422, map[string]string{"error": err.Error()})
	}
}

// Currently there's no way to only update specific columns, update maps struct fields to table :(
func ProductsUpdate(product models.Products, db *hood.Hood, params martini.Params, render render.Render) {
	id, _ := strconv.ParseInt(params["id"], 0, 64)
	product.Id = hood.Id(id)

	_, err := db.Save(&product)

	if err == nil {
		render.JSON(200, product)
	} else {
		render.JSON(422, err.Error())
	}
}

func ProductsDelete(db *hood.Hood, params martini.Params, render render.Render) {
	id, _ := strconv.ParseInt(params["id"], 0, 64)
	_, err := db.Delete(&models.Products{Id: hood.Id(id)})

	if err == nil {
		render.JSON(204, "No content")
	} else {
		render.JSON(404, "Not found") // Let's just say that's the reason
	}
}

func ProductsBulkCreate(collection models.ProductsCollection, db *hood.Hood, render render.Render) {
	_, err := db.SaveAll(&collection.Elements)

	if err == nil {
		render.JSON(201, &collection)
	} else {
		render.JSON(422, err.Error())
	}
}
