#include <opencv2/opencv.hpp>
#include <vector>

using namespace std;

class GcvPoint3f
{
public:
        GcvPoint3f (float x, float y, float z)
                : _data(x, y, z) {};
        ~GcvPoint3f () {};

        cv::Point3f Get() { return _data; }
        void Set(float x, float y, float z) {
                _data = cv::Point3f(x, y, z);
        }
private:
        cv::Point3f _data;
};

class GcvPoint2f
{
public:
        GcvPoint2f (float x, float y)
                : _data(x, y) {};
        ~GcvPoint2f () {};

        cv::Point2f Get() { return _data; }
        void Set(float x, float y) {
                _data = cv::Point2f(x, y);
        }
private:
        cv::Point2f _data;
};
