package opencv

import (
	"errors"
	// "fmt"
	"path"
	"runtime"
	"testing"
)

func GetFocusFromFile(file string) (err error, focus float64) {
	src := LoadImage(file)

	if src == nil {
		return errors.New("LoadImage fail"), focus
	}

	dst := CreateImage(src.Width(), src.Height(), src.Depth(), src.Channels())

	Laplace(src, dst, 3)

	_, sigma := MeanStdDevWithMask(dst, nil)

	focus = sigma.Val()[0] * sigma.Val()[0]
	return err, focus
}

func TestBlurredImgs(t *testing.T) {

	// case different blur on image
	{
		_, currentfile, _, _ := runtime.Caller(0)

		fileBlurryFalse := path.Join(path.Dir(currentfile), "../images/blurry-false.png")
		err, focusBlurryFalse := GetFocusFromFile(fileBlurryFalse)
		if err != nil {
			t.Error("err get focus for blurry-false: ", err)
		}

		fileBlurryTrue := path.Join(path.Dir(currentfile), "../images/blurry-true.png")
		err, focusBlurryTrue := GetFocusFromFile(fileBlurryTrue)
		if err != nil {
			t.Error("err get focus for blurry-true: ", err)
		}

		if focusBlurryTrue > focusBlurryFalse {
			t.Error("err value focuses, blurry-true should be lesser blurry-false: ", err)
		}
		// fmt.Println("blyrry-false: ", focusBlurryFalse)
		// fmt.Println("blyrry-true: ", focusBlurryTrue)

	}
}
