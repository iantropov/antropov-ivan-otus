package queen

import (
	chessPiece "bit-arithmetic/chess-piece"
	countBits "bit-arithmetic/count-bits"
	"testing"
)

const QUEEN_TEST_DIR = "/Users/antropov-ivan/Downloads/0.BITS/5.Bitboard - Ферзь/"

func TestBishopWithCountBitsStraightForward(t *testing.T) {
	chessPiece.TestChessPiece(QUEEN_TEST_DIR, "queen/count-bits-straight-forward", PlaceQueen, countBits.CountBitsStraightforward, t)
}

func TestBishopWithCountBitsBySubstraction(t *testing.T) {
	chessPiece.TestChessPiece(QUEEN_TEST_DIR, "queen/count-bits-straight-forward", PlaceQueen, countBits.CountBitsBySubstraction, t)
}

func TestBishopWithCountBitsWithPrecount(t *testing.T) {
	chessPiece.TestChessPiece(QUEEN_TEST_DIR, "queen/count-bits-with-precount", PlaceQueen, countBits.CountBitsWithPrecount, t)
}
