package controllers

import(
    "net/http"
    "strconv"

    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "github.com/go-xorm/xorm"


    "github.com/joiggama/martini-example/app/models"
)

func ProductsIndex(db *xorm.Engine, params martini.Params, render render.Render, request *http.Request) {
    query     := request.URL.Query()
    limit, _  := strconv.Atoi(query.Get("limit"))
    offset, _ := strconv.Atoi(query.Get("offset"))

    var products []models.Product
    err := db.Limit(limit, offset).Find(&products)

    if err == nil {
        render.JSON(200, &products)

    } else {
        render.JSON(404, "")
    }

}

func ProductsShow(db *xorm.Engine, params martini.Params, render render.Render){
    id, err := strconv.ParseInt(params["id"], 0, 64)

    if err != nil { panic(err) }

    var product = models.Product{Id: id}
    _, err = db.Get(&product)

    if err == nil {
        render.JSON(200, product)
    } else {
        render.JSON(404, map[string]string{ "error" : err.Error() })
    }
}

func ProductsCreate(product models.Product, db *xorm.Engine, render render.Render) {
      _, err := db.Insert(&product)

      if err == nil {
          render.JSON(201, &product)
      } else {
          render.JSON(422, map[string]string{ "error" : err.Error() })
      }
}

// Currently there's no way to only update specific columns, update maps struct fields to table :(
// https://github.com/coopernurse/gorp/issues/92
//func ProductsUpdate(product models.Product, db *xorm.Engine, params martini.Params, render render.Render){
     //id, _ := strconv.ParseInt(params["id"], 0, 64)
     //product.Id = id
     //_, err := db.Update(&product)

     //if err == nil {
         //render.JSON(200, product)
     //} else {
         //render.JSON(422, err.Error())
     //}
//}

//func ProductsDelete(db *xorm.Engine, params martini.Params, render render.Render) {
    //id, _ := strconv.ParseInt(params["id"], 0, 64)
    //_, err := db.Delete(&models.Product{Id: id})

    //if err == nil {
        //render.JSON(204, "No content")
    //} else {
        //render.JSON(404, "Not found") // Let's just say that's the reason
    //}
//}

//func ProductsBulkCreate(products models.Products, db *xorm.Engine, render render.Render){
    //for _, product := range products.Collection {
        //err := db.Insert(&product)
        //if err != nil { panic(err) }
    //}

    //render.JSON(201, products)
//}

