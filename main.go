package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Concat("a", "b"))
}

func Concat(a, b string) string {
	var builder strings.Builder
	builder.Grow(len(a) + len(b))
	builder.WriteString(a)
	builder.WriteString(b)
	return builder.String()
}