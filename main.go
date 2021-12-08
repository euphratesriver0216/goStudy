package main

func cap1(){
	f := make([]func(), 3) //함수 리터럴 3개를 가진 슬라이스
	fmt.Println("캡1")
	for i := 0; i 3; i++ {
		f[i] = func(){
			fmt.Println(i)
		}
	}
	for i :=0; i < 3; i++{
		v := i //v 변수에 i값복사

		f[i] = func(){
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++{
		f[i]()
	}
}

func cap2(){
	f := make([]func(), 3) //함수 리터럴 3개를 가진 슬라이스
	fmt.Println("캡2")
	for i := 0; i 3; i++ {
		f[i] = func(){
			fmt.Println(i)
		}
	}
	for i :=0; i < 3; i++{
		v := i //v 변수에 i값복사

		f[i] = func(){
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++{
		f[i]()
	}
}

func main() {
	//함수 리터럴 외부변수를 내부로 가져오는 것을 캡쳐
	// 캡쳐는 값복사가 아니라 래퍼런스 (참조)형태로 가져온다.
	// 변수의 주소를 포인터 값으로 복사한다.


}
