package main

import "fmt"

func main(){
// define variable a
// define variable b
// define variable c
var a,b,c int;

// read three variable inputs from console
// use fmt.Scan to read 
// use fmt.Println to write
fmt.Scan(&a,&b,&c);
fmt.Println(a,b,c);

// call add function 
// print result here

fmt.Println( add(a,b,c) );

// call average function from second.go
// without returing a variable get the result here
// modifying average functions deinfiniton is allowed

average(a,b,&c);

// print the average here

fmt.Println( c );

// call min_max function from second.go
// print min and max here

min_val , max_val := min_max(2,5,1)
fmt.Println(min_val , max_val)

// create a variable radius and assign value of `b` to it
// call calculate_circle_area from third.go
var radius float64 = float64(b)
calculate_circle_area(radius)
}