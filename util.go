package handlebars

import (
	"path/filepath"
	"strings"

	"github.com/go-mojito/mojito"
	"github.com/go-mojito/mojito/log"
)

var (
	raymondViewTemplateNotFound = "View Template '%s' not found or not readable"
	templateDir                 = ""
)

// normalizeViewPath ensures the path is within bounds and ends with .mojito
func normalizeViewPath(view string) string {
	path := TemplateDir() + view + ".mojito"
	if strings.HasPrefix(path, TemplateDir()) {
		return path
	}
	log.Warnf("Attempted path traversal to " + path)
	return TemplateDir()
}

func TemplateDir() string {
	dir := mojito.ResourcesDir() + "/templates"
	if templateDir != "" {
		dir = templateDir
	}
	path, err := filepath.Abs(dir)
	if err != nil {
		return mojito.ResourcesDir() + "/templates/"
	}
	return path + "/"
}
