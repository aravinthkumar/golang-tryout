package main

func main() {

	slice := []int{1, 2, 3}
	//loop over collection - slice
	for i, r := range slice {
		println(i, r)
	}
}
