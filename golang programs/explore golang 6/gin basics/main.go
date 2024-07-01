package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data,omitempty"`
}

type Parent struct {
	Name  string
	Age   int
	Phone string
}

type Student struct {
	Id     int
	Name   string
	Parent Parent
}

type FileData struct {
	Name    string
	Content string
	Path    string
	Size    string
}

func middleware(ctx *gin.Context) {

	ctx.Set("userid", 20)
	ctx.Set("student", Student{Name: "ketan rathod"})
	ctx.Set("username", "ketan rathod")

	// called inside the middlewares for calling next function
	ctx.Next()
}

func main() {
	r := gin.Default()

	log.Println("Running gin in", gin.Mode(), "mode")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"great":   "wow",
			"object": gin.H{
				"value": 5,
				"good":  10,
			},
		})
	})

	r.GET("/student", func(ctx *gin.Context) {
		var s Student = Student{
			Name: "Manav Vyas",
			Parent: Parent{
				Name:  "Golang",
				Phone: "9099238472",
			},
		}

		ctx.JSON(202, s)
	})

	r.GET("/dir", func(ctx *gin.Context) {
		workingDir, err := os.Getwd()
		if err != nil {
			ctx.JSON(400, Response{Error: "error reading current working directory"})
			ctx.Abort()
			return
		}

		fs := gin.Dir(workingDir, true)
		log.Println("File System", fs)

		file, err := fs.Open("main.go")
		if err != nil {
			log.Println("Error opening file")
			ctx.JSON(400, Response{Error: "error opening file"})
			ctx.Abort()
			return
		}

		var buffer []byte = make([]byte, 1024)
		n, err := file.Read(buffer)
		if err != nil {
			log.Println("error reading file to buffer")
			ctx.JSON(400, Response{Error: "error reading file to buffer"})
			ctx.Abort()
			return
		}
		log.Println("bytes read", n)

		ctx.JSON(200, Response{Data: FileData{Content: string(buffer), Name: "main.go", Size: "10 mb", Path: workingDir}})

	})

	r.GET("/student/:id", func(ctx *gin.Context) {
		// To manually set params for testing
		// TODO: It is not wokring because it is for testing purpose when we want to add Params
		ctx.AddParam("id", "ketan-rathod-testing")
		// Note: For getting params from cotext.
		id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(400, Response{Error: "error getting id param value"})
			ctx.Abort()
			return
		}

		ctx.JSON(200, Response{Data: Student{Id: 0, Name: id, Parent: Parent{Name: "somethigns parent"}}})
	})

	r.GET("/getset", middleware, func(ctx *gin.Context) {
		userid, _ := ctx.Get("userid")
		username, _ := ctx.Get("username")
		student, _ := ctx.Get("student")

		ctx.JSON(200, gin.H{
			"userid":   userid,
			"username": username,
			"student":  student,
		})
	})

	// Asci json
	// serialize as json and sets content type to application/json
	r.GET("ascii", func(ctx *gin.Context) {
		ctx.AsciiJSON(200, Student{Name: "ketone ::;; rathod"})
	})

	r.GET("/ip", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"ip": ctx.ClientIP(),
		})
	})

	// Context.Copy returns a copy of context. This can be used if we want to pass it to a goroutine.

	r.Run() // listen and serve on 0.0.0.0:8080
}

// NOTE:
// func(ctx *gin.Context) can be used for almost all cases
// It can also serves as the middleware and also for the response

// gin.H is for sending response of type map[string]interface{} type to json.
// ctx.Json is used for serializing structs, maps to json with valid status code.

// gin.Default() will return a router or *gin.Engine
// Group on router can be used to group similar routes and apply middlewares on it.
// For eg.
// publicRoute := router.Group("/")
//	publicRoute.Use(gs.VerifyJwtMiddleware)

// TODO: Using context for passing and getting values !
// we can pass values between middlewares and hadlers in gin using context
//
