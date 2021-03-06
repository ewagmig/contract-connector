package server

import (
	_ "github.com/ewagmig/contract-connector/migration/source"
	"github.com/ewagmig/contract-connector/server"
	"github.com/ewagmig/contract-connector/version"
	"github.com/spf13/cobra"
)

func migrateCmd() *cobra.Command {
	return serverMigrateCmd
}

func migrate(mode string) error {
	logger.Infof("Start to migrate in %s mode, with %s", mode, version.Version())
	return server.RunMigration()
}

var serverMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the baas server",
	Long:  "Migrate the baas server to new database or data schema",
	RunE: func(cmd *cobra.Command, args []string) error {
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			return err
		}
		return migrate(mode)
	},
}
