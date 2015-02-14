#include <opencv2/opencv.hpp>
#include <opencv2/core/core.hpp>
#include <iostream>
#include <vector>

#include "gcv_calib3d.hpp"

cv::Mat gcvInitCameraMatrix2D(VecPoint3f objPts, VecPoint2f imgPts) {
        cv::Mat cameraMatrix;

        std::vector<VecPoint3f> *objPtsArr = new std::vector<VecPoint3f>();
        std::vector<VecPoint2f> *imgPtsArr = new std::vector<VecPoint2f>();

        objPtsArr->push_back(objPts);
        imgPtsArr->push_back(imgPts);

        cameraMatrix = cv::initCameraMatrix2D(*objPtsArr, *imgPtsArr, cv::Size(1920, 1080), 1);
        std::cout << cameraMatrix << std::endl;
        return cameraMatrix;
}
