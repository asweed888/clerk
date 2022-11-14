package common

import "os"

type DirectorySpec struct {}

var Directory = &DirectorySpec{}

func (s *DirectorySpec) Create(dirname string) error {
    if err := os.MkdirAll(dirname, os.ModePerm); err != nil {
        return err
    }

    return nil
}
