package controllers

import(
    "os"

    "github.com/go-martini/martini"
    "github.com/martini-contrib/auth"
)

func ApiAuth() martini.Handler{
    return auth.Basic(os.Getenv("API_USERNAME"), os.Getenv("API_PASSWORD"))
}
