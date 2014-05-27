package config

import(
    "database/sql"
    "os"
    "log"

    _ "github.com/lib/pq"
    "github.com/coopernurse/gorp"

    "github.com/joiggama/martini-example/app/models"
)


func DB() *gorp.DbMap {
    connection, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil { panic(err) }
    db := &gorp.DbMap{Db: connection, Dialect: gorp.PostgresDialect{}}

    registerTables(db)
    enableLogging(db)

    return db
}

func enableLogging(db *gorp.DbMap) {
    db.TraceOn("[gorp]", log.New(os.Stdout, "", log.Lmicroseconds))
}

func registerTables(db *gorp.DbMap) {
    db.AddTableWithName(models.Product{}, "products").SetKeys(true, "Id")
}
