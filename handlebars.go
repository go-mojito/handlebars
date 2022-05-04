package handlebars

import (
	"github.com/go-mojito/mojito"
	"github.com/infinytum/injector"
)

func AsDefault() {
	As(injector.DefaultForType)
}

func As(name string) {
	handlebars := &HandlebarsRenderer{}
	injector.Singleton(func() mojito.Renderer {
		return handlebars
	})
	injector.Singleton(func() mojito.FileRenderer {
		return handlebars
	})
}
