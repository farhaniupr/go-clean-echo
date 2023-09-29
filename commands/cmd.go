package commands

import (
	"clean-go-echo/api/middlewares"
	"clean-go-echo/api/routes"
	"clean-go-echo/library"
	"context"
	"log"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

type App struct {
	*cobra.Command
}

var RootApp = NewApp()

var info = &cobra.Command{
	Use:   "clean-echo",
	Short: "Clean Architecture Echo by Papay",
	Long:  "Clean Architecture Echo by Papay",
	Run: func(c *cobra.Command, args []string) {
		opts := fx.Options(
			fx.Invoke(func(
				route routes.Routes,
				router library.RequestHandler,
				middleware middlewares.Middlewares,
			) {

				route.Setup()
				middleware.Setup()

				_ = router.Echo.Start(":" + library.ModuleEnv().ServerPort)
			}),
		)
		ctx := context.Background()
		app := fx.New(CommonModules, opts)
		err := app.Start(ctx)
		defer app.Stop(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	},
}

func NewApp() App {
	cmd := App{
		Command: info,
	}
	// cmd.AddCommand()
	return cmd
}
