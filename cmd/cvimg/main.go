/*
	This is command line tool for converting image file.
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kou164nkn/cvimg"
)

var (
	dir = flag.String("dir", ".", "the name of target dierctory")
	src = flag.String("src", "jpg", "the format before converting")
	dst = flag.String("dst", "png", "the format after converting")
)

func main() {
	flag.Parse()

	c := cvimg.Cvimg{SrcFormat: *src, DstFormat: *dst}

	if !c.ValidFormat() {
		fmt.Fprintln(os.Stderr, errors.New("supported format is: "+strings.Join(cvimg.SupportedFormats, ",")))
		os.Exit(1)
	}

	if err := c.Search(*dir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := c.Convert(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
