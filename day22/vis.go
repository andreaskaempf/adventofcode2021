// Visualization of cubes (incomplete, not used)

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Visualize instructions by drawing overlapping squares
// (not complete)
func visualize(instructions []Instruction) {

	// Dimensions
	var w int64 = 1000
	var h int64 = 1000
	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(w), int(h)}

	// Create image and define colors
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	red := color.RGBA{200, 10, 10, 0xff}
	green := color.RGBA{10, 200, 10, 0xff}

	// Get min/max coords
	xvals := []int64{}
	yvals := []int64{}
	zvals := []int64{}
	for _, i := range instructions {
		xvals = append(xvals, i.x0, i.x1)
		yvals = append(yvals, i.y0, i.y1)
		zvals = append(zvals, i.z0, i.z1)
	}

	// Calculate scaling factor
	minx, maxx := minmax(xvals)
	xrange := maxx - minx + 1
	miny, maxy := minmax(yvals)
	yrange := maxy - miny + 1
	minz, maxz := minmax(zvals)
	zrange := maxz - minz + 1
	scale := float64(min(w, h)) / float64(max(xrange, max(yrange, zrange)))

	// Draw each instruction as a rectangle
	for _, i := range instructions {
		c := red
		if i.onOff == 1 {
			c = green
		}
		line(img, scale, i.x0, i.y0, i.x1, i.y0, c) // top horiz
		line(img, scale, i.x0, i.y1, i.x1, i.y1, c) // bottom horiz
		line(img, scale, i.x0, i.y0, i.x0, i.y1, c) // left
		line(img, scale, i.x1, i.y0, i.x1, i.y1, c) // right
	}

	//img.Set(x, y, red)
	//img.Set(x, y, color.White)

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func line(img *image.RGBA, scale float64, x0, y0, x1, y1 int64, c color.RGBA) {

}
