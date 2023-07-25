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

func Test_Add_Cmd_Add_OK(t *testing.T) {
	tests := []struct {
		name     string
		args     [][]string
		override string
		prepare  []domain.UrlMapJson
		want     []domain.UrlMapJson
		output   string
	}{
		{
			name:     "add if file is empty",
			args:     [][]string{{"google", "https://google.com"}},
			override: "false",
			prepare:  nil,
			want:     []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			output:   "Add successfully google - https://google.com\n",
		},
		{
			name:     "add if target title is not exist and override is off",
			args:     [][]string{{"yahoo", "https://yahoo.com"}},
			override: "false",
			prepare:  []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			want:     []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}, {Title: "yahoo", URL: "https://yahoo.com"}},
			output:   "Add successfully yahoo - https://yahoo.com\n",
		},
		{
			name:     "add if target title is not exist and override is on",
			args:     [][]string{{"yahoo", "https://yahoo.com"}},
			override: "true",
			prepare:  []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			want:     []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}, {Title: "yahoo", URL: "https://yahoo.com"}},
			output:   "Add successfully yahoo - https://yahoo.com\n",
		},
		{
			name:     "not add if target title is already exist and override is off",
			args:     [][]string{{"google", "https://google.com"}},
			override: "false",
			prepare:  []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			want:     []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			output:   "Error: Already exist title: 'google'\n",
		},
		{
			name:     "override if target title is already exist and override is on",
			args:     [][]string{{"google", "https://google1.com"}},
			override: "true",
			prepare:  []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			want:     []domain.UrlMapJson{{Title: "google", URL: "https://google1.com"}},
			output:   "Update successfully\n",
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
			addCmd := newAddCmd(inj)
			_ = addCmd.Flags().Set("override", tt.override)

			for _, arg := range tt.args {
				addCmd.SetArgs([]string{arg[0], arg[1]})
			}
			fnc := addCmd.Execute
			output, err := util.ExtractStdout(t, fnc)
			assert.NoError(t, err)
			if util.RemoveEscape(output) != tt.output {
				t.Errorf("failed to test. got: %q, want: %q", output, tt.output)
			}

			f, err := os.Open(util.TestFilePath)
			if err != nil {
				t.Errorf("newAddCmd() error = %v", err)
			}
			var urlMaps []domain.UrlMapJson
			if err = json.NewDecoder(f).Decode(&urlMaps); err != nil {
				t.Errorf("newAddCmd() error = %v", err)
				return
			}
			if !reflect.DeepEqual(urlMaps, tt.want) {
				t.Errorf("newAddCmd() got = %v, want %v", urlMaps, tt.want)
			}
		})
	}
}
