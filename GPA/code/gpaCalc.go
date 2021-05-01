package main

import (
	"fmt"
	"strings"
)

var gradeInput string
var classInput string

var unweightedSum float64
var weightedSum float64

func main() {

	fmt.Println("How many classes are you taking?")

	// "var", name, type
	var numClasses int

	// taking user input
	fmt.Scanln(&numClasses)

	calculateGPA(numClasses)
}

func calculateGPA(numClasses int) (float64, float64) {

	// initialize grade variables
	unweightedA := 4.0
	unweightedAMinus := 3.7
	unweightedBPlus := 3.3
	unweightedB := 3.0
	unweightedBMinus := 2.7
	unweightedCPlus := 2.3
	unweightedC := 2.0
	unweightedCMinus := 1.7
	unweightedD := 1.0

	weightedA := 4.5
	weightedAMinus := 4.2
	weightedBPlus := 3.8
	weightedB := 3.5
	weightedBMinus := 3.2
	weightedCPlus := 2.8
	weightedC := 2.5
	weightedCMinus := 2.2
	weightedD := 1.5

	f := 0.0

	for i := 0; i < numClasses; i++ {
		fmt.Println("Regular, Honors, or AP?")
		fmt.Scanln(&classInput)
		classType := strings.ToUpper(classInput)

		fmt.Println("What (letter) grade did you receive?")
		fmt.Scanln(&gradeInput)
		grade := strings.ToUpper(gradeInput)

		if classType == "REGULAR" {
			if grade == "A" || grade == "A+" {
				unweightedSum += unweightedA
				weightedSum += unweightedA
			} else if grade == "A-" {
				unweightedSum += unweightedAMinus
				weightedSum += unweightedAMinus
			} else if grade == "B+" {
				unweightedSum += unweightedBPlus
				weightedSum += unweightedBPlus
			} else if grade == "B" {
				unweightedSum += unweightedB
				weightedSum += unweightedB
			} else if grade == "B-" {
				unweightedSum += unweightedBMinus
				weightedSum += unweightedBMinus
			} else if grade == "C+" {
				unweightedSum += unweightedCPlus
				weightedSum += unweightedCPlus
			} else if grade == "C" {
				unweightedSum += unweightedC
				weightedSum += unweightedC
			} else if grade == "C-" {
				unweightedSum += unweightedCMinus
				weightedSum += unweightedCMinus
			} else if grade == "D" {
				unweightedSum += unweightedD
				weightedSum += unweightedD
			} else if grade == "F" {
				unweightedSum += f
				weightedSum += f
			}
		}

		if classType == "HONORS" {
			if grade == "A" || grade == "A+" {
				weightedSum += weightedA
				unweightedSum += unweightedA
			} else if grade == "A-" {
				weightedSum += weightedAMinus
				unweightedSum += unweightedAMinus
			} else if grade == "B+" {
				weightedSum += weightedBPlus
				unweightedSum += unweightedBPlus
			} else if grade == "B" {
				weightedSum += weightedB
				unweightedSum += unweightedB
			} else if grade == "B-" {
				weightedSum += weightedBMinus
				unweightedSum += unweightedBMinus
			} else if grade == "C+" {
				weightedSum += weightedCPlus
				unweightedSum += unweightedCPlus
			} else if grade == "C" {
				weightedSum += weightedC
				unweightedSum += unweightedC
			} else if grade == "C-" {
				weightedSum += weightedCMinus
				unweightedSum += unweightedCMinus
			} else if grade == "D" {
				weightedSum += weightedD
				unweightedSum += unweightedD
			} else if grade == "F" {
				weightedSum += f
				unweightedSum += f
			}
		}

		if classType == "AP" {
			if grade == "A" || grade == "A+" {
				weightedSum += 5.0
				unweightedSum += unweightedA
			} else if grade == "A-" {
				weightedSum += 4.7
				unweightedSum += unweightedAMinus
			} else if grade == "B+" {
				weightedSum += 4.3
				unweightedSum += unweightedBPlus
			} else if grade == "B" {
				weightedSum += 4.0
				unweightedSum += unweightedB
			} else if grade == "B-" {
				weightedSum += 3.7
				unweightedSum += unweightedBMinus
			} else if grade == "C+" {
				weightedSum += 3.3
				unweightedSum += unweightedCPlus
			} else if grade == "C" {
				weightedSum += 3.0
				unweightedSum += unweightedC
			} else if grade == "C-" {
				weightedSum += 2.7
				unweightedSum += unweightedCMinus
			} else if grade == "D" {
				weightedSum += 2.0
				unweightedSum += unweightedD
			} else if grade == "F" {
				weightedSum += f
			}
		}

	}

	fmt.Println()
	unweightedGPA := unweightedSum / float64(numClasses)
	weightedGPA := weightedSum / float64(numClasses)
	fmt.Printf("Unweighted GPA: %.3f", unweightedGPA)
	fmt.Printf("\nWeighted GPA: %.3f\n\n", weightedGPA)

	return unweightedGPA, weightedGPA
}
