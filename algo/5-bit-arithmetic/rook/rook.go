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

	var moveBits uint64 = 0

	var bH uint64 = 0xfefefefefefefefe
	var cH uint64 = 0xfcfcfcfcfcfcfcfc
	var dH uint64 = 0xf8f8f8f8f8f8f8f8
	var eH uint64 = 0xf0f0f0f0f0f0f0f0
	var fH uint64 = 0xe0e0e0e0e0e0e0e0
	var gH uint64 = 0xc0c0c0c0c0c0c0c0
	var h uint64 = 0x8080808080808080

	var aG uint64 = 0x7f7f7f7f7f7f7f7f
	var aF uint64 = 0x3f3f3f3f3f3f3f3f
	var aE uint64 = 0x1f1f1f1f1f1f1f1f
	var aD uint64 = 0xf0f0f0f0f0f0f0f
	var aC uint64 = 0x707070707070707
	var aB uint64 = 0x303030303030303
	var a uint64 = 0x101010101010101

	moveBits = rookPos<<8 | rookPos<<16 | rookPos<<24 | rookPos<<32 | rookPos<<40 | rookPos<<48 | rookPos<<56 |
		rookPos>>8 | rookPos>>16 | rookPos>>24 | rookPos>>32 | rookPos>>40 | rookPos>>48 | rookPos>>56 |
		bH&(rookPos<<1) | cH&(rookPos<<2) | dH&(rookPos<<3) | eH&(rookPos<<4) | fH&(rookPos<<5) | gH&(rookPos<<6) | h&(rookPos<<7) |
		aG&(rookPos>>1) | aF&(rookPos>>2) | aE&(rookPos>>3) | aD&(rookPos>>4) | aC&(rookPos>>5) | aB&(rookPos>>6) | a&(rookPos>>7)

	return moveBits
}

func PlaceRookLikeAPro(pos int) uint64 {
	var a uint64 = 0xFF
	var b uint64 = 0x101010101010101
	return (a << ((pos >> 3) << 3)) ^ (b << (pos & 7))
}
