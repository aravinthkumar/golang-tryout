package main

func main() {

	m := map[int]string{
		1: "abc",
		2: "xyz",
	}
	//loop over collection - map
	for i, r := range m {
		println(i, r)
	}
}
