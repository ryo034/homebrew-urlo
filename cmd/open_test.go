package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
	"urlo/core/domain"
	"urlo/core/infrastructure"
	"urlo/core/infrastructure/injector"
	"urlo/util"
)

func Test_Open_Cmd_Open_OK(t *testing.T) {
	u, _ := domain.NewStrictUrlFromString("https://google.com")
	tests := []struct {
		name    string
		args    [][]string
		prepare []domain.UrlMapJson
		want    []domain.UrlMapJson
		output  string
	}{
		{
			name:    "Open url",
			args:    [][]string{{"google"}},
			prepare: []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			want:    []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			output:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			defer func() {
				if err = os.Remove(util.TestFilePath); err != nil {
					fmt.Println("remove file error")
					panic(err)
				}
			}()

			if tt.prepare != nil {
				f, err := os.Create(util.TestFilePath)
				if err != nil {
					panic(err)
				}
				if err = json.NewEncoder(f).Encode(&tt.prepare); err != nil {
					panic(err)
				}
			}

			mce := new(util.MockCommandExecutor)
			mce.On("Open", u).Return(nil)
			inj := injector.NewInjector(util.TestFilePath, mce, infrastructure.NewPromptExecutor())
			openCmd := newOpenCmd(inj)

			for _, arg := range tt.args {
				openCmd.SetArgs([]string{arg[0]})
			}
			fnc := openCmd.Execute
			output, err := util.ExtractStdout(t, fnc)
			assert.NoError(t, err)
			if util.RemoveEscape(output) != tt.output {
				t.Errorf("failed to test. got: %q, want: %q", output, tt.output)
			}

			mce.AssertCalled(t, "Open", u)

			f, err := os.Open(util.TestFilePath)
			if err != nil {
				t.Errorf("newMockCmd() error = %v", err)
			}
			var urlMaps []domain.UrlMapJson
			if err = json.NewDecoder(f).Decode(&urlMaps); err != nil {
				t.Errorf("newMockCmd() error = %v", err)
				return
			}
			if !reflect.DeepEqual(urlMaps, tt.want) {
				t.Errorf("newMockCmd() got = %v, want %v", urlMaps, tt.want)
			}
		})
	}
}
