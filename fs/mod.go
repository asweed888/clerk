package fs

type clerk struct {
    Directory directory
    ShellFile shellFile
    DotClerkFile dotClerkFile
    CodeFile codeFile
    CodeFileMethod codeFileMethod
}

var Clerk = &clerk{}