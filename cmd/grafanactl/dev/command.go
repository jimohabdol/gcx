package dev

import (
	"embed"

	cmdconfig "github.com/grafana/grafanactl/cmd/grafanactl/config"
	"github.com/grafana/grafanactl/cmd/grafanactl/linter"
	"github.com/grafana/grafanactl/cmd/grafanactl/resources"
	"github.com/spf13/cobra"
)

//go:embed templates/import/*.tmpl templates/scaffold/*.tmpl templates/scaffold/internal/*/*.tmpl templates/scaffold/.github/workflows/*.tmpl templates/generate/*.tmpl
var templatesFS embed.FS

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dev",
		Short: "Manage Grafana resources as code",
		Long:  "Tools for managing Grafana resources as code: scaffold new projects, import existing resources from Grafana, generate typed Go stubs for new resources, lint resources, and serve resources locally.",
	}

	configOpts := &cmdconfig.Options{}
	configOpts.BindFlags(cmd.PersistentFlags())

	cmd.AddCommand(importCmd())
	cmd.AddCommand(scaffoldCmd())
	cmd.AddCommand(generateCmd())
	cmd.AddCommand(linter.Command())
	cmd.AddCommand(resources.ServeCmd(configOpts))

	return cmd
}
