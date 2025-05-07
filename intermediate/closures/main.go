package main

import "fmt"

func main() {
	// Keeps values of variables in memory
	sequence := counter()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	// Starts in the clean scope each time
	fmt.Println(counter()())
	fmt.Println(counter()())
	fmt.Println(counter()())
	fmt.Println(counter()())

	adder := func() func(int) int {
		var total int
		return func(i int) int {
			total += i
			return total
		}
	}

	total, odd, even := adder(), adder(), adder()
	for i := 1; i <= 10; i++ {
		t := total(i)
		fmt.Println("total:", t)
		if i&1 == 1 {
			o := odd(i)
			fmt.Println("odd:", o)
			if t%o == 0 {
				fmt.Println("which divides total, so add to the total:", total(o))
			}
		} else {
			e := even(i)
			fmt.Println("even:", e)
			if t%e == 0 {
				fmt.Println("which divides total, so add to the total:", total(e))
			}
		}
	}
}

func counter() func() int {
	i := 0
	fmt.Println("Start adder")
	return func() int {
		fmt.Println("prev value of i:", i)
		i++
		return i
	}
}
