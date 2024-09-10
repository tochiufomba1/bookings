package main // Can't be called from an external package

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tochiufomba1/bookings/pkg/config"
	"github.com/tochiufomba1/bookings/pkg/handlers"
	"github.com/tochiufomba1/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")

	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	n, err := fmt.Fprintln(w, "Hello, world!")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
// })
