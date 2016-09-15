package opencv

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"runtime"
	"syscall"
	"testing"
	"time"
)

func init() {
	// seed random number generator for
	// creating random mask for TestAddSubWithMask below
	rand.Seed(time.Now().Unix())
}

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
	w, h := 50, 50
	checkVals := func(img *IplImage, val float64, debug string) {
		for i := 0; i < w*h; i++ {
			pix := img.Get1D(i).Val()
			if pix[0] != val || pix[1] != val || pix[2] != val {
				t.Errorf("Unexpeted value for %s: %.1f, %.1f, %.1f. Expected %.1fs",
					debug, pix[0], pix[1], pix[2], val)
				break
			}
		}
	}

	zeroImg := CreateImage(w, h, IPL_DEPTH_8U, 3)
	zeroImg.Zero()

	hundredImg := zeroImg.Clone()
	twoHundredImg := zeroImg.Clone()
	negImage := zeroImg.Clone()
	defer zeroImg.Release()
	defer hundredImg.Release()
	defer twoHundredImg.Release()
	defer negImage.Release()

	hundred := NewScalar(100, 100, 100, 0)

	// 0 + 100 = 100
	AddScalar(zeroImg, hundred, hundredImg)
	checkVals(hundredImg, 100, "AddScalar()")

	// 100 + 100 = 200
	Add(hundredImg, hundredImg, twoHundredImg)
	checkVals(twoHundredImg, 200, "Add()")

	// 200 - 100 = 100
	Subtract(twoHundredImg, hundredImg, hundredImg)
	checkVals(hundredImg, 100, "Sub()")

	// 100 - 100 = 0
	SubScalar(hundredImg, hundred, zeroImg)
	checkVals(zeroImg, 0, "SubScalar()")

	// 100 - 200 = 0 != -100 because it clips
	SubScalarRev(hundred, twoHundredImg, negImage)
	checkVals(negImage, 0, "SubScalarRev()")

	// Uncomment to save these images to disk
	// SaveImage("zeroImg.png", zeroImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("hundredImg.png", hundredImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("twoHundredImg.png", twoHundredImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("negImage.png", negImage, CV_IMWRITE_PNG_COMPRESSION)

}

func TestAddSubWithMask(t *testing.T) {
	w, h := 50, 50
	checkValsWMask := func(img, mask *IplImage, val float64, debug string) {
		for i := 0; i < w*h; i++ {
			pix := img.Get1D(i).Val()
			tst := val
			if mask.Get1D(i).Val()[0] == 0 {
				tst = 0
			}
			if pix[0] != tst || pix[1] != tst || pix[2] != tst {
				t.Errorf("Unexpeted value for %s: %.1f, %.1f, %.1f. Expected %.1fs",
					debug, pix[0], pix[1], pix[2], val)
				break
			}
		}
	}

	zeroImg := CreateImage(w, h, IPL_DEPTH_8U, 3)
	zeroImg.Zero()

	hundredImg := zeroImg.Clone()
	twoHundredImg := zeroImg.Clone()
	negImage := zeroImg.Clone()
	defer zeroImg.Release()
	defer hundredImg.Release()
	defer twoHundredImg.Release()
	defer negImage.Release()

	// generate a random mask
	maskImg := CreateImage(w, h, IPL_DEPTH_8U, 1)
	defer maskImg.Release()
	for i := 0; i < w*h; i++ {
		oneOrZero := float64(rand.Intn(2)) // random number either 1 or 0
		maskImg.Set1D(i, NewScalar(oneOrZero*255, 0, 0, 0))
	}

	hundred := NewScalar(100, 100, 100, 0)

	// 0 + 100 = 100
	AddScalarWithMask(zeroImg, hundred, hundredImg, maskImg)
	checkValsWMask(hundredImg, maskImg, 100, "AddScalarWithMask()")

	// 100 + 100 = 200
	AddWithMask(hundredImg, hundredImg, twoHundredImg, maskImg)
	checkValsWMask(twoHundredImg, maskImg, 200, "AddWithMask()")

	// 200 - 100 = 100
	SubtractWithMask(twoHundredImg, hundredImg, hundredImg, maskImg)
	checkValsWMask(hundredImg, maskImg, 100, "SubtractWithMask()")

	// 100 - 100 = 0
	SubScalarWithMask(hundredImg, hundred, zeroImg, maskImg)
	checkValsWMask(zeroImg, maskImg, 0, "SubScalarWithMask()")

	// 100 - 200 = 0 != -100 because it clips
	SubScalarWithMaskRev(hundred, twoHundredImg, negImage, maskImg)
	checkValsWMask(negImage, maskImg, 0, "SubScalarWithMaskRev()")

	// Uncomment to save these images to disk
	// SaveImage("zeroImgMask.png", zeroImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("hundredImgMask.png", hundredImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("twoHundredImgMask.png", twoHundredImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("negImageMask.png", negImage, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("MaskImg.png", maskImg, CV_IMWRITE_PNG_COMPRESSION)

}
