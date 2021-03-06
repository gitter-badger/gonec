// Package math implements math interface for anko script.
package math

import (
	t "math"

	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("math")
	m.DefineS("Abs", t.Abs)
	m.DefineS("Acos", t.Acos)
	m.DefineS("Acosh", t.Acosh)
	m.DefineS("Asin", t.Asin)
	m.DefineS("Asinh", t.Asinh)
	m.DefineS("Atan", t.Atan)
	m.DefineS("Atan2", t.Atan2)
	m.DefineS("Atanh", t.Atanh)
	m.DefineS("Cbrt", t.Cbrt)
	m.DefineS("Ceil", t.Ceil)
	m.DefineS("Copysign", t.Copysign)
	m.DefineS("Cos", t.Cos)
	m.DefineS("Cosh", t.Cosh)
	m.DefineS("Dim", t.Dim)
	m.DefineS("Erf", t.Erf)
	m.DefineS("Erfc", t.Erfc)
	m.DefineS("Exp", t.Exp)
	m.DefineS("Exp2", t.Exp2)
	m.DefineS("Expm1", t.Expm1)
	m.DefineS("Float32bits", t.Float32bits)
	m.DefineS("Float32frombits", t.Float32frombits)
	m.DefineS("Float64bits", t.Float64bits)
	m.DefineS("Float64frombits", t.Float64frombits)
	m.DefineS("Floor", t.Floor)
	m.DefineS("Frexp", t.Frexp)
	m.DefineS("Gamma", t.Gamma)
	m.DefineS("Hypot", t.Hypot)
	m.DefineS("Ilogb", t.Ilogb)
	m.DefineS("Inf", t.Inf)
	m.DefineS("IsInf", t.IsInf)
	m.DefineS("IsNaN", t.IsNaN)
	m.DefineS("J0", t.J0)
	m.DefineS("J1", t.J1)
	m.DefineS("Jn", t.Jn)
	m.DefineS("Ldexp", t.Ldexp)
	m.DefineS("Lgamma", t.Lgamma)
	m.DefineS("Log", t.Log)
	m.DefineS("Log10", t.Log10)
	m.DefineS("Log1p", t.Log1p)
	m.DefineS("Log2", t.Log2)
	m.DefineS("Logb", t.Logb)
	m.DefineS("Max", t.Max)
	m.DefineS("Min", t.Min)
	m.DefineS("Mod", t.Mod)
	m.DefineS("Modf", t.Modf)
	m.DefineS("NaN", t.NaN)
	m.DefineS("Nextafter", t.Nextafter)
	m.DefineS("Pow", t.Pow)
	m.DefineS("Pow10", t.Pow10)
	m.DefineS("Remainder", t.Remainder)
	m.DefineS("Signbit", t.Signbit)
	m.DefineS("Sin", t.Sin)
	m.DefineS("Sincos", t.Sincos)
	m.DefineS("Sinh", t.Sinh)
	m.DefineS("Sqrt", t.Sqrt)
	m.DefineS("Tan", t.Tan)
	m.DefineS("Tanh", t.Tanh)
	m.DefineS("Trunc", t.Trunc)
	m.DefineS("Y0", t.Y0)
	m.DefineS("Y1", t.Y1)
	m.DefineS("Yn", t.Yn)
	return m
}
