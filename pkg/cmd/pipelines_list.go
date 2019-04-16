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
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1alpha1"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: run,
}

type clients struct {
	PipelineClient v1alpha1.PipelineInterface
	//TaskClient             v1alpha1.TaskInterface
	//TaskRunClient          v1alpha1.TaskRunInterface
	//PipelineRunClient      v1alpha1.PipelineRunInterface
	//PipelineResourceClient v1alpha1.PipelineResourceInterface
}

func run(cmd *cobra.Command, args []string) error {
	var err error
	c := &clients{}

	cs, err := versioned.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Printf("failed to create client from config %s  %s", kubeCfgFile, err)
		return err
	}
	ns := "test"
	c.PipelineClient = cs.TektonV1alpha1().Pipelines(ns)
	//c.TaskClient = cs.TektonV1alpha1().Tasks(namespace)
	//c.TaskRunClient = cs.TektonV1alpha1().TaskRuns(namespace)
	//c.PipelineRunClient = cs.TektonV1alpha1().PipelineRuns(namespace)
	//c.PipelineResourceClient = cs.TektonV1alpha1().PipelineResources(namespace)
	return nil
}

func init() {
	pipelinesCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
