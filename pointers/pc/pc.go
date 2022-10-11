package pc

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPc Pc) PrintBrand() {
	fmt.Println(myPc.Brand)
}

func (myPc *Pc) DuplicateRam() {
	myPc.Ram = myPc.Ram * 2
}

func (myPc Pc) String() string {
	return fmt.Sprintf("%d GB of RAM, %d GB of Disk and it is a %s", myPc.Ram, myPc.Disk, myPc.Brand)
}
