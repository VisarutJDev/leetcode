package hackerrank

func DesignerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	alphabet := map[rune]int32{
		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
		'h': 7,
		'i': 8,
		'j': 9,
		'k': 10,
		'l': 11,
		'm': 12,
		'n': 13,
		'o': 14,
		'p': 15,
		'q': 16,
		'r': 17,
		's': 18,
		't': 19,
		'u': 20,
		'v': 21,
		'w': 22,
		'x': 23,
		'y': 24,
		'z': 25,
	}
	heightmax := int32(-1)
	r := []rune(word)
	for i := range r {
		idx := alphabet[r[i]]
		if heightmax < h[idx] {
			heightmax = h[idx]
		}
	}
	return int32(heightmax) * int32(len(word))
}
