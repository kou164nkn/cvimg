package main_test

import (
	"errors"
	"os"
	"testing"

	"github.com/kou164nkn/cvimg/cmd/cvimg"
)

type MockDir struct {
	os.FileInfo
	dirFlag bool
}

func (m MockDir) IsDir() bool { return m.dirFlag }

func (m MockDir) Stat(dir string) (os.FileInfo, error) {
	if dir == "./errorPath" {
		return nil, errors.New("")
	}
	return m, nil
}

// HACK: Rewrite including the product code
func TestValidArgs(t *testing.T) {
	cases := map[string]struct {
		tgtDir  string
		dirFlag bool
		srcFmt  string
		dstFmt  string

		expectErr [3]error
	}{
		"nothing_wrong":  {"./example", true, "jpg", "png", [3]error{}},
		"osStatFail":     {"./errorPath", true, "gif", "jpeg", [3]error{errors.New("./errorPath: You specified non existing directory")}},
		"DirIsFile":      {"./example", false, "jpeg", "jpg", [3]error{errors.New("./example: You specified non existing directory")}},
		"wrongSrcFormat": {"./example", true, "wrongSrc", "gif", [3]error{nil, errors.New("wrongSrc: You specified invalid image file format"), nil}},
		"wrongDstFormat": {"./example", true, "png", "wrongDst", [3]error{nil, nil, errors.New("wrongDst: You specified invalid image file format")}},
		"allArgsWrong":   {"./errorPath", false, "wrongSrc", "wrongDst", [3]error{errors.New("./errorPath: You specified non existing directory"), errors.New("wrongSrc: You specified invalid image file format"), errors.New("wrongDst: You specified invalid image file format")}},
	}

	for name, tt := range cases {
		var c main.Cvimg
		c.Dir = MockDir{dirFlag: tt.dirFlag}

		t.Run(name, func(t *testing.T) {
			errs := main.ValidArgs(c, tt.tgtDir, tt.srcFmt, tt.dstFmt)

			for i := 0; i < len(errs); i++ {
				switch {
				case errs[i] == nil && tt.expectErr[i] == nil:
					continue
				case errs[i] != nil && tt.expectErr[i] == nil:
					t.Errorf("want %v but got %v", tt.expectErr, errs)
					goto JUMP
				case errs[i] == nil && tt.expectErr[i] != nil:
					t.Errorf("want %v but got %v", tt.expectErr, errs)
					goto JUMP
				case errs[i].Error() != tt.expectErr[i].Error():
					t.Errorf("want %v but got %v", tt.expectErr, errs)
					goto JUMP
				}
			}
		JUMP:
		})
	}
}
