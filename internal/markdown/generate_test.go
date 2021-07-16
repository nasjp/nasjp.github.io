package markdown

import (
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := generate(tt.node)
			if err != nil {
				t.Fatal(err)
			}

			bs, err := io.ReadAll(got)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tt.html, string(bs)); diff != "" {
				t.Errorf("%s mismatch (-want +got):\n%s", tt.name, diff)
			}
		})
	}
}
