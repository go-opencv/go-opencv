#include <opencv2/opencv.hpp>
#include <opencv2/core/core.hpp>
#include <iostream>

#include "gocv_imgproc.hpp"

double GcvThreshold_(cv::Mat src, cv::Mat& dst, double thresh, double maxval, int type) {
        std::cout << "src " << std::endl << src << std::endl;
        std::cout << "thresh " << std::endl << thresh << std::endl;
        std::cout << "maxval " << std::endl << maxval << std::endl;
        std::cout << "type " << std::endl << type << std::endl;

        double rtn;
        rtn = cv::threshold(src, src, 1.0, 3.0, 0);
        return rtn;
}
