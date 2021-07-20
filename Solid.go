// SOLID(객체지향설계) : 객체지향 프로그래밍 및 설계의 다섯가지 기본원칙 5원칙
// Korea-Technology Exchange Blockchain A.I Agent

package main

import (
	"fmt"
)

// 인터페이스 1. 아직 A.I 검증되지 않은 기술에 대한 검증을 준비합니다.
type UnvalidatedTech interface {
	String() string
}

// AIValidator가 UnvalidatedTechnology를 Validating할 수 있도록 인터페이스를 구성해줍니다.
type AIValidator interface {
	Validating() UnvalidatedTech
}

// 일련의 과정을 책임지는 시스템을 ValidateSystem으로 나타냅니다.
type ValidateSystem struct {
	val string
}

// Input_data를 통해 ValidateSystem을 사용합니다.
func (s *ValidateSystem) Input_data(AIValidator AIValidator) {
	validate := AIValidator.Validating()
	s.val += validate.String()
}

// ValidateSystem은 현재 A.I검증중인 기술 목록을 아래와 같이 출력합니다
func (s *ValidateSystem) String() string {
	return "\t [[ A.I 검증중인 기술 목록 ]]\n " + s.val
}

// A.I 검증중인 1번 기술
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

// A.I 검증중인 2번 외부 공개가 꺼려지는 핵심기술은 내부 기능으로 돌려서 보안을 유지합니다. (개방폐쇄원칙)
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

// A.I 검증중인 3번 기술
type Tech_03 struct {
}

func (j *Tech_03) Validating() UnvalidatedTech {
	return &DevelopedTech_03{}
}

type DevelopedTech_03 struct {
}

func (s *DevelopedTech_03) String() string {
	return "3. PoDC 합알고리즘을 활용한 블록체인 기반 전자투표 시스템 <<HOT!>>\n   "
}

// A.I 검증중인 4번 기술
type Tech_04 struct {
}

func (j *Tech_04) Validating() UnvalidatedTech {
	return &DevelopedTech_04{}
}

type DevelopedTech_04 struct {
}

func (s *DevelopedTech_04) String() string {
	return "4. 영지식 증명 기반 투개표 시스템\n   "
}

// A.I 검증중인 5번 기술

type Tech_05 struct {
}

func (j *Tech_05) Validating() UnvalidatedTech {
	return &DevelopedTech_05{}
}

type DevelopedTech_05 struct {
}

func (s *DevelopedTech_05) String() string {
	return "5. 딥러닝 기반 방사선 피폭 진단용 염색체 판독 시스템"
}

/////////////////////////////////////////////////////////////////////////////////////////////

// 인터페이스 3. A.I 검증이 완료된 기술을 나타냅니다.
type ValidatedTech interface {
	String() string
}

// 인터페이스 4. 검증이 완료된 기술을 K_TBA를 통해 거래가능한 목록에 올라갑니다.
type K_TBA interface {
	Trading() ValidatedTech
}

// TradeSystem을 준비합니다
type TradeSystem struct {
	val string
}

func (s *TradeSystem) TradableInput_data(K_TBA K_TBA) {
	tradeit := K_TBA.Trading()
	s.val += tradeit.String()
}

func (s *TradeSystem) String() string {
	return "\t [[ A.I 검증 완료된 거래가능 기술 목록 ]]\n " + s.val
}

// A.I 검증이 완료된 1번 기술
type TradableTech_01 struct {
}

func (j *TradableTech_01) Trading() ValidatedTech {
	return &V_DevelopedTech_01{}
}

type V_DevelopedTech_01 struct {
}

func (s *V_DevelopedTech_01) String() string {
	return "  1. AI DAO <<HOT!>>\n  "
}

// A.I 검증이 완료된 2번 기술
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

// A.I 검증이 완료된 3번 기술
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

// A.I 검증이 완료된 4번 기술
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

// A.I 검증이 완료된 5번 기술
type TradableTech_05 struct {
}

func (j *TradableTech_05) Trading() ValidatedTech {
	return &V_DevelopedTech_05{}
}

type V_DevelopedTech_05 struct {
}

func (s *V_DevelopedTech_05) String() string {
	return " 5. 퍼블릭 블록체인 환경에서 개인정보보호를 위한 거래방법\n  "
}

// A.I 검증이 완료된 6번 기술
type TradableTech_06 struct {
}

func (j *TradableTech_06) Trading() ValidatedTech {
	return &V_DevelopedTech_06{}
}

type V_DevelopedTech_06 struct {
}

func (s *V_DevelopedTech_06) String() string {
	return " 6. 온라인 학습 환경에서 학습자의 시각 행동을 활용한 학습 유형 진단 장치 \n  "
}

// A.I 검증이 완료된 7번 기술
type TradableTech_07 struct {
}

func (j *TradableTech_07) Trading() ValidatedTech {
	return &V_DevelopedTech_07{}
}

type V_DevelopedTech_07 struct {
}

func (s *V_DevelopedTech_07) String() string {
	return " 7. 해양 부유물 수거용 무인로봇 <<기술나눔>>\n  "
}

func main() {
	// 검증시스템과 거래시스템을 정의합니다.
	ValidateSystem := &ValidateSystem{}
	TradeSystem := &TradeSystem{}

	// A.I Validator가 각각의 기술들을 검증합니다.
	AIValidator01 := &Tech_01{}
	AIValidator02 := &tech_02{}
	AIValidator03 := &Tech_03{}
	AIValidator04 := &Tech_04{}
	AIValidator05 := &Tech_05{}

	ValidateSystem.Input_data(AIValidator01)
	ValidateSystem.Input_data(AIValidator02)
	ValidateSystem.Input_data(AIValidator03)
	ValidateSystem.Input_data(AIValidator04)
	ValidateSystem.Input_data(AIValidator05)

	// A.I 검증완료된 기술들을 거래가능 목록으로 나타내도록 합니다.
	TradeTechnology01 := &TradableTech_01{}
	TradeTechnology02 := &TradableTech_02{}
	TradeTechnology03 := &TradableTech_03{}
	TradeTechnology04 := &TradableTech_04{}
	TradeTechnology05 := &TradableTech_05{}
	TradeTechnology06 := &TradableTech_06{}
	TradeTechnology07 := &TradableTech_07{}

	TradeSystem.TradableInput_data(TradeTechnology01)
	TradeSystem.TradableInput_data(TradeTechnology02)
	TradeSystem.TradableInput_data(TradeTechnology03)
	TradeSystem.TradableInput_data(TradeTechnology04)
	TradeSystem.TradableInput_data(TradeTechnology05)
	TradeSystem.TradableInput_data(TradeTechnology06)
	TradeSystem.TradableInput_data(TradeTechnology07)

	// 최종 출력본을 정리합니다.
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	fmt.Println("::::::::##:::'##::::::::::'########:'########:::::'###::::::::::")
	fmt.Println("::::::::##::'##:::::::::::... ##..:: ##.... ##:::'## ##:::::::::")
	fmt.Println("::::::::##:'##::::::::::::::: ##:::: ##:::: ##::'##:. ##::::::::")
	fmt.Println("::::::::#####::::'#######:::: ##:::: ########::'##:::. ##:::::::")
	fmt.Println("::::::::##. ##:::........:::: ##:::: ##.... ##: #########:::::::")
	fmt.Println("::::::::##:. ##:::::::::::::: ##:::: ##:::: ##: ##.... ##:::::::")
	fmt.Println("::::::::##::. ##::::::::::::: ##:::: ########:: ##:::: ##:::::::")
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	fmt.Println("")
	fmt.Println("┍━━━━━━━━━━━━━━━━━━━━━━━━━━━━»•» K-TBA Platform «•«━━━━━━━━━━━━━━━━━┑")
	fmt.Println("")
	fmt.Println(ValidateSystem)
	fmt.Println("")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("")
	fmt.Println(TradeSystem)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("━━━━━━━━━━━━━기술을 선택하시면 추가 상세정보를 확인하실 수 있습니다━━━━━━━━━━━━━━")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("")
	fmt.Println("┕━━━━━━━━━━━━»•» K-TBA Platform «•«━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┙")
}
