package utils_test

import (
	"testing"
	"this_module/internal/pkg/utils"
)

func TestAtoi(t *testing.T) {
	u := utils.New()

	cases := []struct {
		Input    string
		Expected int
	}{
		{
			Input:    "10",
			Expected: 10,
		},
		{
			Input:    "101",
			Expected: 101,
		},
		{
			Input:    "A1",
			Expected: 0,
		},
	}

	for _, v := range cases {
		res := u.Atoi(v.Input)

		if res != v.Expected {
			t.Errorf("expected %d, res: %v", v.Expected, res)
		}

	}
}
