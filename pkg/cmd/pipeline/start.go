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

package pipeline

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tektoncd/cli/pkg/cli"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//const (
//emptyMsg = "No pipelines found"
//header   = "NAME\tAGE\tLAST RUN\tSTARTED\tDURATION\tSTATUS"
//body     = "%s\t%s\t%s\t%s\t%s\t%s\n"
//blank    = "---"
//)

var (
	errNoPipeline      = errors.New("missing pipeline name")
	errInvalidPipeline = errors.New("invalid pipeline name")
)

// NameArg validates that the first argument is a valid pipeline name
func NameArg(p cli.Params) cobra.PositionalArgs {

	return func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errNoPipeline
		}

		c, err := p.Clients()
		if err != nil {
			return err
		}

		name, ns := args[0], p.Namespace()
		_, err = c.Tekton.TektonV1alpha1().Pipelines(ns).Get(name, metav1.GetOptions{})
		if err != nil {
			return errInvalidPipeline
		}

		return nil
	}

}

func startCommand(p cli.Params) *cobra.Command {
	c := &cobra.Command{
		Use:          "start",
		Aliases:      []string{"ls"},
		Short:        "Start pipelines",
		Long:         ``,
		SilenceUsage: true,
		Args:         NameArg(p),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return c
}
