package usecase

import (
	"bytes"
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/timtoronto634/resas-cli/entity"
)

type popArg struct {
	Label         string
	PrefCode      []int
	YearFrom      int
	YearTo        int
	Expected      string
	ExpectedItems []string
}

func TestPrintPopulation(t *testing.T) {
	args := []popArg{
		{
			Label:    "総人口",
			PrefCode: []int{13},
			YearFrom: 1980,
			YearTo:   2020,
			Expected: `東京都,2020,14047594
東京都,2015,13515271
東京都,2010,13159388
東京都,2005,12576601
東京都,2000,12064101
東京都,1995,11773605
東京都,1990,11855563
東京都,1985,11829363
東京都,1980,11618281
`,
		},
		{
			Label:    "老年人口",
			PrefCode: []int{12, 14},
			YearFrom: 2000,
			YearTo:   2010,
			ExpectedItems: []string{`千葉県,2010,1320120
千葉県,2005,1060343
千葉県,2000,837017
`,
				`神奈川県,2010,1819503
神奈川県,2005,1480262
神奈川県,2000,1169528
`,
			},
		},
	}

	ctx := context.Background()
	for _, arg := range args {
		buf := bytes.NewBufferString("")
		PrintPopulations(ctx, buf, arg.Label, arg.PrefCode, arg.YearFrom, arg.YearTo)
		output := buf.String()

		if len(arg.PrefCode) == 1 {
			if output != arg.Expected {
				t.Errorf("Failed test. Expetced %q but got %q", arg.Expected, output)
			}
		} else {
			for _, expected := range arg.ExpectedItems {
				if !strings.Contains(output, expected) {
					t.Errorf("output do not include expected string, expect: %v, got: %v", expected, output)
				}
			}
		}
	}
}
func TestFilterSortWithYear(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name      string
		inputData []*entity.Population
		from      int
		to        int
		want      []*entity.Population
	}{
		{
			name: "filtering functionality",
			inputData: []*entity.Population{
				{Year: 2025, Value: 5},
				{Year: 2020, Value: 2},
				{Year: 2000, Value: 2},
				{Year: 1990, Value: 3},
				{Year: 1985, Value: 10},
			},
			from: 1990,
			to:   2020,
			want: []*entity.Population{
				{Year: 2020, Value: 2},
				{Year: 2000, Value: 2},
				{Year: 1990, Value: 3},
			},
		},
		{
			name: "Out of Range inputs",
			inputData: []*entity.Population{
				{Year: 2020, Value: 2},
				{Year: 2000, Value: 2},
				{Year: 2025, Value: 5},
				{Year: 1985, Value: 10},
				{Year: 1990, Value: 3},
			},
			from: 1900,
			to:   1980,
			want: []*entity.Population{},
		},
		{
			name:      "Empty Input Slice",
			inputData: []*entity.Population{},
			from:      1990,
			to:        2019,
			want:      []*entity.Population{},
		},
		{
			name: "Sorting Functionality",
			inputData: []*entity.Population{
				{Year: 2020, Value: 2},
				{Year: 2000, Value: 2},
				{Year: 2025, Value: 5},
				{Year: 1985, Value: 10},
				{Year: 1990, Value: 3},
			},
			from: 1980,
			to:   2030,
			want: []*entity.Population{
				{Year: 2025, Value: 5},
				{Year: 2020, Value: 2},
				{Year: 2000, Value: 2},
				{Year: 1990, Value: 3},
				{Year: 1985, Value: 10},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := filterSortWithYear(tc.inputData, tc.from, tc.to)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("filterSortWithYear() = %v, want %v", got, tc.want)
			}
		})
	}
}
