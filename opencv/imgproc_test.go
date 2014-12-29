package opencv

import (
	"path"
	"runtime"
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
