package main

/* Explaination of the race condition
1. Two go coroutines will access the same variable X
2. They can run in any order and the interleavings are not deterministic
3. While the value of x is being read and incremented say by coroutine1 it can be prempted
   and coroutine2 can read the old value increment and print. When couroutine1 resumes it will
   the old value and print it. Hence, it will not have access the the right and updated value of X
   so there might be inconsistency due to race-condition

   X=10
   C1 : *x = *x + 1 // Here X=10
   C2 : *x = *x + 1 // Here X=10
   C2 : *x = 10 + 1 // X = 11
   C2 : Prints 11
   C2 : Sleep
   C1 : *x = 10 +1 //Sees the old value
   C2 : Prints 11 // 10 + 1 based on the old value as C1 was pre-empted after it read X and before assigning value
*/

import (
	"fmt"
	"time"
)

func main() {
	v := 0

	go incrementAndPrint1(&v, "t1")
	go incrementAndPrint2(&v, "t2")

	var str string
	fmt.Scanln(&str)
	fmt.Println("Final value of v: ", v)

}

func incrementAndPrint1(x *int, name string) {
	for i := 0; i < 100; i++ {
		*x = *x + 1
		fmt.Println(name, *x)
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Println("End of ", name)
}

func incrementAndPrint2(x *int, name string) {
	for i := 0; i < 100; i++ {
		*x = *x + 1
		fmt.Println(name, *x)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("End of ", name)
}
