package main

import (
	"fmt"
)

type car struct {
	color string
	speed int
}

const road_long int = 10

func main() {
//	x:= 10
//pointerTutorials(&x)
//	fmt.Println(x)

fmt.Println("")
a := car{
	color:"mohamed",speed:10.0,
}
ad :=car{color:"read",speed:200}
fmt.Print(get(a),get(ad))
pointerreciver(&ad)
fmt.Println(ad.speed)
ad.newspeed(10)

slice := [] int {1,2,3}
b := slice[:]
adas :=  slice[2:3]
mk := make([]int,3)
println("\n")
adasasf := append(mk, 1)
println(adasasf)
fmt.Print(adas)
fmt.Print(b)
fmt.Println(slice);


}

func get(c car) int {
	return int(c.speed*road_long)
}
func pointerreciver(car2 *car)  {
car2.speed = 10
}
func pointerTutorials(x* int){
	i := 8
	a := &i
	*a = 10
	fmt.Println(i)
	*x = 11
}
func (cars *car)newspeed(a int){
	cars.speed = a
}
func more(m ... int) int{
	total := 0;
 for _,a := range m{
 	total += a
 }
	return  total;
}
func print(){
	s := 10
	for index := 0; index < s; index++ {
		fmt.Println(index)
	}

}
func print2(){
	m := "mohamed"
	for _, c := range m  {
		fmt.Print(string(c)+"");
	}
}
/*
		if s == 11 {
			fmt.Println(s)
		} else {
			fmt.Println("Not right")
		}


	switch s {
	case 10:
		fmt.Println(s)
	case 11:
		fmt.Println("Wellcome")

	default:
		fmt.Println("Ft")
	}*/
