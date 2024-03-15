package json_test

import (
	"eventi/oss/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJson(t *testing.T) {
	t.Run("Unmarshall", func(t *testing.T) {
		t.Run("Unmarshall to desired type", func(t *testing.T) {
			type Dog struct {
				Name          string
				FavouriteFood string
			}

			require.Equal(
				t,
				&Dog{"Rex", "Bones"},
				json.Unmarshal[Dog]([]byte(`{"Name": "Rex", "FavouriteFood": "Bones"}`)),
			)
		})

		t.Run("Unmarshall to desired string", func(t *testing.T) {
			rocco := "rocco"

			require.Equal(
				t,
				&rocco,
				json.Unmarshal[string]([]byte(`"rocco"`)),
			)
		})
	})
}
