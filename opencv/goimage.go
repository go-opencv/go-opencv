package opencv

import "C"
import (
	"image"
	"image/color"
	"unsafe"
)

/* DecodeImageMem decodes an image from an in memory byte buffer. */
func DecodeImageMem(data []byte) *IplImage {
	buf := CreateMatHeader(1, len(data), CV_8U)
	buf.SetData(unsafe.Pointer(&data[0]), CV_AUTOSTEP)
	defer buf.Release()

	return DecodeImage(unsafe.Pointer(buf), CV_LOAD_IMAGE_UNCHANGED)
}

/* FromImage converts a go image.Image to an opencv.IplImage. */
func FromImage(img image.Image) *IplImage {
	b := img.Bounds()
	model := color.RGBAModel
	dst := CreateImage(
		b.Max.X-b.Min.X,
		b.Max.Y-b.Min.Y,
		IPL_DEPTH_8U, 4)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			px := img.At(x, y)
			c := model.Convert(px).(color.RGBA)

			value := NewScalar(float64(c.B), float64(c.G), float64(c.R), float64(c.A))
			dst.Set2D(x-b.Min.X, y-b.Min.Y, value)
		}
	}

	return dst
}

/* FromImageUnsafe create an opencv.IplImage that shares the buffer with the
go image.RGBA image. All changes made from opencv might affect go! */
func FromImageUnsafe(img *image.RGBA) *IplImage {
	b := img.Bounds()
	buf := CreateImageHeader(
		b.Max.X-b.Min.X,
		b.Max.Y-b.Min.Y,
		IPL_DEPTH_8U, 4)
	dst := CreateImage(
		b.Max.X-b.Min.X,
		b.Max.Y-b.Min.Y,
		IPL_DEPTH_8U, 4)

	buf.SetData(unsafe.Pointer(&img.Pix[0]), CV_AUTOSTEP)
	CvtColor(buf, dst, CV_RGBA2BGRA)
	buf.Release()

	return dst
}

/* ToImage converts a opencv.IplImage to an go image.Image */
func (img *IplImage) ToImage() image.Image {
	var height, width, channels, step int = img.Height(), img.Width(), img.Channels(), img.WidthStep()
	out := image.NewNRGBA(image.Rect(0, 0, width, height))
	if img.Depth() != IPL_DEPTH_8U {
		return nil // TODO return error
	}
	// Turn opencv.Iplimage.imageData(*char) to slice
	var limg *C.char = img.imageData
	var limg_ptr unsafe.Pointer = unsafe.Pointer(limg)
	var data []C.char = (*[1 << 30]C.char)(limg_ptr)[:height*step : height*step]

	c := color.NRGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(255)}
	// Iteratively assign imageData's color to Go's image
	for y := 0; y < height; y++ {
		for x := 0; x < step; x = x + channels {
			c.B = uint8(data[y*step+x])
			c.G = uint8(data[y*step+x+1])
			c.R = uint8(data[y*step+x+2])
			if channels == 4 {
				c.A = uint8(data[y*step+x+3])
			}
			out.SetNRGBA(int(x/channels), y, c)
		}
	}

	return out
}
