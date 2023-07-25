package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"reflect"
	"testing"
	"urlo/core/domain"
	"urlo/core/infrastructure"
	"urlo/core/infrastructure/injector"
	"urlo/util"
)

func Test_List_Cmd_Output_OK(t *testing.T) {
	tests := []struct {
		name       string
		args       [][]string
		json       string
		jsonString string
		prepare    []domain.UrlMapJson
		want       []domain.UrlMapJson
		output     string
	}{
		{
			name:       "records are empty",
			args:       [][]string{},
			json:       "false",
			jsonString: "false",
			prepare:    nil,
			want:       []domain.UrlMapJson{},
			output:     "Error: No URL found\n",
		},
		{
			name:       "output list format",
			args:       [][]string{},
			json:       "false",
			jsonString: "false",
			prepare:    []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			want:       []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			output:     "google - https://google.com\n",
		},
		{
			name:       "output json format",
			args:       [][]string{},
			json:       "true",
			jsonString: "false",
			prepare:    []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			want:       []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			output: `[
  {
    "title": "google",
    "url": "https://google.com"
  }
]
`,
		},
		{
			name:       "output jsonString format",
			args:       [][]string{},
			json:       "false",
			jsonString: "true",
			prepare:    []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			want:       []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			output:     "[{\"title\":\"google\",\"url\":\"https://google.com\"}]\n",
		},
		{
			name:       "output error message when json and jsonString are true",
			args:       [][]string{},
			json:       "true",
			jsonString: "true",
			prepare:    []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			want:       []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			output:     "Can't use both -j and -s\n",
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

			inj := injector.NewInjector(util.TestFilePath, infrastructure.NewCommandExecutor(), infrastructure.NewPromptExecutor())
			listCmd := newListCmd(inj)
			_ = listCmd.Flags().Set("json", tt.json)
			_ = listCmd.Flags().Set("string", tt.jsonString)

			for _, arg := range tt.args {
				listCmd.SetArgs([]string{arg[0], arg[1]})
			}
			fnc := listCmd.Execute
			output, err := util.ExtractStdout(t, fnc)
			assert.NoError(t, err)
			if util.RemoveEscape(output) != tt.output {
				t.Errorf("failed to test. got: %q, want: %q", output, tt.output)
			}

			f, err := os.Open(util.TestFilePath)
			if err != nil {
				t.Errorf("newListCmd() error = %v", err)
			}
			var urlMaps []domain.UrlMapJson
			if err = json.NewDecoder(f).Decode(&urlMaps); err != nil {
				if err == io.EOF {
					return
				}
				t.Errorf("newListCmd() error = %v", err)
				return
			}
			if !reflect.DeepEqual(urlMaps, tt.want) {
				t.Errorf("newListCmd() got = %v, want %v", urlMaps, tt.want)
			}
		})
	}
}
