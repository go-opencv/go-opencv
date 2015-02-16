#include <opencv2/opencv.hpp>
#include <opencv2/core/core.hpp>
#include <iostream>
#include <vector>

#include "gocv_calib3d.hpp"

cv::Mat GcvInitCameraMatrix2D(VecPoint3f objPts, VecPoint2f imgPts) {
        cv::Mat cameraMatrix;

        std::vector<VecPoint3f> objPtsArr;
        std::vector<VecPoint2f> imgPtsArr;

        objPtsArr.push_back(objPts);
        imgPtsArr.push_back(imgPts);

        cameraMatrix = cv::initCameraMatrix2D(objPtsArr, imgPtsArr, cv::Size(1920, 1080), 1);
        std::cout << cameraMatrix.type() << std::endl;
        return cameraMatrix;
}

double GcvCalibrateCamera(VecPoint3f objPts, VecPoint2f imgPts,
                          cv::Size imgSize, cv::Mat cameraMatrix) {
        std::vector<VecPoint3f> objPtsArr;
        std::vector<VecPoint2f> imgPtsArr;
        std::vector<cv::Mat> rvecs, tvecs;
        cv::Mat distCoeffs;

        double rtn;

        objPtsArr.push_back(objPts);
        imgPtsArr.push_back(imgPts);

        std::cout << "init Camera" << cameraMatrix << std::endl;

        rtn = cv::calibrateCamera(objPtsArr, imgPtsArr, imgSize,
                                  cameraMatrix, distCoeffs, rvecs, tvecs);

        std::cout << "final Camera" << cameraMatrix << std::endl;
        std::cout << "final rvecs" << rvecs[0] << std::endl;
        std::cout << "final tvecs" << tvecs[0] << std::endl;

        return rtn;
}
