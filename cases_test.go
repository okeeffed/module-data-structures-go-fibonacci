package fib

type testCase struct {
	description string
	input       uint
	expected    uint
}

var testCases = []testCase{
	{
		description: "should return correct Fib value",
		input:       1,
		expected:    1,
	},
	{
		description: "should return correct Fib value",
		input:       2,
		expected:    1,
	},
	{
		description: "should return correct Fib value",
		input:       6,
		expected:    8,
	},
	{
		description: "should return correct Fib value",
		input:       15,
		expected:    610,
	},
	{
		description: "should return correct Fib value",
		input:       30,
		expected:    832040,
	},
	{
		description: "should return correct Fib value",
		input:       45,
		expected:    1134903170,
	},
}
