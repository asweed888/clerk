package create

import "fmt"

func Mod3Comment(mod1Name string, mod2Name string, mod3Name string, comment string) string {
    f := `/* <location: %s.%s.%s />

%s

*/`
    return fmt.Sprintf(f, mod1Name, mod2Name, mod3Name, comment)
}
