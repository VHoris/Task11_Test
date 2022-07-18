package do

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DoFunction(t *testing.T) {
	tests := map[string]struct {
		s      string
		i      int
		b      bool
		exp    string
		expErr string
	}{
		"success_is_case1":   {s: "a", i: 8, b: false, exp: "[8a]", expErr: ""},
		"success_toUpper_b":   {s: "b", i: 5, b: true, exp: "[5B]", expErr: ""},
		"success_len<2":   {s: "d", i: 21, b: false, exp: "d", expErr: ""},
		"invalid s(i)": {s: "b", i: 35, b: true, exp: "", expErr: "invalid s"},
		"invalid s(s)": {s: "k", i: 1, b: true, exp: "", expErr: "invalid s"},
		"invalid s(s and i)": {s: "k", i: 44, b: true, exp: "", expErr: "invalid s"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Do(tt.s, tt.i, tt.b)
			assert.Equal(t, actual, tt.exp)
			if tt.expErr != "" {
				assert.EqualError(t, err, tt.expErr)
			} else {
				assert.NoError(t, err)
			}
		})

	}

}
