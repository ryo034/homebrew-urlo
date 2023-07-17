package util

import (
	"errors"
	"net/url"
	"reflect"
	"testing"
)

func TestUrlMaps_IsAlreadyExist(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		us   UrlMaps
		args args
		want bool
	}{
		{
			name: "should return true if title already exists",
			us:   UrlMaps{values: []UrlMap{{Title: "Google", URL: &url.URL{}}}},
			args: args{title: "Google"},
			want: true,
		},
		{
			name: "should return false if title does not exist",
			us:   UrlMaps{values: []UrlMap{{Title: "Google", URL: &url.URL{}}}},
			args: args{title: "Yahoo"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.us.IsAlreadyExist(tt.args.title); got != tt.want {
				t.Errorf("IsAlreadyExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrlMaps_TitleMaxLen(t *testing.T) {
	tests := []struct {
		name string
		us   UrlMaps
		want int
	}{
		{
			name: "should return max length of title",
			us:   UrlMaps{values: []UrlMap{{Title: "Google", URL: &url.URL{}}, {Title: "Yahoo", URL: &url.URL{}}}},
			want: 6,
		},
		{
			name: "should return 0 if there is no title",
			us:   UrlMaps{values: []UrlMap{}},
			want: 0,
		},
		{
			name: "should return max length of Japanese title",
			us:   UrlMaps{values: []UrlMap{{Title: "グーグル", URL: &url.URL{}}, {Title: "ヤフー", URL: &url.URL{}}}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.us.TitleMaxLen(); got != tt.want {
				t.Errorf("TitleMaxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToUrlMaps(t *testing.T) {
	type args struct {
		items [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    UrlMaps
		wantErr bool
	}{
		{
			name: "should return UrlMaps",
			args: args{items: [][]string{{"Google", "https://www.google.com/"}, {"Yahoo", "https://www.yahoo.co.jp/"}}},
			want: UrlMaps{[]UrlMap{{Title: "Google", URL: &url.URL{Scheme: "https", Host: "www.google.com", Path: "/"}}, {Title: "Yahoo", URL: &url.URL{Scheme: "https", Host: "www.yahoo.co.jp", Path: "/"}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToUrlMaps(tt.args.items)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToUrlMaps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToUrlMaps() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrlMaps_GetTitles(t *testing.T) {
	tests := []struct {
		name string
		us   UrlMaps
		want []string
	}{
		{
			name: "should return titles",
			us:   UrlMaps{[]UrlMap{{Title: "Google", URL: &url.URL{}}, {Title: "Yahoo", URL: &url.URL{}}}},
			want: []string{"Google", "Yahoo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.us.GetTitles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTitles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrlMaps_GetItemFromLabel(t *testing.T) {
	type args struct {
		label string
	}
	tests := []struct {
		name    string
		us      UrlMaps
		args    args
		want    UrlMap
		wantErr bool
	}{
		{
			name:    "should return UrlMap",
			us:      UrlMaps{[]UrlMap{{Title: "Google", URL: &url.URL{}}, {Title: "Yahoo", URL: &url.URL{}}}},
			args:    args{label: "Google"},
			want:    UrlMap{Title: "Google", URL: &url.URL{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.us.GetItemFromLabel(tt.args.label)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItemFromLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemFromLabel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUrlMaps_OK(t *testing.T) {
	type args struct {
		values []UrlMap
	}
	tests := []struct {
		name    string
		args    args
		want    UrlMaps
		wantErr bool
	}{
		{
			name:    "should return UrlMaps",
			args:    args{values: []UrlMap{{Title: "Google", URL: &url.URL{}}, {Title: "Yahoo", URL: &url.URL{}}}},
			want:    UrlMaps{[]UrlMap{{Title: "Google", URL: &url.URL{}}, {Title: "Yahoo", URL: &url.URL{}}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUrlMaps(tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUrlMaps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUrlMaps() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUrlMaps_Error(t *testing.T) {
	type args struct {
		values []UrlMap
	}
	tests := []struct {
		name    string
		args    args
		want    UrlMaps
		wantErr error
	}{
		{
			name:    "should return error if there is no url",
			args:    args{values: []UrlMap{{Title: "Google", URL: nil}}},
			want:    UrlMaps{[]UrlMap{}},
			wantErr: errors.New("url is required"),
		},
		{
			name:    "should return error if there is no title",
			args:    args{values: []UrlMap{{Title: "", URL: &url.URL{}}}},
			want:    UrlMaps{[]UrlMap{}},
			wantErr: errors.New("title is required"),
		},
		{
			name: "should return UrlMaps",
			args: args{values: []UrlMap{{Title: "Google", URL: &url.URL{}}, {Title: "Google", URL: &url.URL{}}}},
			want: UrlMaps{},
			wantErr: &DuplicationError{
				Values: []DuplicatedUrlMaps{
					{
						Values: UrlMaps{values: []UrlMap{{Title: "Google", URL: &url.URL{}}, {Title: "Google", URL: &url.URL{}}}},
						Title:  "Google",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUrlMaps(tt.args.values)
			if err == nil {
				if tt.wantErr != nil {
					t.Errorf("NewUrlMaps() got = %v, want %v", got, tt.want)
					return
				}
			} else {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("NewUrlMaps() got = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
