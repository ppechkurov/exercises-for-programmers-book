package passval_test

import (
	"exercices/internal/chapter4/passval"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m,
		map[string]func(){
			"passval": passval.Main,
		},
	)
}

func Test_passval(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "script/",
	})
}
