package gui

import (
	"seedhammer.com/bezier"
	"seedhammer.com/engrave"
)

func artworkFlow(ctx *Context, th *Colors) {
	plate, err := toPlate(bitcoinArtwork(ctx.Platform.EngraverParams()), ctx.Platform.EngraverParams())
	if err != nil {
		return
	}
	for {
		if NewEngraveScreen(ctx, plate).Engrave(ctx, th) {
			return
		}
	}
}

// bitcoinArtwork is a line-art rendition of the Bitcoin mark. Coordinates are
// normalized to the 85 mm plate and converted only when the plan is generated.
func bitcoinArtwork(params engrave.Params) engrave.Engraving {
	return func(yield func(engrave.Command) bool) {
		const (
			left   = 15
			top    = 15
			size   = 55
			center = left + size/2
		)
		point := func(x, y int) bezier.Point {
			return bezier.Pt(params.I(x), params.I(y))
		}
		path := func(points ...bezier.Point) bool {
			if !yield(engrave.Move(points[0])) {
				return false
			}
			for _, p := range points[1:] {
				if !yield(engrave.Line(p)) {
					return false
				}
			}
			return true
		}

		// Octagonal border avoids curve support in this initial vector POC.
		if !path(
			point(center-16, top), point(center+16, top),
			point(left+size, center-16), point(left+size, center+16),
			point(center+16, top+size), point(center-16, top+size),
			point(left, center+16), point(left, center-16), point(center-16, top),
		) {
			return
		}
		if !path(
			point(center-9, center-20), point(center+8, center-20),
			point(center+14, center-14), point(center+14, center-4),
			point(center+8, center), point(center-9, center),
			point(center+10, center), point(center+16, center+6),
			point(center+16, center+16), point(center+10, center+20),
			point(center-9, center+20), point(center-9, center-20),
		) {
			return
		}
		if !path(point(center-4, center-25), point(center-4, center+25)) {
			return
		}
		path(point(center+3, center-25), point(center+3, center+25))
	}
}
