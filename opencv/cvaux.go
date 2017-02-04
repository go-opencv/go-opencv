// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

//#include "opencv.h"
//#cgo linux  pkg-config: opencv
//#cgo darwin pkg-config: opencv
//#cgo freebsd pkg-config: opencv
//#cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
import "C"
import (
	"unsafe"
)

/****************************************************************************************\
*                                  Eigen objects                                         *
\****************************************************************************************/

/****************************************************************************************\
*                                       1D/2D HMM                                        *
\****************************************************************************************/

/*********************************** Embedded HMMs *************************************/

/****************************************************************************************\
*               A few functions from old stereo gesture recognition demosions            *
\****************************************************************************************/

/****************************************************************************************\
*                           Additional operations on Subdivisions                        *
\****************************************************************************************/

/****************************************************************************************\
*                           More operations on sequences                                 *
\****************************************************************************************/

/*******************************Stereo correspondence*************************************/

/*****************************************************************************************/
/************ Epiline functions *******************/

/****************************************************************************************\
*                                   Contour Morphing                                     *
\****************************************************************************************/

/****************************************************************************************\
*                                    Texture Descriptors                                 *
\****************************************************************************************/

/****************************************************************************************\
*                                  Face eyes&mouth tracking                              *
\****************************************************************************************/
type HaarCascade struct {
	cascade *C.CvHaarClassifierCascade
}

func LoadHaarClassifierCascade(haar string) *HaarCascade {
	haarCascade := new(HaarCascade)
	haarCascade.cascade = C.cvLoadHaarClassifierCascade(C.CString(haar), C.cvSize(1, 1))
	return haarCascade
}

func (this *HaarCascade) DetectObjects(image *IplImage) []*Rect {
	storage := C.cvCreateMemStorage(C.int(0))
	seq := C.cvHaarDetectObjects(unsafe.Pointer(image), this.cascade, storage, 1.1, 3, C.CV_HAAR_DO_CANNY_PRUNING, C.cvSize(0, 0), C.cvSize(0, 0))
	var faces []*Rect
	for i := 0; i < (int)(seq.total); i++ {
		rect := (*Rect)((*_Ctype_CvRect)(unsafe.Pointer(C.cvGetSeqElem(seq, C.int(i)))))
		rectgc := NewRect(rect.X(),rect.Y(),rect.Width(),rect.Height())
		faces = append(faces, &rectgc)
	}

	storage_c := (*C.CvMemStorage)(storage)
	C.cvReleaseMemStorage(&storage_c)

	return faces
}

func (this *HaarCascade) Release() {
	cascade_c := (*C.CvHaarClassifierCascade)(this.cascade)
	C.cvReleaseHaarClassifierCascade(&cascade_c)
}

/****************************************************************************************\
*                                         3D Tracker                                     *
\****************************************************************************************/

/****************************************************************************************\
*                           Skeletons and Linear-Contour Models                          *
\****************************************************************************************/

/****************************************************************************************\
*                           Background/foreground segmentation                           *
\****************************************************************************************/

/****************************************************************************************\
*                                   Calibration engine                                   *
\****************************************************************************************/

/*****************************************************************************\
*                                 --- END ---                                 *
\*****************************************************************************/
