package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/wildanfaz/simple-golang-echo/configs"
	"github.com/wildanfaz/simple-golang-echo/internal/routers"
	"github.com/wildanfaz/simple-golang-echo/migrations"
)

func InitCmd(ctx context.Context) {
	rootCmd := &cobra.Command{
		Short: "simple-golang-echo",
	}

	rootCmd.AddCommand(startApp)
	rootCmd.AddCommand(migrateTable)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		panic(err)
	}
}

var startApp = &cobra.Command{
	Short: "start app",
	Use:   "start",
	Run: func(cmd *cobra.Command, args []string) {
		routers.InitRouter()
	},
}

var migrateTable = &cobra.Command{
	Short: "migrate table",
	Use:   "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := configs.InitConfig()

		db := configs.InitMySQL(cfg.MySqlDSN)

		migrations.MigrateTable(context.Background(), db)
	},
}
