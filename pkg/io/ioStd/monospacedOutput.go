package ioStd

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func MonospacedPuts(prefix string, values []string) error {
	width := 0
	if w, _, err := term.GetSize(syscall.Stdin); err != nil {
		return err
	} else {
		width = w
	}

	if mono, err := GetMonospacedExps(values, width-getDispLen(prefix)); err != nil {
		return err
	} else {
		fmt.Println(prefix + strings.Join(mono, fmt.Sprintf("\n%s", prefix)))
	}

	return nil
}
