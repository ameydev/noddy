package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-docker",
	Short: "build your docker images directly in the k8s cluster.",
	Long: `docker command helps you to build your docker images directly in  your cluster

  $ kubectl docker init
  > initializes the docker agent in the cluster

  $ kubectl docker build --path /some/to/build-context -t my-app:latest
  > builds the docker image
  
  $ kubectl docker cleanup
  > clean-up.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
