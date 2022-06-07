package cmd

import (
	"context"

	"github.com/bricks-cloud/bricks/pkg/providers/snowflake/importer"
	"github.com/spf13/cobra"
)

type ImportOptions struct {
	Resources   []string
	PathPattern string
	PathOutput  string
	State       string
	Bucket      string
	Profile     string
	Zone        string
	Regions     []string
	Projects    []string
}

const DefaultPathPattern = "{output}/{provider}/{service}/"
const DefaultPathOutput = "generated"
const DefaultState = "local"

var snowflakeOptions = importer.Config{}

func init() {
	importCmd.AddCommand(snowflakeImportCmd)

	snowflakeImportCmd.PersistentFlags().StringVar(&snowflakeOptions.User, "user", "", "")
	snowflakeImportCmd.PersistentFlags().StringVar(&snowflakeOptions.Password, "password", "", "")
	snowflakeImportCmd.PersistentFlags().StringVar(&snowflakeOptions.Database, "database", "", "")
	snowflakeImportCmd.PersistentFlags().StringVar(&snowflakeOptions.Schema, "schema", "", "")
	snowflakeImportCmd.PersistentFlags().StringVar(&snowflakeOptions.Account, "account", "", "")
	snowflakeImportCmd.PersistentFlags().StringVar(&snowflakeOptions.Warehouse, "warehouse", "", "")

}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import existing cloud resources",
	Long:  `Import existing cloud resource to bricks-managed schema and state`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var snowflakeImportCmd = &cobra.Command{
	Use:     "snowflake",
	Aliases: []string{"sf"},
	Short:   "snowflake resources",
	Long:    `Import snowflake resources to bricks-managed schema and state`,
	Run: func(cmd *cobra.Command, args []string) {
		importer.CreateDatabases(context.Background(), &snowflakeOptions)
	},
}
