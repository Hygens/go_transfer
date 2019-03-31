package main

import (
	. "github.com/Hygens/go_transfer/controllers"
	. "github.com/Hygens/go_transfer/models"
	. "github.com/Hygens/go_transfer/utilities"
	"gopkg.in/unrolled/render.v1"
	"net/http"
	"time"
)

func main() {
	// Load Data
	users := GetUsers()

	// Server mux for simple handle routes
	mux := http.NewServeMux()

	// Instance for main controller
	c := &MyController{Users: &users, Render: render.New(render.Options{})}

	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Various sample routes implemented
	mux.Handle("/api/user/", c.Action(c.GetUser))
	mux.Handle("/main", c.Action(c.Transfer))
	mux.Handle("/transfer", c.Action(c.SendFounds))
	server := &http.Server{
		Addr:           Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
