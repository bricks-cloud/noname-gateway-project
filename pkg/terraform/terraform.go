package terraform

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-exec/tfexec"
)

var terraformExecPath string = "terraform"

// func init() {
// 	installer := &releases.ExactVersion{
// 		Product: product.Terraform,
// 		Version: version.Must(version.NewVersion("1.2.0")),
// 	}

// 	var err error
// 	terraformExecPath, err = installer.Install(context.Background())
// 	if err != nil {
// 		log.Fatalf("error installing Terraform: %v", err)
// 	}
// }

func Run(ctx context.Context, tfContent []byte) error {
	workingDir := os.TempDir()
	if err := os.WriteFile(filepath.Join(workingDir, "new.tf.json"), tfContent, 0644); err != nil {
		return err
	}
	tf, err := tfexec.NewTerraform(workingDir, terraformExecPath)
	if err != nil {
		return fmt.Errorf("error running NewTerraform: %s", err)
	}

	err = tf.Init(ctx)
	if err != nil {
		return fmt.Errorf("error running Init: %s", err)
	}
	if err := tf.Apply(ctx); err != nil {
		return err
	}
	return nil
}
