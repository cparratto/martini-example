package config

import (
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"

	"github.com/joiggama/martini-example/app/models"
)

func DB() *xorm.Engine {
	db, err := xorm.NewEngine("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	registerTables(db)
	enableLogging(db)

	return db
}

func enableLogging(db *xorm.Engine) {
	db.ShowSQL = true
}

func registerTables(db *xorm.Engine) {
	db.Sync(new(models.Product))
}
