// Copyright 2014 <mohamed.helala@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Merged with contributions by jrweizhang AT gmail.com.


// Bindings for Intel's OpenCV computer vision library.
package opencv

//#include "opencv.h"
//#cgo linux  pkg-config: opencv
//#cgo darwin pkg-config: opencv
//#cgo freebsd pkg-config: opencv
//#cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
import "C"

import (
	//"errors"
	"unsafe"
)

/*********************** Background statistics accumulation *****************************/

/****************************************************************************************\
*                                    Image Processing                                    *
\****************************************************************************************/

/* creates structuring element used for morphological operations 
CVAPI(IplConvKernel*)  cvCreateStructuringElementEx(
            int cols, int  rows, int  anchor_x, int  anchor_y,
            int shape, int* values CV_DEFAULT(NULL) );*/
func CreateStructuringElementEx(cols, rows, anchor_x, anchor_y, 
								shape int, values *int32) *IplConvKernel{
	kernel := C.cvCreateStructuringElementEx(C.int(cols), C.int(rows),
		C.int(anchor_x), C.int(anchor_y), C.int(shape),
		(*C.int)(values))
	return (*IplConvKernel)(kernel)
}

/* releases structuring element 
CVAPI(void)  cvReleaseStructuringElement( IplConvKernel** element );*/
func ReleaseStructuringElement(kernel *IplConvKernel){
	k := (*C.IplConvKernel)(kernel)
	C.cvReleaseStructuringElement(&k)
}

/* erodes input image (applies minimum filter) one or more times.
   If element pointer is NULL, 3x3 rectangular element is used 
CVAPI(void)  cvErode( const CvArr* src, CvArr* dst,
                      IplConvKernel* element CV_DEFAULT(NULL),
                      int iterations CV_DEFAULT(1) );*/
func Erode(src, dst unsafe.Pointer, kernel *IplConvKernel, 
	iterations int){
	C.cvErode(src, dst, (*C.IplConvKernel)(kernel), C.int(iterations))
}

/* dilates input image (applies maximum filter) one or more times.
   If element pointer is NULL, 3x3 rectangular element is used 
CVAPI(void)  cvDilate( const CvArr* src, CvArr* dst,
                       IplConvKernel* element CV_DEFAULT(NULL),
                       int iterations CV_DEFAULT(1) );*/
func Dilate(src, dst unsafe.Pointer, kernel *IplConvKernel, 
	iterations int){
	C.cvDilate(src, dst, (*C.IplConvKernel)(kernel), C.int(iterations))
}

/* Performs complex morphological transformation
CVAPI(void)  cvMorphologyEx( const CvArr* src, CvArr* dst,
                             CvArr* temp, IplConvKernel* element,
                             int operation, int iterations CV_DEFAULT(1) );*/
func MorphologyEx(src, dst, temp unsafe.Pointer, kernel *IplConvKernel, 
	operation, iterations int){
	C.cvMorphologyEx(src, dst, temp, (*C.IplConvKernel)(kernel), 
		C.int(operation), C.int(iterations))
}

/* Calculates all spatial and central moments up to the 3rd order 
CVAPI(void) cvMoments( const CvArr* arr, CvMoments* moments, int binary CV_DEFAULT(0));*/
func GetMoments(arr unsafe.Pointer, moms *Moments, binary int){
	C.cvMoments(arr, (*C.CvMoments)(moms), C.int(binary))
}

/* Retrieve particular spatial, central or normalized central moments */

/*CVAPI(double)  cvGetSpatialMoment( CvMoments* moments, int x_order, int y_order );*/
func GetSpatialMoment(moms *Moments, x_order, y_order int) float64{
	value := C.cvGetSpatialMoment((*C.CvMoments)(moms), 
		C.int(x_order), C.int(y_order))
	return float64(value)
}

/*CVAPI(double)  cvGetCentralMoment( CvMoments* moments, int x_order, int y_order );*/
func GetCentralMoment(moms *Moments, x_order, y_order int) float64{
	value := C.cvGetCentralMoment((*C.CvMoments)(moms), 
		C.int(x_order), C.int(y_order))
	return float64(value)
}

/*CVAPI(double)  cvGetNormalizedCentralMoment( CvMoments* moments,
                                             int x_order, int y_order );*/
func GetNormalizedCentralMoment(moms *Moments, x_order, y_order int) float64{
	value := C.cvGetNormalizedCentralMoment((*C.CvMoments)(moms), 
		C.int(x_order), C.int(y_order))
	return float64(value)
}

/* Calculates 7 Hu's invariants from precalculated spatial and central moments*/

/*CVAPI(void) cvGetHuMoments( CvMoments*  moments, CvHuMoments*  hu_moments );*/
func GetHuMoments(moms *Moments, hu_moms *HuMoments){
	C.cvGetHuMoments((*C.CvMoments)(moms), 
		(*C.CvHuMoments)(hu_moms))
}

/* CVAPI(void)  cvResize( const CvArr* src, CvArr* dst,
                       int interpolation CV_DEFAULT( CV_INTER_LINEAR ));*/
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
	Copy(src.Ptr(), dest.Ptr(), nil)
	src.ResetROI()

	return dest
}

/*********************************** data sampling **************************************/

/****************************************************************************************\
*                              Contours retrieving                                       *
\****************************************************************************************/

/* Retrieves outer and optionally inner boundaries of white (non-zero) connected
   components in the black (zero) background
   Default Value: header_size = sizeof(CvContour),
   				  mode = CV_RETR_LIST,
   				  method = CV_CHAIN_APPROX_SIMPLE,
   				  offset = cvPoint(0,0))*/
func FindContours(image unsafe.Pointer, storage *MemStorage, first_contour **Seq,
	header_size int, mode int, method int, offset Point) int {
	cvSeq := (*C.CvSeq)(*first_contour)
	r := C.cvFindContours(image, (*C.CvMemStorage)(storage),
		&cvSeq, C.int(header_size), C.int(mode), C.int(method),
		C.cvPoint(C.int(offset.X), C.int(offset.Y)))
	(*first_contour) = (*Seq)(cvSeq)
	return int(r)
}

// CVAPI(int)  cvFindContours( CvArr* image, CvMemStorage* storage, CvSeq** first_contour,
//                             int header_size CV_DEFAULT(sizeof(CvContour)),
//                             int mode CV_DEFAULT(CV_RETR_LIST),
//                             int method CV_DEFAULT(CV_CHAIN_APPROX_SIMPLE),
//                             CvPoint offset CV_DEFAULT(cvPoint(0,0)));

/* Initalizes contour retrieving process.
   Calls cvStartFindContours.
   Calls cvFindNextContour until null pointer is returned
   or some other condition becomes true.
   Calls cvEndFindContours at the end. */
func StartFindContours(image unsafe.Pointer, storage *MemStorage,
	header_size int, mode int, method int, offset Point) *ContourScanner {
	scanner := C.cvStartFindContours(image, (*C.CvMemStorage)(storage),
		C.int(header_size), C.int(mode), C.int(method),
		C.cvPoint(C.int(offset.X), C.int(offset.Y)))
	return (*ContourScanner)(&scanner)
}

// CVAPI(CvContourScanner)  cvStartFindContours( CvArr* image, CvMemStorage* storage,
//                             int header_size CV_DEFAULT(sizeof(CvContour)),
//                             int mode CV_DEFAULT(CV_RETR_LIST),
//                             int method CV_DEFAULT(CV_CHAIN_APPROX_SIMPLE),
//                             CvPoint offset CV_DEFAULT(cvPoint(0,0)));

/* Retrieves next contour */
func FindNextContour(scanner *ContourScanner) *Seq {
	r := C.cvFindNextContour(C.CvContourScanner(*scanner))
	return (*Seq)(r)
}

// CVAPI(CvSeq*)  cvFindNextContour( CvContourScanner scanner );

/* Substitutes the last retrieved contour with the new one
   (if the substitutor is null, the last retrieved contour is removed from the tree) */
func SubstituteContour(scanner *ContourScanner, new_contour *Seq) {
	C.cvSubstituteContour(C.CvContourScanner(*scanner), (*C.CvSeq)(new_contour))
}

// CVAPI(void)   cvSubstituteContour( CvContourScanner scanner, CvSeq* new_contour );

/* Releases contour scanner and returns pointer to the first outer contour */
func EndFindContours(scanner *ContourScanner) *Seq {
	r := C.cvEndFindContours((*C.CvContourScanner)(scanner))
	return (*Seq)(r)
}

// CVAPI(CvSeq*)  cvEndFindContours( CvContourScanner* scanner );

/* Approximates a single Freeman chain or a tree of chains to polygonal curves */
func ApproxChains(src_seq *Seq, storage *MemStorage, method int, parameter float64,
	minimal_perimeter int, recursive int) *Seq {
	r := C.cvApproxChains((*C.CvSeq)(src_seq), (*C.CvMemStorage)(storage), C.int(method),
		C.double(parameter), C.int(minimal_perimeter), C.int(recursive))
	return (*Seq)(r)
}

// CVAPI(CvSeq*) cvApproxChains( CvSeq* src_seq, CvMemStorage* storage,
//                             int method CV_DEFAULT(CV_CHAIN_APPROX_SIMPLE),
//                             double parameter CV_DEFAULT(0),
//                             int  minimal_perimeter CV_DEFAULT(0),
//                             int  recursive CV_DEFAULT(0));

/* Initalizes Freeman chain reader.
   The reader is used to iteratively get coordinates of all the chain points.
   If the Freeman codes should be read as is, a simple sequence reader should be used */
func StartReadChainPoints(chain *Chain, reader *ChainPtReader) {
	C.cvStartReadChainPoints((*C.CvChain)(chain), (*C.CvChainPtReader)(reader))
}

// CVAPI(void) cvStartReadChainPoints( CvChain* chain, CvChainPtReader* reader );

/* Retrieves the next chain point */
func ReadChainPoint(reader *ChainPtReader) *Point {
	r := C.cvReadChainPoint((*C.CvChainPtReader)(reader))
	p := &Point{int(r.x), int(r.y)}
	return p
}

// CVAPI(CvPoint) cvReadChainPoint( CvChainPtReader* reader );

/****************************************************************************************\
*                            Contour Processing and Shape Analysis                       *
\****************************************************************************************/

/* Calculates perimeter of a contour*/
func ContourPerimeter(seq unsafe.Pointer) float64 {
	r := C.cvContourPerimeter(seq)
	return (float64)(r)
}

// CV_INLINE double cvContourPerimeter( const void* contour )
// {
//     return cvArcLength( contour, CV_WHOLE_SEQ, 1 );
// }

/* Approximates a single polygonal curve (contour) or
   a tree of polygonal curves (contours)
   Default Values: recursive = 0*/
func ApproxPoly(src_seq unsafe.Pointer, header_size int, storage *MemStorage,
	method int, eps float64, recursive int) *Seq {
	r := C.cvApproxPoly(src_seq, C.int(header_size), (*C.CvMemStorage)(storage),
		C.int(method), C.double(eps), C.int(recursive))
	return (*Seq)(r)
}

// CVAPI(CvSeq*)  cvApproxPoly( const void* src_seq,
//                              int header_size, CvMemStorage* storage,
//                              int method, double eps,
//                              int recursive CV_DEFAULT(0));

/****************************************************************************************\
*                                  Histogram functions                                   *
\****************************************************************************************/

/* Applies fixed-level threshold to grayscale image.
   This is a basic operation applied before retrieving contours */
func Threshold(src unsafe.Pointer, dst unsafe.Pointer,
	threshold, max_value float64, threshold_type int) float64 {
	r := C.cvThreshold(src, dst, C.double(threshold),
		C.double(max_value), C.int(threshold_type))
	return float64(r)
}

// CVAPI(double)  cvThreshold( const CvArr*  src, CvArr*  dst,
//                             double  threshold, double  max_value,
//                             int threshold_type );

/****************************************************************************************\
*                                  Feature detection                                     *
\****************************************************************************************/
