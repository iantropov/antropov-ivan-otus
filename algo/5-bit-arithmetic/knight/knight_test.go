package knight

import (
	chessPiece "bit-arithmetic/chess-piece"
	countBits "bit-arithmetic/count-bits"
	"testing"
)

const KNIGHT_TEST_DIR = "/Users/antropov-ivan/Downloads/0.BITS/2.Bitboard - Конь/"

func TestKnightWithCountBitsStraightForward(t *testing.T) {
	chessPiece.TestChessPiece(KNIGHT_TEST_DIR, "knight/count-bits-straight-forward", PlaceKnight, countBits.CountBitsStraightforward, t)
}

func TestKnightWithCountBitsBySubstraction(t *testing.T) {
	chessPiece.TestChessPiece(KNIGHT_TEST_DIR, "knight/count-bits-straight-forward", PlaceKnight, countBits.CountBitsBySubstraction, t)
}
