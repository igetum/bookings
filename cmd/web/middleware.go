package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Hit the page")
			next.ServeHTTP(w, r)
		})
}

//adds crsf on all post requests
func NoServe(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)

	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return crsfHandler
}

//Saves and loads session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
