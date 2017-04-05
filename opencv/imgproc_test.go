package opencv

import (
	"os"
	"path"
	"runtime"
	"syscall"
	"testing"
)

func TestGetPerspectiveTransform(t *testing.T) {
	rect := []CvPoint2D32f{
		CvPoint2D32f{265, 284},
		CvPoint2D32f{853, 284},
		CvPoint2D32f{862, 693},
		CvPoint2D32f{264, 708},
	}
	dst := []CvPoint2D32f{
		CvPoint2D32f{0, 0},
		CvPoint2D32f{597, 0},
		CvPoint2D32f{597, 423},
		CvPoint2D32f{0, 423},
	}

	res := GetPerspectiveTransform(rect, dst)

	expectedzerozero := 0.9761296668351994
	zerozero := res.Get(0, 0)
	if zerozero != expectedzerozero {
		t.Fatalf("expected result is %f, returned %f\n", expectedzerozero, zerozero)
	}

	twoone := res.Get(2, 1)
	expectedtwoone := 4.113229363293566e-05
	if twoone != expectedtwoone {
		t.Fatalf("expected result is %f, returned %f\n", expectedtwoone, twoone)
	}
}

func TestWarpPerspective(t *testing.T) {
	_, currentfile, _, _ := runtime.Caller(0)
	filename := path.Join(path.Dir(currentfile), "../images/pic5.png")

	image := LoadImage(filename)

	if image == nil {
		t.Fatal("LoadImage fail")
	}
	defer image.Release()

	warped := image.Clone()
	defer warped.Release()

	rect := []CvPoint2D32f{
		CvPoint2D32f{0, 0},
		CvPoint2D32f{400, 0},
		CvPoint2D32f{400, 300},
		CvPoint2D32f{0, 300},
	}
	dst := []CvPoint2D32f{
		CvPoint2D32f{0, 0},
		CvPoint2D32f{400, 0},
		CvPoint2D32f{200, 160},
		CvPoint2D32f{0, 60},
	}

	M := GetPerspectiveTransform(rect, dst)
	fillVal := ScalarAll(0)
	WarpPerspective(image, warped, M, CV_INTER_LINEAR+CV_WARP_FILL_OUTLIERS, fillVal)
	filename = path.Join(path.Dir(currentfile), "../images/pic5_warped.jpg")
	// Uncomment this code to create the test image "../images/pic5_warped.jpg"
	// It is part of the repo, and what this test compares against
	//
	//SaveImage(filename), warped, nil)

	tempfilename := path.Join(os.TempDir(), "pic5_warped.jpg")
	defer syscall.Unlink(tempfilename)
	SaveImage(tempfilename, warped, nil)

	// Compare actual image with expected image
	same, err := BinaryCompare(filename, tempfilename)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Error("Expected warp file != actual warp file")
	}
}

func TestResize(t *testing.T) {
	_, currentfile, _, _ := runtime.Caller(0)
	filename := path.Join(path.Dir(currentfile), "../images/lena.jpg")

	image := LoadImage(filename)
	if image == nil {
		t.Fatal("LoadImage fail")
	}
	defer image.Release()

	rimage := Resize(image, 10, 10, CV_INTER_LINEAR)
	if rimage == nil {
		t.Fatal("Resize fail")
	}
	defer rimage.Release()

	if rimage.Width() != 10 {
		t.Fatalf("excepted width is 10, returned %d\n", rimage.Width())
	}

	if rimage.Height() != 10 {
		t.Fatalf("excepted width is 10, returned %d\n", rimage.Height())
	}
}

func TestCrop(t *testing.T) {
	_, currentfile, _, _ := runtime.Caller(0)
	filename := path.Join(path.Dir(currentfile), "../images/lena.jpg")

	image := LoadImage(filename)
	if image == nil {
		t.Fatal("LoadImage fail")
	}
	defer image.Release()

	crop := Crop(image, 0, 0, 200, 200)
	if crop == nil {
		t.Fatal("Crop fail")
	}
	defer crop.Release()

	if crop.Width() != 200 {
		t.Fatalf("excepted width is 200, returned %d\n", crop.Width())
	}

	if crop.Height() != 200 {
		t.Fatalf("excepted width is 200, returned %d\n", crop.Height())
	}
}

func TestFindContours(t *testing.T) {
	_, currentfile, _, _ := runtime.Caller(0)
	filename := path.Join(path.Dir(currentfile), "../images/pic5.png")

	image := LoadImage(filename)
	if image == nil {
		t.Fatal("LoadImage fail")
	}
	defer image.Release()

	grayscale := CreateImage(image.Width(), image.Height(), IPL_DEPTH_8U, 1)
	CvtColor(image, grayscale, CV_BGR2GRAY)
	defer grayscale.Release()

	edges := CreateImage(grayscale.Width(), grayscale.Height(), grayscale.Depth(), grayscale.Channels())
	defer edges.Release()
	Canny(grayscale, edges, 50, 200, 3)

	seq := edges.FindContours(CV_RETR_EXTERNAL, CV_CHAIN_APPROX_SIMPLE, Point{0, 0})
	defer seq.Release()

	contours := CreateImage(grayscale.Width(), grayscale.Height(), grayscale.Depth(), grayscale.Channels())
	white := NewScalar(255, 255, 255, 0)
	contours.Set(white)

	black := NewScalar(0, 0, 0, 0)
	red := NewScalar(0, 255, 0, 0)

	for ; seq != nil; seq = seq.HNext() {
		DrawContours(contours, seq, red, black, 0, 2, 8, Point{0, 0})
	}

	filename = path.Join(path.Dir(currentfile), "../images/pic5_contours.png")
	// Uncomment this code to create the test image "../images/shapes_contours.png"
	// It is part of the repo, and what this test compares against
	//
	//SaveImage(filename), contours, nil)

	tempfilename := path.Join(os.TempDir(), "pic5_contours.png")
	defer syscall.Unlink(tempfilename)
	SaveImage(tempfilename, contours, nil)

	// Compare actual image with expected image
	same, err := BinaryCompare(filename, tempfilename)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Error("Expected contour file != actual contour file")
	}

}
