#include <opencv2/opencv.hpp>
#include <opencv2/core/core.hpp>
#include <iostream>
#include <vector>

#include "gocv_calib3d.hpp"

cv::Mat GcvInitCameraMatrix2D_(VecPoint3f objPts, VecPoint2f imgPts, cv::Size imgSize, double aspectRatio) {
        cv::Mat cameraMatrix;

        std::vector<VecPoint3f> objPtsArr;
        std::vector<VecPoint2f> imgPtsArr;

        objPtsArr.push_back(objPts);
        imgPtsArr.push_back(imgPts);

        cameraMatrix = cv::initCameraMatrix2D(objPtsArr, imgPtsArr, imgSize, aspectRatio);
        return cameraMatrix;
}

double GcvCalibrateCamera_(VecPoint3f objPts, VecPoint2f imgPts,
                          cv::Size imgSize, cv::Mat& cameraMatrix, cv::Mat distCoeffs,
                          cv::Mat& rvec, cv::Mat& tvec, int flags) {
        std::vector<VecPoint3f> objPtsArr;
        std::vector<VecPoint2f> imgPtsArr;
        std::vector<cv::Mat> rvecs, tvecs;
        double rtn;

        objPtsArr.push_back(objPts);
        imgPtsArr.push_back(imgPts);

        /* std::cout << "objPts " << std::endl << objPtsArr[0] << std::endl; */
        /* std::cout << "imgPts " << std::endl << imgPtsArr[0] << std::endl; */
        /* std::cout << "imgSize " << std::endl << imgSize << std::endl; */
        /* std::cout << "Before CamMat " << std::endl << cameraMatrix << std::endl; */

        rtn = cv::calibrateCamera(objPtsArr, imgPtsArr, imgSize,
                                  cameraMatrix, distCoeffs,
                                  rvecs, tvecs, flags);

        rvec = rvecs[0];
        tvec = tvecs[0];

        /* std::cout << "After CamMat " << std::endl << cameraMatrix << std::endl; */
        /* std::cout << "distCoeffs " << std::endl << distCoeffs << std::endl; */
        /* std::cout << "rvec " << std::endl << rvec << std::endl; */
        /* std::cout << "tvec " << std::endl << tvec << std::endl; */
        /* std::cout << "rms " << std::endl << rtn << std::endl; */

        return rtn;
}

void GcvRodrigues_(cv::Mat src, cv::Mat& dst) {
        cv::Rodrigues(src, dst);
}
