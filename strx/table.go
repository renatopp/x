package strx

import "strings"

var TableWriterFlavorAscii = NewTableWriterFlavor(
	NewTableWriterRowFlavor("+", "+", "+", "+", "+", "+", "+", "+", "+", "-", "-", "|", "|"),
	NewTableWriterRowFlavor("+", "+", "+", "", "", "", "+", "+", "+", "-", "", "|", "|"),
)
var TableWriterFlavorAsciiSeparated = NewTableWriterFlavor(
	NewTableWriterRowFlavor("+", "+", "+", "+", "+", "+", "+", "+", "+", "=", "=", "|", "|"),
	NewTableWriterRowFlavor("+", "+", "+", "+", "+", "+", "+", "+", "+", "-", "-", "|", "|"),
)
var TableWriterFlavorAsciiCompact = NewTableWriterFlavor(
	NewTableWriterRowFlavor("", "", "", "", " ", "", "", "", "", "", "-", "", " "),
	NewTableWriterRowFlavor("", "", "", "", "", "", "", "", "", "", "", "", " "),
)
var TableWriterFlavorAsciiDots = NewTableWriterFlavor(
	NewTableWriterRowFlavor(".", ".", ".", ":", ":", ":", ":", ":", ":", ".", ".", ":", ":"),
	NewTableWriterRowFlavor(".", ".", ".", "", "", "", ":", ":", ":", ".", "", ":", ":"),
)
var TableWriterFlavorGithub = NewTableWriterFlavor(
	NewTableWriterRowFlavor("", "", "", "|", "|", "|", "", "", "", "", "-", "|", "|"),
	NewTableWriterRowFlavor("", "", "", "", "", "", "", "", "", "", "", "|", "|"),
)
var TableWriterFlavorReddit = NewTableWriterFlavor(
	NewTableWriterRowFlavor("", "", "", "", "|", "", "", "", "", "", "-", "", "|"),
	NewTableWriterRowFlavor("", "", "", "", "", "", "", "", "", "", "", "", "|"),
)
var TableWriterFlavorRestructuredGrid = NewTableWriterFlavor(
	NewTableWriterRowFlavor("+", "+", "+", "+", "+", "+", "+", "+", "+", "-", "=", "|", "|"),
	NewTableWriterRowFlavor("+", "+", "+", "", "", "", "+", "+", "+", "-", "", "|", "|"),
)
var TableWriterFlavorRestructuredSimple = NewTableWriterFlavor(
	NewTableWriterRowFlavor("", " ", "", "", " ", "", "", " ", "", "=", "=", "", " "),
	NewTableWriterRowFlavor("", " ", "", "", "", "", "", " ", "", "=", "", "", " "),
)
var TableWriterFlavorUnicode = NewTableWriterFlavor(
	NewTableWriterRowFlavor("┌", "┬", "┐", "├", "┼", "┤", "└", "┴", "┘", "─", "─", "│", "│"),
	NewTableWriterRowFlavor("┌", "┬", "┐", "", "", "", "└", "┴", "┘", "─", "", "│", "│"),
)
var TableWriterFlavorUnicodeDouble = NewTableWriterFlavor(
	NewTableWriterRowFlavor("╔", "╦", "╗", "╠", "╬", "╣", "╚", "╩", "╝", "═", "═", "║", "║"),
	NewTableWriterRowFlavor("╔", "╦", "╗", "", "", "", "╚", "╩", "╝", "═", "", "║", "║"),
)

//
// a=====b=====c
// $     |     $
// d-----e-----f
// $     |     $
// g=====h=====i
//
// a: topLeft
// b: topMid
// c: topRight
// d: midLeft
// e: midMid
// f: midRight
// g: bottomLeft
// h: bottomMid
// i: bottomRight
// $: vertical external
// |: vertical internal
// =: horizontal external
// -: horizontal internal
//
type TableWriterRowFlavor struct {
	TopLeft            string
	TopMid             string
	TopRight           string
	MidLeft            string
	MidMid             string
	MidRight           string
	BottomLeft         string
	BottomMid          string
	BottomRight        string
	VerticalExternal   string
	VerticalInternal   string
	HorizontalExternal string
	HorizontalInternal string
}

func NewTableWriterRowFlavor(tl, tm, tr, ml, mm, mr, bl, bm, br, he, hi, ve, vi string) TableWriterRowFlavor {
	return TableWriterRowFlavor{
		TopLeft:            tl,
		TopMid:             tm,
		TopRight:           tr,
		MidLeft:            ml,
		MidMid:             mm,
		MidRight:           mr,
		BottomLeft:         bl,
		BottomMid:          bm,
		BottomRight:        br,
		VerticalExternal:   ve,
		VerticalInternal:   vi,
		HorizontalExternal: he,
		HorizontalInternal: hi,
	}
}

// Meta is special rows such as headers, footers, etc.
// Data is the regular rows.
type TableWriterFlavor struct {
	Meta TableWriterRowFlavor
	Data TableWriterRowFlavor
}

func NewTableWriterFlavor(meta, data TableWriterRowFlavor) TableWriterFlavor {
	return TableWriterFlavor{
		Meta: meta,
		Data: data,
	}
}

type TableWriterAlignment int8

const (
	TableWriterAlignmentLeft TableWriterAlignment = iota
	TableWriterAlignmentCenter
	TableWriterAlignmentRight
)

type TableWriterRowConfig struct {
	Alignments []TableWriterAlignment
}

type TableWriterRow struct {
	Config  *TableWriterRowConfig
	Values  []string
	Meta    bool
	Section bool
}

func (r *TableWriterRow) WithAlignments(alignments ...TableWriterAlignment) *TableWriterRow {
	if len(alignments) != len(r.Values) {
		panic("invalid number of alignments")
	}
	if r.Config == nil {
		r.Config = &TableWriterRowConfig{}
	}
	r.Config.Alignments = alignments
	return r
}

type TableWriter struct {
	cols        int
	rows        []*TableWriterRow
	lengths     []int
	totalLength int
	metaConfig  *TableWriterRowConfig
	dataConfig  *TableWriterRowConfig
	flavor      TableWriterFlavor
}

func NewTableWriter(cols int) *TableWriter {
	metaAlignments := make([]TableWriterAlignment, cols)
	dataAlignments := make([]TableWriterAlignment, cols)
	for i := range metaAlignments {
		metaAlignments[i] = TableWriterAlignmentCenter
		dataAlignments[i] = TableWriterAlignmentLeft
	}

	return &TableWriter{
		cols:        cols,
		rows:        make([]*TableWriterRow, 0),
		lengths:     make([]int, cols),
		totalLength: 0,
		metaConfig: &TableWriterRowConfig{
			Alignments: metaAlignments,
		},
		dataConfig: &TableWriterRowConfig{
			Alignments: dataAlignments,
		},
		flavor: TableWriterFlavorAscii,
	}
}

func (t *TableWriter) WithFlavor(flavor TableWriterFlavor) *TableWriter {
	t.flavor = flavor
	return t
}
func (t *TableWriter) Meta(values ...string) *TableWriterRow {
	return t.addRow(true, false, values...)
}
func (t *TableWriter) Data(values ...string) *TableWriterRow {
	return t.addRow(false, false, values...)
}
func (t *TableWriter) MetaSection(value string) *TableWriterRow {
	return t.addRow(true, true, value)
}
func (t *TableWriter) DataSection(value string) *TableWriterRow {
	return t.addRow(false, true, value)
}
func (t *TableWriter) Render() string {
	b := &strings.Builder{}

	t.renderSeparator(b, 0)
	for i := range t.rows {
		t.renderRow(b, i)
		t.renderSeparator(b, i+1)
	}
	return b.String()[:len(b.String())-1] // remove last \n
}

func (t *TableWriter) addRow(meta, section bool, values ...string) *TableWriterRow {
	if !section && len(values) != t.cols {
		panic("invalid number of columns")
	}
	row := &TableWriterRow{
		Values:  values,
		Meta:    meta,
		Section: section,
		Config:  nil,
	}
	t.rows = append(t.rows, row)

	total := 0
	if section {
		total = Length(values[0])
	} else {
		for i, v := range values {
			size := Length(v)
			total += size
			if size > t.lengths[i] {
				t.lengths[i] = size
			}
		}

	}
	if total > t.totalLength {
		t.totalLength = total
	}
	return row
}

func (t *TableWriter) renderSeparator(b *strings.Builder, i int) {
	_, _, section := t.getBorderSeparator(i)
	left, mid, right, h := t.getBorderParts(i)
	if left == "" && mid == "" && right == "" && h == "" {
		return
	}

	if section {
		b.WriteString(Format("%s%s%s", left, strings.Repeat(h, t.totalLength+t.cols*Length(mid)+2), right))
	} else {
		b.WriteString(left)
		for i, length := range t.lengths {
			b.WriteString(strings.Repeat(h, length+2))
			if i < len(t.lengths)-1 {
				b.WriteString(mid)
			}
		}
		b.WriteString(right)
	}
	b.WriteString("\n")
}

func (t *TableWriter) renderRow(b *strings.Builder, i int) {
	flavor := t.flavor.Data
	row := t.rows[i]
	if row.Meta {
		flavor = t.flavor.Meta
	}
	ve := flavor.VerticalExternal
	vi := flavor.VerticalInternal

	alignments := t.dataConfig.Alignments
	if row.Meta {
		alignments = t.metaConfig.Alignments
	}
	if row.Config != nil {
		alignments = row.Config.Alignments
	}

	b.WriteString(ve)
	if row.Section {
		t.renderCell(b, row.Values[0], t.totalLength+t.cols*Length(vi)+2, TableWriterAlignmentCenter)
	} else {
		for c, value := range row.Values {
			align := alignments[c]
			t.renderCell(b, value, t.lengths[c], align)
			if c < len(row.Values)-1 {
				b.WriteString(vi)
			}
		}
	}
	b.WriteString(ve)
	b.WriteString("\n")
}

func (t *TableWriter) renderCell(b *strings.Builder, value string, length int, align TableWriterAlignment) {
	switch align {
	case TableWriterAlignmentLeft:
		b.WriteString(PadRight(Format(" %s ", value), length+2))
	case TableWriterAlignmentCenter:
		b.WriteString(PadCenter(Format(" %s ", value), length+2))
	case TableWriterAlignmentRight:
		b.WriteString(PadLeft(Format(" %s ", value), length+2))
	}
}

// Use i=0 for top border, 1 for the border between the first and second row,
// etc. use i=len(rows) for the bottom border.
func (t *TableWriter) getBorderSeparator(i int) (external, meta, section bool) {
	if i >= len(t.rows) {
		i = len(t.rows) - 1
	}
	j := i - 1
	if i == 0 {
		j = i
	}

	curRow := t.rows[i]
	prvRow := t.rows[j]

	return i == 0 || i == len(t.rows),
		curRow.Meta || prvRow.Meta,
		len(curRow.Values) == 1 || len(prvRow.Values) == 1
}

func (t *TableWriter) getBorderParts(i int) (left, mid, right, h string) {
	_, meta, _ := t.getBorderSeparator(i)
	flavor := t.flavor.Data
	if meta {
		flavor = t.flavor.Meta
	}

	left = flavor.MidLeft
	mid = flavor.MidMid
	right = flavor.MidRight
	h = flavor.HorizontalInternal
	if i == 0 {
		left = flavor.TopLeft
		mid = flavor.TopMid
		right = flavor.TopRight
		h = flavor.HorizontalExternal
	} else if i == len(t.rows) {
		left = flavor.BottomLeft
		mid = flavor.BottomMid
		right = flavor.BottomRight
		h = flavor.HorizontalExternal
	}
	return left, mid, right, h
}
