package cmd

import (
	"os"

	"github.com/bricks-cloud/bricks/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().String("log-format", "auto", "Log format (auto, plain, tty, json)")
	rootCmd.PersistentFlags().StringP("log-level", "l", "info", "Log level")
	rootCmd.PersistentFlags().String("terraform-work-dir", ".workdir", "Terraform work directory path")
	rootCmd.PersistentFlags().String("cue-work-dir", path, "Work directory path")

	rootCmd.AddCommand(
		deployCmd,
	)

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "bricks",
	Short: "Bricks is a Revolutionary Infrastructure As Code Tool",
	Long: `A Revolutionary Infrastructure as Code Tool built with
				  love by Bricks Cloud Technologies, Inc. and friends in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	var (
		lg = logger.New()
	)

	if err := rootCmd.Execute(); err != nil {
		lg.Fatal().Err(err).Msg("failed to execute command")
	}
}

// func excludeFilesAndDirectories(path string, d fs.DirEntry) bool {
// 	name := d.Name()
// 	if strings.Contains(path, "vendor") {
// 		return true
// 	}

// 	if filepath.Ext(name) != ".cue" {
// 		return true
// 	}

// 	if name == "module.cue" {
// 		return true
// 	}

// 	return false
// }

// func findCueFiles(path string, d fs.DirEntry, err error) error {
// 	if err != nil {
// 		return err
// 	}

// 	if excludeFilesAndDirectories(path, d) {
// 		return nil
// 	}

// 	fmt.Fprintln(os.Stderr, d.Name())
// 	fmt.Fprintln(os.Stderr, path)

// 	return nil
// }
