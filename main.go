package main

import "fmt"

func main() {

	// 주석문법

	// 한줄 주석
	/* 여러줄 주석 */

	// 변수문법
	// 이름을 쓰고 타입을 씀.
	// 이와 같은 순서로 변수를 선언하는 이유는 코드를 자연스럽게 읽기 위해서다.

	//• 정수 타입의 제로값: 0
	//• 실수(부동소수점) 타입의 제로값: 0.0
	//• 문자열 타입의 제로값: ""

	//1. 초기화 하지않고 할당하는 방법
	//var 변수명 변수타입
	var a int
	// 초기화 하고 할당하는 방법
	//
	// 명시적인 var 변수명 변수타입 = 초기화값
	var f float32 = 11.
	// 암묵적인 변수명 = 초기화값
	a = 10
	f = 12.0
	// 복수의 변수를 할당하는 방법
	//var i, j, k int
	var i, j, k int = 1, 2, 3

	//타입이 다른 변수 여러 개를 한꺼번에 선언할 때는 소괄호(())로 묶어서 표기한다.
	var (
		name   string
		age    int
		weight float
	)

	//짧은 선언
	//변수 선언과 동시에 값을 할당할 때는 var를 생략할 수 있다. var를 생략할 때는 := 연산자를 사용한다.
	start := 1
	//이미 선언된 변수에 값을 할당할 때는 := 연산자를 사용할 수 없다.
	//:= 연산자는 변수 선언과 동시에 초깃값을 할당할 때만 사용한다.
	//짧은 선언은 함수 안에서만 가능하다.
	//전역 변수는 짧은 선언으로 선언할 수 없다.

	// 상수문법

	// 변수선언 방법에 앞에 const 키워드로 선언함

	const c int = 10
	const s string = "Hi"

	// 여러개 상수들을 묶어서 지정가능

	// 한가지 유용한 팁으로 상수값을 0부터 순차적으로 부여하기 위해 iota 라는 identifier를 사용할 수 있다.
	// 이 경우 iota가 지정된 Apple에는 0이 할당되고, 나머지 상수들을 순서대로 1씩 증가된 값을 부여받는다.
	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)

	// Go언어의 예약된 키워드들
	//break        default      func         interface    select
	//case         defer        go           map          struct
	//chan         else         goto         package      switch
	//const        fallthrough  if           range        type
	//continue     for          import       return       var

	// Go언어의 데이터 타입
	// 01. 부울린
	// 	bool
	// 02. 문자열
	//	string
	// 	string 타입은 한번 생성이되면 수정할수 없는 Immutable타입 임
	// 03. 정수형
	// int, int8, intt16, int32, int64
	// uint, unit8, uint16, uint32, uint64
	// 04. 실수형
	// float32, float64, complex64, complex128
	// 05. 기타타입형
	// byte : uint8과 동일하며, byte코드에 사용
	// rune : int32와 동일하며, 유니코드 코드포인터에 사용

	// 문자열

	// 문자열은 1) Back Quote(``) 혹은 2) Double Quote("")를 사용하여 표현할 수 있고, 두개가 서로 다르다.
	// 1) Back Quote(｀｀)
	// ''로 싸인 문자열은 Raw String Literal이라고하며, 문자열은 별도로 해석되지 않고, Raw String그대로의 값을 갖는다.
	// 예를 들어 '\n'이 개행이 아니라 문자그대로 표현됨. 그리고 개행시, 개행으로 해석되어 여러라인 문자열을 표현할때 자주 사용된다.
	// 2) Double Quote("")

	rawLiteral := `아리랑
아리랑
아라리요'
`
	interLiteral_01 := "아리랑아리랑\n아리리요"
	interLiteral_02 := "아리랑아리랑\n" + "아리리요"

	// 데이타 타입 변환 (Type Conversion)
	// 변환될타입명(변환될값)
	// Go에서 타입간 변환은 명시적으로 지정해 주어야 한다
	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)

	// 배열
	var arry_01 [3]int //정수형 3개 요소를 갖는 배열 a 선언

	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3} //배열크기 자동으로

	// 다차원 배열 선언
	var a = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, //끝에 콤마 추가
	}

	// gofmt 도구를 사용하면 코드의 스타일을 Go에서 사용하는 스타일대로 맞춰준다.
	// $ gofmt -h
	// $ gofmt prog.go
	//-w 옵션을 사용하면 스타일이 맞춰진 코드를 원본 소스 파일에 다시 저장한다.
	// $ gofmt -w prog.go

	fmt.Println("Fuck you man")

}
