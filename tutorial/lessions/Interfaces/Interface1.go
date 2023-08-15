package main

import "fmt"

type Shaper1 interface {
	Area() float32
}

type Square1 struct {
	side float32
}

func (sq *Square1) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square1{5}     // Area() of Square1 needs a pointer
	// shapes := []Shaper1{Shaper1(r), Shaper1(q)}
	// or shorter:
	shapes := []Shaper1{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}

/*
type `Reader` interface {
  Read(p []byte) (n int, err error)
}

var r io.Reader

r = os.Stdin
r = bufio.NewReader(r)
r = new(bytes.Buffer)
f,_ := os.Open("test.txt")
r = bufio.NewReader(f)
*/
