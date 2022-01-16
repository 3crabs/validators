package snils

import "testing"

type snilsArgs struct {
	number string
}

type snilsTest struct {
	name string
	args snilsArgs
	want bool
}

func Test_validateCorrectNumbers(t *testing.T) {
	var tests []snilsTest
	correctNumbers := []string{
		"112-233-445 95",
		"765 925 109 40",
		"385 991 627 56",
		"84752821738",
		"35486451509",
		"65939415653",
		"935 832 714 36",
		"811 081 334 49",
		"323 464 034 43",
		"237-419-636 85",
		"445-711-602 73",
		"529-909-600 30",
		"582 986 318 58",
		"176 613 795 02",
		"310 665 841 53",
		"310 665 841 53",
	}
	for _, number := range correctNumbers {
		tests = append(tests,
			snilsTest{
				name: number,
				args: snilsArgs{number: number},
				want: true,
			},
		)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.number); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateIncorrectNumbers(t *testing.T) {
	var tests []snilsTest
	correctNumbers := []string{
		"+112 233 445 99",
		"112-233-445 99",
		"765 925 109 41",
		"385 992 627 56",
		"84752821730",
		"35482451509",
		"65936415653",
		"935 832 754 37",
		"811 081 333 40",
		"323 464 024 00",
		"237-419-656 81",
		"445511-602 33",
		"569-909-600 20",
		"111-909-600 31",
		"511-109-600 19",
	}
	for _, number := range correctNumbers {
		tests = append(tests,
			snilsTest{
				name: number,
				args: snilsArgs{number: number},
				want: false,
			},
		)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.number); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
