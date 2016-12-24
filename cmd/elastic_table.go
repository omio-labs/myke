package cmd

import (
	"github.com/olekukonko/ts"
	"github.com/tgulacsi/wrap"
	"strings"
	"strconv"
	"sort"
	"fmt"
	"io"
)

const DEFAULT_TERM_WIDTH = 80

type ElasticCol struct {
	index int
	min int
	max int
	weight int
	width int
}

type ElasticTable struct {
	cols []ElasticCol
	header []string
	rows [][]string
}

func (e *ElasticTable) Init(header []string) {
	e.header = header
	e.cols = make([]ElasticCol, len(header))
	e.rows = [][]string{}
	for i, v := range header {
		l, c := len(v), &e.cols[i]
		c.index = i
		c.min, c.max, c.width = l, l, l
		c.weight = 1
	}
}

func (e *ElasticTable) AddRow(row []string) {
	e.rows = append(e.rows, row)
	for i, v := range row {
		l, c := len(v), &e.cols[i]
		if l < c.min {
			c.min = l
		}
		if l > c.max {
			c.max, c.width = l, l
		}
	}
}

func (e *ElasticTable) mapWidths(f func(col ElasticCol) int) ([]int) {
	out := make([]int, len(e.cols))
	for _, v := range e.cols {
		out[v.index] = f(v)
	}
	return out
}

func (e *ElasticTable) optimizedWidths(padding int) ([]int) {
	num := len(e.cols)
	termWidth := termWidth() - (num*padding)
	sort.Sort(elasticSortMax(e.cols))

	minTot, maxTot := 0, 0
	for _, v := range e.cols {
		minTot = minTot + v.min
		maxTot = maxTot + v.width
	}

	OUTER:
	for {
		if minTot > termWidth {
			return e.mapWidths(func(col ElasticCol) int { return col.min })
		} else if maxTot < termWidth {
			break
		}

		for i := 0; i < num-1; i++ {
			curr, next := &e.cols[i], &e.cols[i+1]
			width := curr.max / (curr.weight + 1)
			if width >= next.width {
				maxTot = maxTot - (curr.width - width)
				curr.weight = curr.weight + 1
				curr.width = width
				continue OUTER
			}
		}

		// no further optimizations can be performed
		break
	}

	if balance := termWidth - maxTot; balance > 0 {
		// distribute remaining whitespace to largest column
		e.cols[0].width = e.cols[0].width + balance
	}

	return e.mapWidths(func(col ElasticCol) int { return col.width })
}

func (e *ElasticTable) Render(out io.Writer) {
	widths := e.optimizedWidths(3)
	divider := make([]string, len(widths))
	for i, v := range widths {
		divider[i] = strings.Repeat("-", v)
	}
	printRow(out, e.header, widths, "|", " ")
	printRow(out, divider, widths, "--+", "")
	for _, row := range e.rows {
		printRow(out, row, widths, "|", " ")
	}
}

type elasticSortMax []ElasticCol
func (s elasticSortMax) Len() int {
	return len(s)
}
func (s elasticSortMax) Swap(i int, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s elasticSortMax) Less(i int, j int) bool {
    return s[i].max >= s[j].max
}

func termWidth() (int) {
	termWidth := DEFAULT_TERM_WIDTH
	if termSize, err := ts.GetSize(); err == nil && termSize.Col() > 0 {
		termWidth = termSize.Col()
	}
	return termWidth
}

func printRow(out io.Writer, row []string, widths []int, border string, padding string) {
	colmax := len(row)
	subrows := make([][]string, colmax)
	submax := 1
	for i, w := range widths {
		subrows[i] = strings.Split(wrap.String(row[i], uint(w)), "\n")
		if len(subrows[i]) > submax {
			submax = len(subrows[i])
		}
	}

	for sub := 0; sub < submax; sub++ {
		for i, w := range widths {
			str := ""
			format := padding + "%-" + strconv.Itoa(w) + "s" + padding
			if i < colmax-1 {
				format = format + border
			} else {
				format = format + "\n"
			}
			if sub < len(subrows[i]) {
				str = subrows[i][sub]
			}
			fmt.Fprintf(out, format, str)
		}
	}
}
