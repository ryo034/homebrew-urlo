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

func Test_Delete_Cmd_Delete_OK(t *testing.T) {
	tests := []struct {
		name            string
		prepare         []domain.UrlMapJson
		argTitles       []string
		selectOneReturn string
		want            []domain.UrlMapJson
		output          string
	}{
		{
			name:            "delete url",
			prepare:         []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}, {Title: "yahoo", URL: "https://yahoo.com"}},
			argTitles:       []string{"google", "yahoo"},
			selectOneReturn: "google",
			want:            []domain.UrlMapJson{{Title: "yahoo", URL: "https://yahoo.com"}},
			output:          "Delete successfully google\n",
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
			mpe.On("SelectOne", "Select title", tt.argTitles).Return(tt.selectOneReturn, nil)
			inj := injector.NewInjector(util.TestFilePath, infrastructure.NewCommandExecutor(), mpe)
			deleteCmd := newDeleteCmd(inj)

			fnc := deleteCmd.Execute
			output, err := util.ExtractStdout(t, fnc)
			assert.NoError(t, err)
			if util.RemoveEscape(output) != tt.output {
				t.Errorf("failed to test. got: %q, want: %q", output, tt.output)
			}

			mpe.AssertCalled(t, "SelectOne", "Select title", tt.argTitles)

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
