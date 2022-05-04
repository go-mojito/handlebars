package handlebars

import "github.com/infinytum/raymond/v2"

// helperWhen provides a shorthand if inside expression blocks
func helperWhen(conditional bool, whenTrue interface{}, whenFalse interface{}, options *raymond.Options) interface{} {
	if conditional {
		return whenTrue
	}
	return whenFalse
}
