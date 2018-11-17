// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

/*
#include "opencv.h"
#include <stdlib.h>
#include <string.h>

//-----------------------------------------------------------------------------

// version
// const char* CV_VERSION_ = CV_VERSION;

//-----------------------------------------------------------------------------

// IplImage
static int IPL_IMAGE_MAGIC_VAL_() {
	return IPL_IMAGE_MAGIC_VAL;
}
static const char* CV_TYPE_NAME_IMAGE_() {
	return CV_TYPE_NAME_IMAGE;
}
static int CV_IS_IMAGE_HDR_(void* img) {
	return CV_IS_IMAGE_HDR(img);
}
static int CV_IS_IMAGE_(void* img) {
	return CV_IS_IMAGE(img);
}

//-----------------------------------------------------------------------------

// type filed is go keyword
static int myGetMatType(const CvMat* mat) {
	return mat->type;
}
static char* myGetData(const CvMat* mat) {
	return  (char*)mat->data.ptr;
}
static int myGetMatNDType(const CvMatND* mat) {
	return mat->type;
}
static int myGetSparseMatType(const CvSparseMat* mat) {
	return mat->type;
}
static int myGetTermCriteriaType(const CvTermCriteria* x) {
	return x->type;
}

//-----------------------------------------------------------------------------
*/
import "C"
import (
	"math"
	"unsafe"
)

//-----------------------------------------------------------------------------

// cvver.h

//-----------------------------------------------------------------------------

const (
	CV_MAJOR_VERSION    = int(C.CV_MAJOR_VERSION)
	CV_MINOR_VERSION    = int(C.CV_MINOR_VERSION)
	CV_SUBMINOR_VERSION = int(C.CV_SUBMINOR_VERSION)
)

var (
//CV_VERSION = C.GoString(C.CV_VERSION_)
)

//-----------------------------------------------------------------------------
// cxerror.h
//-----------------------------------------------------------------------------

const (
	CV_StsOk                     = C.CV_StsOk
	CV_StsBackTrace              = C.CV_StsBackTrace
	CV_StsError                  = C.CV_StsError
	CV_StsInternal               = C.CV_StsInternal
	CV_StsNoMem                  = C.CV_StsNoMem
	CV_StsBadArg                 = C.CV_StsBadArg
	CV_StsBadFunc                = C.CV_StsBadFunc
	CV_StsNoConv                 = C.CV_StsNoConv
	CV_StsAutoTrace              = C.CV_StsAutoTrace
	CV_HeaderIsNull              = C.CV_HeaderIsNull
	CV_BadImageSize              = C.CV_BadImageSize
	CV_BadOffset                 = C.CV_BadOffset
	CV_BadDataPtr                = C.CV_BadDataPtr
	CV_BadStep                   = C.CV_BadStep
	CV_BadModelOrChSeq           = C.CV_BadModelOrChSeq
	CV_BadNumChannels            = C.CV_BadNumChannels
	CV_BadNumChannel1U           = C.CV_BadNumChannel1U
	CV_BadDepth                  = C.CV_BadDepth
	CV_BadAlphaChannel           = C.CV_BadAlphaChannel
	CV_BadOrder                  = C.CV_BadOrder
	CV_BadOrigin                 = C.CV_BadOrigin
	CV_BadAlign                  = C.CV_BadAlign
	CV_BadCallBack               = C.CV_BadCallBack
	CV_BadTileSize               = C.CV_BadTileSize
	CV_BadCOI                    = C.CV_BadCOI
	CV_BadROISize                = C.CV_BadROISize
	CV_MaskIsTiled               = C.CV_MaskIsTiled
	CV_StsNullPtr                = C.CV_StsNullPtr
	CV_StsVecLengthErr           = C.CV_StsVecLengthErr
	CV_StsFilterStructContentErr = C.CV_StsFilterStructContentErr
	CV_StsKernelStructContentErr = C.CV_StsKernelStructContentErr
	CV_StsFilterOffsetErr        = C.CV_StsFilterOffsetErr
	CV_StsBadSize                = C.CV_StsBadSize
	CV_StsDivByZero              = C.CV_StsDivByZero
	CV_StsInplaceNotSupported    = C.CV_StsInplaceNotSupported
	CV_StsObjectNotFound         = C.CV_StsObjectNotFound
	CV_StsUnmatchedFormats       = C.CV_StsUnmatchedFormats
	CV_StsBadFlag                = C.CV_StsBadFlag
	CV_StsBadPoint               = C.CV_StsBadPoint
	CV_StsBadMask                = C.CV_StsBadMask
	CV_StsUnmatchedSizes         = C.CV_StsUnmatchedSizes
	CV_StsUnsupportedFormat      = C.CV_StsUnsupportedFormat
	CV_StsOutOfRange             = C.CV_StsOutOfRange
	CV_StsParseError             = C.CV_StsParseError
	CV_StsNotImplemented         = C.CV_StsNotImplemented
	CV_StsBadMemBlock            = C.CV_StsBadMemBlock
	CV_StsAssert                 = C.CV_StsAssert
	//CV_GpuNotSupported=          C.CV_GpuNotSupported
	//CV_GpuApiCallError=          C.CV_GpuApiCallError
	//CV_GpuNppCallError=          C.CV_GpuNppCallError
)

//-----------------------------------------------------------------------------
// cxtypes.h
//-----------------------------------------------------------------------------

type Arr unsafe.Pointer

/*****************************************************************************\
*                      Common macros and inline functions                     *
\*****************************************************************************/

const (
	CV_PI   = 3.1415926535897932384626433832795
	CV_LOG2 = 0.69314718055994530941723212145818
)

func Round(value float64) int {
	rv := C.cvRound(C.double(value))
	return int(rv)
}
func Floor(value float64) int {
	rv := C.cvFloor(C.double(value))
	return int(rv)
}
func Ceil(value float64) int {
	rv := C.cvCeil(C.double(value))
	return int(rv)
}

func IsNaN(value float64) int {
	rv := C.cvIsNaN(C.double(value))
	return int(rv)
}
func IsInf(value float64) int {
	rv := C.cvIsInf(C.double(value))
	return int(rv)
}

/*************** Random number generation *******************/

type RNG C.CvRNG

func NewRNG(seed int64) RNG {
	rv := C.cvRNG(C.int64(seed))
	return RNG(rv)
}
func (rng *RNG) RandInt() uint32 {
	rv := C.cvRandInt((*C.CvRNG)(rng))
	return uint32(rv)
}
func (rng *RNG) RandReal() float64 {
	rv := C.cvRandReal((*C.CvRNG)(rng))
	return float64(rv)
}

/*****************************************************************************\
*                            Image type (IplImage)                            *
\*****************************************************************************/

/*
 * The following definitions (until #endif)
 * is an extract from IPL headers.
 * Copyright (c) 1995 Intel Corporation.
 */
const (
	IPL_DEPTH_SIGN = C.IPL_DEPTH_SIGN

	IPL_DEPTH_1U  = C.IPL_DEPTH_1U
	IPL_DEPTH_8U  = C.IPL_DEPTH_8U
	IPL_DEPTH_16U = C.IPL_DEPTH_16U
	IPL_DEPTH_32F = C.IPL_DEPTH_32F

	IPL_DEPTH_8S  = C.IPL_DEPTH_8S
	IPL_DEPTH_16S = C.IPL_DEPTH_16S
	IPL_DEPTH_32S = C.IPL_DEPTH_32S

	IPL_DATA_ORDER_PIXEL = C.IPL_DATA_ORDER_PIXEL
	IPL_DATA_ORDER_PLANE = C.IPL_DATA_ORDER_PLANE

	IPL_ORIGIN_TL = C.IPL_ORIGIN_TL
	IPL_ORIGIN_BL = C.IPL_ORIGIN_BL

	IPL_ALIGN_4BYTES  = C.IPL_ALIGN_4BYTES
	IPL_ALIGN_8BYTES  = C.IPL_ALIGN_8BYTES
	IPL_ALIGN_16BYTES = C.IPL_ALIGN_16BYTES
	IPL_ALIGN_32BYTES = C.IPL_ALIGN_32BYTES

	IPL_ALIGN_DWORD = C.IPL_ALIGN_DWORD
	IPL_ALIGN_QWORD = C.IPL_ALIGN_QWORD

	IPL_BORDER_CONSTANT  = C.IPL_BORDER_CONSTANT
	IPL_BORDER_REPLICATE = C.IPL_BORDER_REPLICATE
	IPL_BORDER_REFLECT   = C.IPL_BORDER_REFLECT
	IPL_BORDER_WRAP      = C.IPL_BORDER_WRAP
)

type IplImage C.IplImage

// normal fields
func (img *IplImage) Channels() int {
	return int(img.nChannels)
}
func (img *IplImage) Depth() int {
	return int(img.depth)
}
func (img *IplImage) Origin() int {
	return int(img.origin)
}
func (img *IplImage) Width() int {
	return int(img.width)
}
func (img *IplImage) Height() int {
	return int(img.height)
}
func (img *IplImage) WidthStep() int {
	return int(img.widthStep)
}
func (img *IplImage) ImageSize() int {
	return int(img.imageSize)
}
func (img *IplImage) ImageData() unsafe.Pointer {
	return unsafe.Pointer(img.imageData)
}

type IplROI C.IplROI

func (roi *IplROI) Init(coi, xOffset, yOffset, width, height int) {
	roi_c := (*C.IplROI)(roi)
	roi_c.coi = C.int(coi)
	roi_c.xOffset = C.int(xOffset)
	roi_c.yOffset = C.int(yOffset)
	roi_c.width = C.int(width)
	roi_c.height = C.int(height)
}
func (roi *IplROI) Coi() int {
	roi_c := (*C.IplROI)(roi)
	return int(roi_c.coi)
}
func (roi *IplROI) XOffset() int {
	roi_c := (*C.IplROI)(roi)
	return int(roi_c.xOffset)
}
func (roi *IplROI) YOffset() int {
	roi_c := (*C.IplROI)(roi)
	return int(roi_c.yOffset)
}
func (roi *IplROI) Width() int {
	roi_c := (*C.IplROI)(roi)
	return int(roi_c.width)
}
func (roi *IplROI) Height() int {
	roi_c := (*C.IplROI)(roi)
	return int(roi_c.height)
}

type IplConvKernel C.IplConvKernel
type IplConvKernelFP C.IplConvKernelFP

const (
	IPL_IMAGE_HEADER = C.IPL_IMAGE_HEADER
	IPL_IMAGE_DATA   = C.IPL_IMAGE_DATA
	IPL_IMAGE_ROI    = C.IPL_IMAGE_ROI
)

/* extra border mode */
var (
	IPL_IMAGE_MAGIC_VAL = C.IPL_IMAGE_MAGIC_VAL_()
	CV_TYPE_NAME_IMAGE  = C.CV_TYPE_NAME_IMAGE_()
)

func CV_IS_IMAGE_HDR(img unsafe.Pointer) bool {
	rv := C.CV_IS_IMAGE_HDR_(img)
	return (int(rv) != 0)
}
func CV_IS_IMAGE(img unsafe.Pointer) bool {
	rv := C.CV_IS_IMAGE_(img)
	return (int(rv) != 0)
}

const (
	IPL_DEPTH_64F = int(C.IPL_DEPTH_64F)
)

/****************************************************************************************\
*                                  Matrix type (CvMat)                                   *
\****************************************************************************************/

type Mat C.CvMat

func (mat *Mat) Type() int {
	return int(C.myGetMatType((*C.CvMat)(mat)))
}
func (mat *Mat) GetData() []byte {
	return C.GoBytes(unsafe.Pointer(C.myGetData((*C.CvMat)(mat))), C.int(mat.step))
}
func (mat *Mat) Step() int {
	return int(mat.step)
}

func (mat *Mat) Rows() int {
	return int(mat.rows)
}
func (mat *Mat) Cols() int {
	return int(mat.cols)
}

func CV_IS_MAT_HDR(mat interface{}) bool {
	return false
}
func CV_IS_MAT(mat interface{}) bool {
	return false
}
func CV_IS_MASK_ARR() bool {
	return false
}
func CV_ARE_TYPE_EQ() bool {
	return false
}

func (m *Mat) Init(rows, cols int, type_ int, data unsafe.Pointer) {
	return
}
func (m *Mat) Get(row, col int) float64 {
	rv := C.cvmGet((*C.CvMat)(m), C.int(row), C.int(col))
	return float64(rv)
}
func (m *Mat) Set(row, col int, value float64) {
	C.cvmSet((*C.CvMat)(m), C.int(row), C.int(col), C.double(value))
}

func IplDepth(type_ int) int {
	rv := C.cvIplDepth(C.int(type_))
	return int(rv)
}

/****************************************************************************************\
*                       Multi-dimensional dense array (CvMatND)                          *
\****************************************************************************************/

const (
	CV_MATND_MAGIC_VAL = C.CV_MATND_MAGIC_VAL
	CV_TYPE_NAME_MATND = C.CV_TYPE_NAME_MATND

	CV_MAX_DIM      = C.CV_MAX_DIM
	CV_MAX_DIM_HEAP = C.CV_MAX_DIM_HEAP
)

type MatND C.CvMatND

func (m *MatND) Type() int {
	rv := C.myGetMatNDType((*C.CvMatND)(m))
	return int(rv)
}
func (m *MatND) Dims() int {
	rv := m.dims
	return int(rv)
}

/****************************************************************************************\
*                      Multi-dimensional sparse array (CvSparseMat)                      *
\****************************************************************************************/

const (
	CV_SPARSE_MAT_MAGIC_VAL = C.CV_SPARSE_MAT_MAGIC_VAL
	CV_TYPE_NAME_SPARSE_MAT = C.CV_TYPE_NAME_SPARSE_MAT
)

type SparseMat C.CvSparseMat

func (m *SparseMat) Type() int {
	rv := C.myGetSparseMatType((*C.CvSparseMat)(m))
	return int(rv)
}
func (m *SparseMat) Dims() int {
	rv := m.dims
	return int(rv)
}

/**************** iteration through a sparse array *****************/

type SparseNode C.CvSparseNode

func (node *SparseNode) HashVal() uint32 {
	rv := node.hashval
	return uint32(rv)
}
func (node *SparseNode) Next() *SparseNode {
	rv := node.next
	return (*SparseNode)(rv)
}

type SparseMatIterator C.CvSparseMatIterator

func (node *SparseMatIterator) Mat() *SparseMat {
	rv := node.mat
	return (*SparseMat)(rv)
}
func (node *SparseMatIterator) Node() *SparseNode {
	rv := node.node
	return (*SparseNode)(rv)
}
func (node *SparseMatIterator) CurIdx() int {
	rv := node.curidx
	return (int)(rv)
}

/****************************************************************************************\
*                                         Histogram                                      *
\****************************************************************************************/

type HistType C.CvHistType

const (
	CV_HIST_MAGIC_VAL    = C.CV_HIST_MAGIC_VAL
	CV_HIST_UNIFORM_FLAG = C.CV_HIST_UNIFORM_FLAG

	/* indicates whether bin ranges are set already or not */
	CV_HIST_RANGES_FLAG = C.CV_HIST_RANGES_FLAG

	CV_HIST_ARRAY  = C.CV_HIST_ARRAY
	CV_HIST_SPARSE = C.CV_HIST_SPARSE
	CV_HIST_TREE   = C.CV_HIST_TREE

	/* should be used as a parameter only,
	   it turns to CV_HIST_UNIFORM_FLAG of hist->type */
	CV_HIST_UNIFORM = C.CV_HIST_UNIFORM
)

type Histogram C.CvHistogram

func CV_IS_HIST() bool {
	return false
}
func CV_IS_UNIFORM_HIST() bool {
	return false
}
func CV_IS_SPARSE_HIST() bool {
	return false
}
func CV_HIST_HAS_RANGES() bool {
	return false
}

/****************************************************************************************\
*                      Other supplementary data type definitions                         *
\****************************************************************************************/

/*************************************** CvRect *****************************************/

type Rect C.CvRect

func NewRect(x, y, w, h int) Rect {
	return Rect{C.int(x), C.int(y), C.int(w), C.int(h)}
}

func (r *Rect) Init(x, y, w, h int) {
	r.x = C.int(x)
	r.y = C.int(y)
	r.width = C.int(w)
	r.height = C.int(h)
}
func (r *Rect) X() int {
	r_c := (*C.CvRect)(r)
	return int(r_c.x)
}
func (r *Rect) Y() int {
	r_c := (*C.CvRect)(r)
	return int(r_c.y)
}
func (r *Rect) Width() int {
	r_c := (*C.CvRect)(r)
	return int(r_c.width)
}
func (r *Rect) Height() int {
	r_c := (*C.CvRect)(r)
	return int(r_c.height)
}

func (r *Rect) ToROI(coi int) IplROI {
	r_c := (*C.CvRect)(r)
	return (IplROI)(C.cvRectToROI(*r_c, C.int(coi)))
}

func (roi *IplROI) ToRect() Rect {
	r := C.cvRect(
		C.int(roi.XOffset()),
		C.int(roi.YOffset()),
		C.int(roi.Width()),
		C.int(roi.Height()),
	)
	return Rect(r)
}

// Returns the Top-Left Point of the rectangle
func (r *Rect) TL() Point {
	return Point{int(r.x), int(r.y)}
}

// Returns the Bottom-Right Point of the rectangle
func (r *Rect) BR() Point {
	return Point{int(r.x) + int(r.width), int(r.y) + int(r.height)}
}

/*********************************** CvTermCriteria *************************************/

const (
	CV_TERMCRIT_ITER   = C.CV_TERMCRIT_ITER
	CV_TERMCRIT_NUMBER = C.CV_TERMCRIT_NUMBER
	CV_TERMCRIT_EPS    = C.CV_TERMCRIT_EPS
)

type TermCriteria C.CvTermCriteria

func (x *TermCriteria) Init(type_, max_iter int, epsilon float64) {
	rv := C.cvTermCriteria(C.int(type_), C.int(max_iter), C.double(epsilon))
	(*x) = (TermCriteria)(rv)
}

func (x *TermCriteria) Type() int {
	rv := C.myGetTermCriteriaType((*C.CvTermCriteria)(x))
	return int(rv)
}
func (x *TermCriteria) MaxIter() int {
	rv := x.max_iter
	return int(rv)
}
func (x *TermCriteria) Epsilon() float64 {
	rv := x.epsilon
	return float64(rv)
}

/******************************* CvPoint and variants ***********************************/

/********************************* Point Interfaces *************************************/

type Point2D interface {
	Radius() float64   // the radius to the point
	RadiusSq() float64 // the radius to the point squared
	Angle() float64    // the azmuith angle
	ToPoint() Point    // An conversion to an integer Point type
}

type Point3D interface {
	Radius() float64   // the radius to the point
	RadiusSq() float64 // the radius to the point squared
	AzAngle() float64  // the azmuith angle
	IncAngle() float64 // the inclination angle from the z direction
}

/************************************* CvPoint Structs *********************************/

type CvPoint C.CvPoint

// returns a new CvPoint
func NewCvPoint(x, y int) CvPoint {
	return CvPoint{C.int(x), C.int(y)}
}

// Returns a Point
func (p CvPoint) ToPoint() Point {
	return Point{int(p.x), int(p.y)}
}

type CvPoint2D32f C.CvPoint2D32f

// returns a new CvPoint
func NewCvPoint2D32f(x, y float32) CvPoint2D32f {
	return CvPoint2D32f{C.float(x), C.float(y)}
}

// Returns a Point
func (p CvPoint2D32f) ToPoint() Point2D32f {
	return Point2D32f{float32(p.x), float32(p.y)}
}

/************************************* Point *******************************************/

type Point struct {
	X int
	Y int
}

func (p Point) ToCvPoint() CvPoint {
	return NewCvPoint(p.X, p.Y)
}

func (p Point) ToPoint() Point {
	return p
}

func (p Point) Add(p2 Point) Point {
	p.X += p2.X
	p.Y += p2.Y
	return p
}

func (p Point) Sub(p2 Point) Point {
	p.X -= p2.X
	p.Y -= p2.Y
	return p
}

func (p Point) Radius() float64 {
	return math.Sqrt(p.RadiusSq())
}

func (p Point) RadiusSq() float64 {
	return float64(p.X*p.X + p.Y*p.Y)
}

func (p Point) Angle() float64 {
	return math.Atan2(float64(p.Y), float64(p.X))
}

/************************************* Point2D32f **************************************/

type Point2D32f struct {
	X float32
	Y float32
}

func (p Point2D32f) ToPoint() Point {
	return Point{int(p.X), int(p.Y)}
}

func (p Point2D32f) ToCvPoint() CvPoint2D32f {
	return NewCvPoint2D32f(p.X, p.Y)
}

func (p Point2D32f) Add(p2 Point2D32f) Point2D32f {
	p.X += p2.X
	p.Y += p2.Y
	return p
}

func (p Point2D32f) Sub(p2 Point2D32f) Point2D32f {
	p.X -= p2.X
	p.Y -= p2.Y
	return p
}

func (p Point2D32f) Radius() float64 {
	return math.Sqrt(p.RadiusSq())
}

func (p Point2D32f) RadiusSq() float64 {
	return float64(p.X*p.X + p.Y*p.Y)
}

func (p Point2D32f) Angle() float64 {
	return math.Atan2(float64(p.Y), float64(p.X))
}

/************************************* Point3D32f **************************************/

type Point3D32f struct {
	X float32
	Y float32
	Z float32
}

func (p Point3D32f) Add(p2 Point3D32f) Point3D32f {
	p.X += p2.X
	p.Y += p2.Y
	p.Z += p2.Z
	return p
}

func (p Point3D32f) Sub(p2 Point3D32f) Point3D32f {
	p.X -= p2.X
	p.Y -= p2.Y
	p.Z -= p2.Z
	return p
}

func (p Point3D32f) Radius() float64 {
	return math.Sqrt(p.RadiusSq())
}

func (p Point3D32f) RadiusSq() float64 {
	return float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p Point3D32f) AzAngle() float64 {
	return math.Atan2(float64(p.Y), float64(p.X))
}

func (p Point3D32f) IncAngle() float64 {
	if p.Radius() == 0 {
		return 0
	}
	return math.Acos(float64(p.Z) / p.Radius())
}

/************************************* Point2D64f **************************************/

type Point2D64f struct {
	X float64
	Y float64
}

func (p Point2D64f) ToPoint() Point {
	return Point{int(p.X), int(p.Y)}
}

func (p Point2D64f) Add(p2 Point2D64f) Point2D64f {
	p.X += p2.X
	p.Y += p2.Y
	return p
}

func (p Point2D64f) Sub(p2 Point2D64f) Point2D64f {
	p.X -= p2.X
	p.Y -= p2.Y
	return p
}

func (p Point2D64f) Radius() float64 {
	return math.Sqrt(p.RadiusSq())
}

func (p Point2D64f) RadiusSq() float64 {
	return p.X*p.X + p.Y*p.Y
}

func (p Point2D64f) Angle() float64 {
	return math.Atan2(p.Y, p.X)
}

/************************************* Point3D64f **************************************/

type Point3D64f struct {
	X float64
	Y float64
	Z float64
}

func (p Point3D64f) Add(p2 Point3D64f) Point3D64f {
	p.X += p2.X
	p.Y += p2.Y
	p.Z += p2.Z
	return p
}

func (p Point3D64f) Sub(p2 Point3D64f) Point3D64f {
	p.X -= p2.X
	p.Y -= p2.Y
	p.Z -= p2.Z
	return p
}

func (p Point3D64f) Radius() float64 {
	return math.Sqrt(p.RadiusSq())
}

func (p Point3D64f) RadiusSq() float64 {
	return p.X*p.X + p.Y*p.Y + p.Z*p.Z
}

func (p Point3D64f) AzAngle() float64 {
	return math.Atan2(p.Y, p.X)
}

func (p Point3D64f) IncAngle() float64 {
	if p.Radius() == 0 {
		return 0
	}
	return math.Acos(p.Z / p.Radius())
}

/******************************** CvSize's & CvBox **************************************/

type Size struct {
	Width  int
	Height int
}

type Size2D32f struct {
	Width  float32
	Height float32
}

type Box2D struct {
	center Point2D32f
	size   Size2D32f
	angle  float32
}

func (box *Box2D) Size() Size2D32f {
	return box.size
}

func (box *Box2D) Center() Point2D32f {
	return box.center
}

func (box *Box2D) Angle() float32 {
	return box.angle
}

// Returns a CvBox2D
func (box *Box2D) CVBox() C.CvBox2D {
	var cvBox C.CvBox2D
	cvBox.angle = C.float(box.angle)
	cvBox.center.x = C.float(box.center.X)
	cvBox.center.y = C.float(box.center.Y)
	cvBox.size.width = C.float(box.size.Width)
	cvBox.size.height = C.float(box.size.Height)
	return cvBox
}

// Finds box vertices
func (box *Box2D) Points() []Point2D32f {
	var pts [4]C.CvPoint2D32f
	C.cvBoxPoints(
		box.CVBox(),
		(*C.CvPoint2D32f)(unsafe.Pointer(&pts[0])),
	)
	outPts := make([]Point2D32f, 4)
	for i, p := range pts {
		outPts[i].X = float32(p.x)
		outPts[i].Y = float32(p.y)
	}
	return outPts
}

type LineIterator C.CvLineIterator

/************************************* CvSlice ******************************************/

type Slice C.CvSlice

const (
	CV_WHOLE_SEQ_END_INDEX = C.CV_WHOLE_SEQ_END_INDEX
)

/* Equivalent to the C constant CV_WHOLE_SEQ */
func WholeSeq() Slice {
	slice := C.cvSlice(C.int(0), C.CV_WHOLE_SEQ_END_INDEX)
	return (Slice)(slice)
}

/************************************* CvScalar *****************************************/

type Scalar C.CvScalar

func NewScalar(b, g, r, a float64) Scalar {
	rv := C.cvScalar(C.double(b), C.double(g), C.double(r), C.double(a))
	return (Scalar)(rv)
}

func ScalarAll(val0 float64) Scalar {
	rv := C.cvScalarAll(C.double(val0))
	return (Scalar)(rv)
}

/* Val returns an array with the scalars values. */
func (s Scalar) Val() [4]float64 {
	return [4]float64{
		float64(s.val[0]),
		float64(s.val[1]),
		float64(s.val[2]),
		float64(s.val[3]),
	}
}

/****************************************************************************************\
*                                   Dynamic Data structures                              *
\****************************************************************************************/

/******************************** Memory storage ****************************************/

type MemBlock C.CvMemBlock
type MemStorage C.CvMemStorage
type MemStoragePos C.CvMemStoragePos

/*********************************** Sequence *******************************************/

type SeqBlock C.CvSeqBlock
type Seq C.CvSeq

/*************************************** Set ********************************************/

type Set C.CvSet

/************************************* Graph ********************************************/

type GraphEdge C.CvGraphEdge
type GraphVtx C.CvGraphVtx

type GraphVtx2D C.CvGraphVtx2D
type Graph C.CvGraph

/*********************************** Chain/Countour *************************************/

type Chain C.CvChain
type Contour C.CvContour

const (
	CV_RETR_EXTERNAL = C.CV_RETR_EXTERNAL
	CV_RETR_LIST     = C.CV_RETR_LIST
	CV_RETR_CCOMP    = C.CV_RETR_CCOMP
	CV_RETR_TREE     = C.CV_RETR_TREE

	CV_CHAIN_APPROX_NONE      = C.CV_CHAIN_APPROX_NONE
	CV_CHAIN_APPROX_SIMPLE    = C.CV_CHAIN_APPROX_SIMPLE
	CV_CHAIN_APPROX_TC89_L1   = C.CV_CHAIN_APPROX_TC89_L1
	CV_CHAIN_APPROX_TC89_KCOS = C.CV_CHAIN_APPROX_TC89_KCOS
)

/****************************************************************************************\
*                                    Sequence types                                      *
\****************************************************************************************/

/****************************************************************************************/
/*                            Sequence writer & reader                                  */
/****************************************************************************************/

type SeqWriter C.CvSeqWriter
type SeqReader C.CvSeqReader

/****************************************************************************************/
/*                                Operations on sequences                               */
/****************************************************************************************/

/****************************************************************************************\
*             Data structures for persistence (a.k.a serialization) functionality        *
\****************************************************************************************/

/* "black box" file storage */
type FileStorage C.CvFileStorage

/* Storage flags: */
const (
	CV_STORAGE_READ         = C.CV_STORAGE_READ
	CV_STORAGE_WRITE        = C.CV_STORAGE_WRITE
	CV_STORAGE_WRITE_TEXT   = C.CV_STORAGE_WRITE_TEXT
	CV_STORAGE_WRITE_BINARY = C.CV_STORAGE_WRITE_BINARY
	CV_STORAGE_APPEND       = C.CV_STORAGE_APPEND
)

type AttrList C.CvAttrList

/****************************************************************************************/
/*                  Structural Analysis and Shape Descriptors                           */
/****************************************************************************************/

/* For use in ApproxPoly */
const (
	CV_POLY_APPROX_DP = C.CV_POLY_APPROX_DP
)

/*****************************************************************************\
*                                 --- END ---                                 *
\*****************************************************************************/
