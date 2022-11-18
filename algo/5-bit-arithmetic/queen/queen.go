package queen

import (
	"bit-arithmetic/bishop"
	"bit-arithmetic/rook"
)

func PlaceQueen(pos int) uint64 {
	rookMoves := rook.PlaceRook(pos)
	bishopMoves := bishop.PlaceBishop(pos)
	return rookMoves | bishopMoves
}
