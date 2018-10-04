package gotversion

import (
	"errors"
	"sort"
	"sync"
)

// Error conditions
var (
	ErrOutputNotFound = errors.New("output not found")
)

// Outputs holds all the output plugins
type Outputs struct {
	lock     sync.Mutex
	registry map[string]Outputer
}

// Outputer interface
type Outputer interface {
	Output(base *Base) error
}

// NewOutputs returns a new output registry
func NewOutputs() *Outputs {
	return &Outputs{}
}

// Registered enumerates the names of all registered  plugins.
func (o *Outputs) Registered() []string {
	o.lock.Lock()
	defer o.lock.Unlock()
	keys := []string{}
	for k := range o.registry {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Register registers a Outputer by name. This
// is expected to happen during app startup.
func (o *Outputs) Register(name string, output Outputer) {
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.registry != nil {
		_, found := o.registry[name]
		if found {
			return // Registered twice
		}
	} else {
		o.registry = map[string]Outputer{}
	}
	o.registry[name] = output
}

// Output performs the output for the given outputer, if present
func (o *Outputs) Output(name string, base *Base) error {
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.registry != nil {
		output, found := o.registry[name]
		if found {
			return output.Output(base)
		}
	}
	return ErrOutputNotFound
}
