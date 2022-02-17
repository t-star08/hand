package ioStd

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func isDispWidth2(r rune) bool {
    return !(0x0000 <= r && r <= 0x007f || 0xff61 <= r && r <= 0xffdc || 0xffe8 <= r && r <= 0xffee)
}

func getDispLen(s string) int {
	x := 0
	for _, r := range s {
		if isDispWidth2(r) {
			x += 2
		} else {
			x += 1
		}
	}
	return x
}

func buryLeftSpaceUntilDispLen(s string, width int) string {
	diff := width - getDispLen(s)
	if diff <= 0 {
		return s
	}
	return s + strings.Repeat(" ", diff)
}

func GetMonospacedExps(values []string, width int) ([]string, error) {
	dispLenTable := make(map[string]int)
	for _, s := range values {
		dispLenTable[s] = getDispLen(s)
	}

	sort.Strings(values)
	sortedValues := append(make([]string, len(values)), values...)
	sort.SliceStable(
		sortedValues,
		func(i, j int) bool {
			return dispLenTable[sortedValues[i]] > dispLenTable[sortedValues[j]]
		},
	)

	maxDisplayable, baseMargin, p := 0, 2, 0
	for i, s := range sortedValues {
		if p + dispLenTable[s] > width {
			break
		}
		p += dispLenTable[s] + baseMargin
		maxDisplayable = i
	}
	if maxDisplayable == 0 {
		return nil, fmt.Errorf("no enough space")
	}

	rows := int(math.Ceil(float64(len(values)) / float64(maxDisplayable)))
	rawDisp := make([][]string, rows)
	for i, s := range values {
		rawDisp[i % rows] = append(rawDisp[i % rows], s)
	}

	margins := make([]int, len(rawDisp[0]))
	for i := 0; i < len(rawDisp[0]); i++ {
		maxDispLenInCol := 0
		if i == len(rawDisp[0]) - 1 {
			margins[i] = 0
			continue
		}
		for j := 0; j < len(rawDisp); j++ {
			if i > len(rawDisp[j]) - 1 {
				break
			}
			if dispLenTable[rawDisp[j][i]] > maxDispLenInCol {
				maxDispLenInCol = dispLenTable[rawDisp[j][i]]
			}
		}
		margins[i] = maxDispLenInCol + baseMargin
	}

	disp := make([]string, rows)
	for i, d := range rawDisp {
		rowValue := ""
		for j, e := range d {
			rowValue += buryLeftSpaceUntilDispLen(e, margins[j])
		}
		disp[i] = rowValue
	}

	return disp, nil
}
