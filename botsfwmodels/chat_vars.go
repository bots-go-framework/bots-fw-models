package botsfwmodels

import (
	"github.com/strongo/slice"
	"slices"
)

type chatVars struct {
	Vars    map[string]string `dalgo:"vars,omitempty" firestore:"vars,omitempty"`
	changed []string
	deleted []string
}

// GetVar returns a chat variable
func (v *chatVars) GetVar(key string) string {
	if v.Vars == nil {
		return ""
	}
	return v.Vars[key]
}

// SetVar sets a chat variable
func (v *chatVars) SetVar(key, value string) {
	if v.Vars == nil {
		v.Vars = make(map[string]string)
	} else if v.Vars[key] == value {
		return
	}
	v.Vars[key] = value
	slice.RemoveInPlaceByValue(v.deleted, key)
	if !slices.Contains(v.changed, key) {
		v.changed = append(v.changed, key)
	}
}

// DelVar deletes a chat variable
func (v *chatVars) DelVar(key string) {
	if v.Vars == nil {
		return
	}
	if _, ok := v.Vars[key]; !ok {
		return
	}
	delete(v.Vars, key)
	v.changed = slice.RemoveInPlaceByValue(v.changed, key)
	if !slices.Contains(v.deleted, key) {
		v.deleted = append(v.deleted, key)
	}
}

// HasChangedVars returns true if vars have been changed
func (v *chatVars) HasChangedVars() bool {
	return len(v.changed) > 0 || len(v.deleted) > 0
}
