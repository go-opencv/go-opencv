// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

//#include "opencv.h"
//#cgo linux  pkg-config: opencv
//#cgo darwin pkg-config: opencv
//#cgo freebsd pkg-config: opencv
//#cgo windows,pkgconfig pkg-config: opencv
//#cgo windows,!pkgconfig LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
import "C"
import (
	//"errors"
	"unsafe"
)

func init() {
}

const (
	CV_BGR2GRAY  = C.CV_BGR2GRAY
	CV_RGB2GRAY  = C.CV_RGB2GRAY
	CV_GRAY2BGR  = C.CV_GRAY2BGR
	CV_GRAY2RGB  = C.CV_GRAY2RGB
	CV_BGR2BGRA  = C.CV_BGR2BGRA
	CV_RGBA2BGRA = C.CV_RGBA2BGRA

	CV_BLUR_NO_SCALE = C.CV_BLUR_NO_SCALE
	CV_BLUR          = C.CV_BLUR
	CV_GAUSSIAN      = C.CV_GAUSSIAN
	CV_MEDIAN        = C.CV_MEDIAN
	CV_BILATERAL     = C.CV_BILATERAL

	CV_8U  = C.CV_8U
	CV_8S  = C.CV_8S
	CV_16U = C.CV_16U
	CV_16S = C.CV_16S
	CV_32S = C.CV_32S
	CV_32F = C.CV_32F
	CV_64F = C.CV_64F
)

/* Smoothes array (removes noise) */
func Smooth(src, dst *IplImage, smoothtype,
	param1, param2 int, param3, param4 float64) {
	C.cvSmooth(unsafe.Pointer(src), unsafe.Pointer(dst), C.int(smoothtype),
		C.int(param1), C.int(param2), C.double(param3), C.double(param4),
	)
}

//CVAPI(void) cvSmooth( const CvArr* src, CvArr* dst,
//                      int smoothtype CV_DEFAULT(CV_GAUSSIAN),
//                      int param1 CV_DEFAULT(3),
//                      int param2 CV_DEFAULT(0),
//                      double param3 CV_DEFAULT(0),
//                      double param4 CV_DEFAULT(0));

/* Laplace The function calculates the Laplacian */
func Laplace(src, dst *IplImage, aperture_size int) {
	C.cvLaplace(unsafe.Pointer(src), unsafe.Pointer(dst), C.int(aperture_size))
}

//CVAPI(void) cvLaplace( const CvArr* src,
//                       CvArr* dst,
//                       int aperture_size CV_DEFAULT(3) );

/*
ConvertScale converts one image to another with optional linear transformation.
*/
func ConvertScale(a, b *IplImage, scale, shift float64) {
	C.cvConvertScale(unsafe.Pointer(a), unsafe.Pointer(b), C.double(scale), C.double(shift))
}

//CVAPI(void)  cvConvertScale( const CvArr* src,
//                             CvArr* dst,
//                             double scale CV_DEFAULT(1),
//                             double shift CV_DEFAULT(0) );

/* Converts input array pixels from one color space to another */
func CvtColor(src, dst *IplImage, code int) {
	C.cvCvtColor(unsafe.Pointer(src), unsafe.Pointer(dst), C.int(code))
}

//CVAPI(void)  cvCvtColor( const CvArr* src, CvArr* dst, int code );

/* Runs canny edge detector */
func Canny(image, edges *IplImage, threshold1, threshold2 float64, aperture_size int) {
	C.cvCanny(unsafe.Pointer(image), unsafe.Pointer(edges),
		C.double(threshold1), C.double(threshold2),
		C.int(aperture_size),
	)
}

//CVAPI(void)  cvCanny( const CvArr* image, CvArr* edges, double threshold1,
//                      double threshold2, int  aperture_size CV_DEFAULT(3) );

/* Calculates the first, second, third, or mixed image derivatives using an
* extended Sobel operator.  */
func Sobel(src, dst *IplImage, xorder, yorder, aperture_size int) {
	C.cvSobel(unsafe.Pointer(src), unsafe.Pointer(dst),
		C.int(xorder), C.int(yorder),
		C.int(aperture_size),
	)
}

// C: void cvSobel(const CvArr* src, CvArr* dst, int xorder, int yorder, int aperture_size=3 )

const (
	CV_INPAINT_NS    = C.CV_INPAINT_NS
	CV_INPAINT_TELEA = C.CV_INPAINT_TELEA
)

/* Inpaints the selected region in the image */
func Inpaint(src, inpaint_mask, dst *IplImage, inpaintRange float64, flags int) {
	C.cvInpaint(
		unsafe.Pointer(src),
		unsafe.Pointer(inpaint_mask),
		unsafe.Pointer(dst),
		C.double(inpaintRange),
		C.int(flags),
	)
}

//CVAPI(void) cvInpaint( const CvArr* src, const CvArr* inpaint_mask,
//                       CvArr* dst, double inpaintRange, int flags );

const (
	CV_THRESH_BINARY     = C.CV_THRESH_BINARY
	CV_THRESH_BINARY_INV = C.CV_THRESH_BINARY_INV
	CV_THRESH_TRUNC      = C.CV_THRESH_TRUNC
	CV_THRESH_TOZERO     = C.CV_THRESH_TOZERO
	CV_THRESH_TOZERO_INV = C.CV_THRESH_TOZERO_INV
)

/* Applies a fixed-level threshold to each array element. */
func Threshold(src, dst *IplImage, threshold, max_value float64, threshold_type int) {
	C.cvThreshold(
		unsafe.Pointer(src),
		unsafe.Pointer(dst),
		C.double(threshold),
		C.double(max_value),
		C.int(threshold_type),
	)
}

//CVAPI(double) cvThreshold( const CvArr* src, CvArr* dst, double threshold,
//                           double max_value, int threshold_type );

const (
	CV_ADAPTIVE_THRESH_MEAN_C     = C.CV_ADAPTIVE_THRESH_MEAN_C
	CV_ADAPTIVE_THRESH_GAUSSIAN_C = C.CV_ADAPTIVE_THRESH_GAUSSIAN_C
)

/* Applies an adaptive threshold to an array. */
func AdaptiveThreshold(src, dst *IplImage, max_value float64, adaptive_method,
	threshold_type, block_size int, thresh_C float64) {
	C.cvAdaptiveThreshold(
		unsafe.Pointer(src),
		unsafe.Pointer(dst),
		C.double(max_value),
		C.int(adaptive_method),
		C.int(threshold_type),
		C.int(block_size),
		C.double(thresh_C),
	)
}

//CVAPI(void) cvAdaptiveThreshold( const CvArr* src, CvArr* dst, double max_value,
//                                 int adaptive_method=CV_ADAPTIVE_THRESH_MEAN_C,
//                                 int threshold_type=CV_THRESH_BINARY,
//                                 int block_size=3, double param1=5 );

const (
	CV_MORPH_RECT    = 0
	CV_MORPH_ELLIPSE = 1
	CV_MORPH_CROSS   = 2
	//CV_SHAPE_CUSTOM  = 3 // TODO: currently we don't support a fully custom kernel
)

/* Returns a structuring element of the specified size and shape for morphological operations. */
func CreateStructuringElement(cols, rows, anchor_x, anchor_y, shape int) *IplConvKernel {
	return (*IplConvKernel)(C.cvCreateStructuringElementEx(
		C.int(cols),
		C.int(rows),
		C.int(anchor_x),
		C.int(anchor_y),
		C.int(shape),
		nil, // TODO: currently we don't support a fully custom kernel
	))
}

//CVAPI(IplConvKernel*) cvCreateStructuringElementEx( int cols, int rows, int anchor_x,
//                                                    int anchor_y, int shape, int* values=NULL )

/* Releases the structuring element */
func (k *IplConvKernel) ReleaseElement() {
	C.cvReleaseStructuringElement(
		(**C.IplConvKernel)(unsafe.Pointer(&k)),
	)
}

//CVAPI(void) cvReleaseStructuringElement( IplConvKernel** element );

/* Dilates an image by using a specific structuring element. */
func Dilate(src, dst *IplImage, element *IplConvKernel, iterations int) {
	C.cvDilate(
		unsafe.Pointer(src),
		unsafe.Pointer(dst),
		(*C.IplConvKernel)(unsafe.Pointer(element)),
		C.int(iterations),
	)
}

//CVAPI(void) cvDilate( const CvArr* src, CvArr* dst, IplConvKernel* element=NULL,
//                      int iterations=1 );

/* Erodes an image by using a specific structuring element. */
func Erode(src, dst *IplImage, element *IplConvKernel, iterations int) {
	C.cvErode(
		unsafe.Pointer(src),
		unsafe.Pointer(dst),
		(*C.IplConvKernel)(unsafe.Pointer(element)),
		C.int(iterations),
	)
}

//CVAPI(void) cvErode( const CvArr* src, CvArr* dst, IplConvKernel* element=NULL,
//                     int iterations=1 );

const (
	CV_MORPH_OPEN     = C.CV_MOP_OPEN
	CV_MORPH_CLOSE    = C.CV_MOP_CLOSE
	CV_MORPH_GRADIENT = C.CV_MOP_GRADIENT
	CV_MORPH_TOPHAT   = C.CV_MOP_TOPHAT
	CV_MORPH_BLACKHAT = C.CV_MOP_BLACKHAT
)

/* Performs advanced morphological transformations. */
func MorphologyEx(src, dst, temp *IplImage, element *IplConvKernel, operation int, iterations int) {
	C.cvMorphologyEx(
		unsafe.Pointer(src),
		unsafe.Pointer(dst),
		unsafe.Pointer(temp),
		(*C.IplConvKernel)(unsafe.Pointer(element)),
		C.int(operation),
		C.int(iterations),
	)
}

//CVAPI(void) cvMorphologyEx( const CvArr* src, CvArr* dst, CvArr* temp,
//                            IplConvKernel* element, int operation, int iterations=1 );
