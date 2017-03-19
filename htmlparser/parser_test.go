package htmlparser

import (
	"os"
	"reflect"
	"testing"
)

type testcase struct {
	expectedContainer []container
}

type testsetup struct {
	filename  string
	testcases []testcase
}

var tests = []testsetup{
	{
		filename: "fixtures/simple.html",
		testcases: []testcase{
			{
				expectedContainer: []container{
					{
						label: "System",
						/*childs: []container{
							{
								label: "Dev",
								childs: []container{
									{
										label: "Stage A",
										links: []link{
											{label: "DE bar", url: "fooA"},
											{label: "DE baz", url: "googleA"},
											{label: "DE bam", url: "facebookA"},
											{label: "AT a text", url: "alinkA"},
											{label: "CH another text", url: "anotherlinkA"},
										},
									},
									{
										label: "Stage B",
										links: []link{
											{label: "DE bar", url: "fooB"},
											{label: "DE baz", url: "googleB"},
											{label: "DE bam", url: "facebookB"},
											{label: "AT a text", url: "alinkB"},
											{label: "CH another text", url: "anotherlinkB"},
										},
									},
									{
										label: "Stuff for all",
										links: []link{
											{label: "a super text", url: "asuperlink"},
										},
									},
								},
							},
						},*/
					},
				},
			},
		},
	},
}

func TestParseHtml(t *testing.T) {
	for _, testcase := range tests {
		input, err := os.OpenFile(testcase.filename, os.O_RDONLY, 0666)
		defer input.Close()
		if err != nil {
			t.Fatal("error opening test fixture!")
		}
		p, err := New(input)
		if err != nil {
			t.Fatal("error creating parser!", err)
		}

		for i, tt := range testcase.testcases {
			cont := p.ParseHtml()
			t.Fatal("error creating parser!", cont)
			//if len(cont) != len(tt.expectedContainer) {
			if reflect.DeepEqual(cont, tt.expectedContainer) {
				t.Fatalf("tests[%d] - container wrong wrong. expected=%q, got=%q",
					i, tt.expectedContainer, cont)
			}

		}
	}
}
