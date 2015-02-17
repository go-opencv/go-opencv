%{
#include "opencv2/core/core.hpp"
#include "gocv_core.hpp"
%}

%include "std_vector.i"
%include "gocv_core.hpp"

/* Classes defined in core.hpp */
namespace cv {

   template<typename _Tp> class Size_;
   template<typename _Tp> class Point_;
   template<typename _Tp> class Rect_;
   template<typename _Tp, int cn> class Vec;

   //////////////////////////////// Point_ ////////////////////////////////

   /*!
   template 2D point class.

   The class defines a point in 2D space. Data type of the point coordinates is specified
   as a template parameter. There are a few shorter aliases available for user convenience.
   See cv::Point, cv::Point2i, cv::Point2f and cv::Point2d.
   */
   template<typename _Tp> class Point_
   {
   public:
      typedef _Tp value_type;

      // various constructors
      Point_();
      Point_(_Tp _x, _Tp _y);
      Point_(const Point_& pt);
      Point_(const CvPoint& pt);
      Point_(const CvPoint2D32f& pt);
      Point_(const Size_<_Tp>& sz);
      Point_(const Vec<_Tp, 2>& v);

      Point_& operator = (const Point_& pt);
      //! conversion to another data type
      template<typename _Tp2> operator Point_<_Tp2>() const;

      //! conversion to the old-style C structures
      operator CvPoint() const;
      operator CvPoint2D32f() const;
      operator Vec<_Tp, 2>() const;

      //! dot product
      _Tp dot(const Point_& pt) const;
      //! dot product computed in double-precision arithmetics
      double ddot(const Point_& pt) const;
      //! cross-product
      double cross(const Point_& pt) const;
      //! checks whether the point is inside the specified rectangle
      bool inside(const Rect_<_Tp>& r) const;

      _Tp x, y; //< the point coordinates
   };

   /*!
   template 3D point class.

   The class defines a point in 3D space. Data type of the point coordinates is specified
   as a template parameter.

   \see cv::Point3i, cv::Point3f and cv::Point3d
   */
   template<typename _Tp> class Point3_
   {
   public:
      typedef _Tp value_type;

      // various constructors
      Point3_();
      Point3_(_Tp _x, _Tp _y, _Tp _z);
      Point3_(const Point3_& pt);
      explicit Point3_(const Point_<_Tp>& pt);
      Point3_(const CvPoint3D32f& pt);
      Point3_(const Vec<_Tp, 3>& v);

      Point3_& operator = (const Point3_& pt);
      //! conversion to another data type
      template<typename _Tp2> operator Point3_<_Tp2>() const;
      //! conversion to the old-style CvPoint...
      operator CvPoint3D32f() const;
      //! conversion to cv::Vec<>
      operator Vec<_Tp, 3>() const;

      //! dot product
      _Tp dot(const Point3_& pt) const;
      //! dot product computed in double-precision arithmetics
      double ddot(const Point3_& pt) const;
      //! cross product of the 2 3D points
      Point3_ cross(const Point3_& pt) const;

      _Tp x, y, z; //< the point coordinates
   };

   //////////////////////////////// Size_ ////////////////////////////////

   /*!
   The 2D size class

   The class represents the size of a 2D rectangle, image size, matrix size etc.
   Normally, cv::Size ~ cv::Size_<int> is used.
   */
   template<typename _Tp> class Size_
   {
   public:
      typedef _Tp value_type;

      //! various constructors
      Size_();
      Size_(_Tp _width, _Tp _height);
      Size_(const Size_& sz);
      Size_(const CvSize& sz);
      Size_(const CvSize2D32f& sz);
      Size_(const Point_<_Tp>& pt);

      Size_& operator = (const Size_& sz);
      //! the area (width*height)
      _Tp area() const;

      //! conversion of another data type.
      template<typename _Tp2> operator Size_<_Tp2>() const;

      //! conversion to the old-style OpenCV types
      operator CvSize() const;
      operator CvSize2D32f() const;

      _Tp width, height; // the width and the height
   };

   //////////////////////////////// Rect_ ////////////////////////////////

   /*!
   The 2D up-right rectangle class

   The class represents a 2D rectangle with coordinates of the specified data type.
   Normally, cv::Rect ~ cv::Rect_<int> is used.
   */
   template<typename _Tp> class Rect_
   {
   public:
      typedef _Tp value_type;

      //! various constructors
      Rect_();
      Rect_(_Tp _x, _Tp _y, _Tp _width, _Tp _height);
      Rect_(const Rect_& r);
      Rect_(const CvRect& r);
      Rect_(const Point_<_Tp>& org, const Size_<_Tp>& sz);
      Rect_(const Point_<_Tp>& pt1, const Point_<_Tp>& pt2);

      Rect_& operator = ( const Rect_& r );
      //! the top-left corner
      Point_<_Tp> tl() const;
      //! the bottom-right corner
      Point_<_Tp> br() const;

      //! size (width, height) of the rectangle
      Size_<_Tp> size() const;
      //! area (width*height) of the rectangle
      _Tp area() const;

      //! conversion to another data type
      template<typename _Tp2> operator Rect_<_Tp2>() const;
      //! conversion to the old-style CvRect
      operator CvRect() const;

      //! checks whether the rectangle contains the point
      bool contains(const Point_<_Tp>& pt) const;

      _Tp x, y, width, height; //< the top-left corner, as well as width and height of the rectangle
   };


   %template(GcvSize2i) Size_<int>;
   %template(GcvSize2f32_) Size_<float>;
   %template(GcvSize2f64_) Size_<double>;

   %template(GcvRect) Rect_<int>;

   %template(GcvPoint2i) Point_<int>;
   %template(GcvPoint2f32_) Point_<float>;
   %template(GcvPoint2f64_) Point_<double>;

   %template(GcvPoint3i) Point3_<int>;
   %template(GcvPoint3f32_) Point3_<float>;
   %template(GcvPoint3f64_) Point3_<double>;


   /* ----------------- Mat ----------------- */
   %rename(GcvMat) Mat;

   class Mat
   {
   public:
      //! default constructor
      Mat();
      //! constructs 2D matrix of the specified size and type
      // (_type is CV_8UC1, CV_64FC3, CV_32SC(12) etc.)
      Mat(int rows, int cols, int type);
      Mat(cv::Size size, int type);
      //! constucts 2D matrix and fills it with the specified value _s.
      Mat(int rows, int cols, int type, const cv::Scalar& s);
      Mat(cv::Size size, int type, const cv::Scalar& s);

      //! copy constructor
      Mat(const Mat& m);

      //! builds matrix from std::vector with or without copying the data
      template<typename _Tp> explicit Mat(const vector<_Tp>& vec, bool copyData=false);
      //! builds matrix from cv::Vec; the data is copied by default
      template<typename _Tp, int n> explicit Mat(const Vec<_Tp, n>& vec, bool copyData=true);
      //! builds matrix from cv::Matx; the data is copied by default
      template<typename _Tp, int m, int n> explicit Mat(const Matx<_Tp, m, n>& mtx, bool copyData=true);
      //! builds matrix from a 2D point
      template<typename _Tp> explicit Mat(const Point_<_Tp>& pt, bool copyData=true);
      //! builds matrix from a 3D point
      template<typename _Tp> explicit Mat(const Point3_<_Tp>& pt, bool copyData=true);

      //! destructor - calls release()
      ~Mat();

      //! returns a new matrix header for the specified row
      Mat row(int y) const;
      //! returns a new matrix header for the specified column
      Mat col(int x) const;
      //! ... for the specified row span
      Mat rowRange(int startrow, int endrow) const;
      //! ... for the specified column span
      Mat colRange(int startcol, int endcol) const;
      //! ... for the specified diagonal
      // (d=0 - the main diagonal,
      //  >0 - a diagonal from the lower half,
      //  <0 - a diagonal from the upper half)
      Mat diag(int d=0) const;
      //! constructs a square diagonal matrix which main diagonal is vector "d"
      static Mat diag(const Mat& d);

      //! returns deep copy of the matrix, i.e. the data is copied
      Mat clone() const;

      void assignTo( Mat& m, int type=-1 ) const;

      //! creates alternative matrix header for the same data, with different
      // number of channels and/or different number of rows. see cvReshape.
      Mat reshape(int cn, int rows=0) const;
      Mat reshape(int cn, int newndims, const int* newsz) const;

      //! adds element to the end of 1d matrix (or possibly multiple elements when _Tp=Mat)
      template<typename _Tp> void push_back(const _Tp& elem);
      template<typename _Tp> void push_back(const Mat_<_Tp>& elem);
      void push_back(const Mat& m);
      //! removes several hyper-planes from bottom of the matrix
      void pop_back(size_t nelems=1);

      //! special versions for 2D arrays (especially convenient for referencing image pixels)

      template<typename _Tp> _Tp& at(int i0=0);
      template<typename _Tp> const _Tp& at(int i0=0) const;

      template<typename _Tp> _Tp& at(int i0, int i1);
      template<typename _Tp> const _Tp& at(int i0, int i1) const;

      template<typename _Tp> _Tp& at(int i0, int i1, int i2);
      template<typename _Tp> const _Tp& at(int i0, int i1, int i2) const;

      template<typename _Tp> _Tp& at(const int* idx);
      template<typename _Tp> const _Tp& at(const int* idx) const;

      template<typename _Tp, int n> _Tp& at(const Vec<int, n>& idx);
      template<typename _Tp, int n> const _Tp& at(const Vec<int, n>& idx) const;
      template<typename _Tp> _Tp& at(cv::Point pt);
      template<typename _Tp> const _Tp& at(cv::Point pt) const;

      %template(gcvAtf32) at<float>;
      %template(gcvAtf64) at<double>;

      /*! includes several bit-fields:
            - the magic signature
            - continuity flag
            - depth
            - number of channels
      */
      int flags;
      //! the matrix dimensionality, >= 2
      int dims;
      //! the number of rows and columns or (-1, -1) when the matrix has more than 2 dimensions
      int rows, cols;
   };
}

/* Additional STL types */
namespace std {
   %template(GcvPoint3f32Vector) vector<cv::Point3f>;
   %template(GcvPoint2f32Vector) vector<cv::Point2f>;

   %template(GcvIntVector) vector<int>;
   %template(GcvFloat32Vector) vector<float>;
   %template(GcvFloat64Vector) vector<double>;
};
