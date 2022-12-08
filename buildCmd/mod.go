package buildCmd

type clerk struct {
    Golang c_golang
    Modernjs c_modernjs
    Python c_python
    Shell c_shell
}

var Clerk = &clerk{}