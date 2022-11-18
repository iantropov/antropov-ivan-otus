package bishop

func PlaceBishop(pos int) uint64 {
	var bishopPos uint64 = 1 << pos

	var a1h8 uint64 = 0x8040201008040201
	var leftUpperCorner uint64 = 0xff7f3f1f0f070301
	var rightBottomCorder uint64 = 0x80c0e0f0f8fcfeff

	var a8h1 uint64 = 0x102040810204080
	var rightUpperCorner uint64 = 0xfffefcf8f0e0c080
	var leftBottomCorner uint64 = 0x103070f1f3f7fff

	var firstMoveBits, secondMoveBits uint64 = 0, 0

	line := a1h8
	for i := 0; i < 7; i++ {
		if bishopPos&line > 0 {
			firstMoveBits |= line
			break
		}
		line = (line << 1) & rightBottomCorder
	}

	if firstMoveBits == 0 {
		line = a1h8
		for i := 0; i < 7; i++ {
			if bishopPos&line > 0 {
				firstMoveBits |= line
				break
			}
			line = (line >> 1) & leftUpperCorner
		}
	}

	if firstMoveBits == 0 {
		firstMoveBits = bishopPos
	}

	line = a8h1
	for i := 0; i < 7; i++ {
		if bishopPos&line > 0 {
			secondMoveBits |= line
			break
		}
		line = (line << 1) & rightUpperCorner
	}

	if secondMoveBits == 0 {
		line = a8h1
		for i := 0; i < 7; i++ {
			if bishopPos&line > 0 {
				secondMoveBits |= line
				break
			}
			line = (line >> 1) & leftBottomCorner
		}
	}

	if secondMoveBits == 0 {
		secondMoveBits = bishopPos
	}

	moveBits := firstMoveBits ^ secondMoveBits

	return moveBits
}
