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
type Tech_01 struct {
}

func (j *Tech_01) Validating() Person {
	return &DevelopedTech_01{}
}

type DevelopedTech_01 struct {
}

func (s *DevelopedTech_01) String() string {
	return " + 1번"
}

// 2번이 핵심기술이라면 소문자로해서 내부 기능으로 돌리기.(개방폐쇄원칙)
type tech_02 struct {
}

func (j *tech_02) Validating() Person {
	return &Developedtech_02{}
}

type Developedtech_02 struct {
}

func (s *Developedtech_02) String() string {
	return " + 2번"
}

// 3번
type Tech_03 struct {
}

func (j *Tech_03) Validating() Person {
	return &DevelopedTech_03{}
}

type DevelopedTech_03 struct {
}

func (s *DevelopedTech_03) String() string {
	return " + 3번"
}

// 4번
type Tech_04 struct {
}

func (j *Tech_04) Validating() Person {
	return &DevelopedTech_04{}
}

type DevelopedTech_04 struct {
}

func (s *DevelopedTech_04) String() string {
	return " + 4번"
}

func main() {
	Platform := &Platform{}

	Technology01 := &Tech_01{}
	Technology02 := &Tech_03{}
	Technology03 := &tech_02{}
	Technology04 := &Tech_04{}

	Platform.Item(Technology01)
	Platform.Item(Technology02)
	Platform.Item(Technology03)
	Platform.Item(Technology04)

	fmt.Println("----------------Korea Tech eXchange------------------")
	fmt.Println("")
	fmt.Println(Platform)
	fmt.Println("")
}
