package main

import "fmt"

func Burbuja(s []int64) {
	//aux := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				aux := s[i]
				s[i] = s[j]
				s[j] = aux
			}
		}
	}
	fmt.Println(s)
}

func main() {

	//s := []int64{1, 6, 2, 4, 7, 0, -2, -1, 5, 3}
	fmt.Println(s)
	Burbuja(s)

}
