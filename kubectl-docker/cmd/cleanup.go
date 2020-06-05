/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "It cleans up the setup by init command",
	Run: func(cmd *cobra.Command, args []string) {
		cleanup()
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}

func cleanup() {
	// 3. kubectl docker cleanup
	// >  * Delete receiver deployment
	//    * stop port-forwarding

	podName := "build-context-receiver"
	cmd1 := exec.Command("kubectl", "delete", "pod", podName)
	_, err := cmd1.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// check if pod is up
	fmt.Println("Everything is cleaned up!")

}
