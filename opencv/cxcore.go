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

/* dst(idx) = ~src(idx) */
func Not(src, dst *IplImage) {
	C.cvNot(unsafe.Pointer(src), unsafe.Pointer(dst))
}

//CVAPI(void) cvNot( const CvArr* src, CvArr* dst );

/****************************************************************************************\
*                                Math operations                              *
\****************************************************************************************/

/****************************************************************************************\
*                                Matrix operations                            *
\****************************************************************************************/

/****************************************************************************************\
*                                    Array Statistics                         *
\****************************************************************************************/

/****************************************************************************************\
*                      Discrete Linear Transforms and Related Functions       *
\****************************************************************************************/

/****************************************************************************************\
*                              Dynamic data structures                        *
\****************************************************************************************/

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

//CVAPI(void)  cvLine( CvArr* img, CvPoint pt1, CvPoint pt2,
//                     CvScalar color, int thickness CV_DEFAULT(1),
//                     int line_type CV_DEFAULT(8), int shift CV_DEFAULT(0) );

/****************************************************************************************\
*                                    System functions                         *
\****************************************************************************************/

/****************************************************************************************\
*                                    Data Persistence                         *
\****************************************************************************************/
