package main

import "fmt"

//for -> only constracts in go for looping

func while_loop() {
	i := 1
	// this will work as while loop
	for i <= 3 { // while condition while (i<=3) -> do the operstions iside the loop
		fmt.Println(i)
		i++
	}
}

// this the infinite loop syntax
// func infinite_loop() {
// 	for {
// 		println(1)
// 	}
// }

// classic for loop in go
func for_loop() {
	fmt.Println("for loop")
	for i := 1; i <= 3; i++ {
		if i == 2 {
			continue //continue - skip the current iteration
		}
		if i == 3 {
			break //break loop on 3
		}
		fmt.Println(i)
	}
}

func rangeInFor() {
	fmt.Println("range")
	for i := range 10 { // 10 is exclusive here print will be 0 to 9
		fmt.Println(i)
	}
}
func main() {
	while_loop()
	// infinite_loop()
	for_loop()
	rangeInFor()
}
