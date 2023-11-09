package main

import (
	"flag"
	"fmt"

	// user_app "github.com/jamilnoyda/go-server/user_app"
	"html/template"

	logrus "github.com/sirupsen/logrus"

	// "log"
	"net/http"
	"os"
	"path"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	// You might want to move ParseGlob outside of handle so it doesn't
	// re-parse on every http request.
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := ""
	if r.URL.Path == "/" {
		name = "index.html"
	} else {
		name = path.Base(r.URL.Path)
	}

	data := struct {
		Time time.Time
	}{
		Time: time.Now(),
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: helloserver [options]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	greeting = flag.String("g", "Hello", "Greet with `greeting`")
	addr     = flag.String("addr", ":0", "address to serve")
)

func main() {
	// Parse flags.
	flag.Usage = usage
	flag.Parse()
	logrus.Info("hello")

	// Parse and validate arguments (none).
	args := flag.Args()
	if len(args) != 0 {
		usage()
	}
	// log.Printf("serving http://%s\n", *addr)
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	http.HandleFunc("/", handle)
	// http.HandleFunc("/greet", user_app.Greet)
	// http.HandleFunc("/version",user_app.Version)

	http.ListenAndServe(":11", nil)

}
