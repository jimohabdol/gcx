package dev

import (
	"embed"

	"github.com/spf13/cobra"
)

//go:embed templates/import/*.tmpl templates/scaffold/*.tmpl templates/scaffold/internal/*/*.tmpl templates/scaffold/.github/workflows/*.tmpl templates/generate/*.tmpl
var templatesFS embed.FS

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dev",
		Short: "Manage Grafana resources as code",
		Long:  "Tools for managing Grafana resources as code: scaffold new projects, import existing resources from Grafana, and generate typed Go stubs for new resources.",
	}

	cmd.AddCommand(importCmd())
	cmd.AddCommand(scaffoldCmd())
	cmd.AddCommand(generateCmd())

	return cmd
}
