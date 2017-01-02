package pascal

// Triangle computes Pascal's triangle up to a given number of rows.
func Triangle(n int) (tri [][]int) {
	for i := 0; i < n; i++ {
		if i == 0 {
			tri = append(tri, []int{1})
			continue
		}

		tri = append(tri, make([]int, i+1))

		for j := 0; j <= i; j++ {
			left := 0
			if j != 0 {
				left = tri[i-1][j-1]
			}

			right := 0
			if j != i {
				right = tri[i-1][j]
			}

			tri[i][j] = left + right
		}
	}

	return
}
