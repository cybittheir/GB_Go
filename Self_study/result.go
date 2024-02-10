package main

import "fmt"

func main() {
	text := [...]string{"first", "second", "third", "fourth", "fifth"}
	done := make(chan bool)
	for _, v := range text {
		print(v, " :: ")
	}
	print("\n\n")
	for i, _ := range text {
		go func(s *string, f float64) {
			*s = fmt.Sprintf("%f", f)
			done <- true
		}(&text[i], float64(i))
	}
	var g int
	go func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
			fmt.Printf("\n-= %d =-\n", s)
		}
		g = s
		fmt.Printf("\n== %d ==\n", g)
	}(1000)
	for z, _ := range text {
		fmt.Printf(">> %d <<\n", z)
		<-done
	}
	for _, v := range text {
		print(v, " ")
	}
	print("\n")
}
