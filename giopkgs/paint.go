// this file was generated by gomacro command: import _i "gioui.org/op/paint"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	paint "gioui.org/op/paint"
)

// reflection: allow interpreted code to import "gioui.org/op/paint"
func init() {
	imports.Packages["gioui.org/op/paint"] = imports.Package{
	Name: "paint",
	Binds: map[string]r.Value{
		"Fill":	r.ValueOf(paint.Fill),
		"FillShape":	r.ValueOf(paint.FillShape),
		"NewImageOp":	r.ValueOf(paint.NewImageOp),
	}, Types: map[string]r.Type{
		"ColorOp":	r.TypeOf((*paint.ColorOp)(nil)).Elem(),
		"ImageOp":	r.TypeOf((*paint.ImageOp)(nil)).Elem(),
		"LinearGradientOp":	r.TypeOf((*paint.LinearGradientOp)(nil)).Elem(),
		"PaintOp":	r.TypeOf((*paint.PaintOp)(nil)).Elem(),
	}, 
	}
}
