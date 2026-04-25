package driver_test

import (
	"exercices/internal/chapter4/driver"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m,
		map[string]func(){
			"driver": driver.Main,
		},
	)
}

func Test_passval(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "script/",
	})
}
