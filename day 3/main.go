package main

import (
	"fmt"
)

// for loop
// func main(){
// 	for i := 0; i<5; i++{
// 		fmt.Println(i)
// 	}
// }
// func main(){
// 	s := sum(1,2,3,4,5,6,7,8,9)
// 	fmt.Println("The sum is: ", s)
// }

// func sum(values ...int) int{
// 	fmt.Println(values)
// 	results := 0
// 	for  _, v := range values{
// 		results += v
// 	}
// 	return results
// }

// func main(){
// 	d, err := divide(5.0, 2.0)
// 	if err != nil{
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(d)
// }

// func divide (a,b float64)(float64, error){
// 	if b == 0.0{
// 		return 0.0, fmt.Errorf("Cannot divide by Zero")
// 	}
// 	return a/b, nil
// }

// func main() {
// 	func ()  {
// 		fmt.Println("Hello from Anonymous Function")
// 	}()
// }

// func main(){
// 	g := greeter{
// 		greeting: "Hello",
// 		name: "Isomiddin",
// 	}
// 	g.greet()
// }

// type greeter struct{
// 	greeting string
// 	name string
// }
// func (gsome greeter) greet() {
// 	fmt.Println(gsome.greeting, gsome.name)
// }
// func main(){
// 	myInt := IntCounter(0)
// 	var inc Incrementer = &myInt
// 	for i := 0; i<10; i++{
// 		fmt.Println(inc.Increment())
// 	}
// }

// type IntCounter int

// func (ic *IntCounter) Increment() int{
// 	*ic++
// 	return int(*ic)
// }

// type Incrementer interface{
// 	Increment() int
// }

func main(){
	fmt.Println("This is composed inerfaces")
}

type Writer interface{
	Writer([]byte)(int, error)
}

type Closer interface{
	Close() error
}

type WriterCloser interface{
	Writer
	Closer
}
