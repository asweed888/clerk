package fs

type clerk struct {
    Directory directory
    DotClerkFile dotClerkFile
    CodeFile codeFile
    CodeFileMethod codeFileMethod
}

var Clerk = &clerk{}