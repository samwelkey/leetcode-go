func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
    
    if obstacleGrid[0][0] == 1{
        return 0
    }

	colF := false
	rowF := false
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 || colF {
			obstacleGrid[i][0] = 0
			colF = true
		} else {
			obstacleGrid[i][0] = 1
		}
	}

	for j := 1; j < len(obstacleGrid[0]); j++ {

		if obstacleGrid[0][j] == 1 || rowF {
			obstacleGrid[0][j] = 0
			rowF = true
		} else {
			obstacleGrid[0][j] = 1
		}

	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}

		}
	}

	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]

}
