package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Ninja struct {
	Name   string
	Weapon string
}

var ninja Ninja // global variable for just illustration purpose

func getNinja(ctx *fiber.Ctx) error {
	// wallace := Ninja{
	// 	Name:   "Ninja",
	// 	Weapon: "Sword",
	// }

	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

func createNinja(ctx *fiber.Ctx) error {
	body := new(Ninja)
	err := ctx.BodyParser(body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString("Bad Request") // it will return nil hence we should always return the actual error instead of returning this TODO:
		return err
	}

	ninja = Ninja{
		Name:   body.Name,
		Weapon: body.Weapon,
	}

	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

func main() {
	app := fiber.New()

	// Just like express
	app.Use(logger.New())
	app.Use(requestid.New())
	// will get x-request-id in header in response

	// if processes are running concurrently then this can produce an error as we can get side effect due to modifying same mutable value here and there.
	app.Get("/hello", func(ctx *fiber.Ctx) error {
		// Modifying context's values
		ctx.Locals("name", "John")
		return ctx.SendString("Hello, " + ctx.Locals("name").(string) + "!")
	})

	app.Get("/error", func(ctx *fiber.Ctx) error {
		return fiber.NewError(400, "This is unauthorised route")
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		var result = ctx.Params("")

		fmt.Println(result)
		return ctx.SendString("HEllo world")
	})

	// app.Get("/ninja", getNinja)

	// app.Post("/ninja", createNinja)

	// ? let me group above things

	ninjaApp := app.Group("/ninja")

	ninjaApp.Get("/", getNinja)
	ninjaApp.Post("/", createNinja)

	app.Static("/static", "./public")
	/* If you want to have a little bit more control regarding the settings for serving static files. You could use the fiber.Static struct to enable specific settings. */

	// ? Add allows you to specify methods as a value
	app.Add("GET", "/add", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Add is good too")
	})

	// ? app.All - same like app.Use() but not bound to specific prefix

	/*
		? Match any request
		app.Use(func(c *fiber.Ctx) error {
		    return c.Next()
		})

		* Match request starting with /api
		app.Use("/api", func(c *fiber.Ctx) error {
		    return c.Next()
		})

		* Match requests starting with /api or /home (multiple-prefix support)
		app.Use([]string{"/api", "/home"}, func(c *fiber.Ctx) error {
		    return c.Next()
		})

		* Attach multiple handlers
		app.Use("/api", func(c *fiber.Ctx) error {
		  c.Set("X-Custom-Header", random.String(32))
		    return c.Next()
		}, func(c *fiber.Ctx) error {
		    return c.Next()
		}) */

	// ? Mount
	// we can mount fiber instance by creating a *Mount
	// can be usefull when it comes to handling sub apps
	// we can also get mount path for current app instance

	// ! Then what is difference between mount and group

	// ? Group
	// we can group routes based on some prefix.

	// ? Route
	// TODO:

	app.Server().MaxConnsPerIP = 1 // Server returns underlying fasthttp server

	// handlers count // Returns amounts of registered handlers

	// Stack - returns original router stack

	data, _ := json.MarshalIndent(app.Stack(), "", " ")
	fmt.Println(string(data))

	// app.Name assigns name of the route ( latest created one only )

	// ? Get Route // get route by name

	// TODO: http testing request response

	// Hooks TODO:

	app.Listen(":8080")
}
