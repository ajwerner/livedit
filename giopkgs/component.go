// this file was generated by gomacro command: import _i "gioui.org/x/component"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	component "gioui.org/x/component"
)

// reflection: allow interpreted code to import "gioui.org/x/component"
func init() {
	imports.Packages["gioui.org/x/component"] = imports.Package{
	Name: "component",
	Binds: map[string]r.Value{
		"Appearing":	r.ValueOf(component.Appearing),
		"Bottom":	r.ValueOf(component.Bottom),
		"Disappearing":	r.ValueOf(component.Disappearing),
		"Forward":	r.ValueOf(component.Forward),
		"Interpolate":	r.ValueOf(component.Interpolate),
		"Invisible":	r.ValueOf(component.Invisible),
		"ModalNavFrom":	r.ValueOf(component.ModalNavFrom),
		"NewAppBar":	r.ValueOf(component.NewAppBar),
		"NewModal":	r.ValueOf(component.NewModal),
		"NewModalNav":	r.ValueOf(component.NewModalNav),
		"NewModalSheet":	r.ValueOf(component.NewModalSheet),
		"NewNav":	r.ValueOf(component.NewNav),
		"NewSheet":	r.ValueOf(component.NewSheet),
		"Reverse":	r.ValueOf(component.Reverse),
		"SimpleIconAction":	r.ValueOf(component.SimpleIconAction),
		"SimpleIconButton":	r.ValueOf(component.SimpleIconButton),
		"SwapGrounds":	r.ValueOf(component.SwapGrounds),
		"SwapPairs":	r.ValueOf(component.SwapPairs),
		"Top":	r.ValueOf(component.Top),
		"Visible":	r.ValueOf(component.Visible),
		"WithAlpha":	r.ValueOf(component.WithAlpha),
	}, Types: map[string]r.Type{
		"AlphaPalette":	r.TypeOf((*component.AlphaPalette)(nil)).Elem(),
		"AppBar":	r.TypeOf((*component.AppBar)(nil)).Elem(),
		"AppBarAction":	r.TypeOf((*component.AppBarAction)(nil)).Elem(),
		"AppBarContextMenuDismissed":	r.TypeOf((*component.AppBarContextMenuDismissed)(nil)).Elem(),
		"AppBarEvent":	r.TypeOf((*component.AppBarEvent)(nil)).Elem(),
		"AppBarNavigationClicked":	r.TypeOf((*component.AppBarNavigationClicked)(nil)).Elem(),
		"AppBarOverflowActionClicked":	r.TypeOf((*component.AppBarOverflowActionClicked)(nil)).Elem(),
		"C":	r.TypeOf((*component.C)(nil)).Elem(),
		"D":	r.TypeOf((*component.D)(nil)).Elem(),
		"Hoverable":	r.TypeOf((*component.Hoverable)(nil)).Elem(),
		"ModalLayer":	r.TypeOf((*component.ModalLayer)(nil)).Elem(),
		"ModalNavDrawer":	r.TypeOf((*component.ModalNavDrawer)(nil)).Elem(),
		"ModalSheet":	r.TypeOf((*component.ModalSheet)(nil)).Elem(),
		"NavDrawer":	r.TypeOf((*component.NavDrawer)(nil)).Elem(),
		"NavItem":	r.TypeOf((*component.NavItem)(nil)).Elem(),
		"OverflowAction":	r.TypeOf((*component.OverflowAction)(nil)).Elem(),
		"Progress":	r.TypeOf((*component.Progress)(nil)).Elem(),
		"ProgressDirection":	r.TypeOf((*component.ProgressDirection)(nil)).Elem(),
		"Rect":	r.TypeOf((*component.Rect)(nil)).Elem(),
		"Scrim":	r.TypeOf((*component.Scrim)(nil)).Elem(),
		"Sheet":	r.TypeOf((*component.Sheet)(nil)).Elem(),
		"TextField":	r.TypeOf((*component.TextField)(nil)).Elem(),
		"TruncatingLabelStyle":	r.TypeOf((*component.TruncatingLabelStyle)(nil)).Elem(),
		"Validator":	r.TypeOf((*component.Validator)(nil)).Elem(),
		"VerticalAnchorPosition":	r.TypeOf((*component.VerticalAnchorPosition)(nil)).Elem(),
		"VisibilityAnimation":	r.TypeOf((*component.VisibilityAnimation)(nil)).Elem(),
		"VisibilityAnimationState":	r.TypeOf((*component.VisibilityAnimationState)(nil)).Elem(),
	}, Proxies: map[string]r.Type{
		"AppBarEvent":	r.TypeOf((*P_gioui_org_x_component_AppBarEvent)(nil)).Elem(),
	}, Wrappers: map[string][]string{
		"AppBar":	[]string{"Animating","Appear","Clicked","Clicks","Disappear","History","Hours","Microseconds","Milliseconds","Minutes","Nanoseconds","Revealed","Round","Seconds","String","ToggleVisibility","Truncate","Visible",},
		"C":	[]string{"Data","Refs","Reset","Version","Write","Write1","Write2",},
		"Hoverable":	[]string{"Clicked","Clicks","History",},
		"ModalLayer":	[]string{"Animating","Appear","Clicked","Clicks","Disappear","History","Hours","Microseconds","Milliseconds","Minutes","Nanoseconds","Revealed","Round","Seconds","String","ToggleVisibility","Truncate","Visible",},
		"ModalNavDrawer":	[]string{"AddNavItem","CurrentNavDestination","LayoutContents","NavDestinationChanged","SetNavDestination",},
		"ModalSheet":	[]string{"Layout",},
		"Scrim":	[]string{"Clicked","Clicks","History",},
		"TextField":	[]string{"CaretCoords","CaretPos","Delete","Events","Focus","Focused","Insert","Len","Move","NumLines","PaintCaret","PaintText","SetText","Text",},
		"VisibilityAnimation":	[]string{"Hours","Microseconds","Milliseconds","Minutes","Nanoseconds","Round","Seconds","Truncate",},
	}, 
	}
}

// --------------- proxy for gioui.org/x/component.AppBarEvent ---------------
type P_gioui_org_x_component_AppBarEvent struct {
	Object	interface{}
	AppBarEvent_	func(interface{}) 
}
func (P *P_gioui_org_x_component_AppBarEvent) AppBarEvent()  {
	P.AppBarEvent_(P.Object)
}