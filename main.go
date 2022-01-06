package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// read data.txt
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	textSlice := strings.Fields(string(content))
	numbers := atoi(textSlice)
	
	avg := average(numbers)
	mdn := median(numbers)
	vrnc := variance(numbers, avg)
	sDev := standardDeviation(vrnc)

	fmt.Println("Average:", int(math.Round(avg)))
	fmt.Println("Median:", mdn)
	fmt.Println("Variance:", int(math.Round(vrnc)))
	fmt.Println("Standard Deviation:", sDev)
}

func atoi(text []string) []float64{
	var numArray []float64

	for _, val := range text {
		numberInt, _ := strconv.Atoi(val)
		number := float64(numberInt)
		numArray = append(numArray, number)
	}
return numArray
}

// find average
func average(n []float64) float64 {
	var total float64
	for _, num := range n {
		total = total + num
	}
	a := (total / float64(len(n)))
	return a
}

// find median
func median(n []float64) int {
	// first we sort in a growing order the array of numbers
	sort.Float64s(n)
	var m int
	for range n {
		index := len(n) / 2
		// for even amount of numbers we find the median with two middle numbers
		if len(n)%2 == 0 {
			m = int(math.Round((n[index-1] + n[index]) / 2))

			// for odd numbers we find the middle number
		} else {
			m = int(n[index])
		}
	}
	return m
}

// find variance
func variance(n []float64, avrg float64) float64 {
	var sum float64
	for _, num := range n {
		difference := num - avrg
		square := math.Pow(difference, 2)
		sum += square
	}
	v := (sum / float64(len(n)))
	return v
}

// find standard deviation
func standardDeviation(variance float64) int {
	sDfloat := math.Round(math.Sqrt(variance))
	sD := int(sDfloat)
	return sD
}
