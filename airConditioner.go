// package main

// import "fmt"

// func main() {
// 	temp := 19

// 	if temp > 28 {  //if 문의 컨디션 는 boolean type
// 		fmt.Println("에어컨을 켠다")
// 	} else if temp <= 3 {
// 		fmt.Println("히터를 켠다")
// 	} else if temp <= 18 {
// 		fmt.Println("나가자")
// 	} else {
// 		fmt.Println("많이 덥네..")
// 	}

// }

package main

import "fmt"

func main() {

	temp := 40

	if temp > 29 {
		fmt.Println("에어컨을 틀어주세요")
	} else if temp <= 5 {
		fmt.Println("히터를 틀어보자")
	} else if temp <= 18 {
		fmt.Println("날이 적당해서 나가자")
	} else {
		fmt.Println(" 덥다잉")
	}
}

/*

제어문

조건문 - if / switch
반복문 - for  (go 에는 while / do while문이 없다 )


*/
