package main

// find average given function parameters 
// no return statments allowed
func average(a int , b int , c *int){

	average := (a+b)/2
	*c = average
}

// find min and maxmimum from the given function parameters

func min_max(a, b, c int) (int, int) {
	
    min := a
    max := a

    if b < min {
        min = b
    } else if b > max {
        max = b
    }

    if c < min {
        min = c
    } else if c > max {
        max = c
    }

    return min, max
}
