// Copyright 2013 <me@cwchang.me>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

//#include "opencv.h"
import "C"
import (
	//"errors"
	//"log"
	"unsafe"
)

const (
	CV_INTER_NN       = int(C.CV_INTER_NN)
	CV_INTER_LINEAR   = int(C.CV_INTER_LINEAR)
	CV_INTER_CUBIC    = int(C.CV_INTER_CUBIC)
	CV_INTER_AREA     = int(C.CV_INTER_AREA)
	CV_INTER_LANCZOS4 = int(C.CV_INTER_LANCZOS4)
)

// For use with WarpPerspective
const (
	CV_WARP_FILL_OUTLIERS = int(C.CV_WARP_FILL_OUTLIERS)
	CV_WARP_INVERSE_MAP   = int(C.CV_WARP_INVERSE_MAP)
)

// GetPerspectiveTransform calculates a perspective transform from four pairs of the corresponding points.
//
// Parameters:
// 	src – Coordinates of quadrangle vertices in the source image.
// 	dst – Coordinates of the corresponding quadrangle vertices in the destination image.
// 	Returns the computed matrix
func GetPerspectiveTransform(rect, dst []CvPoint2D32f) *Mat {
	mat := CreateMat(3, 3, CV_64F)
	result := C.cvGetPerspectiveTransform(
		(*C.CvPoint2D32f)(&rect[0]),
		(*C.CvPoint2D32f)(&dst[0]),
		(*C.struct_CvMat)(mat))
	return (*Mat)(result)
}

// WarpPerspective applies a perspective transformation to an image.
//
// Parameters:
// 	src - input image
// 	dst – output image
// 	mapMatrix – 3x3 transformation matrix
// 	flags – combination of interpolation methods. In the C version, it is `flags=CV_INTER_LINEAR+CV_WARP_FILL_OUTLIERS` by default
// 	fillVal - In the C version, it is `fillval=(0, 0, 0, 0)` by default
func WarpPerspective(src, dst *IplImage, mapMatrix *Mat, flags int, fillVal Scalar) {
	C.cvWarpPerspective(
		unsafe.Pointer(src),
		unsafe.Pointer(dst),
		(*C.struct_CvMat)(mapMatrix),
		C.int(flags),
		(C.CvScalar)(fillVal))
}

func Resize(src *IplImage, width, height, interpolation int) *IplImage {
	if width == 0 && height == 0 {
		panic("Width and Height cannot be 0 at the same time")
	}
	if width == 0 {
		ratio := float64(height) / float64(src.Height())
		width = int(float64(src.Width()) * ratio)
	} else if height == 0 {
		ratio := float64(width) / float64(src.Width())
		height = int(float64(src.Height()) * ratio)
	}

	dst := CreateImage(width, height, src.Depth(), src.Channels())
	C.cvResize(unsafe.Pointer(src), unsafe.Pointer(dst), C.int(interpolation))
	return dst
}

func Crop(src *IplImage, x, y, width, height int) *IplImage {
	r := C.cvRect(C.int(x), C.int(y), C.int(width), C.int(height))
	rect := Rect(r)

	src.SetROI(rect)
	dest := CreateImage(width, height, src.Depth(), src.Channels())
	Copy(src, dest, nil)
	src.ResetROI()

	return dest
}

/* Returns a Seq of countours in an image, detected according to the parameters.
   Caller must Release() the Seq returned */
func (image *IplImage) FindContours(mode, method int, offset Point) *Seq {
	storage := C.cvCreateMemStorage(0)
	header_size := (C.size_t)(unsafe.Sizeof(C.CvContour{}))
	var seq *C.CvSeq
	C.cvFindContours(
		unsafe.Pointer(image),
		storage,
		&seq,
		C.int(header_size),
		C.int(mode),
		C.int(method),
		C.cvPoint(C.int(offset.X), C.int(offset.Y)))

	return (*Seq)(seq)
}

//cvDrawContours(CvArr* img, CvSeq* contour, CvScalar externalColor, CvScalar holeColor, int maxLevel, int thickness=1, int lineType=8
func DrawContours(image *IplImage, contours *Seq, externalColor, holeColor Scalar, maxLevel, thickness, lineType int, offset Point) {
	C.cvDrawContours(
		unsafe.Pointer(image),
		(*C.CvSeq)(contours),
		(C.CvScalar)(externalColor),
		(C.CvScalar)(holeColor),
		C.int(maxLevel),
		C.int(thickness),
		C.int(lineType),
		C.cvPoint(C.int(offset.X), C.int(offset.Y)))
}

// CvSeq* cvApproxPoly(const void* src_seq, int header_size, CvMemStorage* storage, int method, double eps, int recursive=0 )
func ApproxPoly(src *Seq, header_size int, storage *MemStorage, method int, eps float64, recursive int) *Seq {
	seq := C.cvApproxPoly(
		unsafe.Pointer(src),
		C.int(header_size),
		(*C.CvMemStorage)(storage),
		C.int(method),
		C.double(eps),
		C.int(recursive))
	return (*Seq)(seq)
}

// cvArcLength(const void* curve, CvSlice slice=CV_WHOLE_SEQ, int is_closed=-1 )
func ArcLength(curve *Seq, slice Slice, is_closed bool) float64 {
	is_closed_int := 0
	if is_closed {
		is_closed_int = 1
	}
	return float64(C.cvArcLength(unsafe.Pointer(curve),
		(C.CvSlice)(slice),
		C.int(is_closed_int)))
}

func ContourPerimeter(curve *Seq) float64 {
	return ArcLength(curve, WholeSeq(), true)
}

// double cvContourArea(const CvArr* contour, CvSlice slice=CV_WHOLE_SEQ, int oriented=0 )
func ContourArea(contour *Seq, slice Slice, oriented int) float64 {
	return float64(C.cvContourArea(
		unsafe.Pointer(contour),
		(C.CvSlice)(slice),
		C.int(oriented)))
}

/* points can be either CvSeq* or CvMat* */
func FitEllipse2(points unsafe.Pointer) Box2D {
	box := C.cvFitEllipse2(points)
	center := Point2D32f{float32(box.center.x), float32(box.center.y)}
	size := Size2D32f{float32(box.size.width), float32(box.size.height)}
	angle := float32(box.angle)
	return Box2D{center, size, angle}
}

// Finds a rotated rectangle of the minimum area enclosing the input 2D point set
// points can be either CvSeq* or CvMat*
func MinAreaRect(points unsafe.Pointer) Box2D {
	box := C.cvMinAreaRect2(points, nil)
	center := Point2D32f{float32(box.center.x), float32(box.center.y)}
	size := Size2D32f{float32(box.size.width), float32(box.size.height)}
	angle := float32(box.angle)
	return Box2D{center, size, angle}
}

// Calculates up-right bounding rectangle of point set
// points can be either CvSeq* or CvMat*
func BoundingRect(points unsafe.Pointer) Rect {
	return (Rect)(C.cvBoundingRect(points, C.int(0)))
}
