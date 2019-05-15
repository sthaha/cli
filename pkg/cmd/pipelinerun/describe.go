// Copyright © 2019 The Knative Authors.
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

package pipelinerun

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tektoncd/cli/pkg/cli"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func describeCommand(p cli.Params) *cobra.Command {
	eg := `
# Show details of a pipelinerun including Resources, Tasks, Time taken, etc
tkn pipelinerun -n bar describe run-name

`

	c := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"desc"},
		Short:   "Shows information about pipelineruns",
		Example: eg,
		Args:    cobra.MinimumNArgs(1), // need the piplinerun

		RunE: func(cmd *cobra.Command, args []string) error {
			runName := args[0]

			cs, err := p.Clientset()
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "failed to created tekton client\n")
				return err
			}

			pr, err := getPipelineRun(cs, p.Namespace(), runName)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "Failed to find pipelinerun %q \n", runName)
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Got pipelinerun: %s", pr.Name)
			return nil
		},
	}

	return c
}

func getPipelineRun(cs versioned.Interface, ns, name string) (*v1alpha1.PipelineRun, error) {
	return cs.TektonV1alpha1().PipelineRuns(ns).Get(name, metav1.GetOptions{})
}
