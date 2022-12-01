package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func a() int {

	in, err := ioutil.ReadFile("elves")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	strings := strings.Split(string(in), "\n")
	var maximumCalories int
	var totalCalories int
	for _, s := range strings {
		if s == "" {
			if maximumCalories < totalCalories {
				maximumCalories = totalCalories
			}
			totalCalories = 0
		} else {
			number, err := strconv.Atoi(s)
			if err != nil {
				fmt.Print(err.Error())
			}
			totalCalories = totalCalories + number
		}
	}
	return maximumCalories
}

func b() int {
	in, err := ioutil.ReadFile("elves")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	strings := strings.Split(string(in), "\n")

	var maximumCalories []int
	var counter int
	maximumCalories = append(maximumCalories, 0)
	for _, s := range strings {
		if s == "" {
			counter = counter + 1
			maximumCalories = append(maximumCalories, 0)
		} else {
			number, err := strconv.Atoi(s)
			if err != nil {
				fmt.Print(err.Error())
			}
			maximumCalories[counter] = maximumCalories[counter] + number
		}
	}
	sort.Ints(maximumCalories[:])
	return maximumCalories[len(maximumCalories)-1] + maximumCalories[len(maximumCalories)-2] + maximumCalories[len(maximumCalories)-3]
}
