package handlebars

import (
	"github.com/go-mojito/mojito"
	"github.com/infinytum/injector"
)

func AsDefault() error {
	return As(injector.DefaultForType)
}

func As(name string) error {
	handlebars := &HandlebarsRenderer{}

	if err := injector.Singleton(func() mojito.Renderer {
		return handlebars
	}); err != nil {
		return err
	}

	if err := injector.Singleton(func() mojito.FileRenderer {
		return handlebars
	}); err != nil {
		return err
	}
	return nil
}
