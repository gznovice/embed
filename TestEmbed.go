package main

import(
	"fmt"
)

type animal struct{
	name string
	leg int
}

func (me * animal) info(){
	fmt.Println("I'm " , me.name)
	fmt.Printf("I have %v leg(s)\n", me.leg)
}

type people struct{
	animal	
}

func (me * people) work(){
	fmt.Println("I can work")
}

func main(){
	pig := animal{
		name:"pig",
		leg:4}
	
	pig.info()
	

	/*p1 := people{
		name:"people",
		leg:2}*/// cannot use promoted field animal.name in struct literal of type people
		
	p1 := people{
		animal{
			name:"people",
			leg:2}}
		

	
	p1.info()
	
	p1.work()
}