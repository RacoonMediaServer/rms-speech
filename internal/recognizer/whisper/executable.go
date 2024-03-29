package whisper

import (
	"context"
	"golang.org/x/text/language"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ExecutableRecognizer struct {
	pathToExecutable string
}

func NewExecutableRecognizer(pathToExecutable string) *ExecutableRecognizer {
	return &ExecutableRecognizer{pathToExecutable: pathToExecutable}
}

func (r ExecutableRecognizer) Transcribe(ctx context.Context, inputFile string, lang language.Tag) (string, error) {
	args := []string{"--output_format", "txt", "--output_dir", os.TempDir()}
	if lang.String() != "und" {
		args = append(args, "--language")
		args = append(args, lang.String())
	}
	args = append(args, inputFile)

	cmd := exec.CommandContext(ctx, r.pathToExecutable, args...)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	resultFile := filepath.Join(os.TempDir(), strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))+".txt")
	result, err := os.ReadFile(resultFile)
	if err != nil {
		return "", err
	}
	_ = os.Remove(resultFile)
	return string(result), nil
}
