// this file was generated by gomacro command: import _i "gioui.org/f32"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	f32 "gioui.org/f32"
)

// reflection: allow interpreted code to import "gioui.org/f32"
func init() {
	imports.Packages["gioui.org/f32"] = imports.Package{
	Name: "f32",
	Binds: map[string]r.Value{
		"NewAffine2D":	r.ValueOf(f32.NewAffine2D),
		"Pt":	r.ValueOf(f32.Pt),
		"Rect":	r.ValueOf(f32.Rect),
	}, Types: map[string]r.Type{
		"Affine2D":	r.TypeOf((*f32.Affine2D)(nil)).Elem(),
		"Point":	r.TypeOf((*f32.Point)(nil)).Elem(),
		"Rectangle":	r.TypeOf((*f32.Rectangle)(nil)).Elem(),
	}, 
	}
}
