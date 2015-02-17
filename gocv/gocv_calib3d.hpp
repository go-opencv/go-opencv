#include <opencv2/opencv.hpp>
#include <vector>
#include <iostream>

typedef std::vector<cv::Point3f> VecPoint3f;
typedef std::vector<cv::Point2f> VecPoint2f;

cv::Mat GcvInitCameraMatrix2D_(VecPoint3f objPts, VecPoint2f imgPts);

double GcvCalibrateCamera_(VecPoint3f objPts, VecPoint2f imgPts,
                           cv::Size2i imgSize, cv::Mat& cameraMatrix,
                           cv::Mat& rvec, cv::Mat& tvec);

void GcvRodrigues_(cv::Mat src, cv::Mat& dst);
