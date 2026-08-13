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

// bitcoinArtwork is an outline rendition of the Bitcoin mark from
// https://upload.wikimedia.org/wikipedia/commons/4/46/Bitcoin.svg.
func bitcoinArtwork(params engrave.Params) engrave.Engraving {
	return func(yield func(engrave.Command) bool) {
		const (
			left = 10
			top  = 10
			size = 65
		)
		point := func(x, y int) bezier.Point {
			return bezier.Pt(params.I(left+x*size/64), params.I(top+y*size/64))
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

		// The points trace the two paths in Bitcoin.svg, using short line segments
		// for its curves. This keeps the firmware POC limited to Move and Line.
		if !path(
			point(63, 40), point(62, 44), point(60, 49), point(57, 53),
			point(53, 57), point(49, 60), point(44, 62), point(39, 63),
			point(34, 63), point(29, 62), point(24, 61), point(19, 58),
			point(15, 55), point(11, 51), point(8, 47), point(5, 42),
			point(3, 37), point(2, 32), point(1, 27), point(1, 22),
			point(3, 17), point(5, 13), point(8, 9), point(12, 6),
			point(17, 3), point(22, 1), point(27, 1), point(32, 1),
			point(37, 2), point(42, 4), point(47, 7), point(51, 10),
			point(55, 14), point(58, 18), point(60, 23), point(62, 28),
			point(63, 34), point(63, 40),
		) {
			return
		}
		if !path(
			point(46, 27), point(46, 24), point(44, 22), point(41, 20),
			point(35, 18), point(36, 13), point(33, 13), point(32, 18),
			point(29, 17), point(30, 12), point(27, 11), point(26, 17),
			point(22, 16), point(21, 19), point(24, 20), point(25, 22),
			point(23, 29), point(20, 38), point(19, 39), point(17, 39),
			point(15, 43), point(20, 45), point(23, 46), point(22, 52),
			point(26, 53), point(28, 47), point(31, 48), point(29, 53),
			point(33, 54), point(35, 48), point(41, 49), point(45, 48),
			point(48, 45), point(50, 41), point(50, 38), point(49, 35),
			point(45, 32), point(47, 31), point(46, 27),
		) {
			return
		}
		if !path(
			point(32, 22), point(31, 29), point(35, 30), point(38, 30),
			point(39, 28), point(39, 26), point(37, 24), point(32, 22),
		) {
			return
		}
		path(
			point(29, 32), point(27, 40), point(32, 41), point(37, 42),
			point(39, 39), point(38, 36), point(35, 34), point(29, 32),
		)
	}
}
