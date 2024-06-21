# Explore Golang 2

After covering fundamentals now we will conver some good topics.

## Topics

1. Package net/http
    - http server implementations, servemux
    - http TLS impelementation ???
    - http Client ~ see it
    - Imp Functions
        - MaxBytesReader for limiting size of incoming request
        - NotFound to send 404 not found error
        - Redirect(w, r, url, code) (means page reload right ?)
        - ServeContent ??
        - ServeFile(w, r, name)
        - SetCookie(w, cookie)
        - StatusText(code int)
    - Client
        - PostForm issues a POST to the specified URL, with data's keys and values URL-encoded as the request body.
        - The Content-Type header is set to application/x-www-form-urlencoded. To set other headers, use NewRequest and Client.Do.
    - Server
        - FileServer(root FileSystem) http.Handler for simple static web server or static content.
        - DotFileHinding and StripPrefix see it.
        - Header returns key value pairs in an http header
        - key is case insensitive
    - Request
        - AddCookie
        - BasicAuth
        - ParseForm populates r.Form and r.PostForm

2. GORM
    - gorm.Model with 4 fields
    - gorm:"primaryKey" tag to specify primary key
    - "embedded" tag to embed any struct. can also add embed prefix : embeddedPreffix:author_ . hence Author {Name string} will be now author_name
    - Field Tags 
        - tags are insensitive, camel case is convention
        - column - column db name
        - size
        - primaryKey
        - unique
        - default to specify default value
        - precision ??!!
        - not null
        - autoIncrement
        - check IMP TODO: creates check contraints: https://gorm.io/docs/constraints.html
        - Association Tags
            - 
    - Connecting To database
        - officially supports postgresql, mysql, etc.
        - postgresql gorm.io/driver/postgres"
        - 
    - CRUD
        - Create : db.Create(&user) for multiple records pass slice of users
        - Check result for error, or RowsAffected

        - Create record with selected fields db.Select("Name", "Age").create(&user)
        - For above case we can also use Omit

        - Batch insert : start transaction too. db.CreateinBatches(users, 100) with 100 batch size

    - User defined Hooks ( we can also skip hooks in session mode by marking SkipHooks: true)
        - before create
        - before save
        - after save
        - after create
    
    - 


3. Package sql
    - 


4. Gorrilla Mux
    - Mux means http request multiplexer.
    - Implements request router and dispatcher
    - Imporntant Methods
        - NewRouter() *Router
        - r.HandleFunc(path string, func(w, r){}) *Route
        - r.Handle(path string, handler http.Handler) *Route
        - r.Methods(...string) *Route
        - r.PathPrefix(string) *Route
        - Middleware


