// this file was generated by gomacro command: import _i "gioui.org/x/eventx"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	eventx "gioui.org/x/eventx"
)

// reflection: allow interpreted code to import "gioui.org/x/eventx"
func init() {
	imports.Packages["gioui.org/x/eventx"] = imports.Package{
	Name: "eventx",
	Binds: map[string]r.Value{
		"Combine":	r.ValueOf(eventx.Combine),
		"Enspy":	r.ValueOf(eventx.Enspy),
	}, Types: map[string]r.Type{
		"CombinedQueue":	r.TypeOf((*eventx.CombinedQueue)(nil)).Elem(),
		"EventGroup":	r.TypeOf((*eventx.EventGroup)(nil)).Elem(),
		"Spy":	r.TypeOf((*eventx.Spy)(nil)).Elem(),
	}, 
	}
}
