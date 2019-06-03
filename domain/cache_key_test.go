package domain

import (
	"net/url"
	"testing"
)

func TestCacheKey_Key(t *testing.T) {
	type fields struct {
		URL func() *url.URL
	}
	tests := []struct {
		name   string
		fields fields
		want   Key
	}{
		{
			fields: fields{
				URL: func() *url.URL {
					urlObj, _ := url.ParseRequestURI("https://atcoder.jp/contests/abc119/tasks/abc119_b")
					return urlObj
				},
			},
			want: Key("atcoder.jp/contests/abc119/tasks/abc119_b"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &CacheKey{
				URL: tt.fields.URL(),
			}
			if got := a.Key(); got != tt.want {
				t.Errorf("CacheKey.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
