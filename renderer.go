package handlebars

import (
	"fmt"

	"github.com/go-mojito/mojito/log"
	"github.com/go-mojito/mojito/pkg/renderer"
	"github.com/infinytum/raymond/v2"
)

func init() {
	raymond.ResolvePartial = resolvePartial
	raymond.RegisterHelper("extends", helperExtends)
	raymond.RegisterHelper("view", helperView)
	raymond.RegisterHelper("when", helperWhen)
	raymond.RegisterHelper("formatdate", helperFormatdate)
}

type HandlebarsRenderer struct{}

// SetTemplateDir implements mojito.FileRenderer
func (*HandlebarsRenderer) SetTemplateDir(path string) error {
	templateDir = path
	return nil
}

// TemplateDir implements mojito.FileRenderer
func (*HandlebarsRenderer) TemplateDir() string {
	return templateDir
}

// Render implements mojito.Renderer
func (*HandlebarsRenderer) Render(view string, bag renderer.ViewBag) (string, error) {
	tpl, err := raymond.ParseFile(normalizeViewPath(view))
	if err != nil {
		log.Error(err)
		return fmt.Sprintf(raymondViewTemplateNotFound, view), err
	}
	return tpl.Exec(bag.ToMap())
}

func resolvePartial(view string) *raymond.Partial {
	path := normalizeViewPath(view)
	tpl, err := raymond.ParseFile(path)
	if err != nil {
		log.Error(err)
		return nil
	}
	return raymond.NewPartial(view, path, tpl)
}
