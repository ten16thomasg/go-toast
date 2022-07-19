package main

import (
	"flag"
	"fmt"
	"reflect"

	"github.com/gen2brain/beeep"
)

func main() {
	//---Define the various flags---
	msgPtr := flag.String("msg", "Default BEEP BEEP!",
		"Enter a Custom Message")
	imgPth := flag.String("img", "assets/information.png",
		"Enter a Custom Path")
	titleNm := flag.String("title", "Title",
		"Enter a Custom Title")
	//---parse the command line into the defined flags---
	flag.Parse()
	err := beeep.Notify(*titleNm, *msgPtr, *imgPth)
	if err != nil {
		panic(err)
	}
	for _, arg := range flag.Args() {
		fmt.Print(arg, " ")
		fmt.Println(reflect.TypeOf(arg))
	}
}
