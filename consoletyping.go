/*

콘솔 입력함수
console typing


Scanln  - 공백으로 구분하여 입력
데이터를 입력 받을 때 공백으로만 구분 할 수 있다. 개행하면 입력이 완료된 것으로 처리.

Scan -  공백과 개행으로 구분하여 입력
개행하는 것도 데이터를 구분하는 입력방식으로 인식
ex > Scan(&num1 &num2 &num3)은 숫자를 입력할 때마다 개행해서 입력받을 수 있다.

Scanf - 포멧 지정자를 이용하여 개발자가 원하는 형태로 입력
ex > Scanf("%d-%d", &num1, &num2)
12345-67890
num1 - 12345 입력 / num2 - 67890 이 입력된다

*/
package main

import "fmt"

func main() {
	var name, gen, school string

	fmt.Println("이름 성별 학교를 입력해주세요")
	fmt.Scanln(&name, &gen, &school)

	fmt.Println("이름은", name, "성별은", gen, "학교는", school, "입니다.")

}

/*
C:\Users\happi\go\src>go run consoletyping.go
이름 성별 학교를 입력해주세요
민이 여자 서울
이름은 민이 성별은 여자 학교는 서울 입니다.
*/

/////////////////////////////////////////////////////////////////
// func main() {
// 	var i, j int

// 	fmt.Print("주민등록번호를 -를 이용해 입력하세요 :")
// 	fmt.Scanf("%d-%d", &i, &j)

// 	fmt.Printf("주민등록번호는 %d-%d\n", i, j)
// 	fmt.Println("앞자리는", i)
// 	fmt.Println("뒷자리는", j)
// }

/*
cmd
C:\Users\happi\go\src>go run consoletyping.go
주민등록번호를 -를 이용해 입력하세요 :900211-2011111
주민등록번호는 900211-2011111
앞자리는 900211
뒷자리는 2011111
*/
