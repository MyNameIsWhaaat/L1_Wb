package main

import (
	"fmt"
	"io"
)

type LegacyPrinter struct{}

func (LegacyPrinter) Print(s string) int {
	fmt.Println("Legacy:", s)
	return len(s)
}

type PrinterAdapter struct {
	P LegacyPrinter
}

func (a PrinterAdapter) Write(p []byte) (int, error) {
	n := a.P.Print(string(p))
	if n < len(p) {
		return n, io.ErrShortWrite
	}
	return n, nil
}

func logSomething(w io.Writer) {
	fmt.Fprintln(w, "Hello, adapter!")
}

func main() {
	legacy := LegacyPrinter{}
	adapter := PrinterAdapter{P: legacy}

	logSomething(adapter)
}