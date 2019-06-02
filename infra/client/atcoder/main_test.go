package atcoder

import (
	"reflect"
	"testing"

	"github.com/smith-30/acc/domain"
)

func Test_client_GetTestCase(t *testing.T) {
	type args struct {
		u string
	}
	tests := []struct {
		name    string
		a       *client
		args    args
		want    []domain.TestCase
		wantErr bool
	}{
		{
			a: &client{},
			args: args{
				u: "https://atcoder.jp/contests/abc119/tasks/abc119_a",
			},
			want: []domain.TestCase{
				domain.TestCase{
					Content: "2019/04/30\n",
					Exp:     "Heisei\n",
				},
				domain.TestCase{
					Content: "2019/11/01\n",
					Exp:     "TBD\n",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &client{}
			got, err := a.GetTestCase(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetTestCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.GetTestCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
