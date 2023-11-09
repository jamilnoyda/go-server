package user_app

import (
	"fmt"
	"html"
	"net/http"
	"runtime/debug"
	"strings"
	// "html/template"
	// "path"
	// "time"
)

// Register handlers.

// All requests not otherwise mapped with go to greet.
// /version is mapped specifically to version.

func Version(w http.ResponseWriter, r *http.Request) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		http.Error(w, "no build information available", 500)
		return
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n<pre>\n")
	fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))

}

func Greet(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	if name == "" {
		name = "Gopher"
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n")
	fmt.Fprintf(w, "%s, %s!\n", "dd", html.EscapeString(name))
}
