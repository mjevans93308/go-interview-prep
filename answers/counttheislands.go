package answers

// based on https://leetcode.com/problems/number-of-islands/

func countTheIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}

	numIslands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				numIslands += 1
				clearTouchingLand(grid, i, j)
			}
		}
	}
	return numIslands
}

func clearTouchingLand(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	clearTouchingLand(grid, i+1, j)
	clearTouchingLand(grid, i-1, j)
	clearTouchingLand(grid, i, j+1)
	clearTouchingLand(grid, i, j-1)
}
