// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	var err error

	cs, err := versioned.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Printf("failed to create client from config %s  %s", kubeCfgFile, err)
		return err
	}
	c := cs.TektonV1alpha1().Pipelines(namespace)
	ps, err := c.List(v1.ListOptions{})
	if err != nil {
		fmt.Printf("failed to list pipelines from namespace %s  %s", namespace, err)
		return err
	}

	for _, v := range ps.Items {
		fmt.Printf(v.Name)
	}
	return nil
}

func init() {
	pipelinesCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

}
