package model

type Cwd struct {
    Path string
    Upstream *DeclareUpstream
    CodeFile *DeclareCodeFile
}
