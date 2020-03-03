package main

type DimensionalArrays struct {
	Table [][]int
}

func LogDimensionalArrays() {
	v := &DimensionalArrays{Table: make([][]int, 10)}
	for i := 0; i < 10; i++ {
		v.Table[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			v.Table[i][j] = i * j
		}
	}

	Output(v)
}
