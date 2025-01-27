package main

func main() {
	x := 0
	for x < 5 {
		println(x)
		x++
	}

	for i := 0; i < 5; i++ {
		println(i)
	}

	names := []string{"ted", "fred", "free", "lance"}

	for i := 0; i < len(names); i++ {
		println(names[i])
	}

	for index, name := range names {
		println(index, name)
	}

	for _, name := range names {
		println(name)
	}
}
