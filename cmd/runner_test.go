package cmd

import "testing"

func TestRun(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string
		output string
	}{}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {})
	}
}

func BenchmarkRun_Simple(b *testing.B) {
	input := "./examples/basic.txt"
	output := "./examples/basic_output.txt"
	configs := []string{"./examples/basic.hcl"}
	extras := []string{}
	for i := 0; i < b.N; i++ {
		err := process(input, output, configs, extras)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkRun_Complex(b *testing.B) {
	input := "./examples/input.txt"
	output := "./examples/output.txt"
	configs := []string{"./examples/vars.hcl", "./examples/vars2.hcl"}
	extras := []string{}
	for i := 0; i < b.N; i++ {
		err := process(input, output, configs, extras)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkRun_Extras(b *testing.B) {
	input := "./examples/input.txt"
	output := "./examples/output.txt"
	configs := []string{"./examples/vars.hcl", "./examples/vars2.hcl"}
	extras := []string{`{"i": "value of i"}`}
	for i := 0; i < b.N; i++ {
		err := process(input, output, configs, extras)
		if err != nil {
			b.Error(err)
		}
	}
}
