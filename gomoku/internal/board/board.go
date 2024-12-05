package board

const (
	boardSize = 15
)

type Board struct {
	b      [boardSize][boardSize]uint8
	filled int
}

func (b Board) positionIsValid(x, y int) bool {
	return x >= 0 && x < boardSize && y >= 0 && y < boardSize
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) Set(x, y int, value uint8) bool {
	if !b.positionIsValid(x, y) {
		return false
	}
	b.b[x][y] = value
	b.filled++
	return true
}

func (b Board) Get(x, y int) (uint8, bool) {
	if !b.positionIsValid(x, y) {
		return 0, false
	}
	return b.b[x][y], true
}

func (b *Board) Clear(x, y int) {
	if !b.positionIsValid(x, y) {
		return
	}
	if b.b[x][y] == 0 {
		return
	}
	b.b[x][y] = 0
	b.filled--
}

var empty = [boardSize][boardSize]uint8{}

func (b *Board) Reset() {
	copy(b.b[:], empty[:])
	b.filled = 0
}

func (b *Board) Size() (int, int) {
	return boardSize, boardSize
}

func (b Board) IsWinnerAfterSet(x, y int) (int, int, int, int, bool) {
	// horizontal
	if x0, y0, x1, y1, ok := b.fiveInDir(x, y, 1, 0); ok {
		return x0, y0, x1, y1, true
	}
	// vertical
	if x0, y0, x1, y1, ok := b.fiveInDir(x, y, 0, 1); ok {
		return x0, y0, x1, y1, true
	}
	// diagonal from top-left to bottom-right
	if x0, y0, x1, y1, ok := b.fiveInDir(x, y, 1, 1); ok {
		return x0, y0, x1, y1, true
	}
	// diagonal from bottom-left to top-right
	if x0, y0, x1, y1, ok := b.fiveInDir(x, y, 1, -1); ok {
		return x0, y0, x1, y1, true
	}
	return 0, 0, 0, 0, false
}

func (b Board) fiveInDir(x, y, stepX, stepY int) (int, int, int, int, bool) {
	p, ok := b.Get(x, y)
	if !ok {
		return 0, 0, 0, 0, false
	}
	if p == 0 {
		return 0, 0, 0, 0, false
	}

	count := 0

	x0, y0 := x, y
	i, j := x, y
	for {
		cell, ok := b.Get(i, j)
		if cell != p || !ok {
			break
		}
		x0, y0 = i, j
		count++
		i, j = i+stepX, j+stepY
	}

	x1, y1 := x, y
	i, j = x-stepX, y-stepY
	for {
		cell, ok := b.Get(i, j)
		if cell != p || !ok {
			break
		}
		x1, y1 = i, j
		count++
		i, j = i-stepX, j-stepY
	}

	return x0, y0, x1, y1, count == 5
}

func (b Board) Filled() bool {
	return b.filled == boardSize*boardSize
}
