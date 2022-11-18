package bishop

import (
	chessPiece "bit-arithmetic/chess-piece"
	countBits "bit-arithmetic/count-bits"
	"testing"
)

const BISHOP_TEST_DIR = "/Users/antropov-ivan/Downloads/0.BITS/4.Bitboard - Слон/"

func TestBishopWithCountBitsStraightForward(t *testing.T) {
	chessPiece.TestChessPiece(BISHOP_TEST_DIR, "bishop/count-bits-straight-forward", PlaceBishop, countBits.CountBitsStraightforward, t)
}

func TestBishopWithCountBitsBySubstraction(t *testing.T) {
	chessPiece.TestChessPiece(BISHOP_TEST_DIR, "bishop/count-bits-straight-forward", PlaceBishop, countBits.CountBitsBySubstraction, t)
}

func TestBishopWithCountBitsWithPrecount(t *testing.T) {
	chessPiece.TestChessPiece(BISHOP_TEST_DIR, "bishop/count-bits-with-precount", PlaceBishop, countBits.CountBitsWithPrecount, t)
}
