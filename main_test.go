package main

import "testing"

type CaseIsPolindrome struct {
	input          string
	expectedResult bool
}

func TestIsPalindrome(t *testing.T) {
	testCases := []CaseIsPolindrome{
		{
			input:          "ovo",
			expectedResult: true,
		},
		{
			input:          "reviver",
			expectedResult: true,
		},
		{
			input:          "salas",
			expectedResult: true,
		},
		{
			input:          "correr",
			expectedResult: false,
		},
		{
			input:          "jogador",
			expectedResult: false,
		},
		{
			input:          "",
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		if isPalindrome(testCase.input) != testCase.expectedResult {
			t.Error("FAIL", testCase)
		}
	}
}

func TestIsPalindrome2(t *testing.T) {
	testCases := []CaseIsPolindrome{
		{
			input:          "ovo",
			expectedResult: true,
		},
		{
			input:          "reviver",
			expectedResult: true,
		},
		{
			input:          "salas",
			expectedResult: true,
		},
		{
			input:          "correr",
			expectedResult: false,
		},
		{
			input:          "jogador",
			expectedResult: false,
		},
		{
			input:          "",
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		if isPalindrome2(testCase.input) != testCase.expectedResult {
			t.Error("FAIL", testCase)
		}
	}
}

type CaseInvertText struct {
	input          string
	expectedResult string
}

func TestInvertText(t *testing.T) {
	testCases := []CaseInvertText{
		{
			input:          "ovo",
			expectedResult: "ovo",
		},
		{
			input:          "reviver",
			expectedResult: "reviver",
		},
		{
			input:          "salas",
			expectedResult: "salas",
		},
		{
			input:          "correr",
			expectedResult: "rerroc",
		},
		{
			input:          "jogador",
			expectedResult: "rodagoj",
		},
	}

	for _, testeCase := range testCases {
		if invertText(testeCase.input) != testeCase.expectedResult {
			t.Error("FAIL", testeCase)
		}
	}
}

func BenchmarkInvertText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		invertText("reviver")
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("reviver")
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome2("reviver")
	}
}
