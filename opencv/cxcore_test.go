package opencv

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"syscall"
	"testing"
)

func TestLoadImage2(t *testing.T) {
	// t.Errorf("aaa")
}

func TestInitFont(t *testing.T) {
	// Will assert at the C layer on error
	InitFont(CV_FONT_HERSHEY_DUPLEX, 1, 1, 0, 1, 8)
}

func TestPutText(t *testing.T) {
	_, currentfile, _, _ := runtime.Caller(0)
	filename := path.Join(path.Dir(currentfile), "../images/lena.jpg")

	image := LoadImage(filename)
	if image == nil {
		t.Fatal("LoadImage fail")
	}
	defer image.Release()

	// Write 'Hello' on the image
	font := InitFont(CV_FONT_HERSHEY_DUPLEX, 1, 1, 0, 1, 8)
	color := NewScalar(255, 255, 255, 0)

	pos := Point{image.Width() / 2, image.Height() / 2}
	font.PutText(image, "Hello", pos, color)

	filename = path.Join(path.Dir(currentfile), "../images/lena_with_text.jpg")

	// Uncomment this code to create the test image "../images/lena_with_text.jpg"
	// It is part of the repo, and what this test compares against
	//
	// SaveImage(filename, image, 0)
	// println("Saved file", filename)

	tempfilename := path.Join(os.TempDir(), "lena_with_text.jpg")
	defer syscall.Unlink(tempfilename)

	SaveImage(tempfilename, image, 0)

	// Compare actual image with expected image
	same, err := BinaryCompare(filename, tempfilename)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Error("Actual file differs from expected file with text")
	}
}

// Compare two files, return true if exactly the same
func BinaryCompare(file1, file2 string) (bool, error) {
	f1, err := ioutil.ReadFile(file1)
	if err != nil {
		return false, err
	}

	f2, err := ioutil.ReadFile(file2)
	if err != nil {
		return false, err
	}

	return bytes.Equal(f1, f2), nil
}
