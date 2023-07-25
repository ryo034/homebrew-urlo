package util

import (
	"bytes"
	"github.com/fatih/color"
	"github.com/stretchr/testify/mock"
	"io"
	"os"
	"regexp"
	"testing"
	"urlo/core/domain"
)

var TestFilePath = "test.json"

func ExtractStdout(t *testing.T, fnc func() error) (string, error) {
	t.Helper()

	orgStdout := os.Stdout
	orgColorOutput := color.Output

	pr, pw, _ := os.Pipe()
	os.Stdout = pw
	color.Output = pw

	if err := fnc(); err != nil {
		return "", err
	}

	if err := pw.Close(); err != nil {
		return "", err
	}
	os.Stdout = orgStdout
	color.Output = orgColorOutput

	buf := bytes.Buffer{}
	if _, err := io.Copy(&buf, pr); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type MockPromptExecutor struct {
	mock.Mock
}

func (m *MockPromptExecutor) AskOne(prompt string) (bool, error) {
	args := m.Called(prompt)
	return args.Bool(0), args.Error(1)
}

func (m *MockPromptExecutor) SelectOne(msg string, options []string) (string, error) {
	args := m.Called(msg, options)
	return args.String(0), args.Error(1)
}

type MockCommandExecutor struct {
	mock.Mock
}

func (m *MockCommandExecutor) Open(u domain.StrictUrl) error {
	args := m.Called(u)
	return args.Error(0)
}

// RemoveEscape remove ANSI escape sequences
func RemoveEscape(output string) string {
	re := regexp.MustCompile("\x1b\\[[0-9;]*m")
	return re.ReplaceAllString(output, "")
}
