package day05

func ParseSeat(s string) int {
	row, col := 128, 8
	rowStep, colStep := row/2, col/2

	for _, chr := range s {
		switch string(chr) {
		case "F":
			row -= rowStep
			fallthrough
		case "B":
			rowStep /= 2
		case "L":
			col -= colStep
			fallthrough
		case "R":
			colStep /= 2
		}
	}
	row--
	col--

	return (row * 8) + col
}
