// Copyright 2014 <mohamed.helala@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Bindings for Intel's OpenCV computer vision library.
package opencv

//#include "opencv.h"
//#cgo linux  pkg-config: opencv
//#cgo darwin pkg-config: opencv
//#cgo freebsd pkg-config: opencv
//#cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
import "C"

/* Connected component structure */
type ConnectedComp C.CvConnectedComp

/*
// {
//     double area;     area of the connected component
//     CvScalar value;  average color of the connected component
//     CvRect rect;     ROI of the component
//     CvSeq* contour;  optional component boundary
//                       (the contour might have child contours corresponding to the holes)
// }
*/

// /* Image smooth methods */
const (
	CV_BLUR_NO_SCALE = C.CV_BLUR_NO_SCALE
	CV_BLUR          = C.CV_BLUR
	CV_GAUSSIAN      = C.CV_GAUSSIAN
	CV_MEDIAN        = C.CV_MEDIAN
	CV_BILATERAL     = C.CV_BILATERAL
)

// /* Filters used in pyramid decomposition */
const (
	CV_GAUSSIAN_5x5 = C.CV_GAUSSIAN_5x5
)

// /* Special filters */
const (
	CV_SCHARR          = C.CV_SCHARR
	CV_MAX_SOBEL_KSIZE = C.CV_MAX_SOBEL_KSIZE
)

// /* Constants for color conversion */
const (
	CV_BGR2BGRA = C.CV_BGR2BGRA
	CV_RGB2RGBA = C.CV_RGB2RGBA

	CV_BGRA2BGR = C.CV_BGRA2BGR
	CV_RGBA2RGB = C.CV_RGBA2RGB

	CV_BGR2RGBA = C.CV_BGR2RGBA
	CV_RGB2BGRA = C.CV_RGB2BGRA

	CV_RGBA2BGR = C.CV_RGBA2BGR
	CV_BGRA2RGB = C.CV_BGRA2RGB

	CV_BGR2RGB = C.CV_BGR2RGB
	CV_RGB2BGR = C.CV_RGB2BGR

	CV_BGRA2RGBA = C.CV_BGRA2RGBA
	CV_RGBA2BGRA = C.CV_RGBA2BGRA

	CV_BGR2GRAY  = C.CV_BGR2GRAY
	CV_RGB2GRAY  = C.CV_RGB2GRAY
	CV_GRAY2BGR  = C.CV_GRAY2BGR
	CV_GRAY2RGB  = C.CV_GRAY2RGB
	CV_GRAY2BGRA = C.CV_GRAY2BGRA
	CV_GRAY2RGBA = C.CV_GRAY2RGBA
	CV_BGRA2GRAY = C.CV_BGRA2GRAY
	CV_RGBA2GRAY = C.CV_RGBA2GRAY

	CV_BGR2BGR565  = C.CV_BGR2BGR565
	CV_RGB2BGR565  = C.CV_RGB2BGR565
	CV_BGR5652BGR  = C.CV_BGR5652BGR
	CV_BGR5652RGB  = C.CV_BGR5652RGB
	CV_BGRA2BGR565 = C.CV_BGRA2BGR565
	CV_RGBA2BGR565 = C.CV_RGBA2BGR565
	CV_BGR5652BGRA = C.CV_BGR5652BGRA
	CV_BGR5652RGBA = C.CV_BGR5652RGBA

	CV_GRAY2BGR565 = C.CV_GRAY2BGR565
	CV_BGR5652GRAY = C.CV_BGR5652GRAY

	CV_BGR2BGR555  = C.CV_BGR2BGR555
	CV_RGB2BGR555  = C.CV_RGB2BGR555
	CV_BGR5552BGR  = C.CV_BGR5552BGR
	CV_BGR5552RGB  = C.CV_BGR5552RGB
	CV_BGRA2BGR555 = C.CV_BGRA2BGR555
	CV_RGBA2BGR555 = C.CV_RGBA2BGR555
	CV_BGR5552BGRA = C.CV_BGR5552BGRA
	CV_BGR5552RGBA = C.CV_BGR5552RGBA

	CV_GRAY2BGR555 = C.CV_GRAY2BGR555
	CV_BGR5552GRAY = C.CV_BGR5552GRAY

	CV_BGR2XYZ = C.CV_BGR2XYZ
	CV_RGB2XYZ = C.CV_RGB2XYZ
	CV_XYZ2BGR = C.CV_XYZ2BGR
	CV_XYZ2RGB = C.CV_XYZ2RGB

	CV_BGR2YCrCb = C.CV_BGR2YCrCb
	CV_RGB2YCrCb = C.CV_RGB2YCrCb
	CV_YCrCb2BGR = C.CV_YCrCb2BGR
	CV_YCrCb2RGB = C.CV_YCrCb2RGB

	CV_BGR2HSV = C.CV_BGR2HSV
	CV_RGB2HSV = C.CV_RGB2HSV

	CV_BGR2Lab = C.CV_BGR2Lab
	CV_RGB2Lab = C.CV_RGB2Lab

	CV_BayerBG2BGR = C.CV_BayerBG2BGR
	CV_BayerGB2BGR = C.CV_BayerGB2BGR
	CV_BayerRG2BGR = C.CV_BayerRG2BGR
	CV_BayerGR2BGR = C.CV_BayerGR2BGR

	CV_BayerBG2RGB = C.CV_BayerBG2RGB
	CV_BayerGB2RGB = C.CV_BayerGB2RGB
	CV_BayerRG2RGB = C.CV_BayerRG2RGB
	CV_BayerGR2RGB = C.CV_BayerGR2RGB

	CV_BGR2Luv = C.CV_BGR2Luv
	CV_RGB2Luv = C.CV_RGB2Luv
	CV_BGR2HLS = C.CV_BGR2HLS
	CV_RGB2HLS = C.CV_RGB2HLS

	CV_HSV2BGR = C.CV_HSV2BGR
	CV_HSV2RGB = C.CV_HSV2RGB

	CV_Lab2BGR = C.CV_Lab2BGR
	CV_Lab2RGB = C.CV_Lab2RGB
	CV_Luv2BGR = C.CV_Luv2BGR
	CV_Luv2RGB = C.CV_Luv2RGB
	CV_HLS2BGR = C.CV_HLS2BGR
	CV_HLS2RGB = C.CV_HLS2RGB

	CV_BayerBG2BGR_VNG = C.CV_BayerBG2BGR_VNG
	CV_BayerGB2BGR_VNG = C.CV_BayerGB2BGR_VNG
	CV_BayerRG2BGR_VNG = C.CV_BayerRG2BGR_VNG
	// CV_BayerGR2BGR_VNG =65,

	// CV_BayerBG2RGB_VNG =CV_BayerRG2BGR_VNG,
	// CV_BayerGB2RGB_VNG =CV_BayerGR2BGR_VNG,
	// CV_BayerRG2RGB_VNG =CV_BayerBG2BGR_VNG,
	// CV_BayerGR2RGB_VNG =CV_BayerGB2BGR_VNG,

	// CV_BGR2HSV_FULL = 66,
	// CV_RGB2HSV_FULL = 67,
	// CV_BGR2HLS_FULL = 68,
	// CV_RGB2HLS_FULL = 69,

	// CV_HSV2BGR_FULL = 70,
	// CV_HSV2RGB_FULL = 71,
	// CV_HLS2BGR_FULL = 72,
	// CV_HLS2RGB_FULL = 73,

	// CV_LBGR2Lab     = 74,
	// CV_LRGB2Lab     = 75,
	// CV_LBGR2Luv     = 76,
	// CV_LRGB2Luv     = 77,

	// CV_Lab2LBGR     = 78,
	// CV_Lab2LRGB     = 79,
	// CV_Luv2LBGR     = 80,
	// CV_Luv2LRGB     = 81,

	// CV_BGR2YUV      = 82,
	// CV_RGB2YUV      = 83,
	// CV_YUV2BGR      = 84,
	// CV_YUV2RGB      = 85,

	// CV_BayerBG2GRAY = 86,
	// CV_BayerGB2GRAY = 87,
	// CV_BayerRG2GRAY = 88,
	// CV_BayerGR2GRAY = 89,

	// //YUV 4:2:0 formats family
	// CV_YUV2RGB_NV12 = 90,
	// CV_YUV2BGR_NV12 = 91,
	// CV_YUV2RGB_NV21 = 92,
	// CV_YUV2BGR_NV21 = 93,
	// CV_YUV420sp2RGB = CV_YUV2RGB_NV21,
	// CV_YUV420sp2BGR = CV_YUV2BGR_NV21,

	// CV_YUV2RGBA_NV12 = 94,
	// CV_YUV2BGRA_NV12 = 95,
	// CV_YUV2RGBA_NV21 = 96,
	// CV_YUV2BGRA_NV21 = 97,
	// CV_YUV420sp2RGBA = CV_YUV2RGBA_NV21,
	// CV_YUV420sp2BGRA = CV_YUV2BGRA_NV21,

	// CV_YUV2RGB_YV12 = 98,
	// CV_YUV2BGR_YV12 = 99,
	// CV_YUV2RGB_IYUV = 100,
	// CV_YUV2BGR_IYUV = 101,
	// CV_YUV2RGB_I420 = CV_YUV2RGB_IYUV,
	// CV_YUV2BGR_I420 = CV_YUV2BGR_IYUV,
	// CV_YUV420p2RGB = CV_YUV2RGB_YV12,
	// CV_YUV420p2BGR = CV_YUV2BGR_YV12,

	// CV_YUV2RGBA_YV12 = 102,
	// CV_YUV2BGRA_YV12 = 103,
	// CV_YUV2RGBA_IYUV = 104,
	// CV_YUV2BGRA_IYUV = 105,
	// CV_YUV2RGBA_I420 = CV_YUV2RGBA_IYUV,
	// CV_YUV2BGRA_I420 = CV_YUV2BGRA_IYUV,
	// CV_YUV420p2RGBA = CV_YUV2RGBA_YV12,
	// CV_YUV420p2BGRA = CV_YUV2BGRA_YV12,

	// CV_YUV2GRAY_420 = 106,
	// CV_YUV2GRAY_NV21 = CV_YUV2GRAY_420,
	// CV_YUV2GRAY_NV12 = CV_YUV2GRAY_420,
	// CV_YUV2GRAY_YV12 = CV_YUV2GRAY_420,
	// CV_YUV2GRAY_IYUV = CV_YUV2GRAY_420,
	// CV_YUV2GRAY_I420 = CV_YUV2GRAY_420,
	// CV_YUV420sp2GRAY = CV_YUV2GRAY_420,
	// CV_YUV420p2GRAY = CV_YUV2GRAY_420,

	// //YUV 4:2:2 formats family
	// CV_YUV2RGB_UYVY = 107,
	// CV_YUV2BGR_UYVY = 108,
	// //CV_YUV2RGB_VYUY = 109,
	// //CV_YUV2BGR_VYUY = 110,
	// CV_YUV2RGB_Y422 = CV_YUV2RGB_UYVY,
	// CV_YUV2BGR_Y422 = CV_YUV2BGR_UYVY,
	// CV_YUV2RGB_UYNV = CV_YUV2RGB_UYVY,
	// CV_YUV2BGR_UYNV = CV_YUV2BGR_UYVY,

	// CV_YUV2RGBA_UYVY = 111,
	// CV_YUV2BGRA_UYVY = 112,
	// //CV_YUV2RGBA_VYUY = 113,
	// //CV_YUV2BGRA_VYUY = 114,
	// CV_YUV2RGBA_Y422 = CV_YUV2RGBA_UYVY,
	// CV_YUV2BGRA_Y422 = CV_YUV2BGRA_UYVY,
	// CV_YUV2RGBA_UYNV = CV_YUV2RGBA_UYVY,
	// CV_YUV2BGRA_UYNV = CV_YUV2BGRA_UYVY,

	// CV_YUV2RGB_YUY2 = 115,
	// CV_YUV2BGR_YUY2 = 116,
	// CV_YUV2RGB_YVYU = 117,
	// CV_YUV2BGR_YVYU = 118,
	// CV_YUV2RGB_YUYV = CV_YUV2RGB_YUY2,
	// CV_YUV2BGR_YUYV = CV_YUV2BGR_YUY2,
	// CV_YUV2RGB_YUNV = CV_YUV2RGB_YUY2,
	// CV_YUV2BGR_YUNV = CV_YUV2BGR_YUY2,

	// CV_YUV2RGBA_YUY2 = 119,
	// CV_YUV2BGRA_YUY2 = 120,
	// CV_YUV2RGBA_YVYU = 121,
	// CV_YUV2BGRA_YVYU = 122,
	// CV_YUV2RGBA_YUYV = CV_YUV2RGBA_YUY2,
	// CV_YUV2BGRA_YUYV = CV_YUV2BGRA_YUY2,
	// CV_YUV2RGBA_YUNV = CV_YUV2RGBA_YUY2,
	// CV_YUV2BGRA_YUNV = CV_YUV2BGRA_YUY2,

	// CV_YUV2GRAY_UYVY = 123,
	// CV_YUV2GRAY_YUY2 = 124,
	// //CV_YUV2GRAY_VYUY = CV_YUV2GRAY_UYVY,
	// CV_YUV2GRAY_Y422 = CV_YUV2GRAY_UYVY,
	// CV_YUV2GRAY_UYNV = CV_YUV2GRAY_UYVY,
	// CV_YUV2GRAY_YVYU = CV_YUV2GRAY_YUY2,
	// CV_YUV2GRAY_YUYV = CV_YUV2GRAY_YUY2,
	// CV_YUV2GRAY_YUNV = CV_YUV2GRAY_YUY2,

	// // alpha premultiplication
	// CV_RGBA2mRGBA = 125,
	// CV_mRGBA2RGBA = 126,

	// CV_COLORCVT_MAX  = 127
)

// /* Sub-pixel interpolation methods */
const (
	CV_INTER_NN       = C.CV_INTER_NN
	CV_INTER_LINEAR   = C.CV_INTER_LINEAR
	CV_INTER_CUBIC    = C.CV_INTER_CUBIC
	CV_INTER_AREA     = C.CV_INTER_AREA
	CV_INTER_LANCZOS4 = C.CV_INTER_LANCZOS4
)

// /* ... and other image warping flags */
const (
	CV_WARP_FILL_OUTLIERS = C.CV_WARP_FILL_OUTLIERS
	CV_WARP_INVERSE_MAP   = C.CV_WARP_INVERSE_MAP
)

// /* Shapes of a structuring element for morphological operations */
const (
	CV_SHAPE_RECT    = C.CV_SHAPE_RECT
	CV_SHAPE_CROSS   = C.CV_SHAPE_CROSS
	CV_SHAPE_ELLIPSE = C.CV_SHAPE_ELLIPSE
	CV_SHAPE_CUSTOM  = C.CV_SHAPE_CUSTOM
)

// /* Morphological operations */
const (
	CV_MOP_ERODE    = C.CV_MOP_ERODE
	CV_MOP_DILATE   = C.CV_MOP_DILATE
	CV_MOP_OPEN     = C.CV_MOP_OPEN
	CV_MOP_CLOSE    = C.CV_MOP_CLOSE
	CV_MOP_GRADIENT = C.CV_MOP_GRADIENT
	CV_MOP_TOPHAT   = C.CV_MOP_TOPHAT
	CV_MOP_BLACKHAT = C.CV_MOP_BLACKHAT
)

// /* Spatial and central moments */
type Moments C.CvMoments

// typedef struct CvMoments
// {
//     double  m00, m10, m01, m20, m11, m02, m30, m21, m12, m03; /* spatial moments */
//     double  mu20, mu11, mu02, mu30, mu21, mu12, mu03; /* central moments */
//     double  inv_sqrt_m00; /* m00 != 0 ? 1/sqrt(m00) : 0 */
// }

// /* Hu invariants */
type HuMoments C.CvHuMoments

// typedef struct CvHuMoments
// {
//     double hu1, hu2, hu3, hu4, hu5, hu6, hu7; /* Hu invariants */
// }

// /* Template matching methods */
const (
	CV_TM_SQDIFF        = C.CV_TM_SQDIFF
	CV_TM_SQDIFF_NORMED = C.CV_TM_SQDIFF_NORMED
	CV_TM_CCORR         = C.CV_TM_CCORR
	CV_TM_CCORR_NORMED  = C.CV_TM_CCORR_NORMED
	CV_TM_CCOEFF        = C.CV_TM_CCOEFF
	CV_TM_CCOEFF_NORMED = C.CV_TM_CCOEFF_NORMED
)

// typedef float (CV_CDECL * CvDistanceFunction)( const float* a, const float* b, void* user_param );

/* Contour retrieval modes */
const (
	CV_RETR_EXTERNAL  = C.CV_RETR_EXTERNAL
	CV_RETR_LIST      = C.CV_RETR_LIST
	CV_RETR_CCOMP     = C.CV_RETR_CCOMP
	CV_RETR_TREE      = C.CV_RETR_TREE
	CV_RETR_FLOODFILL = C.CV_RETR_FLOODFILL
)

// /* Contour approximation methods */
const (
	CV_CHAIN_CODE             = C.CV_CHAIN_CODE
	CV_CHAIN_APPROX_NONE      = C.CV_CHAIN_APPROX_NONE
	CV_CHAIN_APPROX_SIMPLE    = C.CV_CHAIN_APPROX_SIMPLE
	CV_CHAIN_APPROX_TC89_L1   = C.CV_CHAIN_APPROX_TC89_L1
	CV_CHAIN_APPROX_TC89_KCOS = C.CV_CHAIN_APPROX_TC89_KCOS
	CV_LINK_RUNS              = C.CV_LINK_RUNS
)

// /*
// Internal structure that is used for sequental retrieving contours from the image.
// It supports both hierarchical and plane variants of Suzuki algorithm.
// */
type ContourScanner C.CvContourScanner

// typedef struct _CvContourScanner* CvContourScanner;

// /* Freeman chain reader state */
type ChainPtReader C.CvChainPtReader

// typedef struct CvChainPtReader
// {
//     CV_SEQ_READER_FIELDS()
//     char      code;
//     CvPoint   pt;
//     schar     deltas[8][2];
// }
// CvChainPtReader;

// /* initializes 8-element array for fast access to 3x3 neighborhood of a pixel */
// #define  CV_INIT_3X3_DELTAS( deltas, step, nch )            \
//     ((deltas)[0] =  (nch),  (deltas)[1] = -(step) + (nch),  \
//      (deltas)[2] = -(step), (deltas)[3] = -(step) - (nch),  \
//      (deltas)[4] = -(nch),  (deltas)[5] =  (step) - (nch),  \
//      (deltas)[6] =  (step), (deltas)[7] =  (step) + (nch))

// /****************************************************************************************\
// *                              Planar subdivisions                                       *
// \****************************************************************************************/

// typedef size_t CvSubdiv2DEdge;

// #define CV_QUADEDGE2D_FIELDS()     \
//     int flags;                     \
//     struct CvSubdiv2DPoint* pt[4]; \
//     CvSubdiv2DEdge  next[4];

// #define CV_SUBDIV2D_POINT_FIELDS()\
//     int            flags;      \
//     CvSubdiv2DEdge first;      \
//     CvPoint2D32f   pt;         \
//     int id;

// #define CV_SUBDIV2D_VIRTUAL_POINT_FLAG (1 << 30)
type QuadEdge2D C.CvQuadEdge2D

// typedef struct CvQuadEdge2D
// {
//     CV_QUADEDGE2D_FIELDS()
// }
// CvQuadEdge2D;

type Subdiv2DPoint C.CvSubdiv2DPoint

// typedef struct CvSubdiv2DPoint
// {
//     CV_SUBDIV2D_POINT_FIELDS()
// }
// CvSubdiv2DPoint;

// #define CV_SUBDIV2D_FIELDS()    \
//     CV_GRAPH_FIELDS()           \
//     int  quad_edges;            \
//     int  is_geometry_valid;     \
//     CvSubdiv2DEdge recent_edge; \
//     CvPoint2D32f  topleft;      \
//     CvPoint2D32f  bottomright;

// typedef struct CvSubdiv2D
// {
//     CV_SUBDIV2D_FIELDS()
// }
// CvSubdiv2D;

// typedef enum CvSubdiv2DPointLocation
// {
//     CV_PTLOC_ERROR = -2,
//     CV_PTLOC_OUTSIDE_RECT = -1,
//     CV_PTLOC_INSIDE = 0,
//     CV_PTLOC_VERTEX = 1,
//     CV_PTLOC_ON_EDGE = 2
// }
// CvSubdiv2DPointLocation;

// typedef enum CvNextEdgeType
// {
//     CV_NEXT_AROUND_ORG   = 0x00,
//     CV_NEXT_AROUND_DST   = 0x22,
//     CV_PREV_AROUND_ORG   = 0x11,
//     CV_PREV_AROUND_DST   = 0x33,
//     CV_NEXT_AROUND_LEFT  = 0x13,
//     CV_NEXT_AROUND_RIGHT = 0x31,
//     CV_PREV_AROUND_LEFT  = 0x20,
//     CV_PREV_AROUND_RIGHT = 0x02
// }
// CvNextEdgeType;

// /* get the next edge with the same origin point (counterwise) */
// #define  CV_SUBDIV2D_NEXT_EDGE( edge )  (((CvQuadEdge2D*)((edge) & ~3))->next[(edge)&3])

// /* Contour approximation algorithms */
const (
	CV_POLY_APPROX_DP = C.CV_POLY_APPROX_DP
)

// /* Shape matching methods */
const (
	CV_CONTOURS_MATCH_I1 = C.CV_CONTOURS_MATCH_I1
	CV_CONTOURS_MATCH_I2 = C.CV_CONTOURS_MATCH_I2
	CV_CONTOURS_MATCH_I3 = C.CV_CONTOURS_MATCH_I3
)

// /* Shape orientation */
const (
	CV_CLOCKWISE         = C.CV_CLOCKWISE
	CV_COUNTER_CLOCKWISE = C.CV_COUNTER_CLOCKWISE
)

// /* Convexity defect */
// typedef struct CvConvexityDefect
// {
//     CvPoint* start; /* point of the contour where the defect begins */
//     CvPoint* end; /* point of the contour where the defect ends */
//     CvPoint* depth_point; /* the farthest from the convex hull point within the defect */
//     float depth; /* distance between the farthest point and the convex hull */
// } CvConvexityDefect;

// /* Histogram comparison methods */
const (
	CV_COMP_CORREL        = C.CV_COMP_CORREL
	CV_COMP_CHISQR        = C.CV_COMP_CHISQR
	CV_COMP_INTERSECT     = C.CV_COMP_INTERSECT
	CV_COMP_BHATTACHARYYA = C.CV_COMP_BHATTACHARYYA
	CV_COMP_HELLINGER     = C.CV_COMP_HELLINGER
)

// /* Mask size for distance transform */
// enum
// {
//     CV_DIST_MASK_3   =3,
//     CV_DIST_MASK_5   =5,
//     CV_DIST_MASK_PRECISE =0
// };

// /* Content of output label array: connected components or pixels */
// enum
// {
//   CV_DIST_LABEL_CCOMP = 0,
//   CV_DIST_LABEL_PIXEL = 1
// };

// /* Distance types for Distance Transform and M-estimators */
// enum
// {
//     CV_DIST_USER    =-1,  /* User defined distance */
//     CV_DIST_L1      =1,   /* distance = |x1-x2| + |y1-y2| */
//     CV_DIST_L2      =2,   /* the simple euclidean distance */
//     CV_DIST_C       =3,   /* distance = max(|x1-x2|,|y1-y2|) */
//     CV_DIST_L12     =4,   /* L1-L2 metric: distance = 2(sqrt(1+x*x/2) - 1)) */
//     CV_DIST_FAIR    =5,   /* distance = c^2(|x|/c-log(1+|x|/c)), c = 1.3998 */
//     CV_DIST_WELSCH  =6,   /* distance = c^2/2(1-exp(-(x/c)^2)), c = 2.9846 */
//     CV_DIST_HUBER   =7    /* distance = |x|<c ? x^2/2 : c(|x|-c/2), c=1.345 */
// };

// /* Threshold types */
const (
	CV_THRESH_BINARY     = C.CV_THRESH_BINARY     /* value = value > threshold ? max_value : 0       */
	CV_THRESH_BINARY_INV = C.CV_THRESH_BINARY_INV /* value = value > threshold ? 0 : max_value       */
	CV_THRESH_TRUNC      = C.CV_THRESH_TRUNC      /* value = value > threshold ? threshold : value   */
	CV_THRESH_TOZERO     = C.CV_THRESH_TOZERO     /* value = value > threshold ? value : 0           */
	CV_THRESH_TOZERO_INV = C.CV_THRESH_TOZERO_INV /* value = value > threshold ? 0 : value           */
	CV_THRESH_MASK       = C.CV_THRESH_MASK
	CV_THRESH_OTSU       = C.CV_THRESH_OTSU /* use Otsu algorithm to choose the optimal threshold value;
	   combine the flag with one of the above CV_THRESH_* values */
)

// /* Adaptive threshold methods */
// enum
// {
//     CV_ADAPTIVE_THRESH_MEAN_C  =0,
//     CV_ADAPTIVE_THRESH_GAUSSIAN_C  =1
// };

// /* FloodFill flags */
// enum
// {
//     CV_FLOODFILL_FIXED_RANGE =(1 << 16),
//     CV_FLOODFILL_MASK_ONLY   =(1 << 17)
// };

// /* Canny edge detector flags */
// enum
// {
//     CV_CANNY_L2_GRADIENT  =(1 << 31)
// };

// /* Variants of a Hough transform */
const (
	CV_HOUGH_STANDARD      = C.CV_HOUGH_STANDARD
	CV_HOUGH_PROBABILISTIC = C.CV_HOUGH_PROBABILISTIC
	CV_HOUGH_MULTI_SCALE   = C.CV_HOUGH_MULTI_SCALE
	CV_HOUGH_GRADIENT      = C.CV_HOUGH_GRADIENT
)

// /* Fast search data structures  */
// struct CvFeatureTree;
// struct CvLSH;
// struct CvLSHOperations;

// #ifdef __cplusplus
// }
// #endif

// #endif
