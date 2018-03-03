package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/olling/slog"
)

func apiAuthorized(token string) (authorized bool) {
	authorized = false
	slog.PrintTrace("Token: " + token)
	for _, t := range CurrentConfiguration.Tokens {
		if token == t {
			authorized = true
		}
	}
	return authorized
}

func httpApi(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")

	if !apiAuthorized(token) {
		http.Error(w, "The token was not accepted", 401)
	}

	prefix := strings.SplitAfterN(r.URL.Path, "/", 2)
	slog.PrintTrace("prefix:", prefix[0])

	switch r.URL.Path {
	case "/api/user":
		httpApiUser(w, r)
	case "/api/remind":
		httpApiRemind(w, r)
	case "/api/accept":
		httpApiAccept(w, r)
	case "/api/swap":
		httpApiSwap(w, r)
	case "/api/skip":
		httpApiSkip(w, r)
	default:
		http.Error(w, "Api function not found", 404)
	}
}

func httpApiSwap(w http.ResponseWriter, r *http.Request) {
	slog.PrintTrace("Func called: httpApiSwap")
	queue := GetCurrentQueue()
	users := strings.Split(r.URL.RawQuery, ":")

	if len(users) != 2 {
		http.Error(w, "Incorrect number of users given", 500)
		return
	}

	user0pos, user0 := queue.GetUser(users[0])
	user1pos, user1 := queue.GetUser(users[1])

	if user0pos != -1 && user1pos != -1 {
		err := queue.SwapUsers(user0, user1)
		if err != nil {
			http.Error(w, "Could not swap users", 500)
			slog.PrintError("Could not swap users", user0, user1, err)
			return
		}

	}
}

func httpApiUser(w http.ResponseWriter, r *http.Request) {
	slog.PrintTrace("Func called: httpApiUser")
	username := r.URL.RawQuery
	queue := GetCurrentQueue()
	userpos, user := queue.GetUser(username)

	switch r.Method {
	case "PUT":
		if userpos != -1 {
			http.Error(w, "User already exists", 500)
			return
		}
		//TODO Implement

	case "DELETE":
		if userpos == -1 {
			http.Error(w, "Could not find the user", 500)
			return
		}
		queue := GetCurrentQueue()
		queue.RemoveUser(user)

		queue.Write()
		return

	case "GET":
		if userpos == -1 {
			http.Error(w, "Could not find the user", 500)
			return
		}
		fmt.Fprint(w, user)
		return
	}
	http.Error(w, "Method not found", 500)
}

func httpApiRemind(w http.ResponseWriter, r *http.Request) {
	slog.PrintTrace("Func called: httpApiRemind")
	queue := GetCurrentQueue()
	user := queue.GetResponsible()
	//TODO Finish this
	Notify("command", user, "Subject", "Message")
}

func httpApiAccept(w http.ResponseWriter, r *http.Request) {
	slog.PrintTrace("Func called: httpApiAccept")
	//TODO IMPLEMENT THIS
}

func httpApiShift(w http.ResponseWriter, r *http.Request) {
	slog.PrintTrace("Func called: httpApiShift")
	queue := GetCurrentQueue()
	queue.MoveFirstToBack()
	queue.Write()
}

func httpApiSkip(w http.ResponseWriter, r *http.Request) {
	slog.PrintTrace("Func called: httpApiSkip")
	//TODO IMPLEMENT THIS
}

// httpHandler replaced by func index()
/*func httpHandler(w http.ResponseWriter, r *http.Request) {
	slog.PrintInfo(r.URL.Path)
	fmt.Fprintln(w, "<!doctype html>")
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, "<p>Helloworld</p>")
	user := r.Header.Get("auth_user")
	fmt.Fprintln(w, user)
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}*/

// index replacing httpHandler
func index(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout", "base.tmpl")
	fp := filepath.Join("templates", "index.tmpl")

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		slog.PrintError(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base.tmpl", nil); err != nil {
		slog.PrintError(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request)  { fmt.Fprint(w, "Running") }
func handleFavicon(w http.ResponseWriter, r *http.Request) {}

func InitializeWebinterface() {
	// Serving static files for web front-end
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/status", handleStatus)
	http.HandleFunc("/api/", httpApi)
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.HandleFunc("/", index)
	slog.PrintInfo("Listening on port: " + strconv.Itoa(8080) + " (http)")
	http.ListenAndServe(":"+strconv.Itoa(8080), nil)
}
