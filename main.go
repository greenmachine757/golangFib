package main

import "fmt"

func fibSeq(x uint) {
	fmt.Println("Printing first " + fmt.Sprint(x) + " entries of the Fibonacci Sequence");
	if (x>=1) { fmt.Println("0") }
	if (x>=2) { fmt.Println("1") }
	x-=2
	pa:=0
	pb:=1
	for x>0 {
		pc:=pa+pb
		fmt.Println(fmt.Sprint(pc))
		pa=pb
		pb=pc
		x--
	}
}

func main() {
	fibSeq(10) // Non-recusion version, can provide that version as well if desired
}
