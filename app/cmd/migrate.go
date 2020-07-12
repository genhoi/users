package cmd

import (
	"github.com/genhoi/users/app"
	"github.com/spf13/cobra"
)

func init() {

	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Запуск миграций",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			container := app.Container()

			err := container.Migrate().Up()
			if err != nil {
				container.Logger().Error(err)
			}
		},
	}

	rootCmd.AddCommand(migrateCmd)
}
