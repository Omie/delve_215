package main

//#cgo CFLAGS: -std=gnu99
//#include <stdio.h>
// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
//		return f();
// }
//
// int fortytwo()
// {
//	    return 42;
// }
import "C"
import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	f := C.intFunc(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))
	// Output: 42
}
