package report

import (
	"time"

	"github.com/weaveworks/scope/common/mtime"
)

// Controls describe the control tags within the Nodes
type Controls map[string]Control

// A Control basically describes an RPC
type Control struct {
	ID    string `json:"id"`
	Human string `json:"human"`
	Icon  string `json:"icon"` // from https://fortawesome.github.io/Font-Awesome/cheatsheet/ please
	Rank  int    `json:"rank"`
}

// Merge merges other with cs, returning a fresh Controls.
func (cs Controls) Merge(other Controls) Controls {
	result := cs.Copy()
	for k, v := range other {
		result[k] = v
	}
	return result
}

// Copy produces a copy of cs.
func (cs Controls) Copy() Controls {
	result := Controls{}
	for k, v := range cs {
		result[k] = v
	}
	return result
}

// AddControl adds c added to cs.
func (cs Controls) AddControl(c Control) {
	cs[c.ID] = c
}

// AddControls adds a collection of controls to cs.
func (cs Controls) AddControls(controls []Control) {
	for _, c := range controls {
		cs[c.ID] = c
	}
}

// NodeControls represent the individual controls that are valid for a given
// node at a given point in time.  It's immutable. A zero-value for Timestamp
// indicated this NodeControls is 'not set'.
type NodeControls struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	Controls  StringSet `json:"controls,omitempty"`
}

// MakeNodeControls makes a new NodeControls
func MakeNodeControls() NodeControls {
	return NodeControls{
		Controls: MakeStringSet(),
	}
}

// Copy is a noop, as NodeControls is immutable
func (nc NodeControls) Copy() NodeControls {
	return nc
}

// Merge returns the newest of the two NodeControls; it does not take the union
// of the valid Controls.
func (nc NodeControls) Merge(other NodeControls) NodeControls {
	if nc.Timestamp.Before(other.Timestamp) {
		return other
	}
	return nc
}

// Add the new control IDs to this NodeControls, producing a fresh NodeControls.
func (nc NodeControls) Add(ids ...string) NodeControls {
	return NodeControls{
		Timestamp: mtime.Now(),
		Controls:  nc.Controls.Add(ids...),
	}
}
