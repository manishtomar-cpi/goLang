package main

import (
	"fmt"
	"slices"
)

// dynamic array like vector
// most used construct in go
// + usedfull methods
func Slice() {

	//uninitialized slice is nil(null)
	var nums []int
	fmt.Println(nums) //[]
	// nums[0] = 1 // we cant do this because size is 0, only append function work
	fmt.Println(nums == nil) //true
	fmt.Println(len(nums))   //0
	nums = append(nums, 1)
	fmt.Println(nums) //[1]

	//declaring size of the slice
	//append -> add the digit at last
	// var digs = make([]int, 2)          // make([]type,size,capacity)
	// fmt.Println(digs)                  //[0,0]
	// fmt.Println(len(digs), " size")    //2
	// fmt.Println(cap(digs), "capacity") //2
	// digs = append(digs, 1)
	// fmt.Println(len(digs), " size")    //3
	// fmt.Println(cap(digs), "capacity") //4

	//capacity -> max numbers of element can fit by default it will as the same size

	// var numbers = make([]int, 2, 5)
	// fmt.Println(len(numbers), " size")    //2
	// fmt.Println(cap(numbers), "capacity") //5
	// numbers = append(numbers, 1)
	// fmt.Println(numbers) // [0,0,1]
	// numbers = append(numbers, 2)
	// numbers = append(numbers, 3)
	// fmt.Println(len(numbers), " size")    //5
	// fmt.Println(cap(numbers), "capacity") //5

	// numbers = append(numbers, 1)
	// fmt.Println(len(numbers), " size")    //6
	// fmt.Println(cap(numbers), "capacity") //10 -> means when we push after max capacity it will be double then existing

	//short hand decleration
	digs := []int{}
	fmt.Println(digs)
	fmt.Println(len(digs), "size")      //0
	fmt.Println(cap(digs), " capacity") //0
	digs = append(digs, 1)
	fmt.Println(digs)
	fmt.Println(len(digs), "size")      //1
	fmt.Println(cap(digs), " capacity") //1

	digs = append(digs, 2)
	fmt.Println(digs)
	fmt.Println(len(digs), "size")      //2
	fmt.Println(cap(digs), " capacity") //2

	digs = append(digs, 3)
	fmt.Println(digs)
	fmt.Println(len(digs), "size")      //3
	fmt.Println(cap(digs), " capacity") //4

}

func slice_methods() {

	fmt.Println("======COPY=======")
	var nums1 = make([]int, 0, 5)
	var nums2 = make([]int, len(nums1))

	nums1 = append(nums1, 1)
	copy(nums2, nums1)
	fmt.Println(nums1, nums2) // [1] [] because the size of nums2 is 0 - for copy the slice the destination size should be > 0

	var nums3 = make([]int, 0, 5)
	nums3 = append(nums3, 1)
	var nums4 = make([]int, len(nums3))
	copy(nums4, nums3)
	fmt.Println(nums3, nums4) // [1] [1] because the length of nums4 is 1

	nums3 = append(nums3, 2)
	copy(nums4, nums3)
	fmt.Println(nums3, nums4) // [1 2] [1] - because the size of nums4 is only 1

	fmt.Println("=====SLICE OPERATOR========")

	var nums = []int{1, 2, 3}
	fmt.Println(nums[0:2]) // 1,2 -> 2 is exclusive
	fmt.Println(nums[:2])  // 1,2 -> from is by default 0
	fmt.Println(nums[1:])  // 2,3 -> to now by default till last , still from is not exclusive

	fmt.Println("=======compare=======")
	var nums5 = []int{1, 2, 3}
	var nums6 = []int{1, 2, 3}

	fmt.Println(slices.Equal(nums5, nums6)) //true -> compare the values at each index not size

	// decleration of slices
	var array []int
	fmt.Println(array)

	var array1 = []int{1, 2, 3}
	fmt.Println(array1)

	var array2 = make([]int, 2, 5)
	fmt.Println(array2)
}

func main() {
	// Slice()
	slice_methods()
}
