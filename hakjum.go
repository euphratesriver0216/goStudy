package main

import "fmt"

func main() {
	var math, english, program, club, avg, sum int
	var result string

	fmt.Println("평균학점계산기 입니다.")
	fmt.Println("수학, 영어, 프로그램, 클럽활동 순으로 학점을 입력해주세요 (100점 만점)")
	fmt.Scan(&math, &english, &program, &club)

	sum = math + english + program + club
	avg = sum / 4

	if avg >= 101 {
		result = "정확한 점수를 입력해주세요"
	} else if avg >= 90 {
		result = "A"
	} else if avg >= 80 {
		result = "B"
	} else if avg >= 70 {
		result = "C"
	} else {
		result = "D 70점 이하의 과목은 재수강하세요"
	}
	fmt.Println("학점", result, "평균", avg)
}
