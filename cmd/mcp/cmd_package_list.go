package main

import (
	"mcp/internal/integrations/sql"

	"github.com/spf13/cobra"
)

var (
	cmdPackageList = &cobra.Command{
		Use:     "list",
		Short:   "List installed packages.",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
			repo, err := sql.NewSQLDatabaseIntegrationsRepository(db)
			cobra.CheckErr(err)

			intalled, err := repo.ListIntegrations(cmd.Context())
			cobra.CheckErr(err)

			if len(intalled) == 0 {
				cmd.PrintErrf("No packages installed\n")
				return
			}

			cmd.PrintErrf("Found %d packages\n", len(intalled))

			for _, integration := range intalled {
				cmd.PrintErrf("\n")
				cmd.PrintErrf("%s %s\n", integration.Manifest.Name, integration.Manifest.Version)
			}
		},
	}
)