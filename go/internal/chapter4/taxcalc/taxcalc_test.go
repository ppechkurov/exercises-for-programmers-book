package taxcalc_test

import (
	"bytes"
	"exercices/internal/taxcalc"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name    string
		r       []byte
		w       []byte
		wantErr bool
	}{
		{
			name:    "aoeu",
			r:       []byte("aoeu\n"),
			w:       []byte{},
			wantErr: true,
		},
		{
			name:    "10",
			r:       []byte("10"),
			w:       []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytes.NewBuffer(nil)
			gotErr := taxcalc.Main(bytes.NewReader(tt.r), w)
			if tt.wantErr {
				require.ErrorIs(t, gotErr, taxcalc.ErrReadOrder)
				return
			}

			require.NoError(t, gotErr)

			actual, err := io.ReadAll(w)
			require.NoError(t, err)
			require.Equal(t, "What's the order? 11", string(actual))
		})
	}
}
