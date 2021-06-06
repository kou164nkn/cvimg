package cvimg

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func SearchAndConvert(dir, src, dst string) {
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
