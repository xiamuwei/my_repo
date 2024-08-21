package main

import "fmt"

func Change_array(array1 [5]string) [5]string{
	for index, value :=range array1 {
		if value == "weak" {
			array1[index] = "strong"
		} else if value == "stupid"{
			array1[index] = "smart"
		}
	}
	fmt.Println(array1)
	return array1

}