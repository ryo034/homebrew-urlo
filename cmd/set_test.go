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

var askMsg = "Are you sure you want to overwrite the existing data?"

func Test_Set_Cmd_Set_OK(t *testing.T) {
	tests := []struct {
		name        string
		args        [][]string
		prepare     []domain.UrlMapJson
		want        []domain.UrlMapJson
		askResponse bool
		input       string
		output      string
	}{
		{
			name:        "set",
			args:        [][]string{{"[{\"title\":\"google\",\"url\":\"https://google.com\"}]"}},
			prepare:     []domain.UrlMapJson{{Title: "yahoo", URL: "https://yahoo.com"}},
			want:        []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			askResponse: true,
			input:       "y\x1b",
			output:      "Set successfully\n",
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

			mpe := new(util.MockPromptExecutor)
			mpe.On("AskOne", askMsg).Return(tt.askResponse, nil)
			inj := injector.NewInjector(util.TestFilePath, infrastructure.NewCommandExecutor(), mpe)
			setCmd := newSetCmd(inj)
			setCmd.SetArgs(tt.args[0])

			fnc := setCmd.Execute
			output, err := util.ExtractStdout(t, fnc)
			assert.NoError(t, err)
			if util.RemoveEscape(output) != tt.output {
				t.Errorf("failed to test. got: %q, want: %q", output, tt.output)
			}

			mpe.AssertCalled(t, "AskOne", askMsg)

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
