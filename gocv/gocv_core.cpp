#include <opencv2/opencv.hpp>
#include <opencv2/core/core.hpp>
#include <iostream>
#include <vector>

#include "gocv_core.hpp"

cv::Mat Mat64ToGcvMat_(int row, int col, std::vector<double> data) {
        assert(row * col == data.size());

        cv::Mat mat = cv::Mat(row, col, CV_64F);

        for (int i = 0; i < row; ++i) {
                for (int j = 0; j < col; ++j) {
                        mat.at<double>(i, j) = data[i*col + j];
                }

        }

        return mat;
}
