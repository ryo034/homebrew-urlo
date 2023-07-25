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

//ファイルが空の場合にデータが追加される
//対象のタイトルがすでに存在する場合にデータが追加されない
//対象のタイトルが存在しない場合にデータが追加される

func Test_Bulk_Add_Cmd_Add_OK(t *testing.T) {
	tests := []struct {
		name    string
		args    [][]string
		prepare []domain.UrlMapJson
		want    []domain.UrlMapJson
		output  string
	}{
		{
			name:    "bulk add if file is empty",
			args:    [][]string{{"[{\"title\":\"google\",\"url\":\"https://google.com\"}]"}},
			prepare: nil,
			want:    []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			output:  "Successfully add all the new URL map.\n",
		},
		{
			name:    "cannot bulk add if target title is already exist",
			args:    [][]string{{"[{\"title\":\"google\",\"url\":\"https://google.com\"}]"}},
			prepare: []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			want:    []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}},
			output:  "Error: already exist: google\n",
		},
		{
			name:    "bulk add if target title is not exist",
			args:    [][]string{{"[{\"title\":\"google\",\"url\":\"https://google.com\"}]"}},
			prepare: []domain.UrlMapJson{{Title: "yahoo", URL: "https://yahoo.com"}},
			want:    []domain.UrlMapJson{{Title: "yahoo", URL: "https://yahoo.com"}, {Title: "google", URL: "https://google.com"}},
			output:  "Successfully add all the new URL map.\n",
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
			bulkAddCmd := newBulkAddCmd(inj)

			for _, arg := range tt.args {
				bulkAddCmd.SetArgs([]string{arg[0]})
			}
			fnc := bulkAddCmd.Execute
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
