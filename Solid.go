// SOLID(객체지향설계) : 객체지향 프로그래밍 및 설계의 다섯가지 기본원칙 5원칙
// K-Technology Exchange

package main

import (
	"fmt"
)

// 인터페이스 1. Seller, Dealer, Buyer를 Person으로 단일책임
type Person interface {
	String() string
}

// 인터페이스 2
type Technology interface {
	Validating() Person
}

type Platform struct {
	val string
}

func (s *Platform) Item(Technology Technology) {
	pickit := Technology.Validating()
	s.val += pickit.String()
}

func (s *Platform) String() string {
	return "핵심 기술" + s.val
}

//////////////////인터페이스 1,2를 활용///////////////////////

// 1번
type FirstTech struct {
}

func (j *FirstTech) Validating() Person {
	return &DevelopedFirstTech{}
}

type DevelopedFirstTech struct {
}

func (s *DevelopedFirstTech) String() string {
	return " + 1번"
}

// 2번이 핵심기술이라면 소문자로해서 내부 기능으로 돌리기.(개방폐쇄원칙)
type secondTech struct {
}

func (j *secondTech) Validating() Person {
	return &DevelopedsecondTech{}
}

type DevelopedsecondTech struct {
}

func (s *DevelopedsecondTech) String() string {
	return " + 2번"
}

// 3번
type ThirdTech struct {
}

func (j *ThirdTech) Validating() Person {
	return &DevelopedThirdTech{}
}

type DevelopedThirdTech struct {
}

func (s *DevelopedThirdTech) String() string {
	return " + 3번"
}

// 4번
type FourthTech struct {
}

func (j *FourthTech) Validating() Person {
	return &DevelopedFourthTech{}
}

type DevelopedFourthTech struct {
}

func (s *DevelopedFourthTech) String() string {
	return " + 4번"
}

func main() {
	Platform := &Platform{}

	Technology01 := &FirstTech{}
	Technology02 := &ThirdTech{}
	Technology03 := &secondTech{}
	Technology04 := &FourthTech{}

	Platform.Item(Technology01)
	Platform.Item(Technology02)
	Platform.Item(Technology03)
	Platform.Item(Technology04)

	fmt.Println("----------------Korea Tech eXchange------------------")
	fmt.Println("")
	fmt.Println(Platform)
	fmt.Println("")
}
