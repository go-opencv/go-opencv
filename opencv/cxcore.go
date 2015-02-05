// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 22/11/2013 Updated by <mohamd.helala@gmail.com>.

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

/*
Reshape changes shape of the image without copying data. A value of `0` means
that channels or rows remain unchanged.
*/
func Reshape(img unsafe.Pointer, header *Mat, channels, rows int) *Mat {
	n := C.cvReshape(img, (*C.CvMat)(header), C.int(channels), C.int(rows))
	return (*Mat)(n)
}


/* Get1D return a specific element from a 1-dimensional matrix. */
func Get1D(img unsafe.Pointer, x int) Scalar {
	ret := C.cvGet1D(img, C.int(x))
	return Scalar(ret)
}
/* Get2D return a specific element from a 2-dimensional matrix. */
func Get2D(img unsafe.Pointer,x, y int) Scalar {
	ret := C.cvGet2D(img, C.int(y), C.int(x))
	return Scalar(ret)
}
/* Get3D return a specific element from a 3-dimensional matrix. */
func Get3D(img unsafe.Pointer,x, y, z int) Scalar {
	ret := C.cvGet3D(img, C.int(z), C.int(y), C.int(x))
	return Scalar(ret)
}
/* Set1D sets a particular element in the image */
func Set1D(img unsafe.Pointer, x int, value Scalar) {
	C.cvSet1D(img, C.int(x), (C.CvScalar)(value))
}
/* Set2D sets a particular element in the image */
func Set2D(img unsafe.Pointer, x, y int, value Scalar) {
	C.cvSet2D(img, C.int(y), C.int(x), (C.CvScalar)(value))
}
/* Set3D sets a particular element in the image */
func Set3D(img unsafe.Pointer, x, y, z int, value Scalar) {
	C.cvSet3D(img, C.int(z), C.int(y), C.int(x), (C.CvScalar)(value))
}


// 
//                        
//                        

/* Converts CvArr (IplImage or CvMat,...) to CvMat.
   If the last parameter is non-zero, function can
   convert multi(>2)-dimensional array to CvMat as long as
   the last array's dimension is continous. The resultant
   matrix will be have appropriate (a huge) number of rows 

   CVAPI(CvMat*) cvGetMat( const CvArr* arr, CvMat* header,
   int* coi CV_DEFAULT(NULL),int allowND CV_DEFAULT(0));*/
func GetMat(arr unsafe.Pointer, header *Mat, coi *int32, allowND int) *Mat{
	m := C.cvGetMat(arr, (*C.CvMat)(header), (*C.int)(coi), C.int(allowND))
	return (*Mat)(m)
}

/* Converts CvArr (IplImage or CvMat) to IplImage 

   CVAPI(IplImage*) cvGetImage( const CvArr* arr, 
   IplImage* image_header );*/
func GetImage(arr unsafe.Pointer, header *IplImage) *IplImage{
	m := C.cvGetImage(arr, (*C.IplImage)(header))
	return (*IplImage)(m)
}



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
func Copy(src, dst, mask unsafe.Pointer) {
	C.cvCopy(src, dst, mask)
}

//CVAPI(void)  cvCopy( const CvArr* src, CvArr* dst,
//                     const CvArr* mask CV_DEFAULT(NULL) );

/* Clears all the array elements (sets them to 0) */
func Zero(img unsafe.Pointer) {
	C.cvSetZero(img)
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

/* Calculates length of sequence slice (with support of negative indices). */

func SliceLength(slice Slice, seq *Seq) {
	C.cvSliceLength(C.CvSlice(slice), (*C.CvSeq)(seq))
}

// CVAPI(int) cvSliceLength( CvSlice slice, const CvSeq* seq );

/* Creates new memory storage.
   block_size == 0 means that default,
   somewhat optimal size, is used (currently, it is 64K) */
func CreateMemStorage(block_size int) *MemStorage {
	block := C.cvCreateMemStorage(C.int(block_size))
	return (*MemStorage)(block)
}

// CVAPI()  cvCreateMemStorage( int block_size CV_DEFAULT(0));

/* Creates a memory storage that will borrow memory blocks from parent storage */
func CreateChildMemStorage(parent *MemStorage) *MemStorage {
	block := C.cvCreateChildMemStorage((*C.CvMemStorage)(parent))
	return (*MemStorage)(block)
}

// CVAPI(CvMemStorage*)  cvCreateChildMemStorage( CvMemStorage* parent );

/* Releases memory storage. All the children of a parent must be released before
   the parent. A child storage returns all the blocks to parent when it is released */
func ReleaseMemStorage(storage *MemStorage) {
	block := (*C.CvMemStorage)(storage)
	C.cvReleaseMemStorage(&block)
}

// CVAPI(void)  cvReleaseMemStorage( CvMemStorage** storage );

/* Clears memory storage. This is the only way(!!!) (besides cvRestoreMemStoragePos)
   to reuse memory allocated for the storage - cvClearSeq,cvClearSet ...
   do not free any memory.
   A child storage returns all the blocks to the parent when it is cleared */
func ClearMemStorage(storage *MemStorage) {
	C.cvClearMemStorage((*C.CvMemStorage)(storage))
}

// CVAPI(void)  cvClearMemStorage( CvMemStorage* storage );

/* Remember a storage "free memory" position */
func SaveMemStoragePos(storage *MemStorage, pos *MemStoragePos) {
	C.cvSaveMemStoragePos((*C.CvMemStorage)(storage), (*C.CvMemStoragePos)(pos))
}

// CVAPI(void)  cvSaveMemStoragePos( const CvMemStorage* storage, CvMemStoragePos* pos );

/* Restore a storage "free memory" position */
func RestoreMemStoragePos(storage *MemStorage, pos *MemStoragePos) {
	C.cvRestoreMemStoragePos((*C.CvMemStorage)(storage), (*C.CvMemStoragePos)(pos))
}

// CVAPI(void)  cvRestoreMemStoragePos( CvMemStorage* storage, CvMemStoragePos* pos );

/* Allocates continuous buffer of the specified size in the storage */
func MemStorageAlloc(storage *MemStorage, size int) {
	C.cvMemStorageAlloc((*C.CvMemStorage)(storage), C.size_t(size))
}

// CVAPI(void*) cvMemStorageAlloc( CvMemStorage* storage, size_t size );

/* Allocates string in memory storage */
func MemStorageAllocString(storage *MemStorage, name string, _len int) {
	C.cvMemStorageAllocString((*C.CvMemStorage)(storage), C.CString(name), C.int(_len))
}

// CVAPI(CvString) cvMemStorageAllocString( CvMemStorage* storage, const char* ptr,
// int len CV_DEFAULT(-1) );

/* Creates new empty sequence that will reside in the specified storage */
func CreateSeq(seq_flags, header_size, elem_size int, storage *MemStorage) {
	C.cvCreateSeq(C.int(seq_flags), C.size_t(header_size), C.size_t(elem_size), (*C.CvMemStorage)(storage))
}

//CVAPI(CvSeq*)  cvCreateSeq( int seq_flags, size_t header_size,
//                            size_t elem_size, CvMemStorage* storage );

/* Removes all the elements from the sequence. The freed memory
   can be reused later only by the same sequence unless cvClearMemStorage
   or cvRestoreMemStoragePos is called */
func ClearSeq(seq *Seq) {
	C.cvClearSeq((*C.CvSeq)(seq))
}

// CVAPI(void)  cvClearSeq( CvSeq* seq );

/* Retrieves pointer to specified sequence element.
   Negative indices are supported and mean counting from the end
   (e.g -1 means the last sequence element) */
func GetSeqElem(seq *Seq, index int) *int8 {
	el := C.cvGetSeqElem((*C.CvSeq)(seq), C.int(index))
	return (*int8)(el)
}

// CVAPI(schar*)  cvGetSeqElem( const CvSeq* seq, int index );

/* Calculates index of the specified sequence element.
   Returns -1 if element does not belong to the sequence */
func SeqElemIdx(seq *Seq, el unsafe.Pointer, seqblock *SeqBlock) int {
	cvseqblock := (*C.CvSeqBlock)(seqblock)
	r := C.cvSeqElemIdx((*C.CvSeq)(seq), el, &cvseqblock)
	return (int)(r)
}

// CVAPI(int)  cvSeqElemIdx( const CvSeq* seq, const void* element,
//                         CvSeqBlock** block CV_DEFAULT(NULL) );

/* Extracts sequence slice (with or without copying sequence elements) */
func SeqSlice(seq *Seq, slice Slice, storage *MemStorage, copy_data int) *Seq {
	r := C.cvSeqSlice((*C.CvSeq)(seq), C.CvSlice(slice), (*C.CvMemStorage)(storage), C.int(copy_data))
	return (*Seq)(r)
}

// CVAPI(CvSeq*) cvSeqSlice( const CvSeq* seq, CvSlice slice,
//                          CvMemStorage* storage CV_DEFAULT(NULL),
//                          int copy_data CV_DEFAULT(0));
/* A wrapper for the somewhat more general routine cvSeqSlice() whuch
creates a deep copy of a sequence and creates another entirely separate
sequence structure.*/
func CloneSeq(seq *Seq, storage *MemStorage) *Seq {
	r := C.cvSeqSlice((*C.CvSeq)(seq), (C.CvSlice)(CV_WHOLE_SEQ), (*C.CvMemStorage)(storage), C.int(1))
	return (*Seq)(r)
}

// CV_INLINE CvSeq* cvCloneSeq( const CvSeq* seq, CvMemStorage* storage CV_DEFAULT(NULL))
// {
//     return cvSeqSlice( seq, CV_WHOLE_SEQ, storage, 1 );
// }

/* Removes sequence slice */
func SeqRemoveSlice(seq *Seq, slice Slice) {
	C.cvSeqRemoveSlice((*C.CvSeq)(seq), C.CvSlice(slice))
}

// CVAPI(void)  cvSeqRemoveSlice( CvSeq* seq, CvSlice slice );

/* Inserts a sequence or array into another sequence */
func SeqInsertSlice(seq *Seq, before_index int, from_arr Arr) {
	C.cvSeqInsertSlice((*C.CvSeq)(seq), C.int(before_index), unsafe.Pointer(from_arr))
}

// CVAPI(void)  cvSeqInsertSlice( CvSeq* seq, int before_index, const CvArr* from_arr );

/* Sorts sequence in-place given element comparison function */
func SeqSort(seq *Seq, f CmpFunc, userdata unsafe.Pointer) {
	fc := func(a C.CVoid, b C.CVoid, data unsafe.Pointer) int {
		return f(unsafe.Pointer(a), unsafe.Pointer(b), data)
	}
	cmpFunc := C.GoOpenCV_CmpFunc(unsafe.Pointer(&fc))
	C.cvSeqSort((*C.CvSeq)(seq), cmpFunc, unsafe.Pointer(userdata))
}

// CVAPI(void) cvSeqSort( CvSeq* seq, CvCmpFunc func, void* userdata CV_DEFAULT(NULL) );

/* Finds element in a [sorted] sequence */
func SeqSearch(seq *Seq, elem unsafe.Pointer, f CmpFunc, is_sorted int,
	lem_idx *int32, userdata unsafe.Pointer) {
	fc := func(a C.CVoid, b C.CVoid, userdata Arr) int {
		return f(unsafe.Pointer(a), unsafe.Pointer(b), unsafe.Pointer(userdata))
	}
	cmpFunc := C.GoOpenCV_CmpFunc(unsafe.Pointer(&fc))
	C.cvSeqSearch((*C.CvSeq)(seq), elem, cmpFunc,
		C.int(is_sorted), (*C.int)(lem_idx), userdata)
}

// CVAPI(schar*) cvSeqSearch( CvSeq* seq, const void* elem, CvCmpFunc func,
//                            int is_sorted, int* elem_idx,
//                            void* userdata CV_DEFAULT(NULL) );

/* Reverses order of sequence elements in-place */
func SeqInvert(seq *Seq) {
	C.cvSeqInvert((*C.CvSeq)(seq))
}

// CVAPI(void) cvSeqInvert( CvSeq* seq );

/* Splits sequence into one or more equivalence classes using the specified criteria */
func SeqPartition(seq *Seq, storage *MemStorage, labels **Seq,
	f CmpFunc, userdata unsafe.Pointer) {
	fc := func(a C.CVoid, b C.CVoid, userdata Arr) int {
		return f(unsafe.Pointer(a), unsafe.Pointer(b), unsafe.Pointer(userdata))
	}
	cmpFunc := C.GoOpenCV_CmpFunc(unsafe.Pointer(&fc))
	cvSeq := (*C.CvSeq)(*labels)
	C.cvSeqPartition((*C.CvSeq)(seq), (*C.CvMemStorage)(storage), &cvSeq,
		cmpFunc, userdata)
}

// CVAPI(int)  cvSeqPartition( const CvSeq* seq, CvMemStorage* storage,
//                             CvSeq** labels, CvCmpFunc is_equal, void*  );

/* Inserts a new element in the middle of sequence.
   cvSeqInsert(seq,0,elem) == cvSeqPushFront(seq,elem) */
func SeqInsert(seq *Seq, before_index int, element unsafe.Pointer) *int8 {
	r := C.cvSeqInsert((*C.CvSeq)(seq), C.int(before_index), element)
	return (*int8)(r)
}

// CVAPI(schar*)  cvSeqInsert( CvSeq* seq, int before_index,
//                            const void* element CV_DEFAULT(NULL));

/* Removes specified sequence element */
func SeqRemove(seq *Seq, index int) {
	C.cvSeqRemove((*C.CvSeq)(seq), C.int(index))
}

// CVAPI(void)  cvSeqRemove( CvSeq* seq, int index );

/* Changes default size (granularity) of sequence blocks.
   The default size is ~1Kbyte */
func SetSeqBlockSize(seq *Seq, delta_elems int) {
	C.cvSetSeqBlockSize((*C.CvSeq)(seq), C.int(delta_elems))
}

// CVAPI(void)  cvSetSeqBlockSize( CvSeq* seq, int delta_elems );

/* Copies sequence content to a continuous piece of memory */
func CvtSeqToArray(seq *Seq, delta_elems unsafe.Pointer, slice Slice) {
	C.cvCvtSeqToArray((*C.CvSeq)(seq), delta_elems, (C.CvSlice)(slice))
}

// CVAPI(void*)  cvCvtSeqToArray( const CvSeq* seq, void* elements,
//                                CvSlice slice CV_DEFAULT(CV_WHOLE_SEQ) );

/* Creates sequence header for array.
   After that all the operations on sequences that do not alter the content
   can be applied to the resultant sequence */
func MakeSeqHeaderForArray(seq_type, header_size, elem_size int,
	elems unsafe.Pointer, total int, seq *Seq, block *SeqBlock) *Seq {
	r := C.cvMakeSeqHeaderForArray(C.int(seq_type), C.int(header_size),
		C.int(elem_size), elems, C.int(total), (*C.CvSeq)(seq),
		(*C.CvSeqBlock)(block))
	return (*Seq)(r)
}

// CVAPI(CvSeq*) cvMakeSeqHeaderForArray( int seq_type, int header_size,
//                                        int elem_size, void* elements, int total,
//                                        CvSeq* seq, CvSeqBlock* block );
/****************************************************************************************\
*                                     Drawing                                 *
\****************************************************************************************/

const CV_FILLED = C.CV_FILLED

/* Draws 4-connected, 8-connected or antialiased line segment connecting two points */
//color Scalar,
func Line(image unsafe.Pointer, pt1, pt2 Point, color Scalar, thickness, line_type, shift int) {
	C.cvLine(image,
		C.cvPoint(C.int(pt1.X), C.int(pt1.Y)),
		C.cvPoint(C.int(pt2.X), C.int(pt2.Y)),
		(C.CvScalar)(color),
		C.int(thickness), C.int(line_type), C.int(shift),
	)
}

//CVAPI(void)  cvLine( CvArr* img, CvPoint pt1, CvPoint pt2,
//                     CvScalar color, int thickness CV_DEFAULT(1),
//                     int line_type CV_DEFAULT(8), int shift CV_DEFAULT(0) );

/* Draws a rectangle given two opposite corners of the rectangle (pt1 & pt2),
   if thickness<0 (e.g. thickness == CV_FILLED), the filled box is drawn */
func Rectangle(image unsafe.Pointer, pt1, pt2 Point, color Scalar, thickness, line_type, shift int) {
	C.cvRectangle(image,
		C.cvPoint(C.int(pt1.X), C.int(pt1.Y)),
		C.cvPoint(C.int(pt2.X), C.int(pt2.Y)),
		(C.CvScalar)(color),
		C.int(thickness), C.int(line_type), C.int(shift),
	)
}

// CVAPI(void)  cvRectangle( CvArr* img, CvPoint pt1, CvPoint pt2,
//                           CvScalar color, int thickness CV_DEFAULT(1),
//                           int line_type CV_DEFAULT(8),
//                           int shift CV_DEFAULT(0));

func Circle(image unsafe.Pointer, pt1 Point, r int, color Scalar, thickness, line_type, shift int) {
	C.cvCircle(image,
		C.cvPoint(C.int(pt1.X), C.int(pt1.Y)),
		C.int(r),
		(C.CvScalar)(color),
		C.int(thickness), C.int(line_type), C.int(shift),
	)
}

/* Draws ellipse outline, filled ellipse, elliptic arc or filled elliptic sector,
   depending on <thickness>, <start_angle> and <end_angle> parameters. The resultant figure
   is rotated by <angle>. All the angles are in degrees */
// CVAPI(void)  cvEllipse( CvArr* img, CvPoint center, CvSize axes,
//                         double angle, double start_angle, double end_angle,
//                         CvScalar color, int thickness CV_DEFAULT(1),
//                         int line_type CV_DEFAULT(8), int shift CV_DEFAULT(0));

func Ellipse(image unsafe.Pointer, center Point, axes Size, angle, start_angle, end_angle float64, 
					color Scalar, thickness, line_type, shift int) {
	C.cvEllipse(image,
		C.cvPoint(C.int(center.X), C.int(center.Y)),
		C.cvSize(C.int(axes.Width), C.int(axes.Height)), C.double(angle),
		C.double(start_angle), C.double(end_angle) ,(C.CvScalar)(color),
		C.int(thickness), C.int(line_type), C.int(shift),
	)
}

// CVAPI(void)  cvDrawContours( CvArr *img, CvSeq* contour,
//                              CvScalar external_color, CvScalar hole_color,
//                              int max_level, int thickness CV_DEFAULT(1),
//                              int line_type CV_DEFAULT(8),
//                              CvPoint offset CV_DEFAULT(cvPoint(0,0)));

/* Draws contour outlines or filled interiors on the image */
func DrawContours(image unsafe.Pointer, contour *Seq, external_color Scalar,
	hole_color Scalar, max_level, thickness, line_type int, offset Point) {
	C.cvDrawContours(image,
		(*C.CvSeq)(contour),
		(C.CvScalar)(external_color),
		(C.CvScalar)(hole_color),
		C.int(max_level), C.int(thickness), C.int(line_type),
		C.cvPoint(C.int(offset.X), C.int(offset.Y)),
	)
}

const (
	CV_FONT_HERSHEY_SIMPLEX  = C.CV_FONT_HERSHEY_SIMPLEX
	CV_FONT_HERSHEY_PLAIN    = C.CV_FONT_HERSHEY_PLAIN
	CV_FONT_HERSHEY_DUPLEX   = C.CV_FONT_HERSHEY_DUPLEX
	CV_FONT_HERSHEY_COMPLEX  = C.CV_FONT_HERSHEY_COMPLEX
	CV_FONT_HERSHEY_TRIPLEX  = C.CV_FONT_HERSHEY_TRIPLEX
	CV_FONT_HERSHEY_COMPLEX_SMALL   = C.CV_FONT_HERSHEY_COMPLEX_SMALL
	CV_FONT_HERSHEY_SCRIPT_SIMPLEX  = C.CV_FONT_HERSHEY_SCRIPT_SIMPLEX
	CV_FONT_HERSHEY_SCRIPT_COMPLEX  = C.CV_FONT_HERSHEY_SCRIPT_COMPLEX
)

/* Renders text stroke with specified font and color at specified location.
   CvFont should be initialized with cvInitFont */
// CVAPI(void)  cvPutText( CvArr* img, const char* text, CvPoint org,
//                         const CvFont* font, CvScalar color );
func PutText(image unsafe.Pointer, text string, org Point, font_face int, 
	 hscale, vscale, shear float64, thickness, line_type int, color Scalar){
	// create CvFont sturucture
	f := &C.CvFont{}
	C.cvInitFont(f, C.int(font_face), C.double(hscale), C.double(vscale), 
		C.double(shear), C.int(thickness), C.int(line_type))
	C.cvPutText(image, C.CString(text), C.cvPoint(C.int(org.X), C.int(org.Y)),
		f, (C.CvScalar)(color))
}

/****************************************************************************************\
*                                    System functions                         *
\****************************************************************************************/

/****************************************************************************************\
*                                    Data Persistence                         *
\****************************************************************************************/

/*********************************** Adding own types ***********************************/

/* universal functions */
func Release(ptr unsafe.Pointer) { C.cvRelease(&ptr) }

// CVAPI(void) cvRelease( void** struct_ptr );