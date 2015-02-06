// Copyright 2014 <me@cwchang.me>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Bindings for OpenCV's Calib3d (Camera Calibration
// and 3D Reconstruction) module
package calib3d

//#include "../opencv.h"
//#cgo linux  pkg-config: opencv
//#cgo darwin pkg-config: opencv
//#cgo freebsd pkg-config: opencv
//#cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
import "C"
import "unsafe"

/*
void cvInitIntrinsicParams2D(
const CvMat* object_points, const CvMat* image_points,
const CvMat* npoints, CvSize image_size, CvMat* camera_matrix,
double aspect_ratio=1. )
*/

func InitIntrinsicParams2D(objPoints, imgPoints, nPoints *Mat, imgWidth, imgHeight int, aspectRatio float64) (cameraMatrix *Mat) {
	cameraMatrix = CreateMat(3, 3, CV_64F)

	size := C.cvSize(C.int(imgWidth), C.int(imgHeight))

	C.cvInitIntrinsicParams2D(
		unsafe.Pointer(objPoints),
		unsafe.Pointer(imgPoints),
		unsafe.Pointer(nPoints),
		size,
		unsafe.Pointer(cameraMatrix),
		aspectRatio)

	return cameraMatrix
}
