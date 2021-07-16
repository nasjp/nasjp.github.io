package markdown

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToken(t *testing.T) {
	t.Parallel()

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(tt.markdown)
			got, err := tokenize(r)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tt.token, got, cmp.AllowUnexported(markdown{}, makdownElement{})); diff != "" {
				t.Errorf("%s mismatch (-want +got):\n%s", tt.name, diff)
			}
		})
	}
}
