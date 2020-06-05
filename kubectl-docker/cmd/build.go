package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var path, tag string

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build flags",
	Short: "Similar to docker build command.",
	Long: ` Just pass the build context path and the -t flag as args
	optional -f.

	kubectl docker build -t my-image:latest .
	
	`,
	PreRunE: func(Cmd *cobra.Command, args []string) error {
		// if len(args) == 0 {
		// 	return fmt.Errorf("specify the path")
		// }
		if path == "" {
			return fmt.Errorf("specify the path -p <context-path>")
		}
		if tag == "" {
			return fmt.Errorf("specify the tag -t <image-name>")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		build(args)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().StringVarP(&tag, "tag", "t", "", "path of build context")
	buildCmd.Flags().StringVarP(&path, "path", "p", "", "path of build context")
}

func build(args []string) {

	// 	kubectl docker build --path /some/to/build-context -t my-app:latest
	// >  * Tar the contents of --path `/some/to/build-context` with name of my-app.tar.gz
	//    * POST request with my-app.tar.gz
	//    * Deploy the DinD container pod (kubectl apply -f docker-agent.yaml)
	//    * Print the node on which DinD, receiver pods are running
	//         (The node has to be added as node-selector in my-app pod spec)
	//    * Delete the DinD pod(the src/ saved on hostPath of node must get clean-up ed).

	// contextPath := path
	fmt.Println(path)
	// for i := 0; i < len(args); i++ {
	// 	fmt.Println(string(args[0]))
	// }

	// Create tarball of contextPath
	// parts := strings.Split(contextPath, "/")
	// dirPath := parts[len(parts)-1]
	buildScript := "/home/amey/go/src/github.com/ameydev/noddy/kubectl-docker/scripts/client.py"

	pyCmd := exec.Command("python3", buildScript, path, tag)
	_, err := pyCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Build context is sent!")

	// kubectl create configmap special-config --from-literal=CONTEXT_PATH=/workspace/test-data/ --from-literal=TAG_NAME=test1
	kCmd := exec.Command("kubectl", "create", "configmap", "docker-context", "--from-literal=CONTEXT_PATH=/workspace/"+tag, "--from-literal=TAG_NAME="+tag)
	_, err1 := kCmd.Output()
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Println("configmap context is created!")

	kCmd = exec.Command("kubectl", "apply", "-f", "/home/amey/go/src/github.com/ameydev/noddy/kubectl-docker/templates/docker-agent.yaml")
	_, err1 = kCmd.Output()
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Println("Started the build")

}

func runDockerAgent(imageName string) {

}

// func createTar(path, tarName string) {

// 	// to be replaced by proggramming way :p
// 	// tar cvzf tarName.tar.gz /home/MyImages
// 	tarCmd := exec.Command("tar", "cvzf", tarName, path)
// 	_, err := tarCmd.Output()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// }
