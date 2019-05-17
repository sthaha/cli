// Copyright © 2019 The Tekton Authors.
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

package taskrun

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tektoncd/cli/pkg/cli"
	"github.com/tektoncd/cli/pkg/flags"
)

func Command(p cli.Params) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "taskrun",
		Aliases: []string{"tr", "taskruns"},
		Short:   "Manage taskruns",

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("taskrun requires a subcommand; see help")
			}
			return nil
		},
	}

	cmd.PersistentPreRunE = flags.InitParams(
		p,
		flags.FromKubeConfig(cmd),
		flags.FromNamespace(cmd, flags.Options{Required: true}),
	)

	cmd.AddCommand(listCommand(p))
	return cmd
}
