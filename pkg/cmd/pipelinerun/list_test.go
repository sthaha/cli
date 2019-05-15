// Copyright © 2019 The tektoncd Authors.
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
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jonboulle/clockwork"
	"github.com/knative/pkg/apis"
	"github.com/spf13/cobra"
	"github.com/tektoncd/cli/pkg/test"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/reconciler/v1alpha1/pipelinerun/resources"
	pipelinetest "github.com/tektoncd/pipeline/test"
	tb "github.com/tektoncd/pipeline/test/builder"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestListPipelineRuns(t *testing.T) {
	now := time.Now()
	prs := pipelineRuns(now)

	tests := []struct {
		name     string
		command  *cobra.Command
		args     []string
		expected []string
	}{
		{
			name:    "by pipeline name",
			command: command(prs, now),
			args:    []string{"list", "bar", "-n", "foo"},
			expected: []string{
				"NAME    STARTED      DURATION   STATUS      ",
				"pr1-1   1 hour ago   1 minute   Succeeded   ",
				"",
			},
		},
		{
			name:    "all in namespace",
			command: command(prs, now),
			args:    []string{"list", "-n", "foo"},
			expected: []string{
				"NAME    STARTED      DURATION   STATUS               ",
				"pr1-1   1 hour ago   1 minute   Succeeded            ",
				"pr2-1   1 hour ago   ---        Succeeded(Running)   ",
				"pr2-2   1 hour ago   1 minute   Failed               ",
				"",
			},
		},
		{
			name:    "print by template",
			command: command(prs, now),
			args:    []string{"list", "-n", "foo", "-o", "jsonpath={range .items[*]}{.metadata.name}{\"\\n\"}{end}"},
			expected: []string{
				"pr1-1",
				"pr2-1",
				"pr2-2",
				"",
			},
		},
		{
			name:     "empty list",
			command:  command(prs, now),
			args:     []string{"list", "-n", "random"},
			expected: []string{msgNoPRsFound, ""},
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			got, err := test.ExecuteCommand(td.command, td.args...)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if d := cmp.Diff(strings.Join(td.expected, "\n"), got); d != "" {
				t.Errorf("Unexpected output mismatch: \n%s\n", d)
			}
		})
	}
}

func command(prs []*v1alpha1.PipelineRun, now time.Time) *cobra.Command {
	// fake clock advanced by 1 hour
	clock := clockwork.NewFakeClockAt(now)
	clock.Advance(time.Duration(60) * time.Minute)

	cs, _ := pipelinetest.SeedTestData(pipelinetest.Data{PipelineRuns: prs})

	p := &test.Params{Client: cs.Pipeline, Clock: clock}

	return Command(p)
}

func pipelineRuns(start time.Time) []*v1alpha1.PipelineRun {
	aMinute, _ := time.ParseDuration("1m")

	prsData := []struct {
		name       string
		ns         string
		pipeline   string
		status     corev1.ConditionStatus
		reason     string
		startTime  time.Time
		finishTime time.Time
	}{
		{
			name:       "pr1-1",
			ns:         "foo",
			pipeline:   "bar",
			status:     corev1.ConditionTrue,
			reason:     resources.ReasonSucceeded,
			startTime:  start,
			finishTime: start.Add(aMinute),
		},
		{
			name:      "pr2-1",
			ns:        "foo",
			pipeline:  "random",
			status:    corev1.ConditionTrue,
			reason:    resources.ReasonRunning,
			startTime: start,
		},
		{
			name:       "pr2-2",
			ns:         "foo",
			pipeline:   "random",
			status:     corev1.ConditionFalse,
			reason:     resources.ReasonFailed,
			startTime:  start,
			finishTime: start.Add(aMinute),
		},
	}

	prs := []*v1alpha1.PipelineRun{}
	for _, data := range prsData {
		pr := tb.PipelineRun(data.name, data.ns,
			tb.PipelineRunLabel("tekton.dev/pipeline", data.pipeline),
			tb.PipelineRunStatus(
				tb.PipelineRunStatusCondition(apis.Condition{
					Status: data.status,
					Reason: data.reason,
				}),
				tb.PipelineRunStartTime(data.startTime),
			),
		)

		pr.Status.CompletionTime = &metav1.Time{Time: data.finishTime}
		prs = append(prs, pr)
	}

	return prs
}
