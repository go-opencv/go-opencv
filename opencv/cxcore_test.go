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

func TestAbsDiff(t *testing.T) {
	_, currentfile, _, _ := runtime.Caller(0)
	filename := path.Join(path.Dir(currentfile), "../images/lena.jpg")

	org := LoadImage(filename)
	modified := LoadImage(filename)
	diff := CreateImage(org.Width(), org.Height(), IPL_DEPTH_8U, 3)

	if org == nil || modified == nil {
		t.Fatal("LoadImage fail")
	}
	defer org.Release()
	defer modified.Release()
	defer diff.Release()

	// Write 'Hello' on the image
	font := InitFont(CV_FONT_HERSHEY_DUPLEX, 1, 1, 0, 1, 8)
	color := NewScalar(255, 255, 255, 0)

	pos := Point{modified.Width() / 2, modified.Height() / 2}
	font.PutText(modified, "Hello", pos, color)

	// diff the images witwh hello on it and the original one
	AbsDiff(org, modified, diff)

	// very basic checking, most of the image should be black and only
	// the "hello" pixels should remain. We should expect this many
	// black pixels = 260766
	black_pixels := 0

	for x := 0; x < diff.Width()-1; x++ {
		for y := 0; y < diff.Height()-1; y++ {
			pixel := diff.Get2D(x, y).Val()

			if pixel[0] == 0.0 && pixel[1] == 0.0 && pixel[2] == 0.0 {
				black_pixels++
			}
		}
	}

	if black_pixels != 260766 {
		t.Error("Unexpected result for AbsDiff")
	}
}

func TestAddSub(t *testing.T) {

	checkVals := func(img *IplImage, val float64, debug string) {
	loop:
		for j := 0; j < img.Height(); j++ {
			for i := 0; i < img.Width(); i++ {
				pix := img.Get2D(i, j).Val()
				if pix[0] != val || pix[1] != val || pix[2] != val {
					t.Errorf("Unexpeted value for %s: %.1f, %.1f, %.1f. Expected %.1fs",
						debug, pix[0], pix[1], pix[2], val)
					break loop
				}
			}
		}
	}

	zeroImg := CreateImage(50, 50, IPL_DEPTH_8U, 3)
	zeroImg.Zero()

	twosImg := zeroImg.Clone()
	foursImg := zeroImg.Clone()
	negTwosImg := zeroImg.Clone()
	defer zeroImg.Release()
	defer twosImg.Release()
	defer foursImg.Release()
	defer negTwosImg.Release()

	two := NewScalar(2, 2, 2, 2)

	// 0 + 2 = 2
	AddScalar(zeroImg, two, twosImg)
	checkVals(twosImg, 2, "AddScalar()")

	// 2 + 2 = 4
	Add(twosImg, twosImg, foursImg)
	checkVals(foursImg, 4, "Add()")

	// 4 - 2 = 2
	Subtract(foursImg, twosImg, twosImg)
	checkVals(twosImg, 2, "Sub()")

	// 2 - 2 = 0
	SubScalar(twosImg, two, zeroImg)
	checkVals(zeroImg, 0, "SubScalar()")

	// 2 - 4 = 0 != -2 because it clips
	SubScalarRev(two, foursImg, negTwosImg)
	checkVals(negTwosImg, 0, "SubScalarRev()")

}
