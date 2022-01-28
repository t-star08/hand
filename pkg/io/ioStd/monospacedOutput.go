package ioStd

import (
	"fmt"
	"math"
	"sort"
	"syscall"

	"golang.org/x/term"
)

func MonospacedPuts(values []string) error {
	sort.Strings(values)

	sortedValues := append(make([]string, len(values)), values...)
	sort.SliceStable(
		sortedValues,
		func(i, j int) bool {
			return len(sortedValues[i]) > len(sortedValues[j])
		},
	)

	width := 0
	if w, _, err := term.GetSize(syscall.Stdin); err != nil {
		return err
	} else {
		width = w
	}
	maxPlace, baseMargin, p := 0, 2, 0
	for i, s := range sortedValues {
		if p + len(s) > width {
			break
		}
		p += len(s) + baseMargin
		maxPlace = i
	}
	if maxPlace == 0 {
		return fmt.Errorf("Not Enough Space")
	}

	rows := int(math.Ceil(float64(len(values)) / float64(maxPlace)))
	disp := make([][]string, rows)
	for i, s := range values {
		disp[i % rows] = append(disp[i % rows], s)
	}

	margins := make([]int, len(disp[0]))
	for i := 0; i < len(disp[0]); i++ {
		maxPlaceInCol := 0
		if i == len(disp[0]) - 1 {
			margins[i] = 0
			continue
		}
		for j := 0; j < len(disp); j++ {
			if i > len(disp[j]) - 1 {
				break
			}
			if len(disp[j][i]) > maxPlaceInCol {
				maxPlaceInCol = len(disp[j][i])
			}
		}
		margins[i] = maxPlaceInCol + baseMargin
	}

	for _, d := range disp {
		row_value := ""
		for i, e := range d {
			row_value += fmt.Sprintf(fmt.Sprintf("%%%ds", -margins[i]), e)
		}
		fmt.Println(row_value)
	}

	return nil
}
