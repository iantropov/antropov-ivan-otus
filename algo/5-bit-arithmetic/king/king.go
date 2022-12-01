package king

func PlaceKing(pos int) uint64 {
	var kingMoves uint64 = 1 << pos

	var nA uint64 = 0xfefefefefefefefe
	var nH uint64 = 0x7f7f7f7f7f7f7f7f

	var moveBits uint64 = (nH & (kingMoves << 7)) |
		kingMoves<<8 |
		(nA & (kingMoves << 9)) |
		(nA & (kingMoves << 1)) |
		(nH & (kingMoves >> 1)) |
		(nA & (kingMoves >> 7)) |
		kingMoves>>8 |
		(nH & (kingMoves >> 9))

	return moveBits
}
