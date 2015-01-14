package opencv

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
	out := image.NewNRGBA(image.Rect(0, 0, img.Width(), img.Height()))
	if img.Depth() != IPL_DEPTH_8U {
		return nil // TODO return error
	}

	for y := 0; y < img.Height(); y++ {
		for x := 0; x < img.Width(); x++ {
			s := img.Get2D(x, y).Val()

			b, g, r := s[0], s[1], s[2]

			c := color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(255)}
			out.Set(x, y, c)
		}
	}

	return out
}
