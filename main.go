//슬라이스 (일종의 자료구조)
package main

import "fmt"

func main() {
	// 모든 언어는 자료구조를 중요하게 여긴다.
	//데이터를 얼마만큼 효율적으로 메모리를 사용하느냐가 중요
	//슬라이스 : 동적배열 (자동으로  배열크기를 증가시켜준다. / 가변적이다)
	//슬라이싱을 이용해 데이터 일부를 불러 올 수 있다.
	//길이가 요소 개수에 따라 자동으로 증가해 관리가 편하다.
	//슬라이싱 기능을 사용해서 배열의 일부를 나타내는 슬라이스를 만들 수 있다.
	// 불편한거를 한방에 처리 할 수 있다.

	//var arr[10] int  10개까지 만듦 //인티저 타입의 배열 10개를 생성// 배열을 0

	// 슬라이스를 초기화 하지 안으면 길이가 0인 슬라이스가 만들어진다.
	var slice []int // 슬라이스 배열 선언 //매우간단함

	if len(slice) == 0 {
		fmt.Println("슬라이스가 비어 있다.", slice)
	}
	// slice[1] = 10 // 할당되지 않음 메모리 공간에 접근  -> 에러가 발생한다.
	// fmt.Println(slice)

	//{}를 사용한 초기화
	//var slice1 = []int {1,2,3}
	var slice2 = []int{1, 5: 2, 10: 3}
	fmt.Println(slice2)

	//var arr = [...]int{1,2,3}//이것이 배열
	//var slice = []int{1,2,3}//슬라이스

	//make() 함수를 이용한 초기화
	//var slice3 = make([]int, 3) //내장되어 있는 함수임 이 함수를 사용해서 첫번째 인수로 타입을 넣어주고, 두번째 인수는 길이를 넣어준다.

	//슬라이스 접근 : 배열과 같다
	var slice4 = make([]int, 3)
	slice4[1] = 5 //[0,5,0]

	var slice5 = make([]int, 5)

	//슬라이스 순회  -> 동적으로 길이가 배열과 동일함.
	//초기화
	for i := 0; i < len(slice5); i++ {
		slice5[i] = i + 1
	}
	//출력
	for _, v := range slice5 {
		fmt.Println((v))
	}

	//요소추가
	//
	var slice6 = []int{1, 2, 3}
	//첫번째 인수 : 추가하려는 슬라이스, 두번째 인수에는 추가하려는 요소
	slice7 := append(slice6, 4) //추가

	fmt.Println(slice6)
	fmt.Println(slice7)
	//append : 첫번째 인수로 들어온 슬라이스이 값을 변경하는게 아니라 요소가 추가된 새로운 슬라이스를 반환
	//기존 슬라이스에 요소를 추가하고 싶다면 append()결과를 기존 슬라이스에 대입해서 변경해야함

	var slice8 []int
	//포문으로 하나씩 추가
	for i := 0; i <= 10; i++ {
		slice8 = append(slice8, i)
	}
	slice8 = append(slice8, 11, 3, 315, 123)
	fmt.Println(slice8)

	// type sliceHeader struct {
	// 	Data uintptr // 실제 배열을 가리키는 포인터
	// 	Len  int     // 요소 개수    (length : len 길이의 줄임말 )
	// 	Cap  int     // 실제 배열의 길이  (capacity: cap 용량의 줄임말)
	// }

	//포인터를 사용하면 연산 속도가 빠르다.

}
