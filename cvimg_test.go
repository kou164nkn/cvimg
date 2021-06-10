package cvimg_test

import (
	"bufio"
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/kou164nkn/cvimg"
)

func TestValidFormat(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		srcFormat string
		dstFormat string
		expect    bool
	}{
		"validFormat1":      {"jpg", "png", true},
		"validFormat2":      {"jpeg", "gif", true},
		"InvalidSrcFormat":  {"wrongFormat", "jpg", false},
		"InvalidDstFormat":  {"png", "wrongFormat", false},
		"InvalidBothFormat": {"wrongFormat", "wrongFormat", false},
	}

	for name, tt := range cases {
		tt := tt

		c := cvimg.Cvimg{SrcFormat: tt.srcFormat, DstFormat: tt.dstFormat}

		t.Run(name, func(t *testing.T) {
			actual := c.ValidFormat()

			if actual != tt.expect {
				t.Errorf("ValidFormat want %v but got %v. (SrcFormat: %v, DstFormat: %v)", tt.expect, actual, tt.srcFormat, tt.dstFormat)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		searchPath  string
		srcFormat   string
		expectPaths []string
	}{
		"jpg":    {"./testdata", "jpg", []string{"./testdata/sample1.jpg", "./testdata/foo/sample1.jpg"}},
		"jpeg":   {"./testdata", "jpeg", []string{"./testdata/sample2.jpeg", "./testdata/foo/sample2.jpeg"}},
		"png":    {"./testdata", "png", []string{"./testdata/sample3.png", "./testdata/foo/sample3.png"}},
		"gif":    {"./testdata", "gif", []string{"./testdata/sample4.gif", "./testdata/foo/sample4.gif"}},
		"noData": {"./cmd", "jpg", []string{}},
	}

	for name, tt := range cases {
		tt := tt

		c := cvimg.Cvimg{SrcFormat: tt.srcFormat}

		t.Run(name, func(t *testing.T) {
			actualPaths := c.Search(tt.searchPath)

			if reflect.DeepEqual(actualPaths, tt.expectPaths) {
				t.Errorf("Search want %v but got %v. (SrcFormat: %v)", tt.expectPaths, actualPaths, tt.srcFormat)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		targetPath string
		srcFormat  string
		dstFormat  string

		resultFilePath string
		expectFilePath string
	}{
		"jpgToPng": {"./testdata/sample1.jpg", "jpg", "png", "./testdata/sample1.png", "./testdata/sample1.png.golden"},
		"jpgToGif": {"./testdata/sample1.jpg", "jpg", "gif", "./testdata/sample1.gif", "./testdata/sample1.gif.golden"},
		"pngToJpg": {"./testdata/sample3.png", "png", "jpg", "./testdata/sample3.jpg", "./testdata/sample3.jpg.golden"},
		"pngToGif": {"./testdata/sample3.png", "png", "gif", "./testdata/sample3.gif", "./testdata/sample3.gif.golden"},
		"gifToJpg": {"./testdata/sample4.gif", "gif", "jpg", "./testdata/sample4.jpg", "./testdata/sample4.jpg.golden"},
		"gifToPng": {"./testdata/sample4.gif", "gif", "png", "./testdata/sample4.png", "./testdata/sample4.png.golden"},
	}

	for name, tt := range cases {
		tt := tt

		c := cvimg.Cvimg{SrcFormat: tt.srcFormat, DstFormat: tt.dstFormat, TargetPaths: []string{tt.targetPath}}

		t.Run(name, func(t *testing.T) {
			if err := c.Convert(); err != nil {
				t.Fatal(err)
			}

			resultImageData := getImageData(t, tt.resultFilePath)
			expectImageData := getImageData(t, tt.expectFilePath)

			if !bytes.Equal(resultImageData, expectImageData) {
				t.Errorf("Convert produced different image data than expected")
			}
		})
		deleteresultFile(t, tt.resultFilePath)
	}
}

func getImageData(t *testing.T, filePath string) []byte {
	t.Helper()

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	return bytes
}

func deleteresultFile(t *testing.T, filePath string) {
	t.Helper()

	if err := os.Remove(filePath); err != nil {
		t.Fatal(err)
	}
}
