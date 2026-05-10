package strx

import "strings"

type TableAlignment int8

const (
	TableAuto TableAlignment = iota
	TableLeft
	TableCenter
	TableRight
)

var TableStyleAscii = TableStyle{
	"+", "+", "+", "+", "+", "+", "+", "+", "+", "-", "-", "|", "|",
	"+", "+", "+", "", "", "", "+", "+", "+", "-", "", "|", "|",
}
var TableStyleAsciiSeparated = TableStyle{
	"+", "+", "+", "+", "+", "+", "+", "+", "+", "=", "=", "|", "|",
	"+", "+", "+", "+", "+", "+", "+", "+", "+", "-", "-", "|", "|",
}
var TableStyleAsciiCompact = TableStyle{
	"", "", "", "", " ", "", "", "", "", "", "-", "", " ",
	"", "", "", "", "", "", "", "", "", "", "", "", " ",
}
var TableStyleAsciiDots = TableStyle{
	".", ".", ".", ":", ":", ":", ":", ":", ":", ".", ".", ":", ":",
	".", ".", ".", "", "", "", ":", ":", ":", ".", "", ":", ":",
}
var TableStyleGithub = TableStyle{
	"", "", "", "|", "|", "|", "", "", "", "", "-", "|", "|",
	"", "", "", "", "", "", "", "", "", "", "", "|", "|",
}
var TableStyleReddit = TableStyle{
	"", "", "", "", "|", "", "", "", "", "", "-", "", "|",
	"", "", "", "", "", "", "", "", "", "", "", "", "|",
}
var TableStyleRestructuredGrid = TableStyle{
	"+", "+", "+", "+", "+", "+", "+", "+", "+", "-", "=", "|", "|",
	"+", "+", "+", "", "", "", "+", "+", "+", "-", "", "|", "|",
}
var TableStyleRestructuredSimple = TableStyle{
	"", " ", "", "", " ", "", "", " ", "", "=", "=", "", " ",
	"", " ", "", "", "", "", "", " ", "", "=", "", "", " ",
}
var TableStyleUnicode = TableStyle{
	"┌", "┬", "┐", "├", "┼", "┤", "└", "┴", "┘", "─", "─", "│", "│",
	"┌", "┬", "┐", "", "", "", "└", "┴", "┘", "─", "", "│", "│",
}
var TableStyleUnicodeDouble = TableStyle{
	"╔", "╦", "╗", "╠", "╬", "╣", "╚", "╩", "╝", "═", "═", "║", "║",
	"╔", "╦", "╗", "", "", "", "╚", "╩", "╝", "═", "", "║", "║",
}

type TableStyle struct {
	MTL, MTM, MTR, MML, MMM, MMR, MBL, MBM, MBR, MHE, MHI, MVE, MVI string
	DTL, DTM, DTR, DML, DMM, DMR, DBL, DBM, DBR, DHE, DHI, DVE, DVI string
}

type TableRow struct {
	values     []string
	alignments []TableAlignment
	meta       bool
	section    bool
}

func (t *TableRow) ToLeft() *TableRow {
	for i := range t.values {
		t.alignments[i] = TableLeft
	}
	return t
}
func (t *TableRow) ToCenter() *TableRow {
	for i := range t.values {
		t.alignments[i] = TableCenter
	}
	return t
}
func (t *TableRow) ToRight() *TableRow {
	for i := range t.values {
		t.alignments[i] = TableRight
	}
	return t
}
func (t *TableRow) WithAlignments(alignments ...TableAlignment) *TableRow {
	t.alignments = alignments[:len(t.values)]
	return t
}

type Table struct {
	cols    int
	rows    []*TableRow
	lengths []int
	style   TableStyle
	align   TableAlignment
}

func NewTable() *Table {
	return &Table{
		cols:    0,
		rows:    make([]*TableRow, 0),
		lengths: make([]int, 0),
		style:   TableStyleAscii,
		align:   TableAuto,
	}
}

func (t *Table) Meta(values ...any) *TableRow {
	return t.addRow(true, false, values...)
}
func (t *Table) MetaSection(values any) *TableRow {
	return t.addRow(true, true, values)
}
func (t *Table) Data(values ...any) *TableRow {
	return t.addRow(false, false, values...)
}
func (t *Table) DataSection(values any) *TableRow {
	return t.addRow(false, true, values)
}

func (t *Table) WithStyle(style TableStyle) *Table {
	t.style = style
	return t
}

func (t *Table) WithAlignAuto() *Table {
	t.align = TableAuto
	return t
}

func (t *Table) WithAlignLeft() *Table {
	t.align = TableLeft
	return t
}
func (t *Table) WithAlignCenter() *Table {
	t.align = TableCenter
	return t
}

func (t *Table) WithAlignRight() *Table {
	t.align = TableRight
	return t
}

func (t *Table) Render() string {
	b := &strings.Builder{}
	t.renderBorder(b, 0)
	for i := range t.rows {
		t.renderRow(b, i)
		t.renderBorder(b, i)
	}
	return b.String()
}

func (t *Table) addRow(meta bool, section bool, values ...any) *TableRow {
	row := &TableRow{
		values:     make([]string, len(values)),
		alignments: make([]TableAlignment, len(values)),
		meta:       meta,
		section:    section,
	}
	for i := range row.values {
		v, a := t.toString(values[i])
		row.values[i] = v
		row.alignments[i] = a
		if t.align != TableAuto {
			row.alignments[i] = t.align
		} else if meta || section {
			row.alignments[i] = TableCenter
		}
		if i >= len(t.lengths) {
			t.lengths = append(t.lengths, 0)
		}
		t.lengths[i] = max(t.lengths[i], Length(v))
	}
	t.cols = max(len(values), t.cols)
	t.rows = append(t.rows, row)
	return row
}

func (t *Table) toString(value any) (string, TableAlignment) {
	switch v := value.(type) {
	case string:
		return v, TableLeft
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return Format("%d", v), TableRight
	case float32, float64:
		return Format("%.2f", v), TableRight
	default:
		return Format("%v", v), TableLeft
	}
}

func (t *Table) renderBorder(b *strings.Builder, i int) {
	left, mid, right, h := t.getBorderParts(i)
	if left == "" && mid == "" && right == "" && h == "" {
		return
	}

	rowAbove, rowBelow := t.getBorderRows(i)
	basedRow := rowAbove
	if (rowAbove == nil || rowAbove.section) && rowBelow != nil {
		basedRow = rowBelow
	}

	if basedRow.section {
		length := t.getTotalLength()
		dividers := Length(mid) * t.cols
		paddings := 2 * t.cols
		b.WriteString(Format("%s%s%s", left, strings.Repeat(h, length+dividers+paddings), right))
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

// TODO: consider multi-line cells when rendering borders and rows
func (t *Table) renderRow(b *strings.Builder, i int) {
	row := t.rows[i]
	ve, vi := t.getVerticalSeparator(i)
	if row.section {
		length := t.getTotalLength()
		dividers := Length(vi) * t.cols
		paddings := 2 * t.cols
		value := t.alignValue(row.values[0], row.alignments[0], length+dividers+paddings)
		b.WriteString(Format("%s%s%s\n", ve, value, ve))
	} else {
		b.WriteString(ve)
		for c, value := range row.values {
			align := row.alignments[c]
			length := t.lengths[c]
			value = t.alignValue(value, align, length)
			b.WriteString(Format(" %s ", value))
			if c < len(row.values)-1 {
				b.WriteString(vi)
			}
		}
		b.WriteString(ve)
		b.WriteString("\n")
	}
}

func (t *Table) renderCell(b *strings.Builder, row *TableRow, col int) {
	value := row.values[col]
	align := row.alignments[col]
	length := t.lengths[col]
	value = t.alignValue(value, align, length)
	b.WriteString(Format(" %s ", value))
}

func (t *Table) getBorderRows(i int) (above, below *TableRow) {
	var rowAbove, rowBelow *TableRow
	if i > 0 {
		rowAbove = t.rows[i-1]
	}
	if i < len(t.rows) {
		rowBelow = t.rows[i]
	}
	return rowAbove, rowBelow
}

func (t *Table) getBorderParts(i int) (left, mid, right, h string) {
	rowAbove, rowBelow := t.getBorderRows(i)
	isMeta := rowAbove != nil && rowAbove.meta || rowBelow != nil && rowBelow.meta

	switch {
	case rowAbove == nil && isMeta:
		return t.style.MTL, t.style.MTM, t.style.MTR, t.style.MHE
	case rowAbove == nil && !isMeta:
		return t.style.DBL, t.style.DTM, t.style.DTR, t.style.DHE
	case rowBelow == nil && isMeta:
		return t.style.MBL, t.style.MBM, t.style.MBR, t.style.MHE
	case rowBelow == nil && !isMeta:
		return t.style.DBL, t.style.DBM, t.style.DBR, t.style.DHE
	case isMeta:
		switch {
		case rowAbove.section:
			return t.style.MML, t.style.MTM, t.style.MMR, t.style.MHI
		case rowBelow.section:
			return t.style.MML, t.style.MTM, t.style.MMR, t.style.MHI
		default:
			return t.style.MML, t.style.MMM, t.style.MMR, t.style.MHI
		}
	default:
		switch {
		case rowAbove.section:
			return t.style.DML, t.style.DTM, t.style.DMR, t.style.DHI
		case rowBelow.section:
			return t.style.DML, t.style.DTM, t.style.DMR, t.style.DHI
		default:
			return t.style.DML, t.style.DMM, t.style.DMR, t.style.DHI
		}
	}
}

func (t *Table) getVerticalSeparator(i int) (ve, vi string) {
	if t.rows[i].meta {
		return t.style.MVE, t.style.MVI
	} else {
		return t.style.DVE, t.style.DVI
	}
}

func (t *Table) getTotalLength() int {
	total := 0
	for _, length := range t.lengths {
		total += length
	}
	return total
}

func (t *Table) alignValue(v string, align TableAlignment, length int) string {
	switch align {
	case TableLeft:
		return PadRight(v, length)
	case TableCenter:
		return PadCenter(v, length)
	case TableRight:
		return PadLeft(v, length)
	default:
		return v
	}
}
