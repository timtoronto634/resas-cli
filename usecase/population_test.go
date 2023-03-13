package usecase

import (
	"bytes"
	"context"
	"testing"
)

type popArg struct {
	Label    string
	PrefCode []string
	YearFrom int
	YearTo   int
	Expected string
}

func TestPrintPopulation(t *testing.T) {
	args := []popArg{
		{
			Label:    "総人口",
			PrefCode: []string{"13"},
			YearFrom: 1980,
			YearTo:   2020,
			Expected: `東京都,1980,11618281
東京都,1985,11829363
東京都,1990,11855563
東京都,1995,11773605
東京都,2000,12064101
東京都,2005,12576601
東京都,2010,13159388
東京都,2015,13515271
東京都,2020,14047594
`,
		},
	}

	ctx := context.Background()
	for _, arg := range args {
		buf := bytes.NewBufferString("")
		PrintPopulation(ctx, buf, arg.Label, arg.PrefCode, arg.YearFrom, arg.YearTo)
		output := buf.String()

		if output != arg.Expected {
			t.Errorf("Failed test. Expetced %q but got %q", arg.Expected, output)
		}
	}
}
