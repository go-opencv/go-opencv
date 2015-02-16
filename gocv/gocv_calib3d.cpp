#include <opencv2/opencv.hpp>
#include <opencv2/core/core.hpp>
#include <iostream>
#include <vector>

#include "gocv_calib3d.hpp"

cv::Mat GcvInitCameraMatrix2D_(VecPoint3f objPts, VecPoint2f imgPts) {
        cv::Mat cameraMatrix;

        std::vector<VecPoint3f> objPtsArr;
        std::vector<VecPoint2f> imgPtsArr;

        objPtsArr.push_back(objPts);
        imgPtsArr.push_back(imgPts);

        cameraMatrix = cv::initCameraMatrix2D(objPtsArr, imgPtsArr, cv::Size(1920, 1080), 1);
        return cameraMatrix;
}

double GcvCalibrateCamera_(VecPoint3f objPts, VecPoint2f imgPts,
                          cv::Size imgSize, cv::Mat& cameraMatrix,
                          cv::Mat& rvec, cv::Mat& tvec) {
        std::vector<VecPoint3f> objPtsArr;
        std::vector<VecPoint2f> imgPtsArr;
        std::vector<cv::Mat> rvecs, tvecs;
        cv::Mat distCoeffs;
        double rtn;

        objPtsArr.push_back(objPts);
        imgPtsArr.push_back(imgPts);

        rtn = cv::calibrateCamera(objPtsArr, imgPtsArr, imgSize,
                                  cameraMatrix, distCoeffs,
                                  rvecs, tvecs);

        rvec = rvecs[0];
        tvec = tvecs[0];

        return rtn;
}
