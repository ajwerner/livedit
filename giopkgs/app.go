// this file was generated by gomacro command: import _i "gioui.org/app"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	app "gioui.org/app"
)

// reflection: allow interpreted code to import "gioui.org/app"
func init() {
	imports.Packages["gioui.org/app"] = imports.Package{
	Name: "app",
	Binds: map[string]r.Value{
		"DataDir":	r.ValueOf(app.DataDir),
		"Main":	r.ValueOf(app.Main),
		"MaxSize":	r.ValueOf(app.MaxSize),
		"MinSize":	r.ValueOf(app.MinSize),
		"NewWindow":	r.ValueOf(app.NewWindow),
		"Size":	r.ValueOf(app.Size),
		"Title":	r.ValueOf(app.Title),
	}, Types: map[string]r.Type{
		"Option":	r.TypeOf((*app.Option)(nil)).Elem(),
		"Window":	r.TypeOf((*app.Window)(nil)).Elem(),
	}, 
	}
}