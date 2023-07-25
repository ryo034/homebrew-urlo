package gateway

import (
	"reflect"
	"testing"
	"urlo/core/domain"
)

func Test_gatewayAdapter_Adapt(t *testing.T) {
	type args struct {
		v domain.UrlMapJson
	}

	ti, _ := domain.NewTitle("google")
	u, _ := domain.NewStrictUrlFromString("https://google.com")
	want := domain.NewUrlMap(ti, u)

	tests := []struct {
		name    string
		args    args
		want    domain.UrlMap
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				v: domain.UrlMapJson{
					Title: "google",
					URL:   "https://google.com",
				},
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gatewayAdapter{}
			got, err := g.Adapt(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Adapt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Adapt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
