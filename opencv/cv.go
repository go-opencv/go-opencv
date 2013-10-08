// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

//#include "opencv.h"
//#cgo linux  pkg-config: opencv
//#cgo darwin pkg-config: opencv
//#cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
import "C"
import (
	//"errors"
	"unsafe"
)

func init() {
}

const (
	CV_BGR2GRAY = C.CV_BGR2GRAY

	CV_BLUR = C.CV_BLUR
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


const (
	CV_INPAINT_NS = C.CV_INPAINT_NS
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



