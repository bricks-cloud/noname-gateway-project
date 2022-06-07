package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"cuelang.org/go/cue/cuecontext"
	"github.com/bricks-cloud/bricks/logger"
	"github.com/spf13/cobra"
)

var (
	ErrInvalidModuleName = errors.New("please input a module name in the following format: <host>/repo/path/within/repo")
	domainNameRegex      = "^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\\.[a-zA-Z]{2,3})"
	pkgNameRegex         = "^[a-zA-Z][a-zA-Z0-9_-]*([a-zA-Z0-9_-]+)*[a-zA-Z0-9]*$"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inititialize the directory as a Bricks project",
	Long:  `Prepare the directory with necessary resources to run Bricks`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("cannot accept more than 1 arg")
		}

		if len(args) < 1 {
			return ErrInvalidModuleName
		}

		err := validatePackageName(args[0])
		if err != nil {
			return err
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			lg = logger.New()
		)

		err := createCueMod(args[0])
		if err != nil {
			lg.Fatal().Err(err).Msg("Failed to initialize a Bricks project")
		}

		lg.Info().Msg("Successfully initialized a Bricks project!")
	},
}

func createCueMod(pkgName string) error {
	if err := os.Mkdir("cue.mod", os.ModePerm); err != nil {
		return err
	}

	cueMod, err := os.Create("./cue.mod/module.cue")
	if err != nil {
		return err
	}

	c := cuecontext.New()
	v, err := c.Encode(fmt.Sprintf(
		`module: "%s"`, pkgName,
	)).Bytes()
	if err != nil {
		return err
	}

	_, err = cueMod.Write(v)
	if err != nil {
		return err
	}

	return nil
}

func validatePackageName(pkgName string) error {
	if strings.HasSuffix(pkgName, "/") {
		return ErrInvalidModuleName
	}

	parts := strings.Split(pkgName, "/")
	if len(parts) < 2 {
		return ErrInvalidModuleName
	}

	for index, part := range parts {
		if index == 0 {
			r, _ := regexp.Compile(domainNameRegex)
			if !r.MatchString(part) {
				return ErrInvalidModuleName
			}
			continue
		}

		r, _ := regexp.Compile(pkgNameRegex)
		if !r.MatchString(part) {
			return ErrInvalidModuleName
		}
	}

	return nil
}
