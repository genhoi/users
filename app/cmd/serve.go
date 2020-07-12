package cmd

import (
	"context"
	"github.com/genhoi/users/app"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func init() {

	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Запуск сервера",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			container := app.Container()

			air := container.Air()
			air.GET("/search", container.SearchAction().Get)
			air.HEAD("/search", container.SearchAction().Head)
			air.GET("/*", container.UiAction().Get)

			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

			go func() {
				<-c
				air.Shutdown(context.Background())
			}()

			air.Serve()
		},
	}

	rootCmd.AddCommand(serveCmd)
}
