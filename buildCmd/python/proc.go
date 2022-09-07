package python

import (
	"github.com/asweed888/clerk/schema"
)

func Proc(scm *schema.ClerkYaml) error {

    var modFilePath string

    for _, clerk := range scm.Spec {
        location := clerk.Location

        for _, mod1 := range clerk.Modules {
            modFilePath := "./%s/%s/__init__.py"
        }
    }

    return nil
}
