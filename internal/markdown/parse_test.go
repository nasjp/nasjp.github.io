package markdown

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	t.Parallel()

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := parse(tt.token)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tt.node, got, cmp.AllowUnexported(node{})); diff != "" {
				t.Errorf("%s mismatch (-want +got):\n%s", tt.name, diff)
			}
		})
	}
}
