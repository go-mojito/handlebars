package handlebars

import (
	"fmt"

	"github.com/go-mojito/mojito/log"
	"github.com/infinytum/raymond/v2"
)

// helperExtends provides the template extension feature
func helperExtends(view string, options *raymond.Options) raymond.SafeString {
	tpl, err := raymond.ParseFile(normalizeViewPath(view))
	if err != nil {
		log.Error("Failed to parse template", "error", err, "view", view)
		return raymond.SafeString(fmt.Sprintf(raymondViewTemplateNotFound, view))
	}

	bag, ok := options.Ctx().(map[string]interface{})
	if !ok {
		return raymond.SafeString("Unable to convert ViewBag")
	}
	bag["subview"] = options.FnWith(options.Ctx())
	return raymond.SafeString(tpl.MustExec(bag))
}
