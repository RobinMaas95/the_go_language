// Surface computes an SVG rendering of a 3D surface function.
// Task:
// Follwoing the approach of the Lissajous example in Section 1.7, construct approach
// a webserver taht computes surfaces and writes SVG data to the client. The server
// must set the content-type header like this ('w.Header().Set("Content-Type", "image/svg+xml")')
// Allow the client to specify values like height width and color as HTTP request parameters.

package main
