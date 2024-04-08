package core

import (
	"fmt"
	"html/template"
	"path"
	"path/filepath"
	"strings"
)

type TemplateManager struct {
	baseDir      string
	commonPrefix string

	common *template.Template
	tmpl   map[string]*template.Template
}

func NewTemplateManager(templateDir string) *TemplateManager {
	return &TemplateManager{
		baseDir:      templateDir,
		commonPrefix: "_",
		common:       template.New("common"),
		tmpl:         make(map[string]*template.Template),
	}
}

func (tm *TemplateManager) Load() error {
	if tm.common == nil {
		tm.common = template.New("common")
	}

	pattern := path.Join(tm.baseDir, "*.html")
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	if len(filenames) == 0 {
		return fmt.Errorf("html/template: pattern matches no files: %#q", pattern)
	}

	commons := make([]string, 0)
	templates := make([]string, 0)

	for _, filename := range filenames {
		if tm.isCommonTemplate(filename) {
			commons = append(commons, filename)
		} else {
			templates = append(templates, filename)
		}
	}

	for _, filename := range commons {
		if _, err := tm.common.ParseFiles(filename); err != nil {
			return err
		}
	}

	for _, filename := range templates {
		err := tm.Add(filename)
		if err != nil {
			return err
		}
	}

	return nil
}

func (tm *TemplateManager) Add(filename string) error {
	if tm.tmpl == nil {
		tm.tmpl = make(map[string]*template.Template)
	}

	if tm.common == nil {
		tm.common = template.New("common")
	}

	t, err := tm.common.Clone()
	if err != nil {
		return err
	}

	t, err = t.ParseFiles(filename)
	if err != nil {
		return err
	}

	tm.tmpl[path.Base(filename)] = t

	printTemplateContent(path.Base(filename), t)

	return nil
}

func (t *TemplateManager) Get(name string) *template.Template {
	nt, ok := t.tmpl[name]
	if ok {
		return nt
	}

	return t.common
}

func (t *TemplateManager) Templates() map[string]*template.Template {
	return t.tmpl
}

func (tm *TemplateManager) isCommonTemplate(filename string) bool {
	filename, err := filepath.Rel(tm.baseDir, filename)
	if err != nil {
		return false
	}

	for _, part := range filepath.SplitList(filename) {
		if strings.HasPrefix(part, tm.commonPrefix) {
			return true
		}
	}

	return false
}

func printTemplateContent(key string, t *template.Template) {
	names := make([]string, 0, len(t.Templates()))
	for _, tmpl := range t.Templates() {
		names = append(names, tmpl.Name())
	}

	fmt.Printf("%s: %s\n", key, strings.Join(names, ", "))
}
