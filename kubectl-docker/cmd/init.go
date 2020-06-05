package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		setUp()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func setUp() {
	// 	1. kubectl docker init
	// >  * deploy the receiver pod (kubectl apply -f noddy/services/receiver/deployment.yaml)
	//    * port-forward the pod in background (kubectl port-forward 6000:5000 &)
	//    * Print build-context receiver is ready..

	podName := "build-context-receiver"
	cmd1 := exec.Command("kubectl", "apply", "-f", "/home/amey/go/src/github.com/ameydev/noddy/receiver/deployment.yaml")
	_, err := cmd1.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// check if pod is up
	for checkPodStatus(podName) != true {
		time.Sleep(2 * time.Second)
		fmt.Println("waiting for pod ", podName)
	}

	fmt.Println("Build context receiver is ready!")

	// Lets port-forward it on localhost:6000
	// build-context-receiver 6000:5000
	cmd1 = exec.Command("kubectl", "port-forward", "build-context-receiver", "6000:5000", "&")
	_, _ = cmd1.Output()
	fmt.Println("receiver is ready.")

}

func checkPodStatus(podName string) bool {
	key := "status"
	cmd1 := exec.Command("kubectl", "get", "pod", podName)
	stdout, err := cmd1.Output()
	if err != nil {
		return false
	}
	components := strings.Fields(string(stdout))

	keyIndex := getIndex(components, key)
	status := components[keyIndex]
	if status == "Running" {
		return true
	} else {
		return false
	}

}

func getIndex(components []string, key string) int {
	key = strings.ToUpper(key)
	var index int
	for i := 0; i < len(components); i++ {
		if components[i] == key {
			index = i + (len(components) / 2)
		}
	}
	return index
}
