package config

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/asweed888/clerk/domain/model"
)

func pkgname(path string) string {
    idx := strings.LastIndexByte(path, '/') + 1
    return path[idx:]
}


func isDDD(c *model.TacitConfig) bool {
    return c.Arch == "ddd"
}

func isDomainModel(c *model.TacitConfig, path string) bool {
    return isDDD(c) && strings.Contains(path, "/domain/model")
}

func isDomainRepository(c *model.TacitConfig, path string) bool {
    return isDDD(c) && strings.Contains(path, "/domain/repository")
}

func isDomainService(c *model.TacitConfig, path string) bool {
    return isDDD(c) && strings.Contains(path, "/domain/service")
}

func isInfra(c *model.TacitConfig, path string) bool {
    return isDDD(c) && strings.Contains(path, "/infra")
}

func isUseCase(c *model.TacitConfig, path string) bool {
    return isDDD(c) && strings.Contains(path, "/usecase")
}

func isPresentation(c *model.TacitConfig, path string) bool {
    return isDDD(c) && strings.Contains(path, "/presentation")
}

func isContains2Codefile(codeFileFullPath string, targetStr string) bool {
    file, _ := os.Open(codeFileFullPath)
    defer file.Close()

    b, _ := ioutil.ReadAll(file)
    return strings.Contains(string(b), targetStr)
}
