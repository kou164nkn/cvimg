package cvimg

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func SearchAndConvert(dir, src, dst string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+src {
			err := convertImg(path, src, dst)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func convertImg(path, src, dst string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var image image.Image

	switch src {
	case "jpg", "jpeg":
		image, err = jpeg.Decode(file)
		if err != nil {
			return err
		}
	case "png":
		image, err = png.Decode(file)
		if err != nil {
			return err
		}
	case "gif":
		image, err = gif.Decode(file)
		if err != nil {
			return err
		}
	default:
	}

	fileName := strings.TrimRight(path, src)
	dstFileName := fileName + dst

	dstFile, err := os.Create(dstFileName)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	switch dst {
	case "jpg", "jpeg":
		if err := jpeg.Encode(dstFile, image, nil); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(dstFile, image); err != nil {
			return err
		}
	case "gif":
		if err := gif.Encode(dstFile, image, nil); err != nil {
			return err
		}
	default:
	}

	return nil
}
