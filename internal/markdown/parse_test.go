package markdown

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	t.Parallel()

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(tt.markdown)
			got, err := parse(r)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tt.block, got, cmp.AllowUnexported(block{}, inline{})); diff != "" {
				t.Errorf("%s mismatch (-want +got):\n%s", tt.name, diff)
			}
		})
	}
}
