// SOLID(객체지향설계) : 객체지향 프로그래밍 및 설계의 다섯가지 기본원칙 5원칙
// Korea-Technology Exchange

package main

import (
	"fmt"
)

// 인터페이스 1. Seller, Dealer, Buyer를 UnvalidatedTeh으로 단일책임
type UnvalidatedTech interface {
	String() string
}

type AIValidator interface {
	Validating() UnvalidatedTech
}

type ValidateSystem struct {
	val string
}

func (s *ValidateSystem) Item(AIValidator AIValidator) {
	pickit := AIValidator.Validating()
	s.val += pickit.String()
}

func (s *ValidateSystem) String() string {
	return "\t [[ A.I 검증중인 기술 목록 ]]\n " + s.val
}

// 1번 기술
type Tech_01 struct {
}

func (j *Tech_01) Validating() UnvalidatedTech {
	return &DevelopedTech_01{}
}

type DevelopedTech_01 struct {
}

func (s *DevelopedTech_01) String() string {
	return "  1. 영지식 증명을 활용한 블라인드 전자계약서\n  "
}

// 2번 외부 공개가 꺼려지는 핵심기술은 내부 기능으로 돌려서 보안 유지(개방폐쇄원칙)
type tech_02 struct {
}

func (j *tech_02) Validating() UnvalidatedTech {
	return &Developedtech_02{}
}

type Developedtech_02 struct {
}

func (s *Developedtech_02) String() string {
	return " 2. Blinded Technology #1\n   "
}

// 3번 기술
type Tech_03 struct {
}

func (j *Tech_03) Validating() UnvalidatedTech {
	return &DevelopedTech_03{}
}

type DevelopedTech_03 struct {
}

func (s *DevelopedTech_03) String() string {
	return "3. PoDC 합알고리즘을 활용한 블록체인 기반 전자투표 시스템 <HOT!>\n   "
}

// 4번 기술
type Tech_04 struct {
}

func (j *Tech_04) Validating() UnvalidatedTech {
	return &DevelopedTech_04{}
}

type DevelopedTech_04 struct {
}

func (s *DevelopedTech_04) String() string {
	return "4. 영지식 증명 기반 투개표 시스템"
}

/////////////////////////////////////////////////

type ValidatedTech interface {
	String() string
}

type KTexchange interface {
	Trading() ValidatedTech
}

type TradeSystem struct {
	val string
}

func (s *TradeSystem) TradableItem(KTexchange KTexchange) {
	tradeit := KTexchange.Trading()
	s.val += tradeit.String()
}

func (s *TradeSystem) String() string {
	return "\t [[ A.I 검증 완료된 거래가능 기술 목록 ]]\n " + s.val
}

// 1번
type TradableTech_01 struct {
}

func (j *TradableTech_01) Trading() ValidatedTech {
	return &V_DevelopedTech_01{}
}

type V_DevelopedTech_01 struct {
}

func (s *V_DevelopedTech_01) String() string {
	return "  1. 블록체인 기반 코로나19 백신 접종 증명시스템\n  "
}

// 2번
type TradableTech_02 struct {
}

func (j *TradableTech_02) Trading() ValidatedTech {
	return &V_DevelopedTech_02{}
}

type V_DevelopedTech_02 struct {
}

func (s *V_DevelopedTech_02) String() string {
	return " 2. DID 비대칭키기술을 활용한 신원증명 시스템\n  "
}

// 3번
type TradableTech_03 struct {
}

func (j *TradableTech_03) Trading() ValidatedTech {
	return &V_DevelopedTech_03{}
}

type V_DevelopedTech_03 struct {
}

func (s *V_DevelopedTech_03) String() string {
	return " 3. 물리적 복제 방지 기능을 가진 난수 발생 장치\n  "
}

// 4번
type TradableTech_04 struct {
}

func (j *TradableTech_04) Trading() ValidatedTech {
	return &V_DevelopedTech_04{}
}

type V_DevelopedTech_04 struct {
}

func (s *V_DevelopedTech_04) String() string {
	return " 4. 딥러닝을 이용한 선박 충돌 방지 자율회피 시스템\n  "
}

func main() {
	ValidateSystem := &ValidateSystem{}
	TradeSystem := &TradeSystem{}

	AIValidator01 := &Tech_01{}
	AIValidator02 := &tech_02{}
	AIValidator03 := &Tech_03{}
	AIValidator04 := &Tech_04{}

	TradeTechnology01 := &TradableTech_01{}
	TradeTechnology02 := &TradableTech_02{}
	TradeTechnology03 := &TradableTech_03{}
	TradeTechnology04 := &TradableTech_04{}

	ValidateSystem.Item(AIValidator01)
	ValidateSystem.Item(AIValidator02)
	ValidateSystem.Item(AIValidator03)
	ValidateSystem.Item(AIValidator04)

	TradeSystem.TradableItem(TradeTechnology01)
	TradeSystem.TradableItem(TradeTechnology02)
	TradeSystem.TradableItem(TradeTechnology03)
	TradeSystem.TradableItem(TradeTechnology04)

	fmt.Println("--------------------Korea Tech Exchange----------------------")
	fmt.Println("")
	fmt.Println(ValidateSystem)
	fmt.Println("")
	fmt.Println("--------------------Korea Tech Exchange----------------------")
	fmt.Println("")
	fmt.Println(TradeSystem)
	fmt.Println("--------------------Korea Tech Exchange----------------------")

}
