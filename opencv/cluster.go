// Copyright 2014 <t.kastner@cumulo.at>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

//#include "opencv.h"
import "C"
import "unsafe"

const (
	/* Select random initial centers in each attempt. */
	KMEANS_RANDOM_CENTERS = 0

	/* Use kmeans++ center initialization by Arthur and Vassilvitskii [Arthur2007]. */
	KMEANS_PP_CENTERS = 2
)

/*
KMeans finds centers of k clusters in data and groups input samples around
the clusters. It returns a matrix that stores the cluster indices for every
sample, and a matrix that stores the cluster centers.
*/
func KMeans(data *Mat, k int, termcrit TermCriteria, attempts int, rng RNG, flags int) (labels, centers *Mat) {
	var compactness C.double

	labels = CreateMat(data.Rows(), 1, CV_32S)
	centers = CreateMat(k, 1, data.Type())

	C.cvKMeans2(
		unsafe.Pointer(data),
		C.int(k),
		unsafe.Pointer(labels),
		(C.CvTermCriteria)(termcrit),
		C.int(attempts),
		(*C.CvRNG)(&rng),
		C.int(flags),
		unsafe.Pointer(centers),
		&compactness)

	return labels, centers
}
