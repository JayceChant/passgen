package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var length int
	var chars string

	flag.IntVar(&length, "length", 20, "code length")
	flag.StringVar(&chars, "chars", "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz", "characters")
	flag.Parse()

	base := int64(len(chars))
	src := rand.NewSource(time.Now().UnixNano())
	builder := strings.Builder{}
	builder.Grow(length)
	for i := 0; i < length; i++ {
		builder.WriteByte(chars[src.Int63()%base])
	}
	fmt.Println(builder.String())
}
