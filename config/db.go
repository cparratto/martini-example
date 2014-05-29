package config

import (
	//"database/sql"
	//"log"
	"os"

	"github.com/joiggama/hood"
	_ "github.com/lib/pq"

	//"github.com/joiggama/martini-example/app/models"
)

func DB() *hood.Hood {
	db, err := hood.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	//registerTables(db)
	//enableLogging(db)

	return db
}

//func enableLogging(db *gorp.DbMap) {
	//db.TraceOn("[gorp]", log.New(os.Stdout, "", log.Lmicroseconds))
//}

//func registerTables(db *gorp.DbMap) {
	//db.AddTableWithName(models.Product{}, "products").SetKeys(true, "Id")
//}
