package problems

func fill(image [][]int, sr int, sc int, newCol, sourceCol int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) ||
		image[sr][sc] != sourceCol {
		return
	}

	image[sr][sc] = newCol
	fill(image, sr+1, sc, newCol, sourceCol)
	fill(image, sr, sc+1, newCol, sourceCol)
	fill(image, sr-1, sc, newCol, sourceCol)
	fill(image, sr, sc-1, newCol, sourceCol)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}

	fill(image, sr, sc, newColor, image[sr][sc])
	return image
}
