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

package cli

import "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"

// Params interface provides
type Params interface {
	// SetKubeConfigPath uses the kubeconfig path to instantiate clientset
	// returned by Clientset function
	SetKubeConfigPath(string)
	Clientset() versioned.Interface

	// SetNamespace can be used to store the namespace parameter that is needed
	// by most commands
	SetNamespace(string)
	Namespace() string
}
