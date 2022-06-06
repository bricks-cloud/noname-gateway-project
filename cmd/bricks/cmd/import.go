package cmd

import (
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

func init() {
	importCmd.AddCommand(snowflakeImportCmd)
}

var importCmd = &cobra.Command{
	Use:           "import",
	Short:         "Import existing cloud resources",
	Long:          `Import existing cloud resource to bricks-managed schema and state`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var snowflakeImportCmd = &cobra.Command{
	Use:     "snowflake",
	Aliases: []string{"sf"},
	Short:   "snowflake resources",
	Long:    `Import snowflake resources to bricks-managed schema and state`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
