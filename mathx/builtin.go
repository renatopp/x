package mathx

import (
	"cmp"
	"math"
)

const (
	E   = math.E
	Pi  = math.Pi
	Phi = math.Phi

	Sqrt2   = math.Sqrt2
	SqrtE   = math.SqrtE
	SqrtPi  = math.SqrtPi
	SqrtPhi = math.SqrtPhi

	Ln2    = math.Ln2
	Log2E  = math.Log2E
	Ln10   = math.Ln10
	Log10E = math.Log10E

	MaxFloat32             = math.MaxFloat32
	SmallestNonzeroFloat32 = math.SmallestNonzeroFloat32
	MaxFloat64             = math.MaxFloat64
	SmallestNonzeroFloat64 = math.SmallestNonzeroFloat64

	MaxInt    = math.MaxInt
	MinInt    = math.MinInt
	MaxInt8   = math.MaxInt8
	MinInt8   = math.MinInt8
	MaxInt16  = math.MaxInt16
	MinInt16  = math.MinInt16
	MaxInt32  = math.MaxInt32
	MinInt32  = math.MinInt32
	MaxInt64  = math.MaxInt64
	MinInt64  = math.MinInt64
	MaxUint   = math.MaxUint
	MaxUint8  = math.MaxUint8
	MaxUint16 = math.MaxUint16
	MaxUint32 = math.MaxUint32
	MaxUint64 = math.MaxUint64
)

func Abs[T Number](x T) T {
	if cmp.Less(x, 0) {
		return -x
	}
	return x
}

func Acos[T Number](x T) float64          { return math.Acos(float64(x)) }
func Acosh[T Number](x T) float64         { return math.Acosh(float64(x)) }
func Asin[T Number](x T) float64          { return math.Asin(float64(x)) }
func Asinh[T Number](x T) float64         { return math.Asinh(float64(x)) }
func Atan[T Number](x T) float64          { return math.Atan(float64(x)) }
func Atan2[Y, X Number](y Y, x X) float64 { return math.Atan2(float64(y), float64(x)) }
func Atanh[T Number](x T) float64         { return math.Atanh(float64(x)) }
func Cbrt[T Number](x T) float64          { return math.Cbrt(float64(x)) }
func Ceil[T Float](x T) float64           { return math.Ceil(float64(x)) }
func CeilInt[T Float](x T) int            { return int(math.Ceil(float64(x))) }
func CeilTo[T Float, I Number](x T) I     { return I(math.Ceil(float64(x))) }
func CopySign[T Number](x, sign T) T {
	if (x < 0) != (sign < 0) {
		return -x
	}
	return x
}
func Cos[T Number](x T) float64  { return math.Cos(float64(x)) }
func Cosh[T Number](x T) float64 { return math.Cosh(float64(x)) }
func Dim[T Number](x, y T) T {
	if cmp.Less(x, y) {
		return 0
	}
	return x - y
}
func Erf[T Number](x T) float64     { return math.Erf(float64(x)) }
func Erfc[T Number](x T) float64    { return math.Erfc(float64(x)) }
func Erfcinv[T Number](x T) float64 { return math.Erfcinv(float64(x)) }
func Erfinv[T Number](x T) float64  { return math.Erfinv(float64(x)) }
func Exp[T Number](x T) float64     { return math.Exp(float64(x)) }
func Exp2[T Number](x T) float64    { return math.Exp2(float64(x)) }
func Expm1[T Number](x T) float64   { return math.Expm1(float64(x)) }
func FMA[X, Y, Z Number](x X, y Y, z Z) float64 {
	return math.FMA(float64(x), float64(y), float64(z))
}
func Float32bits(f float32) uint32        { return math.Float32bits(f) }
func Float32frombits(b uint32) float32    { return math.Float32frombits(b) }
func Float64bits(f float64) uint64        { return math.Float64bits(f) }
func Float64frombits(b uint64) float64    { return math.Float64frombits(b) }
func Floor[T Float](x T) float64          { return math.Floor(float64(x)) }
func FloorInt[T Float](x T) int           { return int(math.Floor(float64(x))) }
func FloorTo[T Float, I Number](x T) I    { return I(math.Floor(float64(x))) }
func Frexp[T Number](f T) (float64, int)  { return math.Frexp(float64(f)) }
func Gamma[T Number](x T) float64         { return math.Gamma(float64(x)) }
func Hypot[X, Y Number](x X, y Y) float64 { return math.Hypot(float64(x), float64(y)) }
func Ilogb[T Number](x T) int             { return math.Ilogb(float64(x)) }
func Inf[T Number](sign T) float64 {
	if sign < 0 {
		return math.Inf(-1)
	}
	return math.Inf(1)
}
func IsInf[T Number](f T, sign T) bool {
	if sign < 0 {
		return math.IsInf(float64(f), -1)
	}
	return math.IsInf(float64(f), 1)
}
func IsNaN[T Number](f T) bool                { return math.IsNaN(float64(f)) }
func J0[T Number](x T) float64                { return math.J0(float64(x)) }
func J1[T Number](x T) float64                { return math.J1(float64(x)) }
func Jn[T Number](n int, x T) float64         { return math.Jn(n, float64(x)) }
func Ldexp[T Number](frac T, exp int) float64 { return math.Ldexp(float64(frac), exp) }
func Lgamma[T Number](x T) (float64, int)     { return math.Lgamma(float64(x)) }
func Log[T Number](x T) float64               { return math.Log(float64(x)) }
func Log1p[T Number](x T) float64             { return math.Log1p(float64(x)) }
func Log2[T Number](x T) float64              { return math.Log2(float64(x)) }
func Log10[T Number](x T) float64             { return math.Log10(float64(x)) }
func Logb[T Number](x T) float64              { return math.Logb(float64(x)) }
func Max[T Number](x, y T) T {
	if cmp.Less(x, y) {
		return y
	}
	return x
}
func Min[T Number](x, y T) T {
	if cmp.Less(x, y) {
		return x
	}
	return y
}
func Mod[X, Y Number](x X, y Y) float64       { return math.Mod(float64(x), float64(y)) }
func Modf[T Number](f T) (float64, float64)   { return math.Modf(float64(f)) }
func NaN() float64                            { return math.NaN() }
func Nextafter[T Number](x, y T) float64      { return math.Nextafter(float64(x), float64(y)) }
func Nextafter32(x, y float32) float32        { return math.Nextafter32(x, y) }
func Pow[X, Y Number](x X, y Y) float64       { return math.Pow(float64(x), float64(y)) }
func Pow10[T Number](x T) float64             { return math.Pow10(int(x)) }
func Remainder[X, Y Number](x X, y Y) float64 { return math.Remainder(float64(x), float64(y)) }
func Round[T Float](x T) float64              { return math.Round(float64(x)) }
func RoundInt[T Float](x T) int               { return int(math.Round(float64(x))) }
func RoundTo[T Float, I Number](x T) I        { return I(math.Round(float64(x))) }
func RoundToEven[T Float](x T) float64        { return math.RoundToEven(float64(x)) }
func RoundToEvenInt[T Float](x T) int         { return int(math.RoundToEven(float64(x))) }
func RoundToEvenTo[T Float, I Number](x T) I  { return I(math.RoundToEven(float64(x))) }
func Signbit[T Number](x T) bool              { return math.Signbit(float64(x)) }
func Sin[T Number](x T) float64               { return math.Sin(float64(x)) }
func Sincos[T Number](x T) (sin, cos float64) { return math.Sincos(float64(x)) }
func Sinh[T Number](x T) float64              { return math.Sinh(float64(x)) }
func Sqrt[T Number](x T) float64              { return math.Sqrt(float64(x)) }
func Tan[T Number](x T) float64               { return math.Tan(float64(x)) }
func Tanh[T Number](x T) float64              { return math.Tanh(float64(x)) }
func Trunc[T Float](x T) float64              { return math.Trunc(float64(x)) }
func TruncInt[T Float](x T) int               { return int(math.Trunc(float64(x))) }
func TruncTo[T Float, I Number](x T) I        { return I(math.Trunc(float64(x))) }
func Y0[T Number](x T) float64                { return math.Y0(float64(x)) }
func Y1[T Number](x T) float64                { return math.Y1(float64(x)) }
func Yn[T Number](n int, x T) float64         { return math.Yn(n, float64(x)) }
