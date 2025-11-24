package handlebars

import (
	"fmt"

	"github.com/go-mojito/mojito/log"
	"github.com/infinytum/raymond/v2"
)

// helperView provides dynamic view injection & rendering
func helperView(view string, bag interface{}) raymond.SafeString {
	tpl, err := raymond.ParseFile(normalizeViewPath(view))
	if err != nil {
		log.Error("Failed to parse template", "error", err, "view", view)
		return raymond.SafeString(fmt.Sprintf(raymondViewTemplateNotFound, view))
	}
	return raymond.SafeString(tpl.MustExec(bag))
}
