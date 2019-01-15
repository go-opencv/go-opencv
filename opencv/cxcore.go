// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

//#include "opencv.h"
import "C"
import (
	//"errors"
	"unsafe"
)

func init() {
}

/****************************************************************************************\
* Array allocation, deallocation, initialization and access to elements       *
\****************************************************************************************/

func Alloc(size int) unsafe.Pointer {
	return unsafe.Pointer(C.cvAlloc(C.size_t(size)))
}
func Free(p unsafe.Pointer) {
	C.cvFree_(p)
}

/* Allocates and initializes IplImage header */
func CreateImageHeader(w, h, depth, channels int) *IplImage {
	hdr := C.cvCreateImageHeader(
		C.cvSize(C.int(w), C.int(h)),
		C.int(depth),
		C.int(channels),
	)
	return (*IplImage)(hdr)
}

/* Inializes IplImage header */
func (img *IplImage) InitHeader(w, h, depth, channels, origin, align int) {
	C.cvInitImageHeader(
		(*C.IplImage)(img),
		C.cvSize(C.int(w), C.int(h)),
		C.int(depth),
		C.int(channels),
		C.int(origin),
		C.int(align),
	)
}

/* Creates IPL image (header and data) */
func CreateImage(w, h, depth, channels int) *IplImage {
	size := C.cvSize(C.int(w), C.int(h))
	img := C.cvCreateImage(size, C.int(depth), C.int(channels))
	return (*IplImage)(img)
}

// Merge creates one multichannel array out of several single-channel ones.
func Merge(imgBlue, imgGreen, imgRed, imgAlpha, dst *IplImage) {
	C.cvMerge(
		unsafe.Pointer(imgBlue),
		unsafe.Pointer(imgGreen),
		unsafe.Pointer(imgRed),
		unsafe.Pointer(imgAlpha),
		unsafe.Pointer(dst),
	)
}

// Split divides a multi-channel array into several single-channel arrays.
func Split(src, imgBlue, imgGreen, imgRed, imgAlpha *IplImage) {
	C.cvSplit(
		unsafe.Pointer(src),
		unsafe.Pointer(imgBlue),
		unsafe.Pointer(imgGreen),
		unsafe.Pointer(imgRed),
		unsafe.Pointer(imgAlpha),
	)
}

// AddWeighted calculates the weighted sum of two images.
func AddWeighted(src1 *IplImage, alpha float64, src2 *IplImage, beta float64, gamma float64, dst *IplImage) {
	C.cvAddWeighted(
		unsafe.Pointer(src1),
		C.double(alpha),
		unsafe.Pointer(src2),
		C.double(beta),
		C.double(gamma),
		unsafe.Pointer(dst),
	)
}

/* SetData assigns user data to the image header */
func (img *IplImage) SetData(data unsafe.Pointer, step int) {
	C.cvSetData(unsafe.Pointer(img), data, C.int(step))
}

/* Releases (i.e. deallocates) IPL image header */
func (img *IplImage) ReleaseHeader() {
	img_c := (*C.IplImage)(img)
	C.cvReleaseImageHeader(&img_c)
}

/* Releases IPL image header and data */
func (img *IplImage) Release() {
	img_c := (*C.IplImage)(img)
	C.cvReleaseImage(&img_c)
}

func (img *IplImage) Zero() {
	C.cvSetZero(unsafe.Pointer(img))
}

/* Creates a copy of IPL image (widthStep may differ) */
func (img *IplImage) Clone() *IplImage {
	p := C.cvCloneImage((*C.IplImage)(img))
	return (*IplImage)(p)
}

/* Sets a Channel Of Interest (only a few functions support COI) -
   use cvCopy to extract the selected channel and/or put it back */
func (img *IplImage) SetCOI(coi int) {
	C.cvSetImageCOI((*C.IplImage)(img), C.int(coi))
}

/* Retrieves image Channel Of Interest */
func (img *IplImage) GetCOI() int {
	coi := C.cvGetImageCOI((*C.IplImage)(img))
	return int(coi)
}

/* Sets image ROI (region of interest) (COI is not changed) */
func (img *IplImage) SetROI(rect Rect) {
	C.cvSetImageROI((*C.IplImage)(img), C.CvRect(rect))
}

/* Resets image ROI and COI */
func (img *IplImage) ResetROI() {
	C.cvResetImageROI((*C.IplImage)(img))
}

/* Retrieves image ROI */
func (img *IplImage) GetROI() Rect {
	r := C.cvGetImageROI((*C.IplImage)(img))
	return Rect(r)
}

/*
Reshape changes shape of the image without copying data. A value of `0` means
that channels or rows remain unchanged.
*/
func (img *IplImage) Reshape(channels, rows, _type int) *Mat {
	total := img.Width() * img.Height()
	header := CreateMat(rows, total/rows, _type)
	n := C.cvReshape(unsafe.Pointer(img), (*C.CvMat)(header), C.int(channels), C.int(rows))
	return (*Mat)(n)
}

/* Get1D return a specific element from a 1-dimensional matrix. */
func (img *IplImage) Get1D(x int) Scalar {
	ret := C.cvGet1D(unsafe.Pointer(img), C.int(x))
	return Scalar(ret)
}

/* Get2D return a specific element from a 2-dimensional matrix. */
func (img *IplImage) Get2D(x, y int) Scalar {
	ret := C.cvGet2D(unsafe.Pointer(img), C.int(y), C.int(x))
	return Scalar(ret)
}

/* Get3D return a specific element from a 3-dimensional matrix. */
func (img *IplImage) Get3D(x, y, z int) Scalar {
	ret := C.cvGet3D(unsafe.Pointer(img), C.int(z), C.int(y), C.int(x))
	return Scalar(ret)
}

/* Sets every element of an array to a given value. */
func (img *IplImage) Set(value Scalar) {
	C.cvSet(unsafe.Pointer(img), (C.CvScalar)(value), nil)
}

/* Set1D sets a particular element in the image */
func (img *IplImage) Set1D(x int, value Scalar) {
	C.cvSet1D(unsafe.Pointer(img), C.int(x), (C.CvScalar)(value))
}

/* Set2D sets a particular element in the image */
func (img *IplImage) Set2D(x, y int, value Scalar) {
	C.cvSet2D(unsafe.Pointer(img), C.int(y), C.int(x), (C.CvScalar)(value))
}

/* Set3D sets a particular element in the image */
func (img *IplImage) Set3D(x, y, z int, value Scalar) {
	C.cvSet3D(unsafe.Pointer(img), C.int(z), C.int(y), C.int(x), (C.CvScalar)(value))
}

/* GetMat returns the matrix header for an image.*/
func (img *IplImage) GetMat() *Mat {
	var null C.int
	tmp := CreateMat(img.Height(), img.Width(), CV_32S)
	m := C.cvGetMat(unsafe.Pointer(img), (*C.CvMat)(tmp), &null, C.int(0))
	return (*Mat)(m)
}

// mat step
const (
	CV_AUTOSTEP = C.CV_AUTOSTEP
)

/* Allocates and initalizes CvMat header */
func CreateMatHeader(rows, cols, type_ int) *Mat {
	mat := C.cvCreateMatHeader(
		C.int(rows), C.int(cols), C.int(type_),
	)
	return (*Mat)(mat)
}

/* Allocates and initializes CvMat header and allocates data */
func CreateMat(rows, cols, type_ int) *Mat {
	mat := C.cvCreateMat(
		C.int(rows), C.int(cols), C.int(type_),
	)
	return (*Mat)(mat)
}

/* Initializes CvMat header */
func (mat *Mat) InitHeader(rows, cols, type_ int, data unsafe.Pointer, step int) {
	C.cvInitMatHeader(
		(*C.CvMat)(mat),
		C.int(rows),
		C.int(cols),
		C.int(type_),
		data,
		C.int(step),
	)
}

/* SetData assigns user data to the matrix header. */
func (mat *Mat) SetData(data unsafe.Pointer, step int) {
	C.cvSetData(unsafe.Pointer(mat), data, C.int(step))
}

/* Releases CvMat header and deallocates matrix data
   (reference counting is used for data) */
func (mat *Mat) Release() {
	mat_c := (*C.CvMat)(mat)
	C.cvReleaseMat(&mat_c)
}

/* Decrements CvMat data reference counter and deallocates the data if
   it reaches 0 */
func DecRefData(arr Arr) {
	C.cvDecRefData(unsafe.Pointer(arr))
}

/* Increments CvMat data reference counter */
func IncRefData(arr Arr) {
	C.cvIncRefData(unsafe.Pointer(arr))
}

/* Creates an exact copy of the input matrix (except, may be, step value) */
func (mat *Mat) Clone() *Mat {
	mat_new := C.cvCloneMat((*C.CvMat)(mat))
	return (*Mat)(mat_new)
}

func (mat *Mat) Zero() {
	C.cvSetZero(unsafe.Pointer(mat))
}

/*
Reshape changes shape of the matrix without copying data. A value of `0` means
that channels or rows remain unchanged.
*/
func (m *Mat) Reshape(channels, rows int) *Mat {
	total := m.Cols() * m.Rows()
	n := CreateMat(rows, total/rows, m.Type())
	C.cvReshape(unsafe.Pointer(m), (*C.CvMat)(n), C.int(channels), C.int(rows))
	return n
}

/* Makes a new matrix from <rect> subrectangle of input array.
   No data is copied */
func GetSubRect(arr Arr, submat *Mat, rect Rect) *Mat {
	mat_new := C.cvGetSubRect(
		unsafe.Pointer(arr),
		(*C.CvMat)(submat),
		(C.CvRect)(rect),
	)
	return (*Mat)(mat_new)
}

//#define cvGetSubArr cvGetSubRect

/* Selects row span of the input array: arr(start_row:delta_row:end_row,:)
   (end_row is not included into the span). */
func GetRows(arr Arr, submat *Mat, start_row, end_row, delta_row int) *Mat {
	mat_new := C.cvGetRows(
		unsafe.Pointer(arr),
		(*C.CvMat)(submat),
		C.int(start_row),
		C.int(end_row),
		C.int(delta_row),
	)
	return (*Mat)(mat_new)
}
func GetRow(arr Arr, submat *Mat, row int) *Mat {
	mat_new := C.cvGetRow(
		unsafe.Pointer(arr),
		(*C.CvMat)(submat),
		C.int(row),
	)
	return (*Mat)(mat_new)
}

/* Selects column span of the input array: arr(:,start_col:end_col)
   (end_col is not included into the span) */
func GetCols(arr Arr, submat *Mat, start_col, end_col int) *Mat {
	mat_new := C.cvGetCols(
		unsafe.Pointer(arr),
		(*C.CvMat)(submat),
		C.int(start_col),
		C.int(end_col),
	)
	return (*Mat)(mat_new)
}
func GetCol(arr Arr, submat *Mat, col int) *Mat {
	mat_new := C.cvGetCol(
		unsafe.Pointer(arr),
		(*C.CvMat)(submat),
		C.int(col),
	)
	return (*Mat)(mat_new)
}

/* Select a diagonal of the input array.
   (diag = 0 means the main diagonal, >0 means a diagonal above the main one,
   <0 - below the main one).
   The diagonal will be represented as a column (nx1 matrix). */
func GetDiag(arr Arr, submat *Mat, diag int) *Mat {
	mat_new := C.cvGetDiag(
		unsafe.Pointer(arr),
		(*C.CvMat)(submat),
		C.int(diag),
	)
	return (*Mat)(mat_new)
}

/* Get1D return a specific element from a 1-dimensional matrix. */
func (m *Mat) Get1D(x int) Scalar {
	ret := C.cvGet1D(unsafe.Pointer(m), C.int(x))
	return Scalar(ret)
}

/* Get2D return a specific element from a 2-dimensional matrix. */
func (m *Mat) Get2D(x, y int) Scalar {
	ret := C.cvGet2D(unsafe.Pointer(m), C.int(x), C.int(y))
	return Scalar(ret)
}

/* Get3D return a specific element from a 3-dimensional matrix. */
func (m *Mat) Get3D(x, y, z int) Scalar {
	ret := C.cvGet3D(unsafe.Pointer(m), C.int(x), C.int(y), C.int(z))
	return Scalar(ret)
}

/* Set1D sets a particular element in them matrix */
func (m *Mat) Set1D(x int, value Scalar) {
	C.cvSet1D(unsafe.Pointer(m), C.int(x), (C.CvScalar)(value))
}

/* Set2D sets a particular element in them matrix */
func (m *Mat) Set2D(x, y int, value Scalar) {
	C.cvSet2D(unsafe.Pointer(m), C.int(x), C.int(y), (C.CvScalar)(value))
}

/* Set3D sets a particular element in them matrix */
func (m *Mat) Set3D(x, y, z int, value Scalar) {
	C.cvSet3D(unsafe.Pointer(m), C.int(x), C.int(y), C.int(z), (C.CvScalar)(value))
}

/* GetImage returns the image header for the matrix. */
func (m *Mat) GetImage(channels int) *IplImage {
	tmp := CreateImage(m.Cols(), m.Rows(), m.Type(), channels)
	img := C.cvGetImage(unsafe.Pointer(m), (*C.IplImage)(tmp))

	return (*IplImage)(img)
}

/* low-level scalar <-> raw data conversion functions */
func ScalarToRawData(scalar *Scalar, data unsafe.Pointer, type_, extend_to_12 int) {
	C.cvScalarToRawData(
		(*C.CvScalar)(scalar),
		data,
		C.int(type_),
		C.int(extend_to_12),
	)
}
func RawDataToScalar(data unsafe.Pointer, type_ int, scalar *Scalar) {
	C.cvRawDataToScalar(
		data,
		C.int(type_),
		(*C.CvScalar)(scalar),
	)
}

/* Allocates and initializes CvMatND header */
func CreateMatNDHeader(sizes []int, type_ int) *MatND {
	dims := C.int(len(sizes))
	sizes_c := make([]C.int, len(sizes))
	for i := 0; i < len(sizes); i++ {
		sizes_c[i] = C.int(sizes[i])
	}

	mat := C.cvCreateMatNDHeader(
		dims, (*C.int)(&sizes_c[0]), C.int(type_),
	)
	return (*MatND)(mat)
}

/* Allocates and initializes CvMatND header and allocates data */
func CreateMatND(sizes []int, type_ int) *MatND {
	dims := C.int(len(sizes))
	sizes_c := make([]C.int, len(sizes))
	for i := 0; i < len(sizes); i++ {
		sizes_c[i] = C.int(sizes[i])
	}

	mat := C.cvCreateMatND(
		dims, (*C.int)(&sizes_c[0]), C.int(type_),
	)
	return (*MatND)(mat)
}

/* Initializes preallocated CvMatND header */
func (mat *MatND) InitMatNDHeader(sizes []int, type_ int, data unsafe.Pointer) {
	dims := C.int(len(sizes))
	sizes_c := make([]C.int, len(sizes))
	for i := 0; i < len(sizes); i++ {
		sizes_c[i] = C.int(sizes[i])
	}

	C.cvInitMatNDHeader(
		(*C.CvMatND)(mat),
		dims, (*C.int)(&sizes_c[0]), C.int(type_),
		data,
	)
}

/* Releases CvMatND */
func (mat *MatND) Release() {
	mat_c := (*C.CvMatND)(mat)
	C.cvReleaseMatND(&mat_c)
}

/* Creates a copy of CvMatND (except, may be, steps) */
func (mat *MatND) Clone() *MatND {
	mat_c := (*C.CvMatND)(mat)
	mat_ret := C.cvCloneMatND(mat_c)
	return (*MatND)(mat_ret)
}

/* Allocates and initializes CvSparseMat header and allocates data */
func CreateSparseMat(sizes []int, type_ int) *SparseMat {
	dims := C.int(len(sizes))
	sizes_c := make([]C.int, len(sizes))
	for i := 0; i < len(sizes); i++ {
		sizes_c[i] = C.int(sizes[i])
	}

	mat := C.cvCreateSparseMat(
		dims, (*C.int)(&sizes_c[0]), C.int(type_),
	)
	return (*SparseMat)(mat)
}

/* Releases CvSparseMat */
func (mat *SparseMat) Release() {
	mat_c := (*C.CvSparseMat)(mat)
	C.cvReleaseSparseMat(&mat_c)
}

/* Creates a copy of CvSparseMat (except, may be, zero items) */
func (mat *SparseMat) Clone() *SparseMat {
	mat_c := (*C.CvSparseMat)(mat)
	mat_ret := C.cvCloneSparseMat(mat_c)
	return (*SparseMat)(mat_ret)
}

/* Initializes sparse array iterator
   (returns the first node or NULL if the array is empty) */
func (mat *SparseMat) InitSparseMatIterator(iter *SparseMatIterator) *SparseNode {
	mat_c := (*C.CvSparseMat)(mat)
	node := C.cvInitSparseMatIterator(mat_c, (*C.CvSparseMatIterator)(iter))
	return (*SparseNode)(node)
}

// returns next sparse array node (or NULL if there is no more nodes)
func (iter *SparseMatIterator) Next() *SparseNode {
	node := C.cvGetNextSparseNode((*C.CvSparseMatIterator)(iter))
	return (*SparseNode)(node)
}

/******** matrix iterator: used for n-ary operations on dense arrays *********/

// P290

/* Returns width and height of array in elements */
func GetSizeWidth(img *IplImage) int {
	size := C.cvGetSize(unsafe.Pointer(img))
	w := int(size.width)
	return w
}
func GetSizeHeight(img *IplImage) int {
	size := C.cvGetSize(unsafe.Pointer(img))
	w := int(size.height)
	return w
}
func GetSize(img *IplImage) Size {
	sz := C.cvGetSize(unsafe.Pointer(img))
	return Size{int(sz.width), int(sz.height)}

}

/* Copies source array to destination array */
func Copy(src, dst, mask *IplImage) {
	C.cvCopy(unsafe.Pointer(src), unsafe.Pointer(dst), unsafe.Pointer(mask))
}

//CVAPI(void)  cvCopy( const CvArr* src, CvArr* dst,
//                     const CvArr* mask CV_DEFAULT(NULL) );

/* Clears all the array elements (sets them to 0) */
func Zero(img *IplImage) {
	C.cvSetZero(unsafe.Pointer(img))
}

//CVAPI(void)  cvSetZero( CvArr* arr );
//#define cvZero  cvSetZero

/****************************************************************************************\
*                   Arithmetic, logic and comparison operations               *
\****************************************************************************************/

/****************************************************************************************\
*                                Logic operations                             *
\****************************************************************************************/

// Inverts every bit of an array.
func Not(src, dst *IplImage) {
	C.cvNot(
		unsafe.Pointer(src),
		unsafe.Pointer(dst),
	)
}

// Calculates the per-element bit-wise conjunction of two arrays.
func And(src1, src2, dst *IplImage) {
	AndWithMask(src1, src2, dst, nil)
}

// Calculates the per-element bit-wise conjunction of two arrays with a mask.
func AndWithMask(src1, src2, dst, mask *IplImage) {
	C.cvAnd(
		unsafe.Pointer(src1),
		unsafe.Pointer(src2),
		unsafe.Pointer(dst),
		unsafe.Pointer(mask),
	)
}

// Calculates the per-element bit-wise conjunction of an array and a scalar.
func AndScalar(src *IplImage, value Scalar, dst *IplImage) {
	AndScalarWithMask(src, value, dst, nil)
}

// Calculates the per-element bit-wise conjunction of an array and a scalar with a mask.
func AndScalarWithMask(src *IplImage, value Scalar, dst, mask *IplImage) {
	C.cvAndS(
		unsafe.Pointer(src),
		(C.CvScalar)(value),
		unsafe.Pointer(dst),
		unsafe.Pointer(mask),
	)
}

// Calculates the per-element bit-wise disjunction of two arrays.
func Or(src1, src2, dst *IplImage) {
	OrWithMask(src1, src2, dst, nil)
}

// Calculates the per-element bit-wise disjunction of two arrays with a mask.
func OrWithMask(src1, src2, dst, mask *IplImage) {
	C.cvOr(
		unsafe.Pointer(src1),
		unsafe.Pointer(src2),
		unsafe.Pointer(dst),
		unsafe.Pointer(mask),
	)
}

// Calculates the per-element bit-wise disjunction of an array and a scalar.
func OrScalar(src *IplImage, value Scalar, dst *IplImage) {
	OrScalarWithMask(src, value, dst, nil)
}

// Calculates the per-element bit-wise disjunction of an array and a scalar with a mask.
func OrScalarWithMask(src *IplImage, value Scalar, dst, mask *IplImage) {
	C.cvOrS(
		unsafe.Pointer(src),
		(C.CvScalar)(value),
		unsafe.Pointer(dst),
		unsafe.Pointer(mask),
	)
}

// Calculates the per-element bit-wise “exclusive or” operation on two arrays.
func Xor(src1, src2, dst *IplImage) {
	XorWithMask(src1, src2, dst, nil)
}

// Calculates the per-element bit-wise “exclusive or” operation on two arrays with a mask.
func XorWithMask(src1, src2, dst, mask *IplImage) {
	C.cvXor(
		unsafe.Pointer(src1),
		unsafe.Pointer(src2),
		unsafe.Pointer(dst),
		unsafe.Pointer(mask),
	)
}

// Calculates the per-element bit-wise “exclusive or” operation on an array and a scalar.
func XorScalar(src *IplImage, value Scalar, dst *IplImage) {
	XorScalarWithMask(src, value, dst, nil)
}

// Calculates the per-element bit-wise “exclusive or” operation on an array and a scalar with a mask.
func XorScalarWithMask(src *IplImage, value Scalar, dst, mask *IplImage) {
	C.cvXorS(
		unsafe.Pointer(src),
		(C.CvScalar)(value),
		unsafe.Pointer(dst),
		unsafe.Pointer(mask),
	)
}

/****************************************************************************************\
*                                Math operations                              *
\****************************************************************************************/

// Calculates the per-element sum of two arrays.
//   dst = src1 + src2
func Add(src1, src2, dst *IplImage) {
	AddWithMask(src1, src2, dst, nil)
}

// Calculates the per-element sum of two arrays with a mask.
//   dst = src1 + src2
func AddWithMask(src1, src2, dst, mask *IplImage) {
	C.cvAdd(
		unsafe.Pointer(src1),
		unsafe.Pointer(src2),
		unsafe.Pointer(dst),
		unsafe.Pointer(mask),
	)
}

// Calculates the per-element sum of an array and a scalar.
//   dst = src + value
func AddScalar(src *IplImage, value Scalar, dst *IplImage) {
	AddScalarWithMask(src, value, dst, nil)
}

// Calculates the per-element sum of an array and a scalar with a mask.
//   dst = src + value
func AddScalarWithMask(src *IplImage, value Scalar, dst, mask *IplImage) {
	C.cvAddS(
		unsafe.Pointer(src),
		(C.CvScalar)(value),
		unsafe.Pointer(dst),
		unsafe.Pointer(mask),
	)
}

// Calculates the per-element difference between two arrays.
//   dst = src1 - src2
func Subtract(src1, src2, dst *IplImage) {
	SubtractWithMask(src1, src2, dst, nil)
}

// Calculates the per-element difference between two arrays with a mask.
//   dst = src1 - src2
func SubtractWithMask(src1, src2, dst, mask *IplImage) {
	C.cvSub(
		unsafe.Pointer(src1),
		unsafe.Pointer(src2),
		unsafe.Pointer(dst),
		unsafe.Pointer(mask),
	)
}

// Calculates the per-element difference between an array and a scalar.
//   dst = src - value
func SubScalar(src *IplImage, value Scalar, dst *IplImage) {
	SubScalarWithMask(src, value, dst, nil)
}

// Calculates the per-element difference between an array and a scalar with a mask.
//   dst = src - value
func SubScalarWithMask(src *IplImage, value Scalar, dst, mask *IplImage) {
	C.cvSubS(
		unsafe.Pointer(src),
		(C.CvScalar)(value),
		unsafe.Pointer(dst),
		unsafe.Pointer(mask),
	)
}

// Calculates the per-element difference between a scalar and an array.
//   dst = value - src
func SubScalarRev(value Scalar, src, dst *IplImage) {
	SubScalarWithMaskRev(value, src, dst, nil)
}

// Calculates the per-element difference between a scalar and an array with a mask.
//   dst = value - src
func SubScalarWithMaskRev(value Scalar, src, dst, mask *IplImage) {
	C.cvSubRS(
		unsafe.Pointer(src),
		(C.CvScalar)(value),
		unsafe.Pointer(dst),
		unsafe.Pointer(mask),
	)
}

// Calculates the per-element absolute difference between two arrays.
func AbsDiff(src1, src2, dst *IplImage) {
	C.cvAbsDiff(
		unsafe.Pointer(src1),
		unsafe.Pointer(src2),
		unsafe.Pointer(dst),
	)
}

// Calculates the per-element absolute difference between an array and a scalar.
func AbsDiffScalar(src *IplImage, value Scalar, dst *IplImage) {
	C.cvAbsDiffS(
		unsafe.Pointer(src),
		unsafe.Pointer(dst),
		(C.CvScalar)(value),
	)
}

/****************************************************************************************\
*                                Matrix operations                            *
\****************************************************************************************/

/****************************************************************************************\
*                                    Array Statistics                         *
\****************************************************************************************/
// CvScalar cvAvg(const CvArr* arr, const CvArr* mask=NULL )
func (src *IplImage) Avg(mask *IplImage) Scalar {
	return (Scalar)(C.cvAvg(unsafe.Pointer(src), unsafe.Pointer(mask)))
}

// cvEqualizeHist(const CvArr* src, CvArr* dst)
func (src *IplImage) EqualizeHist(dst *IplImage) {
	C.cvEqualizeHist(unsafe.Pointer(src), unsafe.Pointer(dst))
}

// MeanStdDev alculates mean and standard deviation of pixel values
func (src *IplImage) MeanStdDev() (Scalar, Scalar) {
	return MeanStdDevWithMask(src, nil)
}

// MeanStdDevWithMask calculates mean and standard deviation of pixel values with mask
func MeanStdDevWithMask(src, mask *IplImage) (Scalar, Scalar) {
	var mean, stdDev Scalar
	C.cvAvgSdv(
		unsafe.Pointer(src),
		(*C.CvScalar)(&mean),
		(*C.CvScalar)(&stdDev),
		unsafe.Pointer(mask),
	)

	return mean, stdDev
}

/****************************************************************************************\
*                      Discrete Linear Transforms and Related Functions       *
\****************************************************************************************/

/****************************************************************************************\
*                              Dynamic data structures                        *
\****************************************************************************************/

const (
	// different sequence flags to use in CreateSeq()
	CV_SEQ_ELTYPE_POINT   = C.CV_SEQ_ELTYPE_POINT
	CV_32FC2              = C.CV_32FC2
	CV_SEQ_ELTYPE_POINT3D = C.CV_SEQ_ELTYPE_POINT3D
)

// Creates a new sequence.
func CreateSeq(seq_flags, elem_size int) *Seq {
	return (*Seq)(C.cvCreateSeq(
		C.int(seq_flags),
		C.size_t(unsafe.Sizeof(Seq{})),
		C.size_t(elem_size),
		C.cvCreateMemStorage(C.int(0)),
	))
}

// Adds an element to the sequence end.
// Returns a pointer to the element added.
func (seq *Seq) Push(element unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.cvSeqPush((*C.struct_CvSeq)(seq), element))
}

// Removes element from the sequence end.
// Copies the element into the paramter element.
func (seq *Seq) Pop(element unsafe.Pointer) {
	C.cvSeqPop((*C.struct_CvSeq)(seq), element)
}

// Adds an element to the sequence beginning.
// Returns a pointer to the element added.
func (seq *Seq) PushFront(element unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer((C.cvSeqPushFront((*C.struct_CvSeq)(seq), element)))
}

// Removes element from the sequence beginning.
// Copies the element into the paramter element.
func (seq *Seq) PopFront(element unsafe.Pointer) {
	C.cvSeqPopFront((*C.struct_CvSeq)(seq), element)
}

// Releases the sequence storage.
func (seq *Seq) Release() {
	C.cvReleaseMemStorage(&seq.storage)
}

// Gets the total number of elements in the sequence
func (seq *Seq) Total() int {
	return (int)(seq.total)
}

// Gets a pointer to the next sequence
func (seq *Seq) HNext() *Seq {
	return (*Seq)(seq.h_next)
}

// Gets a pointer to the previous sequence
func (seq *Seq) HPrev() *Seq {
	return (*Seq)(seq.h_prev)
}

// Gets a pointer to the 2nd next sequence
func (seq *Seq) VNext() *Seq {
	return (*Seq)(seq.v_next)
}

// Gets a pointer to the 2nd previous sequence
func (seq *Seq) VPrev() *Seq {
	return (*Seq)(seq.v_prev)
}

// Gets a pointer to the element at the index
func (seq *Seq) GetElemAt(index int) unsafe.Pointer {
	return (unsafe.Pointer)(C.cvGetSeqElem(
		(*C.struct_CvSeq)(seq),
		C.int(index),
	))
}

// Removes an element from the middle of a sequence.
func (seq *Seq) RemoveAt(index int) {
	C.cvSeqRemove((*C.struct_CvSeq)(seq), C.int(index))
}

// Removes all elements from the sequence.
// Does not release storage, do that by calling Release().
func (seq *Seq) Clear() {
	C.cvClearSeq((*C.struct_CvSeq)(seq))
}

// Gets a pointer to the storage
func (seq *Seq) Storage() *MemStorage {
	return (*MemStorage)(seq.storage)
}

/****************************************************************************************\
*                                     Drawing                                 *
\****************************************************************************************/

/* Draws 4-connected, 8-connected or antialiased line segment connecting two points */
//color Scalar,
func Line(image *IplImage, pt1, pt2 Point, color Scalar, thickness, line_type, shift int) {
	C.cvLine(
		unsafe.Pointer(image),
		C.cvPoint(C.int(pt1.X), C.int(pt1.Y)),
		C.cvPoint(C.int(pt2.X), C.int(pt2.Y)),
		(C.CvScalar)(color),
		C.int(thickness), C.int(line_type), C.int(shift),
	)
}
func Rectangle(image *IplImage, pt1, pt2 Point, color Scalar, thickness, line_type, shift int) {
	C.cvRectangle(
		unsafe.Pointer(image),
		C.cvPoint(C.int(pt1.X), C.int(pt1.Y)),
		C.cvPoint(C.int(pt2.X), C.int(pt2.Y)),
		(C.CvScalar)(color),
		C.int(thickness), C.int(line_type), C.int(shift),
	)
}
func Circle(image *IplImage, pt1 Point, radius int, color Scalar, thickness, line_type, shift int) {
	C.cvCircle(
		unsafe.Pointer(image),
		C.cvPoint(C.int(pt1.X), C.int(pt1.Y)),
		C.int(radius),
		(C.CvScalar)(color),
		C.int(thickness), C.int(line_type), C.int(shift),
	)
}

const (
	CV_FONT_HERSHEY_SIMPLEX        = int(C.CV_FONT_HERSHEY_SIMPLEX)
	CV_FONT_HERSHEY_PLAIN          = int(C.CV_FONT_HERSHEY_PLAIN)
	CV_FONT_HERSHEY_DUPLEX         = int(C.CV_FONT_HERSHEY_DUPLEX)
	CV_FONT_HERSHEY_COMPLEX        = int(C.CV_FONT_HERSHEY_COMPLEX)
	CV_FONT_HERSHEY_TRIPLEX        = int(C.CV_FONT_HERSHEY_TRIPLEX)
	CV_FONT_HERSHEY_COMPLEX_SMALL  = int(C.CV_FONT_HERSHEY_COMPLEX_SMALL)
	CV_FONT_HERSHEY_SCRIPT_SIMPLEX = int(C.CV_FONT_HERSHEY_SCRIPT_SIMPLEX)
	CV_FONT_HERSHEY_SCRIPT_COMPLEX = int(C.CV_FONT_HERSHEY_SCRIPT_COMPLEX)
	CV_FONT_ITALIC                 = int(C.CV_FONT_ITALIC)
)

type Font struct {
	font C.CvFont
}

//void cvInitFont(CvFont* font, int font_face, double hscale, double vscale, double shear=0, int thickness=1, int line_type=8 )
func InitFont(fontFace int, hscale, vscale, shear float32, thickness, lineType int) *Font {
	font := new(Font)
	C.cvInitFont(
		&font.font,
		C.int(fontFace),
		C.double(hscale),
		C.double(vscale),
		C.double(shear),
		C.int(thickness),
		C.int(lineType),
	)
	return font
}

// void cvPutText(CvArr* img, const char* text, CvPoint org, const CvFont* font, CvScalar color)
func (this *Font) PutText(image *IplImage, text string, pt1 Point, color Scalar) {
	C.cvPutText(
		unsafe.Pointer(image),
		C.CString(text),
		C.cvPoint(C.int(pt1.X), C.int(pt1.Y)),
		&this.font,
		(C.CvScalar)(color),
	)
}

//CVAPI(void)  cvLine( CvArr* img, CvPoint pt1, CvPoint pt2,
//                     CvScalar color, int thickness CV_DEFAULT(1),
//                     int line_type CV_DEFAULT(8), int shift CV_DEFAULT(0) );

/****************************************************************************************\
*                                    System functions                         *
\****************************************************************************************/

/****************************************************************************************\
*                                    Data Persistence                         *
\****************************************************************************************/
