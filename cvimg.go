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

var SupportedFormats = []string{"jpg", "jpeg", "png", "gif"}

type Cvimg struct {
	SrcFormat   string
	DstFormat   string
	TargetPaths []string
}

func (c Cvimg) ValidFormat() bool {
	var validSrc, validDst bool

	for _, v := range SupportedFormats {
		if v == c.SrcFormat {
			validSrc = true
		}
		if v == c.DstFormat {
			validDst = true
		}
	}

	return validSrc && validDst
}

func (c *Cvimg) Search(file string) error {
	err := filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+c.SrcFormat {
			c.TargetPaths = append(c.TargetPaths, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c Cvimg) Convert() error {
	for _, path := range c.TargetPaths {
		err := convertImage(path, c.SrcFormat, c.DstFormat)
		if err != nil {
			return err
		}
	}

	return nil
}

func convertImage(path, src, dst string) error {
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
