package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("Nhao vao canh cua tam giac:")
	fmt.Scanln(&a, &b, &c)
	if a > 0 && b > 0 && c > 0 && a < b+c && b < a+c && c < a+b {
		a_mu := math.Pow(a, 2)
		b_mu := math.Pow(b, 2)
		c_mu := math.Pow(c, 2)
		if a_mu == b_mu+c_mu || b_mu == a_mu+c_mu || c_mu == a_mu+b_mu {
			fmt.Println("Tam Giac Vuong")
		}
		if a == b || b == c || c == a {
			fmt.Println("Tam Giac Can")
		}
		if a == b && b == c {
			fmt.Println("Tam Giac Deu")
		}
	} else {
		fmt.Println("Khong a b c khong phai la 3 canh cua  1 tam giac")
	}
}
