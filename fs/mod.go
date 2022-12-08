package fs

type clerk struct {
    Directory c_directory
    ShellFile c_shellFile
    DotClerkFile c_dotClerkFile
    CodeFile c_codeFile
    CodeFileMethod c_codeFileMethod
}

var Clerk = &clerk{}