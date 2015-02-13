#include <opencv2/opencv.hpp>
#include "go_calib3d.hpp"
#include "iostream"
#include "vector"

using namespace std;

void
GoCalib3d::foo() {
        cout << "Hello there" << endl;
        vector< vector<cv::Point3f>> objPts;
        vector< vector<cv::Point2f>> imgPts;

        objPts.push_back(vector<cv::Point3f>());
        imgPts.push_back(vector<cv::Point2f>());

        objPts[0].push_back(cv::Point3f(0, 25, 0));
        objPts[0].push_back(cv::Point3f(0, -25, 0));
        objPts[0].push_back(cv::Point3f(-47, 25, 0));
        objPts[0].push_back(cv::Point3f(-47, -25, 0));

        imgPts[0].push_back(cv::Point2f(1136.4140625, 1041.89208984));
        imgPts[0].push_back(cv::Point2f(1845.33190918, 671.39581299));
        imgPts[0].push_back(cv::Point2f(302.73373413, 634.79998779));
        imgPts[0].push_back(cv::Point2f(1051.46154785, 352.76107788));

        cv::Mat cameraMatrix = cv::initCameraMatrix2D(objPts, imgPts, cv::Size(1920, 1080), 1);
        std::cout << cameraMatrix << std::endl;
}
