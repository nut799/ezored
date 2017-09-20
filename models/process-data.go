package models

import (
	"fmt"
	"github.com/ezored/ezored/constants"
	"github.com/ezored/ezored/utils/os-utils"
	"os"
	"path/filepath"
	"strings"
)

type ProcessData struct {
	ProjectName    string
	TargetName     string
	DependencyName string
	ProjectRootDir string
	TargetDir      string
	FullTargetDir  string
	TempDir        string
	BuildDir       string
	VendorDir      string
	DirSep         string
}

func (This *ProcessData) GetEnviron() []string {
	env := os.Environ()

	env = append(env, fmt.Sprintf("%sPROJECT_ROOT=%s", constants.ENV_VAR_PREFIX, This.ProjectRootDir))
	env = append(env, fmt.Sprintf("%sTEMP_DIR=%s", constants.ENV_VAR_PREFIX, This.TempDir))
	env = append(env, fmt.Sprintf("%sTARGET_DIR=%s", constants.ENV_VAR_PREFIX, This.TargetDir))
	env = append(env, fmt.Sprintf("%sFULL_TARGET_DIR=%s", constants.ENV_VAR_PREFIX, This.FullTargetDir))
	env = append(env, fmt.Sprintf("%sBUILD_DIR=%s", constants.ENV_VAR_PREFIX, This.BuildDir))
	env = append(env, fmt.Sprintf("%sVENDOR_DIR=%s", constants.ENV_VAR_PREFIX, This.VendorDir))
	env = append(env, fmt.Sprintf("%sTARGET_NAME=%s", constants.ENV_VAR_PREFIX, This.TargetName))
	env = append(env, fmt.Sprintf("%sDEPENDENCY_NAME=%s", constants.ENV_VAR_PREFIX, This.DependencyName))
	env = append(env, fmt.Sprintf("%sDS=%s", constants.ENV_VAR_PREFIX, This.DirSep))

	return env
}

func (This *ProcessData) Reset() {
	This.TargetName = ""
	This.DependencyName = ""
	This.DirSep = string(filepath.Separator)
	This.ProjectRootDir = osutils.GetCurrentDir()
	This.TargetDir = filepath.Join(osutils.GetCurrentDir(), "target")
	This.FullTargetDir = ""
	This.BuildDir = filepath.Join(osutils.GetCurrentDir(), "build")
	This.TempDir = filepath.Join(osutils.GetCurrentDir(), "temp")
	This.VendorDir = filepath.Join(osutils.GetCurrentDir(), "vendor")
}

func (This *ProcessData) SetTargetName(name string) {
	This.TargetName = name
	This.FullTargetDir = filepath.Join(This.TargetDir, This.TargetName)
}

func (This *ProcessData) ParseStringList(data []string) []string {
	for index, item := range data {
		data[index] = This.ParseString(item)
	}

	return data
}

func (This *ProcessData) ParseCopyFileList(data []*CopyFile) []*CopyFile {
	for index, item := range data {
		data[index].From = This.ParseString(item.From)
		data[index].To = This.ParseString(item.To)
	}

	return data
}

func (This *ProcessData) ParseString(data string) string {
	data = strings.Replace(data, fmt.Sprintf("${%sPROJECT_NAME}", constants.ENV_VAR_PREFIX), This.ProjectName, -1)
	data = strings.Replace(data, fmt.Sprintf("${%sDS}", constants.ENV_VAR_PREFIX), This.DirSep, -1)

	data = strings.Replace(data, fmt.Sprintf("${%sPROJECT_ROOT}", constants.ENV_VAR_PREFIX), This.ProjectRootDir, -1)
	data = strings.Replace(data, fmt.Sprintf("${%sTEMP_DIR}", constants.ENV_VAR_PREFIX), This.TempDir, -1)
	data = strings.Replace(data, fmt.Sprintf("${%sTARGET_DIR}", constants.ENV_VAR_PREFIX), This.TargetDir, -1)
	data = strings.Replace(data, fmt.Sprintf("${%sBUILD_DIR}", constants.ENV_VAR_PREFIX), This.BuildDir, -1)
	data = strings.Replace(data, fmt.Sprintf("${%sVENDOR_DIR}", constants.ENV_VAR_PREFIX), This.VendorDir, -1)
	data = strings.Replace(data, fmt.Sprintf("${%sTARGET_NAME}", constants.ENV_VAR_PREFIX), This.TargetName, -1)
	data = strings.Replace(data, fmt.Sprintf("${%sDEPENDENCY_NAME}", constants.ENV_VAR_PREFIX), This.DependencyName, -1)
	data = strings.Replace(data, fmt.Sprintf("${%sFULL_TARGET_DIR}", constants.ENV_VAR_PREFIX), This.FullTargetDir, -1)

	return data
}