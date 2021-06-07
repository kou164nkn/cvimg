/*
	This is command line tool for converting image file.
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/kou164nkn/cvimg"
)

var (
	dir       = flag.String("dir", ".", "the name of target dierctory")
	src       = flag.String("src", "jpg", "the format before converting")
	dst       = flag.String("dst", "png", "the format after converting")
	imgFormat = [...]string{"jpg", "jpeg", "png", "gif"}
)

// consider testability for getting FileInfo
type Dir interface {
	Stat(string) (os.FileInfo, error)
}

type DirFunc func() (os.FileInfo, error)

func (f DirFunc) Stat(dir string) (os.FileInfo, error) {
	return f()
}

type Cvimg struct {
	Dir Dir
}

func (c Cvimg) stat(dir string) (os.FileInfo, error) {
	if c.Dir == nil {
		return os.Stat(dir)
	}
	return c.Dir.Stat(dir)
}

func main() {
	flag.Parse()

	var c Cvimg

	if errs := ValidArgs(c, *dir, *src, *dst); len(errs) > 0 {
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}

	cvimg.SearchAndConvert(*dir, *src, *dst)
}

func ValidArgs(c Cvimg, dir, src, dst string) [3]error {
	var errs [3]error

	if f, err := c.stat(dir); err != nil || !f.IsDir() {
		errs[0] = errors.New(dir + ": You specified non existing directory")
	}

	if err := validFormat(src); err != nil {
		errs[1] = err
	}

	if err := validFormat(dst); err != nil {
		errs[2] = err
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
