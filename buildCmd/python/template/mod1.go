package template


type Mod1 struct {}

type Mod1Comment struct {}

type Mod1Clerk struct {}


type Mod1MethodStruct struct {}
var Mod1Method = &Mod1MethodStruct{}

func (m *Mod1MethodStruct) Get() string {
    return `

def _%s():
    print("this is clerk's default return value")`
}
