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

func PlaceRook2(pos int) uint64 {
	var rookPos uint64 = 1 << pos

	// var horizontal uint64 = 0xff
	// var vertical uint64 = 0x101010101010101
	var n1 uint64 = 0x101010101010101
	var moveBits uint64 = 0

	moveBits = rookPos |
		rookPos<<8 | rookPos<<16 | rookPos<<24 | rookPos<<32 | rookPos<<40 | rookPos<<48 | rookPos<<56 |
		rookPos>>8 | rookPos>>16 | rookPos>>24 | rookPos>>32 | rookPos>>40 | rookPos>>48 | rookPos>>56 |
		((rookPos << 1) & n1) | ((rookPos << 2) & (n1 | n1<<1)) | ((rookPos << 3) & (n1 | n1<<1 | n1<<2))

	moveBits ^= rookPos

	return moveBits
}
