/*
	This is command line tool for converting image file.
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	dir       = flag.String("dir", ".", "the name of target dierctory")
	src       = flag.String("src", "jpg", "the format before converting")
	dst       = flag.String("dst", "png", "the format after converting")
	imgFormat = [...]string{"jpg", "jpeg", "png", "gif"}
)

func main() {
	flag.Parse()

	if errs := validArgs(*dir, *src, *dst); len(errs) > 0 {
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}

	searchAndConvert(*dir, *src, *dst)
}

func validArgs(dir, src, dst string) []error {
	var errs []error

	if f, err := os.Stat(dir); err != nil || !f.IsDir() {
		errs = append(errs, errors.New(dir+": You specified non existing directory"))
	}

	if err := validFormat(src); err != nil {
		errs = append(errs, err)
	}

	if err := validFormat(dst); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func validFormat(ext string) error {
	for _, v := range imgFormat {
		if v == ext {
			return nil
		}
	}
	return errors.New(ext + ": You specified invalid image file format")
}
