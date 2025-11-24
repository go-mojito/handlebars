package handlebars

import (
	"github.com/go-mojito/mojito"
	"github.com/infinytum/injector"
)

type HandlebarsOption func(*HandlebarsRenderer) error

func AsDefault(opts ...HandlebarsOption) error {
	return As(injector.DefaultForType, opts...)
}

func As(name string, opts ...HandlebarsOption) error {
	handlebars := &HandlebarsRenderer{}

	for _, opt := range opts {
		if err := opt(handlebars); err != nil {
			return err
		}
	}

	if err := injector.Singleton(func() mojito.Renderer {
		return handlebars
	}, name); err != nil {
		return err
	}

	if err := injector.Singleton(func() mojito.FileRenderer {
		return handlebars
	}, name); err != nil {
		return err
	}
	return nil
}

func WithTemplateDir(dir string) HandlebarsOption {
	return func(handlebars *HandlebarsRenderer) error {
		return handlebars.SetTemplateDir(dir)
	}
}
