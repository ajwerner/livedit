// this file was generated by gomacro command: import _i "gioui.org/gpu/backend"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	image "image"
	time "time"
	backend "gioui.org/gpu/backend"
)

// reflection: allow interpreted code to import "gioui.org/gpu/backend"
func init() {
	imports.Packages["gioui.org/gpu/backend"] = imports.Package{
	Name: "backend",
	Binds: map[string]r.Value{
		"AccessRead":	r.ValueOf(backend.AccessRead),
		"AccessWrite":	r.ValueOf(backend.AccessWrite),
		"BlendFactorDstColor":	r.ValueOf(backend.BlendFactorDstColor),
		"BlendFactorOne":	r.ValueOf(backend.BlendFactorOne),
		"BlendFactorOneMinusSrcAlpha":	r.ValueOf(backend.BlendFactorOneMinusSrcAlpha),
		"BlendFactorZero":	r.ValueOf(backend.BlendFactorZero),
		"BufferBindingFramebuffer":	r.ValueOf(backend.BufferBindingFramebuffer),
		"BufferBindingIndices":	r.ValueOf(backend.BufferBindingIndices),
		"BufferBindingShaderStorage":	r.ValueOf(backend.BufferBindingShaderStorage),
		"BufferBindingTexture":	r.ValueOf(backend.BufferBindingTexture),
		"BufferBindingUniforms":	r.ValueOf(backend.BufferBindingUniforms),
		"BufferBindingVertices":	r.ValueOf(backend.BufferBindingVertices),
		"DataTypeFloat":	r.ValueOf(backend.DataTypeFloat),
		"DataTypeInt":	r.ValueOf(backend.DataTypeInt),
		"DataTypeShort":	r.ValueOf(backend.DataTypeShort),
		"DepthFuncGreater":	r.ValueOf(backend.DepthFuncGreater),
		"DepthFuncGreaterEqual":	r.ValueOf(backend.DepthFuncGreaterEqual),
		"DrawModeTriangleStrip":	r.ValueOf(backend.DrawModeTriangleStrip),
		"DrawModeTriangles":	r.ValueOf(backend.DrawModeTriangles),
		"FeatureCompute":	r.ValueOf(backend.FeatureCompute),
		"FeatureFloatRenderTargets":	r.ValueOf(backend.FeatureFloatRenderTargets),
		"FeatureTimers":	r.ValueOf(backend.FeatureTimers),
		"FilterLinear":	r.ValueOf(backend.FilterLinear),
		"FilterNearest":	r.ValueOf(backend.FilterNearest),
		"TextureFormatFloat":	r.ValueOf(backend.TextureFormatFloat),
		"TextureFormatRGBA8":	r.ValueOf(backend.TextureFormatRGBA8),
		"TextureFormatSRGB":	r.ValueOf(backend.TextureFormatSRGB),
		"UploadImage":	r.ValueOf(backend.UploadImage),
	}, Types: map[string]r.Type{
		"AccessBits":	r.TypeOf((*backend.AccessBits)(nil)).Elem(),
		"BlendFactor":	r.TypeOf((*backend.BlendFactor)(nil)).Elem(),
		"Buffer":	r.TypeOf((*backend.Buffer)(nil)).Elem(),
		"BufferBinding":	r.TypeOf((*backend.BufferBinding)(nil)).Elem(),
		"Caps":	r.TypeOf((*backend.Caps)(nil)).Elem(),
		"DataType":	r.TypeOf((*backend.DataType)(nil)).Elem(),
		"DepthFunc":	r.TypeOf((*backend.DepthFunc)(nil)).Elem(),
		"Device":	r.TypeOf((*backend.Device)(nil)).Elem(),
		"DrawMode":	r.TypeOf((*backend.DrawMode)(nil)).Elem(),
		"Features":	r.TypeOf((*backend.Features)(nil)).Elem(),
		"Framebuffer":	r.TypeOf((*backend.Framebuffer)(nil)).Elem(),
		"InputDesc":	r.TypeOf((*backend.InputDesc)(nil)).Elem(),
		"InputLayout":	r.TypeOf((*backend.InputLayout)(nil)).Elem(),
		"InputLocation":	r.TypeOf((*backend.InputLocation)(nil)).Elem(),
		"Program":	r.TypeOf((*backend.Program)(nil)).Elem(),
		"ShaderSources":	r.TypeOf((*backend.ShaderSources)(nil)).Elem(),
		"StorageBufferBinding":	r.TypeOf((*backend.StorageBufferBinding)(nil)).Elem(),
		"Texture":	r.TypeOf((*backend.Texture)(nil)).Elem(),
		"TextureBinding":	r.TypeOf((*backend.TextureBinding)(nil)).Elem(),
		"TextureFilter":	r.TypeOf((*backend.TextureFilter)(nil)).Elem(),
		"TextureFormat":	r.TypeOf((*backend.TextureFormat)(nil)).Elem(),
		"Timer":	r.TypeOf((*backend.Timer)(nil)).Elem(),
		"UniformBlock":	r.TypeOf((*backend.UniformBlock)(nil)).Elem(),
		"UniformLocation":	r.TypeOf((*backend.UniformLocation)(nil)).Elem(),
		"UniformsReflection":	r.TypeOf((*backend.UniformsReflection)(nil)).Elem(),
	}, Proxies: map[string]r.Type{
		"Buffer":	r.TypeOf((*P_gioui_org_gpu_backend_Buffer)(nil)).Elem(),
		"Device":	r.TypeOf((*P_gioui_org_gpu_backend_Device)(nil)).Elem(),
		"Framebuffer":	r.TypeOf((*P_gioui_org_gpu_backend_Framebuffer)(nil)).Elem(),
		"InputLayout":	r.TypeOf((*P_gioui_org_gpu_backend_InputLayout)(nil)).Elem(),
		"Program":	r.TypeOf((*P_gioui_org_gpu_backend_Program)(nil)).Elem(),
		"Texture":	r.TypeOf((*P_gioui_org_gpu_backend_Texture)(nil)).Elem(),
		"Timer":	r.TypeOf((*P_gioui_org_gpu_backend_Timer)(nil)).Elem(),
	}, 
	}
}

// --------------- proxy for gioui.org/gpu/backend.Buffer ---------------
type P_gioui_org_gpu_backend_Buffer struct {
	Object	interface{}
	Download_	func(_proxy_obj_ interface{}, data []byte) error
	Release_	func(interface{}) 
	Upload_	func(_proxy_obj_ interface{}, data []byte) 
}
func (P *P_gioui_org_gpu_backend_Buffer) Download(data []byte) error {
	return P.Download_(P.Object, data)
}
func (P *P_gioui_org_gpu_backend_Buffer) Release()  {
	P.Release_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Buffer) Upload(data []byte)  {
	P.Upload_(P.Object, data)
}

// --------------- proxy for gioui.org/gpu/backend.Device ---------------
type P_gioui_org_gpu_backend_Device struct {
	Object	interface{}
	BeginFrame_	func(interface{}) 
	BindFramebuffer_	func(_proxy_obj_ interface{}, f backend.Framebuffer) 
	BindImageTexture_	func(_proxy_obj_ interface{}, unit int, texture backend.Texture, access backend.AccessBits, format backend.TextureFormat) 
	BindIndexBuffer_	func(_proxy_obj_ interface{}, b backend.Buffer) 
	BindInputLayout_	func(_proxy_obj_ interface{}, i backend.InputLayout) 
	BindProgram_	func(_proxy_obj_ interface{}, p backend.Program) 
	BindTexture_	func(_proxy_obj_ interface{}, unit int, t backend.Texture) 
	BindVertexBuffer_	func(_proxy_obj_ interface{}, b backend.Buffer, stride int, offset int) 
	BlendFunc_	func(_proxy_obj_ interface{}, sfactor backend.BlendFactor, dfactor backend.BlendFactor) 
	Caps_	func(interface{}) backend.Caps
	Clear_	func(_proxy_obj_ interface{}, r float32, g float32, b float32, a float32) 
	ClearDepth_	func(_proxy_obj_ interface{}, d float32) 
	CurrentFramebuffer_	func(interface{}) backend.Framebuffer
	DepthFunc_	func(_proxy_obj_ interface{}, f backend.DepthFunc) 
	DepthMask_	func(_proxy_obj_ interface{}, mask bool) 
	DispatchCompute_	func(_proxy_obj_ interface{}, x int, y int, z int) 
	DrawArrays_	func(_proxy_obj_ interface{}, mode backend.DrawMode, off int, count int) 
	DrawElements_	func(_proxy_obj_ interface{}, mode backend.DrawMode, off int, count int) 
	EndFrame_	func(interface{}) 
	IsTimeContinuous_	func(interface{}) bool
	MemoryBarrier_	func(interface{}) 
	NewBuffer_	func(_proxy_obj_ interface{}, typ backend.BufferBinding, size int) (backend.Buffer, error)
	NewComputeProgram_	func(_proxy_obj_ interface{}, shader backend.ShaderSources) (backend.Program, error)
	NewFramebuffer_	func(_proxy_obj_ interface{}, tex backend.Texture, depthBits int) (backend.Framebuffer, error)
	NewImmutableBuffer_	func(_proxy_obj_ interface{}, typ backend.BufferBinding, data []byte) (backend.Buffer, error)
	NewInputLayout_	func(_proxy_obj_ interface{}, vertexShader backend.ShaderSources, layout []backend.InputDesc) (backend.InputLayout, error)
	NewProgram_	func(_proxy_obj_ interface{}, vertexShader backend.ShaderSources, fragmentShader backend.ShaderSources) (backend.Program, error)
	NewTexture_	func(_proxy_obj_ interface{}, format backend.TextureFormat, width int, height int, minFilter backend.TextureFilter, magFilter backend.TextureFilter, bindings backend.BufferBinding) (backend.Texture, error)
	NewTimer_	func(interface{}) backend.Timer
	SetBlend_	func(_proxy_obj_ interface{}, enable bool) 
	SetDepthTest_	func(_proxy_obj_ interface{}, enable bool) 
	Viewport_	func(_proxy_obj_ interface{}, x int, y int, width int, height int) 
}
func (P *P_gioui_org_gpu_backend_Device) BeginFrame()  {
	P.BeginFrame_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Device) BindFramebuffer(f backend.Framebuffer)  {
	P.BindFramebuffer_(P.Object, f)
}
func (P *P_gioui_org_gpu_backend_Device) BindImageTexture(unit int, texture backend.Texture, access backend.AccessBits, format backend.TextureFormat)  {
	P.BindImageTexture_(P.Object, unit, texture, access, format)
}
func (P *P_gioui_org_gpu_backend_Device) BindIndexBuffer(b backend.Buffer)  {
	P.BindIndexBuffer_(P.Object, b)
}
func (P *P_gioui_org_gpu_backend_Device) BindInputLayout(i backend.InputLayout)  {
	P.BindInputLayout_(P.Object, i)
}
func (P *P_gioui_org_gpu_backend_Device) BindProgram(p backend.Program)  {
	P.BindProgram_(P.Object, p)
}
func (P *P_gioui_org_gpu_backend_Device) BindTexture(unit int, t backend.Texture)  {
	P.BindTexture_(P.Object, unit, t)
}
func (P *P_gioui_org_gpu_backend_Device) BindVertexBuffer(b backend.Buffer, stride int, offset int)  {
	P.BindVertexBuffer_(P.Object, b, stride, offset)
}
func (P *P_gioui_org_gpu_backend_Device) BlendFunc(sfactor backend.BlendFactor, dfactor backend.BlendFactor)  {
	P.BlendFunc_(P.Object, sfactor, dfactor)
}
func (P *P_gioui_org_gpu_backend_Device) Caps() backend.Caps {
	return P.Caps_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Device) Clear(r float32, g float32, b float32, a float32)  {
	P.Clear_(P.Object, r, g, b, a)
}
func (P *P_gioui_org_gpu_backend_Device) ClearDepth(d float32)  {
	P.ClearDepth_(P.Object, d)
}
func (P *P_gioui_org_gpu_backend_Device) CurrentFramebuffer() backend.Framebuffer {
	return P.CurrentFramebuffer_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Device) DepthFunc(f backend.DepthFunc)  {
	P.DepthFunc_(P.Object, f)
}
func (P *P_gioui_org_gpu_backend_Device) DepthMask(mask bool)  {
	P.DepthMask_(P.Object, mask)
}
func (P *P_gioui_org_gpu_backend_Device) DispatchCompute(x int, y int, z int)  {
	P.DispatchCompute_(P.Object, x, y, z)
}
func (P *P_gioui_org_gpu_backend_Device) DrawArrays(mode backend.DrawMode, off int, count int)  {
	P.DrawArrays_(P.Object, mode, off, count)
}
func (P *P_gioui_org_gpu_backend_Device) DrawElements(mode backend.DrawMode, off int, count int)  {
	P.DrawElements_(P.Object, mode, off, count)
}
func (P *P_gioui_org_gpu_backend_Device) EndFrame()  {
	P.EndFrame_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Device) IsTimeContinuous() bool {
	return P.IsTimeContinuous_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Device) MemoryBarrier()  {
	P.MemoryBarrier_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Device) NewBuffer(typ backend.BufferBinding, size int) (backend.Buffer, error) {
	return P.NewBuffer_(P.Object, typ, size)
}
func (P *P_gioui_org_gpu_backend_Device) NewComputeProgram(shader backend.ShaderSources) (backend.Program, error) {
	return P.NewComputeProgram_(P.Object, shader)
}
func (P *P_gioui_org_gpu_backend_Device) NewFramebuffer(tex backend.Texture, depthBits int) (backend.Framebuffer, error) {
	return P.NewFramebuffer_(P.Object, tex, depthBits)
}
func (P *P_gioui_org_gpu_backend_Device) NewImmutableBuffer(typ backend.BufferBinding, data []byte) (backend.Buffer, error) {
	return P.NewImmutableBuffer_(P.Object, typ, data)
}
func (P *P_gioui_org_gpu_backend_Device) NewInputLayout(vertexShader backend.ShaderSources, layout []backend.InputDesc) (backend.InputLayout, error) {
	return P.NewInputLayout_(P.Object, vertexShader, layout)
}
func (P *P_gioui_org_gpu_backend_Device) NewProgram(vertexShader backend.ShaderSources, fragmentShader backend.ShaderSources) (backend.Program, error) {
	return P.NewProgram_(P.Object, vertexShader, fragmentShader)
}
func (P *P_gioui_org_gpu_backend_Device) NewTexture(format backend.TextureFormat, width int, height int, minFilter backend.TextureFilter, magFilter backend.TextureFilter, bindings backend.BufferBinding) (backend.Texture, error) {
	return P.NewTexture_(P.Object, format, width, height, minFilter, magFilter, bindings)
}
func (P *P_gioui_org_gpu_backend_Device) NewTimer() backend.Timer {
	return P.NewTimer_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Device) SetBlend(enable bool)  {
	P.SetBlend_(P.Object, enable)
}
func (P *P_gioui_org_gpu_backend_Device) SetDepthTest(enable bool)  {
	P.SetDepthTest_(P.Object, enable)
}
func (P *P_gioui_org_gpu_backend_Device) Viewport(x int, y int, width int, height int)  {
	P.Viewport_(P.Object, x, y, width, height)
}

// --------------- proxy for gioui.org/gpu/backend.Framebuffer ---------------
type P_gioui_org_gpu_backend_Framebuffer struct {
	Object	interface{}
	Invalidate_	func(interface{}) 
	ReadPixels_	func(_proxy_obj_ interface{}, src image.Rectangle, pixels []byte) error
	Release_	func(interface{}) 
}
func (P *P_gioui_org_gpu_backend_Framebuffer) Invalidate()  {
	P.Invalidate_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Framebuffer) ReadPixels(src image.Rectangle, pixels []byte) error {
	return P.ReadPixels_(P.Object, src, pixels)
}
func (P *P_gioui_org_gpu_backend_Framebuffer) Release()  {
	P.Release_(P.Object)
}

// --------------- proxy for gioui.org/gpu/backend.InputLayout ---------------
type P_gioui_org_gpu_backend_InputLayout struct {
	Object	interface{}
	Release_	func(interface{}) 
}
func (P *P_gioui_org_gpu_backend_InputLayout) Release()  {
	P.Release_(P.Object)
}

// --------------- proxy for gioui.org/gpu/backend.Program ---------------
type P_gioui_org_gpu_backend_Program struct {
	Object	interface{}
	Release_	func(interface{}) 
	SetFragmentUniforms_	func(_proxy_obj_ interface{}, buf backend.Buffer) 
	SetStorageBuffer_	func(_proxy_obj_ interface{}, binding int, buf backend.Buffer) 
	SetVertexUniforms_	func(_proxy_obj_ interface{}, buf backend.Buffer) 
}
func (P *P_gioui_org_gpu_backend_Program) Release()  {
	P.Release_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Program) SetFragmentUniforms(buf backend.Buffer)  {
	P.SetFragmentUniforms_(P.Object, buf)
}
func (P *P_gioui_org_gpu_backend_Program) SetStorageBuffer(binding int, buf backend.Buffer)  {
	P.SetStorageBuffer_(P.Object, binding, buf)
}
func (P *P_gioui_org_gpu_backend_Program) SetVertexUniforms(buf backend.Buffer)  {
	P.SetVertexUniforms_(P.Object, buf)
}

// --------------- proxy for gioui.org/gpu/backend.Texture ---------------
type P_gioui_org_gpu_backend_Texture struct {
	Object	interface{}
	Release_	func(interface{}) 
	Upload_	func(_proxy_obj_ interface{}, offset image.Point, size image.Point, pixels []byte) 
}
func (P *P_gioui_org_gpu_backend_Texture) Release()  {
	P.Release_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Texture) Upload(offset image.Point, size image.Point, pixels []byte)  {
	P.Upload_(P.Object, offset, size, pixels)
}

// --------------- proxy for gioui.org/gpu/backend.Timer ---------------
type P_gioui_org_gpu_backend_Timer struct {
	Object	interface{}
	Begin_	func(interface{}) 
	Duration_	func(interface{}) (time.Duration, bool)
	End_	func(interface{}) 
	Release_	func(interface{}) 
}
func (P *P_gioui_org_gpu_backend_Timer) Begin()  {
	P.Begin_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Timer) Duration() (time.Duration, bool) {
	return P.Duration_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Timer) End()  {
	P.End_(P.Object)
}
func (P *P_gioui_org_gpu_backend_Timer) Release()  {
	P.Release_(P.Object)
}