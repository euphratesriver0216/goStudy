//상수 설명
package main

import (
	"fmt"
)

const (
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	SmallRoom
)

//불을 켜는 함수
func SetLight(rooms, room uint8) uint8 {
	return rooms | room //OR
}

//불을 끄는 함수
func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room //XOR
}

//어떤 룸이 불이 켜져있는지 확인하는 함수
func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room
}

//
func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("안방에 불을 켠다")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("거실에 불을 켠다")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("화장실에 불을 켠다")
	}
	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("작은방에 불을 켠다")
	}
}

func main() {
	var rooms uint8 = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, SmallRoom)

	rooms = ResetLight(rooms, SmallRoom)

	TurnLights(rooms)

}
