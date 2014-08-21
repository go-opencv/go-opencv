package opencv

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"reflect"
	"testing"
)

func TestFromImage(t *testing.T) {
	fh, err := os.Open("../images/pic5.png")
	if err != nil {
		t.Fatal(err)
	}

	img, _, err := image.Decode(fh)
	if err != nil {
		t.Fatal(err)
	}

	ocv := FromImage(img)
	if ocv == nil {
		t.Fatal("failed to convert image")
	}

	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	if ocv.Width() != width || ocv.Height() != height {
		t.Fatalf("loaded image has wrong dimensions: got %dx%d, expected %dx%d", ocv.Width(), ocv.Height(), width, height)
	}

	ex := [4]float64{98, 139, 26, 255}
	px := ocv.Get2D(50, 90).Val()
	if !reflect.DeepEqual(px, ex) {
		t.Fatalf("wrong color @50,90: got %v, expected %v", px, ex)
	}
}

func TestFromImageUnsafe(t *testing.T) {
	fh, err := os.Open("../images/pic5.png")
	if err != nil {
		t.Fatal(err)
	}

	img, _, err := image.Decode(fh)
	if err != nil {
		t.Fatal(err)
	}

	rgba, ok := img.(*image.RGBA)
	if !ok {
		t.Fatal("image is no RGBA image")
	}

	ocv := FromImageUnsafe(rgba)
	if ocv == nil {
		t.Fatal("failed to convert image")
	}

	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	if ocv.Width() != width || ocv.Height() != height {
		t.Fatalf("loaded image has wrong dimensions: got %dx%d, expected %dx%d", ocv.Width(), ocv.Height(), width, height)
	}

	ex := [4]float64{98, 139, 26, 255}
	px := ocv.Get2D(50, 90).Val()
	if !reflect.DeepEqual(px, ex) {
		t.Fatalf("wrong color @50,90: got %v, expected %v", px, ex)
	}
}

func BenchmarkFromImage(b *testing.B) {
	fh, _ := os.Open("../images/pic5.png")
	img, _, _ := image.Decode(fh)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ocv := FromImage(img)
		ocv.Release()
	}
}

func BenchmarkFromImageUnsafe(b *testing.B) {
	fh, _ := os.Open("../images/pic5.png")
	img, _, _ := image.Decode(fh)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rgba := img.(*image.RGBA)
		ocv := FromImageUnsafe(rgba)
		ocv.Release()
	}
}
