package rook

func PlaceRook(pos int) uint64 {
	var rookPos uint64 = 1 << pos

	var horizontal uint64 = 0xff
	var vertical uint64 = 0x101010101010101
	var moveBits uint64 = 0

	for i := 0; i < 8; i++ {
		if rookPos&horizontal > 0 {
			moveBits |= horizontal
			break
		}
		horizontal <<= 8
	}

	for i := 0; i < 8; i++ {
		if rookPos&vertical > 0 {
			moveBits ^= vertical
			break
		}
		vertical <<= 1
	}

	return moveBits
}
