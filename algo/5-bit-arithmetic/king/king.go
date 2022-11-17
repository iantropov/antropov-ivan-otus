package king

func PlaceKing(pos int) (uint64, uint64) {
	var kingMoves uint64 = 1 << pos

	var nA uint64 = 0xfefefefefefefefe
	var nH uint64 = 0x7f7f7f7f7f7f7f7f
	var h uint64 = 0x8080808080808080
	var a uint64 = 0x101010101010101

	var moveBits uint64 = (nH & kingMoves << 7) |
		kingMoves<<8 |
		(nA & kingMoves << 9) |
		(nA & kingMoves << 1) |
		(nH & kingMoves >> 1) |
		(nA & kingMoves >> 7) |
		kingMoves>>8 |
		(nA & kingMoves >> 9)

	return kingMoves, moveBits
}
