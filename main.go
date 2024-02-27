package main

import (
	"fmt"
	"math"
)

//1. Write a variadic function in Go called multiply that takes an arbitrary number of integers as arguments and returns their product.
func multiply(numbers ...int) int {
	product := 1
	for _, num := range numbers {
		product *= num 
	}
	return product
}


//2.Write a variadic function in Go called average that takes an arbitrary number of floats as arguments and returns their average.
func average(numbers ...float64) float64 {
 	total := 0.0
	for _, num := range numbers {
	 total += num
	}
 	return total / float64(len(numbers))
 }

 //3.Write a variadic function in Go called min that takes an arbitrary number of integers as arguments and returns the smallest number.
 func min(numbers... int) int {
	min := math.MaxInt64
    for _, num := range numbers {
        if num < min {
            min = num
        }
    }
    return min
 }

 //4. Write a variadic function in Go called max that takes an arbitrary number of integers as arguments and returns the largest number.
 func max(values ...int) int {

	maxValue := math.MinInt64
	for _, value := range values {
	 if value > maxValue {
		maxValue = value
	 }
	}
	return maxValue
 }

//5. Write a variadic function in Go called concat that takes an arbitrary number of strings as arguments and returns their concatenation
func concat(values ...string) string {
		var result string

		for _, str := range values {
			result += str
		}
	return result
}

//6. Write a variadic function in Go called countOdd that takes an arbitrary number of integers as arguments and returns the count of odd numbers.
func countOdd(values ...int) int {
	count := 0
    for _, num := range values {
        if num%2 != 0 {
            count++
        }
    }
    return count

}

//7. Write a variadic function in Go called sumSquares that takes an arbitrary number of integers as arguments and returns the sum of their squares.
func sumSquares(nums ...int) int {
	sum := 0
	for _, num := range nums {
			sum += num * num
	}
	return sum
}

//8. Write a variadic function in Go called sumEven that takes an arbitrary number of integers as arguments and returns the sum of even numbers.
func sumEven(nums ...int) int {
	sum := 0
	for _, num := range nums {
			if num%2 == 0 {
					sum += num
			}
	}
	return sum
}

//9. Write a variadic function in Go called productNonZero that takes an arbitrary number of integers as arguments and returns the product of non-zero numbers
func productNonZero(nums ...int) int {
	product := 1
	for _, num := range nums {
			if num != 0 {
					product *= num
			}
	}
	return product
}

//10. Write a variadic function in Go called countPositive that takes an arbitrary number of integers as arguments and returns the count of positive numbers.
func countPositive(nums ...int) int {
	count := 0
	for _, num := range nums {
			if num > 0 {
					count++
			}
	}
	return count
}

func main() {
	result := multiply(1,5,7,8)
	fmt.Println(result)

	result2:= average(5,100)
	fmt.Println(result2)

	result3 := min(8,9,7)
	fmt.Println(result3)

	result4 := max(10,50,75,33)
	fmt.Println(result4)

	result5 := concat("Olá Isabela ", "que bom que você veio ", "Te amo!")
	fmt.Println(result5)

	result6 := countOdd(12,7,9,5,63,44)
	println(result6)

	result7 := sumSquares(5, 8, 55, 2)
  fmt.Println("Sum of squares:", result7) 

	result8 := sumEven(4, 62, 7, 8, 9)
	fmt.Println("Sum of even numbers:", result8) 

	result9 := productNonZero(1,22, 35, 0, 4, 5)
	fmt.Println("Product of non-zero numbers:", result9)

	result10 := countPositive(10, -2, 31, -4, -8, 9)
	fmt.Println("Count of positive numbers:", result10) 

}





