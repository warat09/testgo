package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	fmt.Print("Enter the number of elements: ")
	fmt.Scanln(&n1)

	arr1 := make([]string, n1)
	var dup []string

	fmt.Println("Enter the elements:")
	for i := 0; i < n1; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scanln(&arr1[i])
	}
	fmt.Print("Enter the number of elements2: ")
	fmt.Scanln(&n2)

	arr2 := make([]string, n2)
	for i := 0; i < n2; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scanln(&arr2[i])
	}

	var arr3 = arr1
	var check bool
	fmt.Println(check)

	for i := 0; i < n1; i++ {
		check = false
		for j := 0; j < n2; j++ {
			if arr1[i] == arr2[j] {
				dup = append(dup, arr2[j])
				check = true
			}

		}
		if check == false {
			arr3 = append(arr3, arr2[i])
		}
	}

	fmt.Println("The elements1 you entered are ", arr1)
	fmt.Println("The elements2 you entered are ", arr2)
	fmt.Println("Merge is ", arr3)
	fmt.Println("Duplicate is ", dup)
}
