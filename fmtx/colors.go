package fmtx

import (
	"fmt"
	"strings"
)

type ColorMode uint8

const (
	ColorModeAnsi ColorMode = iota
	ColorMode256
	ColorModeTrue
)

type StyleColor struct {
	mode ColorMode
	r    uint8 // Used for all modes
	g    uint8 // Not used for ansi or 256 modes
	b    uint8 // Not used for ansi or 256 modes
}

func NewColor(r, g, b uint8) StyleColor {
	return StyleColor{mode: ColorModeTrue, r: r, g: g, b: b}
}

func NewColorAnsi(code uint8) StyleColor {
	return StyleColor{mode: ColorModeAnsi, r: code}
}

func NewColor256(code uint8) StyleColor {
	return StyleColor{mode: ColorMode256, r: code}
}

func (c *StyleColor) ToAnsi() StyleColor {
	if c.mode == ColorModeAnsi {
		return *c
	}
	if c.mode == ColorMode256 {
		return StyleColor{mode: ColorModeAnsi, r: c.r}
	}
	// Convert true color to 256 color
	r := int(c.r) * 5 / 255
	g := int(c.g) * 5 / 255
	b := int(c.b) * 5 / 255
	code := 16 + r*36 + g*6 + b
	return StyleColor{mode: ColorModeAnsi, r: uint8(code)}
}

func (c *StyleColor) Code() string {
	switch c.mode {
	case ColorModeAnsi:
		return fmt.Sprintf("%d", c.r)
	case ColorMode256:
		return fmt.Sprintf("8;5;%d", c.r)
	case ColorModeTrue:
		return fmt.Sprintf("38;2;%d;%d;%d", c.r, c.g, c.b)
	default:
		return "0"
	}
}

const (
	modBold            = "1"
	modDim             = "2"
	modItalic          = "3"
	modUnderline       = "4"
	modSlowBlink       = "5"
	modRapidBlink      = "6"
	modReverse         = "7"
	modHidden          = "8"
	modStrikeThrough   = "9"
	modFraktur         = "20"
	modDoubleUnderline = "21"
	modFramed          = "51"
	modEncircled       = "52"
	modOverline        = "53"
)

type Style struct {
	modifiers  []string
	foreground *StyleColor
	background *StyleColor
}

func NewStyle() *Style {
	return &Style{}
}

func (s *Style) WithBold() *Style {
	s.modifiers = append(s.modifiers, modBold)
	return s
}
func (s *Style) WithDim() *Style {
	s.modifiers = append(s.modifiers, modDim)
	return s
}
func (s *Style) WithItalic() *Style {
	s.modifiers = append(s.modifiers, modItalic)
	return s
}
func (s *Style) WithUnderline() *Style {
	s.modifiers = append(s.modifiers, modUnderline)
	return s
}
func (s *Style) WithSlowBlink() *Style {
	s.modifiers = append(s.modifiers, modSlowBlink)
	return s
}
func (s *Style) WithRapidBlink() *Style {
	s.modifiers = append(s.modifiers, modRapidBlink)
	return s
}
func (s *Style) WithReverse() *Style {
	s.modifiers = append(s.modifiers, modReverse)
	return s
}
func (s *Style) WithHidden() *Style {
	s.modifiers = append(s.modifiers, modHidden)
	return s
}
func (s *Style) WithStrikeThrough() *Style {
	s.modifiers = append(s.modifiers, modStrikeThrough)
	return s
}
func (s *Style) WithFraktur() *Style {
	s.modifiers = append(s.modifiers, modFraktur)
	return s
}
func (s *Style) WithDoubleUnderline() *Style {
	s.modifiers = append(s.modifiers, modDoubleUnderline)
	return s
}
func (s *Style) WithFramed() *Style {
	s.modifiers = append(s.modifiers, modFramed)
	return s
}
func (s *Style) WithEncircled() *Style {
	s.modifiers = append(s.modifiers, modEncircled)
	return s
}
func (s *Style) WithOverline() *Style {
	s.modifiers = append(s.modifiers, modOverline)
	return s
}
func (s *Style) WithForeground(color StyleColor) *Style {
	s.foreground = &color
	return s
}
func (s *Style) WithBackground(color StyleColor) *Style {
	s.background = &color
	return s
}

func (s *Style) Apply(text string) string {
	codes := strings.Join(s.modifiers, ";")
	if s.foreground != nil {
		codes += ";" + s.foreground.Code()
	}
	if s.background != nil {
		codes += ";" + s.background.Code()
	}
	return "\x1b[" + codes + "m" + text + "\x1b[0m"
}

var (
	styleBold            = NewStyle().WithBold()
	styleDim             = NewStyle().WithDim()
	styleItalic          = NewStyle().WithItalic()
	styleUnderline       = NewStyle().WithUnderline()
	styleSlowBlink       = NewStyle().WithSlowBlink()
	styleRapidBlink      = NewStyle().WithRapidBlink()
	styleReverse         = NewStyle().WithReverse()
	styleHidden          = NewStyle().WithHidden()
	styleStrikeThrough   = NewStyle().WithStrikeThrough()
	styleFraktur         = NewStyle().WithFraktur()
	styleDoubleUnderline = NewStyle().WithDoubleUnderline()
	styleFramed          = NewStyle().WithFramed()
	styleEncircled       = NewStyle().WithEncircled()
	styleOverline        = NewStyle().WithOverline()
	styleBlack           = NewStyle().WithForeground(NewColorAnsi(30))
	styleRed             = NewStyle().WithForeground(NewColorAnsi(31))
	styleGreen           = NewStyle().WithForeground(NewColorAnsi(32))
	styleYellow          = NewStyle().WithForeground(NewColorAnsi(33))
	styleBlue            = NewStyle().WithForeground(NewColorAnsi(34))
	styleMagenta         = NewStyle().WithForeground(NewColorAnsi(35))
	styleCyan            = NewStyle().WithForeground(NewColorAnsi(36))
	styleWhite           = NewStyle().WithForeground(NewColorAnsi(37))
	styleDefault         = NewStyle().WithForeground(NewColorAnsi(39))
	styleBgBlack         = NewStyle().WithBackground(NewColorAnsi(40))
	styleBgRed           = NewStyle().WithBackground(NewColorAnsi(41))
	styleBgGreen         = NewStyle().WithBackground(NewColorAnsi(42))
	styleBgYellow        = NewStyle().WithBackground(NewColorAnsi(43))
	styleBgBlue          = NewStyle().WithBackground(NewColorAnsi(44))
	styleBgMagenta       = NewStyle().WithBackground(NewColorAnsi(45))
	styleBgCyan          = NewStyle().WithBackground(NewColorAnsi(46))
	styleBgWhite         = NewStyle().WithBackground(NewColorAnsi(47))
	styleBgDefault       = NewStyle().WithBackground(NewColorAnsi(49))
	styleBrightBlack     = NewStyle().WithForeground(NewColorAnsi(90))
	styleBrightRed       = NewStyle().WithForeground(NewColorAnsi(91))
	styleBrightGreen     = NewStyle().WithForeground(NewColorAnsi(92))
	styleBrightYellow    = NewStyle().WithForeground(NewColorAnsi(93))
	styleBrightBlue      = NewStyle().WithForeground(NewColorAnsi(94))
	styleBrightMagenta   = NewStyle().WithForeground(NewColorAnsi(95))
	styleBrightCyan      = NewStyle().WithForeground(NewColorAnsi(96))
	styleBrightWhite     = NewStyle().WithForeground(NewColorAnsi(97))
	styleBgBrightBlack   = NewStyle().WithBackground(NewColorAnsi(100))
	styleBgBrightRed     = NewStyle().WithBackground(NewColorAnsi(101))
	styleBgBrightGreen   = NewStyle().WithBackground(NewColorAnsi(102))
	styleBgBrightYellow  = NewStyle().WithBackground(NewColorAnsi(103))
	styleBgBrightBlue    = NewStyle().WithBackground(NewColorAnsi(104))
	styleBgBrightMagenta = NewStyle().WithBackground(NewColorAnsi(105))
	styleBgBrightCyan    = NewStyle().WithBackground(NewColorAnsi(106))
	styleBgBrightWhite   = NewStyle().WithBackground(NewColorAnsi(107))
)

func Bold(s string) string            { return styleBold.Apply(s) }
func Dim(s string) string             { return styleDim.Apply(s) }
func Italic(s string) string          { return styleItalic.Apply(s) }
func Underline(s string) string       { return styleUnderline.Apply(s) }
func SlowBlink(s string) string       { return styleSlowBlink.Apply(s) }
func RapidBlink(s string) string      { return styleRapidBlink.Apply(s) }
func Reverse(s string) string         { return styleReverse.Apply(s) }
func Hidden(s string) string          { return styleHidden.Apply(s) }
func StrikeThrough(s string) string   { return styleStrikeThrough.Apply(s) }
func Fraktur(s string) string         { return styleFraktur.Apply(s) }
func DoubleUnderline(s string) string { return styleDoubleUnderline.Apply(s) }
func Framed(s string) string          { return styleFramed.Apply(s) }
func Encircled(s string) string       { return styleEncircled.Apply(s) }
func Overline(s string) string        { return styleOverline.Apply(s) }
func Black(s string) string           { return styleBlack.Apply(s) }
func Red(s string) string             { return styleRed.Apply(s) }
func Green(s string) string           { return styleGreen.Apply(s) }
func Yellow(s string) string          { return styleYellow.Apply(s) }
func Blue(s string) string            { return styleBlue.Apply(s) }
func Magenta(s string) string         { return styleMagenta.Apply(s) }
func Cyan(s string) string            { return styleCyan.Apply(s) }
func White(s string) string           { return styleWhite.Apply(s) }
func Default(s string) string         { return styleDefault.Apply(s) }
func BgBlack(s string) string         { return styleBgBlack.Apply(s) }
func BgRed(s string) string           { return styleBgRed.Apply(s) }
func BgGreen(s string) string         { return styleBgGreen.Apply(s) }
func BgYellow(s string) string        { return styleBgYellow.Apply(s) }
func BgBlue(s string) string          { return styleBgBlue.Apply(s) }
func BgMagenta(s string) string       { return styleBgMagenta.Apply(s) }
func BgCyan(s string) string          { return styleBgCyan.Apply(s) }
func BgWhite(s string) string         { return styleBgWhite.Apply(s) }
func BgDefault(s string) string       { return styleBgDefault.Apply(s) }
func BrightBlack(s string) string     { return styleBrightBlack.Apply(s) }
func BrightRed(s string) string       { return styleBrightRed.Apply(s) }
func BrightGreen(s string) string     { return styleBrightGreen.Apply(s) }
func BrightYellow(s string) string    { return styleBrightYellow.Apply(s) }
func BrightBlue(s string) string      { return styleBrightBlue.Apply(s) }
func BrightMagenta(s string) string   { return styleBrightMagenta.Apply(s) }
func BrightCyan(s string) string      { return styleBrightCyan.Apply(s) }
func BrightWhite(s string) string     { return styleBrightWhite.Apply(s) }
func BgBrightBlack(s string) string   { return styleBgBrightBlack.Apply(s) }
func BgBrightRed(s string) string     { return styleBgBrightRed.Apply(s) }
func BgBrightGreen(s string) string   { return styleBgBrightGreen.Apply(s) }
func BgBrightYellow(s string) string  { return styleBgBrightYellow.Apply(s) }
func BgBrightBlue(s string) string    { return styleBgBrightBlue.Apply(s) }
func BgBrightMagenta(s string) string { return styleBgBrightMagenta.Apply(s) }
func BgBrightCyan(s string) string    { return styleBgBrightCyan.Apply(s) }
func BgBrightWhite(s string) string   { return styleBgBrightWhite.Apply(s) }
