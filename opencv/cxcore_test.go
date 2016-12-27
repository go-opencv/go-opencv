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
	"unsafe"
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
	filename := path.Join(path.Dir(currentfile), "../images/pic3.png")

	image := LoadImage(filename)
	if image == nil {
		t.Fatal("LoadImage fail")
	}
	defer image.Release()

	// Write 'Hello' on the image
	font := InitFont(CV_FONT_HERSHEY_DUPLEX, 1, 1, 0, 1, 8)
	color := NewScalar(0, 0, 0, 0)

	pos := Point{image.Width() / 2, image.Height() / 2}
	font.PutText(image, "Hello", pos, color)

	filename = path.Join(path.Dir(currentfile), "../images/pic3_with_text.png")

	// Uncomment this code to create the test image "../images/pic3_with_text.jpg"
	// It is part of the repo, and what this test compares against
	//
	// SaveImage(filename, image, 0)
	// println("Saved file", filename)

	tempfilename := path.Join(os.TempDir(), "pic3_with_text.png")
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

func checkValsWMask(t *testing.T, img, mask *IplImage, val float64, debug string) {
	for i := 0; i < img.Width()*img.Height(); i++ {
		pix := img.Get1D(i).Val()
		tst := val
		if mask != nil && mask.Get1D(i).Val()[0] == 0 {
			tst = 0
		}
		if pix[0] != tst || pix[1] != tst || pix[2] != tst {
			t.Errorf("Unexpeted value for %s: %f, %f, %f. Expected %fs",
				debug, pix[0], pix[1], pix[2], val)
			break
		}
	}
}

func TestAddSub(t *testing.T) {
	w, h := 50, 50

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
	checkValsWMask(t, hundredImg, nil, 100, "AddScalar()")

	// 100 + 100 = 200
	Add(hundredImg, hundredImg, twoHundredImg)
	checkValsWMask(t, twoHundredImg, nil, 200, "Add()")

	// 200 - 100 = 100
	Subtract(twoHundredImg, hundredImg, hundredImg)
	checkValsWMask(t, hundredImg, nil, 100, "Sub()")

	// 100 - 100 = 0
	SubScalar(hundredImg, hundred, zeroImg)
	checkValsWMask(t, zeroImg, nil, 0, "SubScalar()")

	// 100 - 200 = 0 != -100 because it clips
	SubScalarRev(hundred, twoHundredImg, negImage)
	checkValsWMask(t, negImage, nil, 0, "SubScalarRev()")

	// Uncomment to save these images to disk
	// SaveImage("zeroImg.png", zeroImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("hundredImg.png", hundredImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("twoHundredImg.png", twoHundredImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("negImage.png", negImage, CV_IMWRITE_PNG_COMPRESSION)

}

func TestAddSubWithMask(t *testing.T) {
	w, h := 50, 50
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
	checkValsWMask(t, hundredImg, maskImg, 100, "AddScalarWithMask()")

	// 100 + 100 = 200
	AddWithMask(hundredImg, hundredImg, twoHundredImg, maskImg)
	checkValsWMask(t, twoHundredImg, maskImg, 200, "AddWithMask()")

	// 200 - 100 = 100
	SubtractWithMask(twoHundredImg, hundredImg, hundredImg, maskImg)
	checkValsWMask(t, hundredImg, maskImg, 100, "SubtractWithMask()")

	// 100 - 100 = 0
	SubScalarWithMask(hundredImg, hundred, zeroImg, maskImg)
	checkValsWMask(t, zeroImg, maskImg, 0, "SubScalarWithMask()")

	// 100 - 200 = 0 != -100 because it clips
	SubScalarWithMaskRev(hundred, twoHundredImg, negImage, maskImg)
	checkValsWMask(t, negImage, maskImg, 0, "SubScalarWithMaskRev()")

	// Uncomment to save these images to disk
	// SaveImage("zeroImgMask.png", zeroImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("hundredImgMask.png", hundredImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("twoHundredImgMask.png", twoHundredImg, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("negImageMask.png", negImage, CV_IMWRITE_PNG_COMPRESSION)
	// SaveImage("MaskImg.png", maskImg, CV_IMWRITE_PNG_COMPRESSION)

}

func TestLogicAndAbsDiff(t *testing.T) {
	w, h := 50, 50

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
	checkValsWMask(t, outImg, maskImg, 0, "AndWithMask()")
	// 0 = 1 & 0
	AndWithMask(oneImg, zeroImg, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 0, "AndWithMask()")
	// 1 = 1 & 1
	AndWithMask(oneImg, oneImg, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 1, "AndWithMask()")

	// 0 = 0 | 0
	OrWithMask(zeroImg, zeroImg, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 0, "OrWithMask()")
	// 1 = 1 | 0
	OrWithMask(oneImg, zeroImg, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 1, "OrWithMask()")
	// 1 = 1 | 1
	OrWithMask(oneImg, oneImg, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 1, "OrWithMask()")

	// 0 = 0 ^ 0
	XorWithMask(zeroImg, zeroImg, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 0, "XOrWithMask()")
	// 1 = 1 ^ 0
	XorWithMask(oneImg, zeroImg, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 1, "XOrWithMask()")
	// 0 = 1 ^ 1
	XorWithMask(oneImg, oneImg, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 0, "XOrWithMask()")

	// 0 = 0 & 0
	AndScalarWithMask(zeroImg, zero, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 0, "AndScalarWithMask()")
	// 0 = 1 & 0
	AndScalarWithMask(oneImg, zero, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 0, "AndScalarWithMask()")
	// 1 = 1 & 1
	AndScalarWithMask(oneImg, one, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 1, "AndScalarWithMask()")

	// 0 = 0 | 0
	OrScalarWithMask(zeroImg, zero, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 0, "OrScalarWithMask()")
	// 1 = 1 | 0
	OrScalarWithMask(oneImg, zero, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 1, "OrScalarWithMask()")
	// 1 = 1 | 1
	OrScalarWithMask(oneImg, one, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 1, "OrScalarWithMask()")

	// 0 = 0 ^ 0
	XorScalarWithMask(zeroImg, zero, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 0, "XorScalarWithMask()")
	// 1 = 1 ^ 0
	XorScalarWithMask(oneImg, zero, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 1, "XorScalarWithMask()")
	// 0 = 1 ^ 1
	XorScalarWithMask(oneImg, one, outImg, maskImg)
	checkValsWMask(t, outImg, maskImg, 0, "XorScalarWithMask()")

	// 1 = |1-0|
	AbsDiff(oneImg, zeroImg, outImg)
	checkValsWMask(t, outImg, nil, 1, "AbsDiff")

	// 1 = |0-1|
	AbsDiff(zeroImg, oneImg, outImg)
	checkValsWMask(t, outImg, nil, 1, "AbsDiff")

	// 1 = |1-0|
	AbsDiffScalar(oneImg, zero, outImg)
	checkValsWMask(t, outImg, nil, 1, "AbsDiffScalar")

	// 1 = |0-1|
	AbsDiffScalar(zeroImg, one, outImg)
	checkValsWMask(t, outImg, nil, 1, "AbsDiffScalar")

}

func TestPointRadiusAngleHelpers(t *testing.T) {

	// test Point
	test2D := func(p Point2D, rad, az float64) {
		if p.Radius() != rad {
			t.Errorf("%v, Radius() Fail: Expected: %f, Received: %f", p, rad, p.Radius())
		}
		if p.Angle() != az {
			t.Errorf("%v, Angle() Fail: Expected: %f, Received: %f", p, az, p.Angle())
		}
	}

	test3D := func(p Point3D, rad, az, inc float64) {
		if p.Radius() != rad {
			t.Errorf("%v, Radius() Fail: Expected: %f, Received: %f", p, rad, p.Radius())
		}
		if p.AzAngle() != az {
			t.Errorf("%v, Angle() Fail: Expected: %f, Received: %f", p, az, p.AzAngle())
		}
		if p.IncAngle() != inc {
			t.Errorf("%v, IncAngle() Fail: Expected: %f, Received: %f", p, inc, p.IncAngle())
		}
	}

	// Test Point: {0,1,0}
	zeroOne2DTests := []Point2D{
		Point{0, 1},
		Point2D32f{0, 1},
		Point2D64f{0, 1},
	}
	zeroOne3DTests := []Point3D{
		Point3D32f{0, 1, 0},
		Point3D64f{0, 1, 0},
	}
	for _, zeroOne := range zeroOne2DTests {
		test2D(zeroOne, 1, math.Pi/2)
	}
	for _, zeroOne := range zeroOne3DTests {
		test3D(zeroOne, 1, math.Pi/2, math.Pi/2)
	}

	// Test Point: {5,5,0}
	fiveFive2DTests := []Point2D{
		Point{5, 5},
		Point2D32f{5, 5},
		Point2D64f{5, 5},
	}
	fiveFive3DTests := []Point3D{
		Point3D32f{5, 5, 0},
		Point3D64f{5, 5, 0},
	}
	for _, fiveFive := range fiveFive2DTests {
		test2D(fiveFive, math.Sqrt(50), math.Pi/4)
	}
	for _, fiveFive := range fiveFive3DTests {
		test3D(fiveFive, math.Sqrt(50), math.Pi/4, math.Pi/2)
	}

	// Test Point: {-3,4,0}
	negThreeFour2DTests := []Point2D{
		Point{-3, 4},
		Point2D32f{-3, 4},
		Point2D64f{-3, 4},
	}
	negThreeFour3DTests := []Point3D{
		Point3D32f{-3, 4, 0},
		Point3D64f{-3, 4, 0},
	}
	for _, negThreeFour := range negThreeFour2DTests {
		test2D(negThreeFour, 5, math.Pi-math.Atan(4.0/3.0))
	}
	for _, negThreeFour := range negThreeFour3DTests {
		test3D(negThreeFour, 5, math.Pi-math.Atan(4.0/3.0), math.Pi/2)
	}

	// Test Point: {1,1,1}
	oneOneOneTests := []Point3D{
		Point3D32f{1, 1, 1},
		Point3D64f{1, 1, 1},
	}
	for _, oneoneone := range oneOneOneTests {
		test3D(oneoneone, math.Sqrt(3), math.Pi/4, math.Acos(1/math.Sqrt(3)))
	}
}

func TestPointAddSub(t *testing.T) {

	// test Point.Add
	p1 := Point{0, 1}
	out1 := p1.Add(Point{1, 1})
	if out1.X != 1 || out1.Y != 2 {
		t.Error("Unexpected result from Point.Add()")
	}

	// test Point.Sub
	p2 := Point{0, 1}
	out2 := p2.Sub(Point{1, 1})
	if out2.X != -1 || out2.Y != 0 {
		t.Error("Unexpected result from Point.Sub()")
	}

	// test Point2D32f.Add
	p3 := Point2D32f{0, 1}
	out3 := p3.Add(Point2D32f{1, 1})
	if out3.X != 1 || out3.Y != 2 {
		t.Error("Unexpected result from Point2D32f.Add()")
	}

	// test Point2D32f.Sub
	p4 := Point2D32f{0, 1}
	out4 := p4.Sub(Point2D32f{1, 1})
	if out4.X != -1 || out4.Y != 0 {
		t.Error("Unexpected result from Point2D32f.Sub()")
	}

	// test Point2D64f.Add
	p5 := Point2D64f{0, 1}
	out5 := p5.Add(Point2D64f{1, 1})
	if out5.X != 1 || out5.Y != 2 {
		t.Error("Unexpected result from Point2D64f.Add()")
	}

	// test Point2D64f.Sub
	p6 := Point2D64f{0, 1}
	out6 := p6.Sub(Point2D64f{1, 1})
	if out6.X != -1 || out6.Y != 0 {
		t.Error("Unexpected result from Point2D64f.Sub()")
	}

	// test Point3D32f.Add
	p7 := Point3D32f{0, 1, 2}
	out7 := p7.Add(Point3D32f{1, 1, 3})
	if out7.X != 1 || out7.Y != 2 || out7.Z != 5 {
		t.Error("Unexpected result from Point3D32f.Add()")
	}

	// test Point3D32f.Sub
	p8 := Point3D32f{0, 1, 2}
	out8 := p8.Sub(Point3D32f{1, 1, 3})
	if out8.X != -1 || out8.Y != 0 || out8.Z != -1 {
		t.Error("Unexpected result from Point3D32f.Sub()")
	}

	// test Point3D64f.Add
	p9 := Point3D64f{0, 1, 2}
	out9 := p9.Add(Point3D64f{1, 1, 3})
	if out9.X != 1 || out9.Y != 2 || out9.Z != 5 {
		t.Error("Unexpected result from Point3D64f.Add()")
	}

	// test Point3D64f.Sub
	p10 := Point3D64f{0, 1, 2}
	out10 := p10.Sub(Point3D64f{1, 1, 3})
	if out10.X != -1 || out10.Y != 0 || out10.Z != -1 {
		t.Error("Unexpected result from Point3D64f.Sub()")
	}
}

func TestSeq(t *testing.T) {
	// tests the following functions and methods:
	//		NewCvPoint()
	//		CreateSeq()
	//		Seq.Push()
	//		Seq.PushFront()
	//		Seq.Pop()
	//		Seq.PopFront()
	//		Seq.Total()
	//		Seq.GetElemAt()
	//		Seq.RemoveAt()

	// create an empty sequence of points
	seqOfPoints := CreateSeq(
		CV_SEQ_ELTYPE_POINT,
		int(unsafe.Sizeof(CvPoint{})),
	)
	defer seqOfPoints.Release()

	// push some points onto the sequence.
	zerozero := NewCvPoint(0, 0)
	zeroone := NewCvPoint(0, 1)
	onezero := NewCvPoint(1, 0)
	oneone := NewCvPoint(1, 1)
	seqOfPoints.Push(unsafe.Pointer(&zerozero))
	seqOfPoints.Push(unsafe.Pointer(&zeroone))
	seqOfPoints.Push(unsafe.Pointer(&onezero))     // this will be the last element
	seqOfPoints.PushFront(unsafe.Pointer(&oneone)) // this will be the first element

	// the sequence total should be 4
	if seqOfPoints.Total() != 4 {
		t.Error("seq Total() should be 4!")
	}

	testPtVal := func(result, expected CvPoint, debug string) {
		if result.x != expected.x || result.y != expected.y {
			t.Errorf("%s: Result: %v.  Expected: %v", debug, result, expected)
		}
	}

	// test the access to points:
	elem := (*CvPoint)(seqOfPoints.GetElemAt(0))
	testPtVal(*elem, oneone, "GetElemAt(0)")
	elem = (*CvPoint)(seqOfPoints.GetElemAt(1))
	testPtVal(*elem, zerozero, "GetElemAt(1)")
	elem = (*CvPoint)(seqOfPoints.GetElemAt(2))
	testPtVal(*elem, zeroone, "GetElemAt(2)")
	elem = (*CvPoint)(seqOfPoints.GetElemAt(3))
	testPtVal(*elem, onezero, "GetElemAt(3)")

	// pop some points off from the front and back of the sequence
	var firstPt, lastPt CvPoint
	seqOfPoints.PopFront(unsafe.Pointer(&firstPt))
	seqOfPoints.Pop(unsafe.Pointer(&lastPt))

	// check the values of the popped points
	testPtVal(firstPt, oneone, "PopFront()")
	testPtVal(lastPt, onezero, "Pop()")

	// the sequence total should only be two now.
	if seqOfPoints.Total() != 2 {
		t.Error("seq Total() should be 2!")
	}

	// push the last point again...
	seqOfPoints.Push(unsafe.Pointer(&lastPt))

	// remove the middle element
	seqOfPoints.RemoveAt(1)

	// the sequence total should be back at 2
	if seqOfPoints.Total() != 2 {
		t.Error("seq Total() should be 2!")
	}

	// verify that the middle was removed by
	// looking at the two values that are left:
	elem = (*CvPoint)(seqOfPoints.GetElemAt(0))
	testPtVal(*elem, zerozero, "GetElemAt(0)")
	elem = (*CvPoint)(seqOfPoints.GetElemAt(1))
	testPtVal(*elem, onezero, "GetElemAt(1)")

}

func TestBoundingRectInt(t *testing.T) {

	nTests := 20        // number of tests to run
	nPointsInSeq := 15  // number of points in each tests
	maxXYAllowed := 100 // the max X and Y value of each random point generated

	// create an empty sequence of points
	seqOfPoints := CreateSeq(
		CV_SEQ_ELTYPE_POINT,
		int(unsafe.Sizeof(CvPoint{})),
	)
	defer seqOfPoints.Release()

	for iTest := 0; iTest < nTests; iTest++ {

		seqOfPoints.Clear()

		minP := Point{maxXYAllowed + 1, maxXYAllowed + 1}
		maxP := Point{-1, -1}

		// fill the sequence with random points
		for i := 0; i < nPointsInSeq; i++ {
			x, y := rand.Intn(maxXYAllowed), rand.Intn(maxXYAllowed)
			cvP := NewCvPoint(x, y)
			seqOfPoints.Push(unsafe.Pointer(&cvP))
			// keep track of minimum and maximum values
			if x < minP.X {
				minP.X = x
			}
			if y < minP.Y {
				minP.Y = y
			}
			if x > maxP.X {
				maxP.X = x
			}
			if y > maxP.Y {
				maxP.Y = y
			}
		}

		// get the bounding rectangle:
		b := BoundingRect(unsafe.Pointer(seqOfPoints))

		// calculate expected rectangle:
		e := NewRect(minP.X, minP.Y, maxP.X-minP.X+1, maxP.Y-minP.Y+1)

		// verify rectangle:
		if b.X() != e.X() || b.Y() != e.Y() ||
			b.Width() != e.Width() || b.Height() != e.Height() {
			t.Errorf("BoundingRect (%d): %v doesn't match expected rect: %v", iTest, b, e)
		}
	}
}

func TestBoundingRectFloat32(t *testing.T) {

	nTests := 20        // number of tests to run
	nPointsInSeq := 15  // number of points in each test
	maxXYAllowed := 100 // the max X and y value of each random point generated

	// create an empty sequence of CvPoint2D32f's
	seqOfPoints := CreateSeq(
		CV_32FC2,
		int(unsafe.Sizeof(CvPoint2D32f{})),
	)
	defer seqOfPoints.Release()

	for iTest := 0; iTest < nTests; iTest++ {

		seqOfPoints.Clear()

		minP := Point2D32f{float32(maxXYAllowed + 1), float32(maxXYAllowed + 1)}
		maxP := Point2D32f{-1, -1}

		// fill the sequence with random points, keep track of minimum and maximum values
		for i := 0; i < nPointsInSeq; i++ {
			x := rand.Float32() * float32(maxXYAllowed)
			y := rand.Float32() * float32(maxXYAllowed)

			cvP := NewCvPoint2D32f(x, y)
			seqOfPoints.Push(unsafe.Pointer(&cvP))

			if x < minP.X {
				minP.X = x
			}
			if y < minP.Y {
				minP.Y = y
			}
			if x > maxP.X {
				maxP.X = x
			}
			if y > maxP.Y {
				maxP.Y = y
			}
		}

		// get the bounding rectangle:
		b := BoundingRect(unsafe.Pointer(seqOfPoints))

		// calculate expected rectangle:
		e := NewRect(int(minP.X), int(minP.Y), int(maxP.X)-int(minP.X)+1, int(maxP.Y)-int(minP.Y)+1)

		// verify rectangle:
		if b.X() != e.X() || b.Y() != e.Y() ||
			b.Width() != e.Width() || b.Height() != e.Height() {
			t.Errorf("BoundingRectFloat32 (%d): %v doesn't match expected rect: %v", iTest, b, e)
		}
	}
}

func TestMeanStdDev(t *testing.T) {
	_, currentfile, _, _ := runtime.Caller(0)
	filename := path.Join(path.Dir(currentfile), "../images/lena.jpg")
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}

	image := LoadImage(filename)
	if image == nil {
		t.Fatal("LoadImage fail")
	}
	defer image.Release()

	mean, stdDev := image.MeanStdDev()

	width := image.Width()
	height := image.Height()

	// Calculated mean and standard deviation
	for i := 0; i < 3; i++ {

		total := 0.0
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				total += image.Get2D(x, y).Val()[i]
			}
		}

		average := total / float64(width*height)

		total = 0.0
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				deviation := image.Get2D(x, y).Val()[i] - average
				total += math.Pow(deviation, 2)
			}
		}

		variance := total / float64(width*height)
		stdDevCalculated := math.Sqrt(variance)

		stdDevPass := stdDevCalculated-stdDev.Val()[i] < 0.0000000001
		meanPass := average-mean.Val()[i] < 0.0000000001
		if !stdDevPass {
			t.Error("Standard deviation calculation fail")
		}
		if !meanPass {
			t.Error("Mean calculation fail")
		}
	}

}
