package testutil

import (
	"github.com/jonboulle/clockwork"
	"github.com/tektoncd/cli/pkg/cli"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
)

type TestParams struct {
	ns, kc string
	Client versioned.Interface
	Clock clockwork.Clock
}

var _ cli.Params = &TestParams{}

func (p *TestParams) SetNamespace(ns string) {
	p.ns = ns
}
func (p *TestParams) Namespace() string {
	return p.ns
}

func (p *TestParams) SetKubeConfigPath(path string) {
	p.kc = path
}

func (p *TestParams) KubeConfigPath() string {
	return p.kc
}

func (p *TestParams) Clientset() (versioned.Interface, error) {
	return p.Client, nil
}

func (p *TestParams) Time() clockwork.Clock {
	if p.Clock == nil {
		p.Clock = clockwork.NewFakeClock()
	}
	
	return p.Clock
}
