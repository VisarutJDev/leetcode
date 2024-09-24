package hackerrank

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// Write your code here
	// b amount of desired black gifts
	// w amount of desired white gifts
	// each b cost bc units
	// each w cost wc units
	// convert z
	// case 1 same cost b * bc + w * wc
	standard := (int64(b) * int64(bc)) + (int64(w) * int64(wc))
	if bc == wc {
		return standard
	}
	// case 2 convert and shows to expensive  more whan actual price b * bc + w * wc
	if bc < wc {
		if wc > bc+z {
			// int64((b * (bc+z)) + (w*wc))
			return ((int64(w) + int64(b)) * int64(bc)) + (int64(w) * int64(z))
			// return int64(((w+b) * bc) + (w*z))
		}
		return standard
	} else {
		if bc > wc+z {
			// buy more white to decrease cost of black
			// (b+w) * wc) => mean amount of white plus black buy use white cost
			// (b*z) => mean price of convert
			return ((int64(b) + int64(w)) * int64(wc)) + (int64(b) * int64(z))
			// return int64(((b+w) * wc) + (b*z))
		}
		return standard
	}
}
