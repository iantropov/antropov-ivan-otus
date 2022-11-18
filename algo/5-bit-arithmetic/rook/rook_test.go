package rook

import (
	chessPiece "bit-arithmetic/chess-piece"
	countBits "bit-arithmetic/count-bits"
	"testing"
)

const ROOK_TEST_DIR = "/Users/antropov-ivan/Downloads/0.BITS/3.Bitboard - Ладья/"

func TestRookWithCountBitsStraightForward(t *testing.T) {
	chessPiece.TestChessPiece(ROOK_TEST_DIR, "rook/count-bits-straight-forward", PlaceRook, countBits.CountBitsStraightforward, t)
}

func TestRookWithCountBitsBySubstraction(t *testing.T) {
	chessPiece.TestChessPiece(ROOK_TEST_DIR, "rook/count-bits-straight-forward", PlaceRook, countBits.CountBitsBySubstraction, t)
}

func TestRookWithCountBitsWithPrecount(t *testing.T) {
	chessPiece.TestChessPiece(ROOK_TEST_DIR, "rook/count-bits-with-precount", PlaceRook, countBits.CountBitsWithPrecount, t)
}
