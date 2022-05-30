package cmd

import (
	"context"
	"os"

	"github.com/bricks-cloud/bricks/compiler"
	"github.com/bricks-cloud/bricks/logger"
	"github.com/bricks-cloud/bricks/pkg/terraform"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy cloud infrastructures",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			lg = logger.New()
		)

		cueWorkDir := viper.GetString("cue-work-dir")
		v, err := compiler.Build(context.Background(), cueWorkDir, nil)
		if err != nil {
			lg.Fatal().Msg("cannot load")
		}

		tfWorkDir := viper.GetString("terraform-work-dir")
		if err := os.MkdirAll(tfWorkDir, os.ModePerm); err != nil {
			lg.Fatal().Msgf("Failed to create work dir: %v", err)
		}

		if err := terraform.Run(context.Background(), tfWorkDir, v.JSON()); err != nil {
			lg.Fatal().Err(err).Msg("failed to deploy via Terraform")
		}

		lg.Info().Msg("Successful Deploy!")
	},
}
