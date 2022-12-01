package knight

func PlaceKnight(pos int) uint64 {
	var knightBits uint64 = 1 << pos

	var nA uint64 = 0xFeFeFeFeFeFeFeFe
	var nAB uint64 = 0xFcFcFcFcFcFcFcFc
	var nH uint64 = 0x7f7f7f7f7f7f7f7f
	var nGH uint64 = 0x3f3f3f3f3f3f3f3f

	var movesBits uint64 = nGH&(knightBits<<6|knightBits>>10) | // на b5 и b3
		nH&(knightBits<<15|knightBits>>17) | // на c6 и c2
		nA&(knightBits<<17|knightBits>>15) | // на e6 и e2
		nAB&(knightBits<<10|knightBits>>6) // на f5 и f3;

	return movesBits
}
