package godynamo

import (
	"testing"

	dynamopb "github.com/decade-eng/buf-public-apis/gen/go/dynamo"
)

func TestBuildKeyTag(t *testing.T) {
	cases := []struct {
		name string
		cfg  *dynamopb.KeyConfig
		want string
	}{
		{
			name: "hash key",
			cfg:  &dynamopb.KeyConfig{Type: dynamopb.KeyType_KEY_TYPE_HASH, ColumnName: "hash_key"},
			want: `dynamo:"hash_key,hash"`,
		},
		{
			name: "range key",
			cfg:  &dynamopb.KeyConfig{Type: dynamopb.KeyType_KEY_TYPE_RANGE, ColumnName: "created_at"},
			want: `dynamo:"created_at,range"`,
		},
		{
			name: "column name only",
			cfg:  &dynamopb.KeyConfig{ColumnName: "email"},
			want: `dynamo:"email"`,
		},
		{
			name: "omit empty on a plain column",
			cfg:  &dynamopb.KeyConfig{ColumnName: "deleted_at", OmitEmpty: true},
			want: `dynamo:"deleted_at,omitempty"`,
		},
		{
			name: "omit empty alongside a key flag",
			cfg:  &dynamopb.KeyConfig{Type: dynamopb.KeyType_KEY_TYPE_RANGE, ColumnName: "created_at", OmitEmpty: true},
			want: `dynamo:"created_at,range,omitempty"`,
		},
		{
			name: "omit empty with no column name falls back to the field name",
			cfg:  &dynamopb.KeyConfig{OmitEmpty: true},
			want: `dynamo:",omitempty"`,
		},
		{
			name: "a field asking for nothing keeps no tag",
			cfg:  &dynamopb.KeyConfig{},
			want: "",
		},
		{
			name: "hash key with no column name",
			cfg:  &dynamopb.KeyConfig{Type: dynamopb.KeyType_KEY_TYPE_HASH},
			want: `dynamo:",hash"`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := buildKeyTag(c.cfg); got != c.want {
				t.Errorf("buildKeyTag() = %q, want %q", got, c.want)
			}
		})
	}
}
