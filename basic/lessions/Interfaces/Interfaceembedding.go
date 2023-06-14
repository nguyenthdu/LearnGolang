package main

// type de tranh loi
type Buffer struct {
}

type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

// v := varI.(T) // unchecked type assertion
/*
if v, ok := varI.(T); ok { // checked type assertion
Process(v)
return
}
// here varI is not of type T

=============
if v, ok := varI.(T); ok {
// ...
}
*/
func main() {

}
