package config

import (
	"os"

	"github.com/joiggama/hood"
	_ "github.com/lib/pq"
)

func DB() *hood.Hood {
	db, err := hood.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	return db
}
