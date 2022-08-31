package note

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const templatesDirectoryPath = "./templates/"

func findTemplate(title string) (string, error) {
	files, err := os.ReadDir(templatesDirectoryPath)
	if err != nil {
		return "", ErrFindTemplatesReadDir
	}

	templates := []string{}

	for _, file := range files {
		if !file.Type().IsRegular() {
			continue
		}

		if filepath.Ext(file.Name()) != ".md" {
			continue
		}

		withoutExtension := strings.TrimSuffix(
			file.Name(), filepath.Ext(file.Name()),
		)

		templates = append(templates, withoutExtension)
	}

	sort.Slice(
		templates,
		func(i, j int) bool {
			x := len(strings.Split(templates[i], "."))
			y := len(strings.Split(templates[j], "."))
			return x > y
		},
	)

	for _, t := range templates {
		if strings.HasPrefix(title, t) {
			return fmt.Sprintf("%s%s.md", templatesDirectoryPath, t), nil
		}
	}

	return "", ErrMissingTemplates
}

func readTemplate(p string) (string, error) {
	f, err := os.Open(p) //nolint:gosec
	if err != nil {
		return "", fmt.Errorf("failed to open template file: %w", err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("failed to read contents of template file: %w", err)
	}

	return string(b), nil
}
