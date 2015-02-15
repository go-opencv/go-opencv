#include <opencv2/opencv.hpp>
#include <vector>

using namespace std;

cv::Point3f GetPoint3f(float x, float y, float z) {
        return cv::Point3f(x, y, z);
}

cv::Point2f GetPoint2f(float x, float y) {
        return cv::Point2f(x, y);
}
