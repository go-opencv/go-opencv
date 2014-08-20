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

	return DecodeImage(buf, CV_LOAD_IMAGE_UNCHANGED)
}

/* From Image converts a go image.Image to an opencv.IplImage. */
func FromImage(img image.Image) *IplImage {
	b := img.Bounds()
	model := color.NRGBAModel
	dst := CreateImage(
		b.Max.Y-b.Min.Y,
		b.Max.X-b.Min.X,
		IPL_DEPTH_8U, 4)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			px := img.At(x, y)
			c := model.Convert(px).(color.NRGBA)

			value := NewScalar(float64(c.B), float64(c.G), float64(c.R), float64(c.A))
			dst.Set2D(x, y, value)
		}
	}

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
			b, g, r, a := s[2], s[1], s[0], s[3]

			c := color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			out.Set(x, y, c)
		}
	}

	return out
}
