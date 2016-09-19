package opencv

import (
	"bytes"
	"io/ioutil"
	"math"
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

func TestLogic(t *testing.T) {
	w, h := 50, 50
	checkValsWMask := func(img, mask *IplImage, val float64, debug string) {
		for i := 0; i < w*h; i++ {
			pix := img.Get1D(i).Val()
			tst := val
			if mask.Get1D(i).Val()[0] == 0 {
				tst = 0
			}
			if pix[0] != tst || pix[1] != tst || pix[2] != tst {
				t.Errorf("Unexpeted value for %s: %f, %f, %f. Expected %fs",
					debug, pix[0], pix[1], pix[2], val)
				break
			}
		}
	}

	zero := NewScalar(0, 0, 0, 0)
	one := NewScalar(1, 1, 1, 0)

	zeroImg := CreateImage(w, h, IPL_DEPTH_8U, 3)
	zeroImg.Set(zero)
	oneImg := CreateImage(w, h, IPL_DEPTH_8U, 3)
	oneImg.Set(one)
	outImg := zeroImg.Clone()

	defer zeroImg.Release()
	defer oneImg.Release()
	defer outImg.Release()

	// generate a random mask
	maskImg := CreateImage(w, h, IPL_DEPTH_8U, 1)
	defer maskImg.Release()
	for i := 0; i < w*h; i++ {
		oneOrZero := float64(rand.Intn(2)) // random number either 1 or 0
		maskImg.Set1D(i, NewScalar(oneOrZero*255, 0, 0, 0))
	}

	// 0 = 0 & 0
	AndWithMask(zeroImg, zeroImg, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 0, "AndWithMask()")
	// 0 = 1 & 0
	AndWithMask(oneImg, zeroImg, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 0, "AndWithMask()")
	// 1 = 1 & 1
	AndWithMask(oneImg, oneImg, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 1, "AndWithMask()")

	// 0 = 0 | 0
	OrWithMask(zeroImg, zeroImg, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 0, "OrWithMask()")
	// 1 = 1 | 0
	OrWithMask(oneImg, zeroImg, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 1, "OrWithMask()")
	// 1 = 1 | 1
	OrWithMask(oneImg, oneImg, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 1, "OrWithMask()")

	// 0 = 0 ^ 0
	XorWithMask(zeroImg, zeroImg, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 0, "XOrWithMask()")
	// 1 = 1 ^ 0
	XorWithMask(oneImg, zeroImg, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 1, "XOrWithMask()")
	// 0 = 1 ^ 1
	XorWithMask(oneImg, oneImg, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 0, "XOrWithMask()")

	// 0 = 0 & 0
	AndScalarWithMask(zeroImg, zero, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 0, "AndScalarWithMask()")
	// 0 = 1 & 0
	AndScalarWithMask(oneImg, zero, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 0, "AndScalarWithMask()")
	// 1 = 1 & 1
	AndScalarWithMask(oneImg, one, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 1, "AndScalarWithMask()")

	// 0 = 0 | 0
	OrScalarWithMask(zeroImg, zero, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 0, "OrScalarWithMask()")
	// 1 = 1 | 0
	OrScalarWithMask(oneImg, zero, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 1, "OrScalarWithMask()")
	// 1 = 1 | 1
	OrScalarWithMask(oneImg, one, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 1, "OrScalarWithMask()")

	// 0 = 0 ^ 0
	XorScalarWithMask(zeroImg, zero, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 0, "XorScalarWithMask()")
	// 1 = 1 ^ 0
	XorScalarWithMask(oneImg, zero, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 1, "XorScalarWithMask()")
	// 0 = 1 ^ 1
	XorScalarWithMask(oneImg, one, outImg, maskImg)
	checkValsWMask(outImg, maskImg, 0, "XorScalarWithMask()")
}

func TestPointRadiusAngleHelpers(t *testing.T) {

	// test Point
	testPoint := func(p PointIntrfc, rad, az, inc float64) {
		if p.Radius() != rad {
			t.Errorf("%v, Radius() Fail: Expected: %f, Received: %f", p, rad, p.Radius())
		}
		if p.Angle() != az {
			t.Errorf("%v, Angle() Fail: Expected: %f, Received: %f", p, az, p.Angle())
		}
		if p.IncAngle() != inc {
			t.Errorf("%v, IncAngle() Fail: Expected: %f, Received: %f", p, inc, p.IncAngle())
		}
	}

	zeroOneTests := []PointIntrfc{
		Point{0, 1},
		Point2D32f{0, 1},
		Point2D64f{0, 1},
		Point3D32f{0, 1, 0},
		Point3D64f{0, 1, 0},
	}

	for _, zeroOne := range zeroOneTests {
		testPoint(
			zeroOne,
			1,         // radius should be 1
			math.Pi/2, // angle should be pi/2
			math.Pi/2, // inc angle should also be pi/2
		)
	}

	fiveFiveTests := []PointIntrfc{
		Point{5, 5},
		Point2D32f{5, 5},
		Point2D64f{5, 5},
		Point3D32f{5, 5, 0},
		Point3D64f{5, 5, 0},
	}
	for _, fiveFive := range fiveFiveTests {
		testPoint(
			fiveFive,
			math.Sqrt(50), // radius
			math.Pi/4,     // angle
			math.Pi/2,     // inc angle
		)
	}

	negThreeFourTests := []PointIntrfc{
		Point{-3, 4},
		Point2D32f{-3, 4},
		Point2D64f{-3, 4},
		Point3D32f{-3, 4, 0},
		Point3D64f{-3, 4, 0},
	}
	for _, negThreeFour := range negThreeFourTests {
		testPoint(
			negThreeFour,
			5, // radius
			math.Pi-math.Atan(4.0/3.0), // az angle
			math.Pi/2,                  // inc angle
		)
	}

	oneOneOneTests := []PointIntrfc{
		Point3D32f{1, 1, 1},
		Point3D64f{1, 1, 1},
	}
	for _, oneoneone := range oneOneOneTests {
		testPoint(
			oneoneone,
			math.Sqrt(3),              // radius
			math.Pi/4,                 // az angle
			math.Acos(1/math.Sqrt(3)), // inc angle
		)
	}
}

func TestPointAddSub(t *testing.T) {

	// test Point.Add
	p1 := Point{0, 1}
	p1.Add(Point{1, 1})
	if p1.X != 1 || p1.Y != 2 {
		t.Error("Unexpected result from Point.Add()")
	}

	// test Point.Sub
	p2 := Point{0, 1}
	p2.Sub(Point{1, 1})
	if p2.X != -1 || p2.Y != 0 {
		t.Error("Unexpected result from Point.Sub()")
	}

	// test Point2D32f.Add
	p3 := Point2D32f{0, 1}
	p3.Add(Point2D32f{1, 1})
	if p3.X != 1 || p3.Y != 2 {
		t.Error("Unexpected result from Point2D32f.Add()")
	}

	// test Point2D32f.Sub
	p4 := Point2D32f{0, 1}
	p4.Sub(Point2D32f{1, 1})
	if p4.X != -1 || p4.Y != 0 {
		t.Error("Unexpected result from Point2D32f.Sub()")
	}

	// test Point2D64f.Add
	p5 := Point2D64f{0, 1}
	p5.Add(Point2D64f{1, 1})
	if p5.X != 1 || p5.Y != 2 {
		t.Error("Unexpected result from Point2D64f.Add()")
	}

	// test Point2D64f.Sub
	p6 := Point2D64f{0, 1}
	p6.Sub(Point2D64f{1, 1})
	if p6.X != -1 || p6.Y != 0 {
		t.Error("Unexpected result from Point2D64f.Sub()")
	}

	// test Point3D32f.Add
	p7 := Point3D32f{0, 1, 2}
	p7.Add(Point3D32f{1, 1, 3})
	if p7.X != 1 || p7.Y != 2 || p7.Z != 5 {
		t.Error("Unexpected result from Point3D32f.Add()")
	}

	// test Point3D32f.Sub
	p8 := Point3D32f{0, 1, 2}
	p8.Sub(Point3D32f{1, 1, 3})
	if p8.X != -1 || p8.Y != 0 || p8.Z != -1 {
		t.Error("Unexpected result from Point3D32f.Sub()")
	}

	// test Point3D64f.Add
	p9 := Point3D64f{0, 1, 2}
	p9.Add(Point3D64f{1, 1, 3})
	if p9.X != 1 || p9.Y != 2 || p9.Z != 5 {
		t.Error("Unexpected result from Point3D64f.Add()")
	}

	// test Point3D64f.Sub
	p10 := Point3D64f{0, 1, 2}
	p10.Sub(Point3D64f{1, 1, 3})
	if p10.X != -1 || p10.Y != 0 || p10.Z != -1 {
		t.Error("Unexpected result from Point3D64f.Sub()")
	}
}
