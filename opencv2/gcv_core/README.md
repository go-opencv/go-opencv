Gcv Core
========

Wrap the core types in OpenCV.

## Supporting Types and Examples

| OpenCV C++    | Go OpenCV     | Constructor                   |
|---------------|---------------|-------------------------------|
| `cv::Point2i` | `GcvPoint2i`  | `NewGcvPoint2i(x, y int)`     |
| `cv::Point2f` | `GcvPoint2f_` | `NewGcvPoint2f(x, y float32)` |
| `cv::Point2d` | `GcvPoint2d_` | `NewGcvPoint2d(x, y float64)` |
| `cv::Point3i` | `GcvPoint3i`  | `NewGcvPoint3i(x, y, z int)`     |
| `cv::Point3f` | `GcvPoint3f_` | `NewGcvPoint3f(x, y, z float32)` |
| `cv::Point3d` | `GcvPoint3d_` | `NewGcvPoint3d(x, y, z float64)` |
| `cv::Size2i`  | `GcvSize2i`   | `NewGcvSize2i(x, y int)`     |
| `cv::Size2f`  | `GcvSize2f_`  | `NewGcvSize2f(x, y float32)` |
| `cv::Size2d`  | `GcvSize2d_`  | `NewGcvSize2d(x, y float64)` |

----------

### Note for Renamed Types

Some of the types are renamed to `*_`. The reason is that we'd like to wrap a better interface for them.  
For example, the original `NewPoint2f` takes strictly two `float32`, and we are not able to pass `float64` or `int`, which doesn't make too much sense.  
After wrapping an extra level, we are now able to pass `int`, `float32`, and `float64` to these methods.  
Also note that **renaming doesn't affect any usage**, except you are manipulating the types yourself.
