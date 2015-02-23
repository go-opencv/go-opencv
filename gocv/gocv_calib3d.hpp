#include <opencv2/opencv.hpp>
#include <vector>
#include <iostream>

typedef std::vector<cv::Point3f> VecPoint3f;
typedef std::vector<cv::Point2f> VecPoint2f;

cv::Mat GcvInitCameraMatrix2D_(VecPoint3f objPts, VecPoint2f imgPts,
                               cv::Size imgSize, double aspectRatio);

double GcvCalibrateCamera_(VecPoint3f objPts, VecPoint2f imgPts,
                          cv::Size imgSize, cv::Mat& cameraMatrix, cv::Mat distCoeffs,
                          cv::Mat& rvec, cv::Mat& tvec, int flags);

void GcvRodrigues_(cv::Mat src, cv::Mat& dst);
