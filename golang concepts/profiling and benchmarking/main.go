package main

func MemoryWaste(n int) []int {
	var data []int

	for i := 0; i < n; i++ {
		data = append(data, i)
	}

	return data
}
