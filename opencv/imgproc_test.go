package opencv

import (
	"os"
	"path"
	"runtime"
	"syscall"
	"testing"
)

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
	//SaveImage(filename), contours, 0)

	tempfilename := path.Join(os.TempDir(), "pic5_contours.png")
	defer syscall.Unlink(tempfilename)
	SaveImage(tempfilename, contours, 0)

	// Compare actual image with expected image
	same, err := BinaryCompare(filename, tempfilename)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Error("Expected contour file != actual contour file")
	}

}
