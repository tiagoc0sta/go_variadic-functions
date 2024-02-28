/*package main

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

}*/

/*/// Recursion examples

package main

import "fmt"

func factorial(n int) int {

 if n == 0 {
  return 1 //base case: 0! is 1
 }

 return n * factorial(n-1) //recursive call to calculate (n-1)

}

func main() {

 result := factorial(5)
 fmt.Println("Factorial of 5:", result)
 result = factorial(10)
 fmt.Println("Factorial of 10:", result)
}*/

/*package main

import "fmt"

//Recursive function to find the sum of array elements

func sumArray(arr []int, n int) int {

 if n <= 0 {

  return 0 //base case: array is empty

 }

 return sumArray(arr, n-1) + arr[n-1] //recursive call: sum of n-1 element plus nth element

}

func main() {

 arr := []int{1, 2, 3, 4, 5}
 result := sumArray(arr, len(arr))

 fmt.Println("Sum of array elements: ", result)

}
*/

/*package main

import "fmt"

func reverse(s string) string {

 if len(s) == 0 {

  return s //base case: empty string

 }

 //recursive call: reverse substring excluding first character, then add first character to the end

 return reverse(s[1:]) + string(s[0])

}

func main() {

 str := "hello"

 result := reverse(str)

 fmt.Println("Reversed string: ", result)

}*/

//////////////

package main

import (
	"fmt"
)

//Function declaration and argument
//calculate total cost of trip
func calculateTotal(accommodationCost, transportCost, mealCost, discount float64) (totalCost float64, finalCostAfterDiscount float64) {
    totalCost = accommodationCost + transportCost + mealCost
    finalCostAfterDiscount = totalCost - (totalCost * discount)
    return
}
//Variadic functions
//Calculate accommodation cost for variable number of nights and rates
func calculateAccommodationCost(nightsRates ...float64) float64 {
    totalCost := 0.0
    for _, rate := range nightsRates {
        totalCost += rate
    }
    return totalCost
}
//Recursion
//Calculate discount for group bookings: 5% for each traveller after the first, capped at 20%
func calculateGroupDiscount(travelers int) float64 {
    if travelers <= 1 {
        return 0.0
    } else if travelers >= 5 {
        return 0.20 //maximum discount cap
    }
    return 0.05 + calculateGroupDiscount(travelers-1)
}
func main() {
    //Anonymous function
    //Apply a special offer if travelling to a certain destination with more than 3 travelers
    specialOffer := func(destination string, travelers int, baseCost float64) float64 {
        if destination == "Paris" && travelers > 3 {
            return baseCost * 0.9 //10% off for groups larger than 3 to paris
        }
        return baseCost
    }
    //Example booking: 3 nights in Paris with variable rates, 4 travellers
    nightsRates := []float64{100.0, 120.0, 110.0} //different rates per night
    accommodationCost := calculateAccommodationCost(nightsRates...)
    transportCost := 500.0 //flat transport cost
    mealCost := 200.0 //estimated meal cost
    travelers := 4
    destination := "Paris"
    discount := calculateGroupDiscount(travelers)
    totalCost, finalCostAfterDiscount := calculateTotal(accommodationCost, transportCost, mealCost, discount)
    finalCostAfterSpecialOffer := specialOffer(destination, travelers, finalCostAfterDiscount)
    fmt.Printf("Total cost before discount: $%.2f\n", totalCost)
    fmt.Printf("Total cost after discount: $%.2f\n", finalCostAfterDiscount)
    fmt.Printf("Total cost after special offer: $%.2f\n", finalCostAfterSpecialOffer)
}


 