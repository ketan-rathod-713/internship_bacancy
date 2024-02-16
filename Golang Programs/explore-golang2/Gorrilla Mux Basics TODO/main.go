package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/gorilla/mux"
)

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)

	fmt.Println(path)

	file, err := os.Stat(path)

	if os.IsNotExist(err) || file.IsDir() {
		// file does not exist or path is a directory, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Serving Default Ha ha")
	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	fmt.Println(vars["id"], reflect.TypeOf(vars))
	// 	w.Write([]byte(vars["id"]))
	// })

	productRouter := r.PathPrefix("/products").Methods("GET").Subrouter()

	productRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars["id"], reflect.TypeOf(vars))
		w.Write([]byte(vars["id"]))
	}).Methods("POST")
	// We can define Methods of specific route.

	// Router ke upar ham HandleFuc laga sakte he
	productRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WElCOME TO PRODUCTS HOME PAGE"))
	}).Methods("GET")

	// Hence matching ke liye hamare pass Methods,Queries,Headers,etc chise he agar fir bhi hame match me problem he then ham khudka matcher function bhi define kar sakatte he

	bookRouter := r.PathPrefix("/book").Subrouter()
	bookRouter.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return r.ProtoMajor == 0
	})

	// Routes are tested in the order they were added to the router. If two routes match, the first one wins:

	// r.HandleFunc("/specific", specificHandler)
	// r.PathPrefix("/").Handler(catchAllHandler)

	// We create subrouter whenever we want to. It also optimises query as subrouter is tested first.

	// There's one more thing about subroutes. When a subrouter has a path prefix, the inner routes use it as base for their paths:

	/*
		Static Files
		Note that the path provided to PathPrefix() represents a "wildcard": calling PathPrefix("/static/").Handler(...) means that the handler will be passed any request that matches "/static/*". This makes it easy to serve static files with mux:

	*/

	// This will serve files under http://localhost:8000/static/<filename>
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(dir))))

	// spa // single page applications

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	spa := spaHandler{staticPath: "static/", indexPath: "index.html"}
	r.PathPrefix("/good").Handler(spa)

	/* Registered URLs with Name : we can name any route and then can generate url based on it */
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Guys"))
	}).Name("article")

	url, err := r.Get("article").URL("category", "tech", "id", "56")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(url)

	// above same thing works for host and query too
	r.Host("{subdomain}.example.com").
		Path("/articles/{category}/{id:[0-9]+}").
		Queries("filter", "{filter}").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello"))
		}).
		Name("article")

	// url.String() will be "http://news.example.com/articles/technology/42?filter=gorilla"
	url, err = r.Get("article").URL("subdomain", "news",
		"category", "technology",
		"id", "42",
		"filter", "gorilla")

	if err != nil {
		log.Fatal(err)
	}
	log.Println(url)

	// Regex is also supported for header values
	r.Headers("Content-Type", "application/(text|json)")

	/*
		Walking Routes
		The Walk function on mux.Router can be used to visit all of the routes that are registered on a router. For example, the following prints all of the registered routes:
	*/

	/*
		Graceful Shutdowns
	*/

	/*
		Middleware
		Mux supports the addition of middlewares to a Router, which are executed in the order they are added if a match is found, including its subrouters. Middlewares are (typically) small pieces of code which take one request, do something with it, and pass it down to another middleware or the final handler. Some common use cases for middleware are request logging, header manipulation, or ResponseWriter hijacking.

		Mux middlewares are defined using the de facto standard type:

		Typically, the returned handler is a closure which does something with the http.ResponseWriter and http.Request passed to it, and then calls the handler passed as parameter to the MiddlewareFunc. This takes advantage of closures being able access variables from the context where they are created, while retaining the signature enforced by the receivers.
	*/

	// it takes next handler in input and outputs handler function which is made using http.HandleFuc
	type MiddlewareFunc func(http.Handler) http.Handler

	// To use middleware r.Use is used.

	r.Use(loggingMiddleware)

	http.ListenAndServe(":8080", r)
}

// Understand Closure Function
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
