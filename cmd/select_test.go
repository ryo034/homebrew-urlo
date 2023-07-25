package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
	"urlo/core/domain"
	"urlo/core/infrastructure/injector"
	"urlo/util"
)

func Test_Select_Cmd_Select_OK(t *testing.T) {
	tests := []struct {
		name            string
		query           string
		prepare         []domain.UrlMapJson
		argTitles       []string
		selectOneReturn string
		openURL         string
		want            []domain.UrlMapJson
		output          string
	}{
		{
			name:            "select url",
			query:           "",
			prepare:         []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}, {Title: "yahoo", URL: "https://yahoo.com"}},
			argTitles:       []string{"google", "yahoo"},
			selectOneReturn: "google",
			openURL:         "https://google.com",
			want:            []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}, {Title: "yahoo", URL: "https://yahoo.com"}},
			output:          "Selected google\n",
		},
		{
			name:            "select url with query",
			query:           "y",
			prepare:         []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}, {Title: "yahoo", URL: "https://yahoo.com"}},
			argTitles:       []string{"yahoo"},
			selectOneReturn: "yahoo",
			openURL:         "https://yahoo.com",
			want:            []domain.UrlMapJson{{Title: "google", URL: "https://google.com"}, {Title: "yahoo", URL: "https://yahoo.com"}},
			output:          "Selected yahoo\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := domain.NewStrictUrlFromString(tt.openURL)

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

			mpe := new(util.MockPromptExecutor)
			mpe.On("SelectOne", "Select title", tt.argTitles).Return(tt.selectOneReturn, nil)

			inj := injector.NewInjector(util.TestFilePath, mce, mpe)
			selectCmd := newSelectCmd(inj)
			if tt.query != "" {
				_ = selectCmd.Flags().Set("query", tt.query)
			}

			fnc := selectCmd.Execute
			output, err := util.ExtractStdout(t, fnc)
			assert.NoError(t, err)
			if util.RemoveEscape(output) != tt.output {
				t.Errorf("failed to test. got: %q, want: %q", output, tt.output)
			}

			mce.AssertCalled(t, "Open", u)
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
