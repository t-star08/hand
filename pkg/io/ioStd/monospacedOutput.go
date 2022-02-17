package ioStd

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func MonospacedPuts(values []string) error {
	width := 0
	if w, _, err := term.GetSize(syscall.Stdin); err != nil {
		return err
	} else {
		width = w
	}

	if mono, err := GetMonospacedExps(values, width); err != nil {
		return err
	} else {
		fmt.Println(strings.Join(mono, "\n"))
	}

	return nil
}
