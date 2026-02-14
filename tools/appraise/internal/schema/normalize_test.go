package schema

import "testing"

func TestNormalizeQuery(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "no dots, no change",
			input:  `list(modules)`,
			expect: `list(modules)`,
		},
		{
			name:   "dotted module.function gets quoted",
			input:  `calc(pricing.bvr, input="data.json")`,
			expect: `calc("pricing.bvr", input="data.json")`,
		},
		{
			name:   "already quoted, no change",
			input:  `calc("pricing.bvr", input="data.json")`,
			expect: `calc("pricing.bvr", input="data.json")`,
		},
		{
			name:   "dotted value after equals",
			input:  `calc(pricing.bvr, path=some/file.json)`,
			expect: `calc("pricing.bvr", path="some/file.json")`,
		},
		{
			name:   "batch with dots",
			input:  `calc(pricing.bvr, input="a.json"); calc(financial.clv, input="b.json")`,
			expect: `calc("pricing.bvr", input="a.json"); calc("financial.clv", input="b.json")`,
		},
		{
			name:   "field projection untouched",
			input:  `calc(pricing.bvr, input="a.json") { bvr interpretation }`,
			expect: `calc("pricing.bvr", input="a.json") { bvr interpretation }`,
		},
		{
			name:   "schema with no args",
			input:  `schema()`,
			expect: `schema()`,
		},
		{
			name:   "summary with no args",
			input:  `summary()`,
			expect: `summary()`,
		},
		{
			name:   "list functions with module filter",
			input:  `list(functions, module=pricing)`,
			expect: `list(functions, module=pricing)`,
		},
		{
			name:   "deep dotted path",
			input:  `calc(a.b.c, input="d.json")`,
			expect: `calc("a.b.c", input="d.json")`,
		},
		{
			name:   "slashes in path value",
			input:  `calc(pricing.bvr, input=path/to/file)`,
			expect: `calc("pricing.bvr", input="path/to/file")`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NormalizeQuery(tc.input)
			if got != tc.expect {
				t.Errorf("NormalizeQuery(%q)\n  got:    %q\n  expect: %q", tc.input, got, tc.expect)
			}
		})
	}
}
