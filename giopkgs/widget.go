// this file was generated by gomacro command: import _i "gioui.org/widget"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	widget "gioui.org/widget"
)

// reflection: allow interpreted code to import "gioui.org/widget"
func init() {
	imports.Packages["gioui.org/widget"] = imports.Package{
	Name: "widget",
	Binds: map[string]r.Value{
		"NewIcon":	r.ValueOf(widget.NewIcon),
	}, Types: map[string]r.Type{
		"Bool":	r.TypeOf((*widget.Bool)(nil)).Elem(),
		"Border":	r.TypeOf((*widget.Border)(nil)).Elem(),
		"ChangeEvent":	r.TypeOf((*widget.ChangeEvent)(nil)).Elem(),
		"Click":	r.TypeOf((*widget.Click)(nil)).Elem(),
		"Clickable":	r.TypeOf((*widget.Clickable)(nil)).Elem(),
		"Editor":	r.TypeOf((*widget.Editor)(nil)).Elem(),
		"EditorEvent":	r.TypeOf((*widget.EditorEvent)(nil)).Elem(),
		"Enum":	r.TypeOf((*widget.Enum)(nil)).Elem(),
		"Float":	r.TypeOf((*widget.Float)(nil)).Elem(),
		"Icon":	r.TypeOf((*widget.Icon)(nil)).Elem(),
		"Image":	r.TypeOf((*widget.Image)(nil)).Elem(),
		"Label":	r.TypeOf((*widget.Label)(nil)).Elem(),
		"Press":	r.TypeOf((*widget.Press)(nil)).Elem(),
		"SelectEvent":	r.TypeOf((*widget.SelectEvent)(nil)).Elem(),
		"SubmitEvent":	r.TypeOf((*widget.SubmitEvent)(nil)).Elem(),
	}, 
	}
}
