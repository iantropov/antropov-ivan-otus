package king

import (
	chessPiece "bit-arithmetic/chess-piece"
	countBits "bit-arithmetic/count-bits"
	"testing"
)

const KING_TEST_DIR = "/Users/antropov-ivan/Downloads/0.BITS/1.Bitboard - Король/"

func TestKingWithCountBitsStraightForward(t *testing.T) {
	chessPiece.TestChessPiece(KING_TEST_DIR, "king/count-bits-straight-forward", PlaceKing, countBits.CountBitsStraightforward, t)
}

func TestKingWithCountBitsBySubstraction(t *testing.T) {
	chessPiece.TestChessPiece(KING_TEST_DIR, "king/count-bits-straight-forward", PlaceKing, countBits.CountBitsBySubstraction, t)
}

func TestKingWithCountBitsWithPrecount(t *testing.T) {
	chessPiece.TestChessPiece(KING_TEST_DIR, "king/count-bits-with-precount", PlaceKing, countBits.CountBitsWithPrecount, t)
}