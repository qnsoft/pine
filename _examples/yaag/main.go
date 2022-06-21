package main

import (
	"github.com/betacraft/yaag/yaag"
	"github.com/qnsoft/pine"
	"github.com/qnsoft/pine/middlewares/yaag_pine"
)

func main() {
	app := pine.New()

	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "Pine",
		DocPath:  "./index.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(yaag_pine.New())

	app.Static("/", "./")
	app.GET("/json", func(ctx *pine.Context) {
		ctx.Render().JSON(pine.H{"result": "Hello World!"})
	})
	app.GET("/plain", func(ctx *pine.Context) {
		ctx.Render().Text("Hello World!")
	})

	app.Run(pine.Addr(":9528"))
}
