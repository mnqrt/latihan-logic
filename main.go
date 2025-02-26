package main

import (
	"github.com/mnqrt/go-print-slice/printslice"
	"github.com/mnqrt/latihan-logic/lat1"
	"github.com/mnqrt/latihan-logic/lat2"
	"github.com/mnqrt/latihan-logic/lat3"
)

func main() {
	printslice.PrintSlice(lat1.No1(10))
	printslice.PrintSlice(lat1.No5(10))
	printslice.PrintSlice(lat1.No9(10))
	
	printslice.Print2DSlice(lat2.No1(9))
	printslice.Print2DSlice(lat2.No2(9))
	printslice.Print2DSlice(lat2.No3(9))
	printslice.Print2DSlice(lat2.No4(9))
	printslice.Print2DSlice(lat2.No6(9))
	printslice.Print2DSlice(lat2.No9(9))
	printslice.Print2DSlice(lat2.No10(9))
	printslice.Print2DSlice(lat2.No12(9))
	printslice.Print2DSlice(lat2.No13(9))

	printslice.Print2DSlice(lat3.No1(9))
	printslice.Print2DSlice(lat3.No2(9))
	printslice.Print2DSlice(lat3.No3(9))
	printslice.Print2DSlice(lat3.No4(9))
	printslice.Print2DSlice(lat3.No5(9))
	printslice.Print2DSlice(lat3.No6(9))
	printslice.Print2DSlice(lat3.No7(9))
	printslice.Print2DSlice(lat3.No8(9))
	printslice.Print2DSlice(lat3.No9(9))
	printslice.Print2DSlice(lat3.No10(9))
	printslice.Print2DSlice(lat3.No11(9))
	printslice.Print2DSlice(lat3.No11B(9))
	printslice.Print2DSlice(lat3.No12(9))
	printslice.Print2DSlice(lat3.No12B(9))
	printslice.Print2DSlice(lat3.No14(5))
}