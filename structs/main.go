package main

import "fmt"

type contactinfo struct {
	email   string
	pincode int
}

type person struct {
	fname   string
	lname   string
	contact contactinfo
}

type personB struct {
	fname string
	lname string
	contactinfo
}

func main() {
	blex := person{
		fname: "Blex",
		lname: "Banderson",
		contact: contactinfo{
			email:   "blex@gmail.com",
			pincode: 100015,
		},
	}

	// kalex := personB{
	// 	fname: "kalex",
	// 	lname: "kanderson",
	// 	contactinfo: contactinfo{
	// 		email:   "blex@gmail.com",
	// 		pincode: 100015,
	// 	},
	// }

	blex.print()
	// blex.updatename("Billy")
	blexPointer := &blex
	blexPointer.updatenameUsingPtr("calex")
	// blex.updatenameUsingPtr("calex")		: does the same thing as statement above, since go sees the
	// receiver type of the function being called and gets the address of the calling struct implicitly

	var age int
	age = 54
	update(&age)
	print(age)
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updatename(new_fname string) {
	// p.fname = new_fname	: this will not update the value within the calling struct,
	// since p is a copy of the calling struct not the struct itself
}

func (pointerTOp *person) updatenameUsingPtr(new_fname string) {
	pointerTOp.print()
	(*pointerTOp).fname = new_fname
	// pointerTOp.fname = new_fname		: this gave the same result as the line above,
	// since a var of type struct is anyways a pointer(i.e. it stores an address), so the statements
	// (*pointerTop) and pointerTOp both point to the same location(***** i think)
}

func update(a *int) {
	fmt.Print(a)
	(*a) += 1
}
