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
			return bezier.Pt(
				params.I(left)+(x*params.I(size)+3200)/6400,
				params.I(top)+(y*params.I(size)+3200)/6400,
			)
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

		// Cubic paths from bitcoin-outline.svg flattened to 0.01 SVG units.
		// The firmware receives only line segments.
		if !path(
			point(2475, 6354), point(2112, 6209), point(1815, 6085), point(1537, 5935),
			point(1278, 5759), point(1039, 5561), point(822, 5342), point(628, 5103),
			point(457, 4848), point(312, 4578), point(193, 4294), point(101, 3999),
			point(38, 3696), point(5, 3384), point(3, 3068), point(32, 2748),
			point(96, 2427), point(191, 2113), point(315, 1817), point(465, 1538),
			point(641, 1279), point(839, 1040), point(1058, 823), point(1297, 629),
			point(1552, 458), point(1822, 313), point(2106, 193), point(2400, 102),
			point(2704, 39), point(3015, 5), point(3332, 3), point(3652, 33),
			point(3973, 96), point(4286, 191), point(4583, 315), point(4861, 466),
			point(5121, 641), point(5359, 840), point(5576, 1059), point(5771, 1297),
			point(5941, 1553), point(6087, 1823), point(6206, 2107), point(6298, 2401),
			point(6361, 2705), point(6394, 3016), point(6396, 3333), point(6366, 3653),
			point(6303, 3974), point(6208, 4288), point(6084, 4584), point(5934, 4863),
			point(5758, 5122), point(5560, 5361), point(5340, 5578), point(5102, 5772),
			point(4847, 5943), point(4576, 6088), point(4293, 6207), point(3998, 6299),
			point(3694, 6362), point(3383, 6395), point(3066, 6397), point(2746, 6368),
			point(2475, 6354),
		) {
			return
		}
		if !path(
			point(4661, 2794), point(4617, 2594), point(4590, 2460), point(4532, 2342),
			point(4449, 2238), point(4341, 2147), point(4212, 2068), point(3907, 1937),
			point(4050, 1360), point(3699, 1273), point(3559, 1834), point(3278, 1768),
			point(3419, 1203), point(3068, 1115), point(2924, 1692), point(2700, 1639),
			point(2700, 1637), point(2216, 1516), point(2123, 1891), point(2134, 1894),
			point(2163, 1901), point(2252, 1922), point(2340, 1943), point(2368, 1951),
			point(2378, 1955), point(2464, 1991), point(2514, 2042), point(2538, 2101),
			point(2541, 2159), point(2377, 2816), point(2414, 2828), point(2377, 2819),
			point(2147, 3739), point(2098, 3806), point(2051, 3826), point(1986, 3823),
			point(1976, 3821), point(1948, 3815), point(1860, 3793), point(1772, 3770),
			point(1742, 3762), point(1731, 3759), point(1557, 4161), point(2014, 4275),
			point(2264, 4339), point(2119, 4923), point(2469, 5010), point(2613, 4433),
			point(2893, 4506), point(2750, 5080), point(3101, 5168), point(3246, 4585),
			point(3464, 4620), point(3666, 4635), point(3851, 4628), point(4019, 4594),
			point(4167, 4529), point(4295, 4430), point(4402, 4292), point(4485, 4111),
			point(4527, 3956), point(4543, 3817), point(4533, 3694), point(4499, 3584),
			point(4443, 3486), point(4367, 3401), point(4162, 3260), point(4322, 3201),
			point(4455, 3099), point(4553, 2949), point(4661, 2794),
		) {
			return
		}
		if !path(
			point(3808, 3869), point(3740, 4003), point(3627, 4084), point(3481, 4121),
			point(3316, 4126), point(3145, 4107), point(2981, 4075), point(2838, 4039),
			point(2728, 4011), point(2920, 3238), point(3031, 3263), point(3175, 3297),
			point(3336, 3344), point(3496, 3405), point(3640, 3485), point(3751, 3587),
			point(3812, 3714), point(3808, 3869),
		) {
			return
		}
		path(
			point(3917, 2738), point(3857, 2860), point(3761, 2935), point(3638, 2972),
			point(3500, 2978), point(3220, 2938), point(3100, 2908), point(3009, 2884),
			point(3183, 2183), point(3276, 2204), point(3396, 2232), point(3663, 2324),
			point(3782, 2394), point(3873, 2484), point(3923, 2597), point(3917, 2738),
		)
	}
}
