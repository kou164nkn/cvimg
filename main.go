package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
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

// The following code is converting image function
func searchAndConvert(dir, src, dst string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+src {
			convertImg(path, src, dst)
		}
		return nil
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func convertImg(path, src, dst string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	var image image.Image

	switch src {
	case "jpg", "jpeg":
		image, err = jpeg.Decode(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	case "png":
		image, err = png.Decode(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	case "gif":
		image, err = gif.Decode(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	default:
	}

	fileName := strings.TrimRight(path, src)
	dstFileName := fileName + dst

	dstFile, err := os.Create(dstFileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer dstFile.Close()

	switch dst {
	case "jpg", "jpeg":
		if err := jpeg.Encode(dstFile, image, nil); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	case "png":
		if err := png.Encode(dstFile, image); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	case "gif":
		if err := gif.Encode(dstFile, image, nil); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	default:
	}

	return
}
