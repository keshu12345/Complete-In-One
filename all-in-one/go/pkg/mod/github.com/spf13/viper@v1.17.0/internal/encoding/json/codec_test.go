package json

import (
	"reflect"
	"testing"
)

// encoded form of the data
const encoded = `{
  "key": "value",
  "list": [
    "item1",
    "item2",
    "item3"
  ],
  "map": {
    "key": "value"
  },
  "nested_map": {
    "map": {
      "key": "value",
      "list": [
        "item1",
        "item2",
        "item3"
      ]
    }
  }
}`

// Viper's internal representation
var data = map[string]any{
	"key": "value",
	"list": []any{
		"item1",
		"item2",
		"item3",
	},
	"map": map[string]any{
		"key": "value",
	},
	"nested_map": map[string]any{
		"map": map[string]any{
			"key": "value",
			"list": []any{
				"item1",
				"item2",
				"item3",
			},
		},
	},
}

func TestCodec_Encode(t *testing.T) {
	codec := Codec{}

	b, err := codec.Encode(data)
	if err != nil {
		t.Fatal(err)
	}

	if encoded != string(b) {
		t.Fatalf("decoded value does not match the expected one\nactual:   %#v\nexpected: %#v", string(b), encoded)
	}
}

func TestCodec_Decode(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		codec := Codec{}

		v := map[string]any{}

		err := codec.Decode([]byte(encoded), v)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(data, v) {
			t.Fatalf("decoded value does not match the expected one\nactual:   %#v\nexpected: %#v", v, data)
		}
	})

	t.Run("InvalidData", func(t *testing.T) {
		codec := Codec{}

		v := map[string]any{}

		err := codec.Decode([]byte(`invalid data`), v)
		if err == nil {
			t.Fatal("expected decoding to fail")
		}

		t.Logf("decoding failed as expected: %s", err)
	})
}