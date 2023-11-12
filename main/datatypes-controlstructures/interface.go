package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

type Tesla struct {
	TeslaModel         string
	IsTeslaSelfDriving bool
}

type Xaiomi struct {
	XaiomiModel         string
	IsXaiomiSelfDriving bool
}

func (t *Tesla) Drive() {
	fmt.Println("go-go-tesla")
	fmt.Println(t.IsTeslaSelfDriving)
}

func (x *Xaiomi) Drive() {
	fmt.Println("go-go-xaiomi")
	fmt.Println(x.IsXaiomiSelfDriving)
}

func main() {
	//interfaces
	fmt.Println("*************interfaces***************")

	t := Tesla{"Model y", true}
	x := Xaiomi{"Mi xDrive", false}
	t.Drive()
	x.Drive()

}
