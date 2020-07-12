package cmd

import (
	"github.com/genhoi/users/app"
	"github.com/spf13/cobra"
)

func init() {
	importCmd := &cobra.Command{
		Use:   "import",
		Short: "Импорт пользователей из файлов",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			container := app.Container()

			importer := container.UserImporter()
			err := importer.Import(args...)
			if err != nil {
				container.Logger().Error(err)
			}
		},
	}

	rootCmd.AddCommand(importCmd)
}
