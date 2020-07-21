package adapters

import (
	"github.com/cli-playground/devfile-parser/pkg/devfile/adapters/common"
)

type PlatformAdapter interface {
	Push(parameters common.PushParameters) error
	Build(parameters common.BuildParameters) error
	Deploy(parameters common.DeployParameters) error
	DoesComponentExist(cmpName string) bool
	Delete(labels map[string]string) error
	DeployDelete(manifest []byte) error
}
