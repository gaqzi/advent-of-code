package day03

const tree = "#"

func Traverse(geography []string, right, down int64) int64 {
	var treesHit, posX, posY int64
	rows := int64(len(geography)) - 1
	cols := int64(len(geography[0]))

	for posY < rows {
		posX = (posX + right) % cols
		posY += down

		if hasTree(geography, posX, posY) {
			treesHit++
		}
	}

	return treesHit
}

func hasTree(geography []string, x, y int64) bool {
	return string(geography[y][x]) == tree
}
