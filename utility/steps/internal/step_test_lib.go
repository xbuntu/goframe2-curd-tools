package internal

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// ============================ 开始 ==【不要在意下面这些细节】============================

// StartStep 生成的文件名、项目模块名、是否删除创建
func StartStep(fileName string, moduleName string, isDelStep bool) {
	//根据文件名，获得模板变量名
	varName := wordToCamel(fileName)
	for _, item := range fileAndPathList {
		saveFile := item.path + fileName + ".go"
		//如果是删除文件，则删除已创建的文件，并结束创建
		if isDelStep {
			err := gfile.Remove(saveFile)
			if err != nil {
				fmt.Printf("文件：%s删除失败\n", saveFile)
			} else {
				fmt.Printf("文件：%s删除成功\n", saveFile)
			}
			continue
		}
		//模板变量替换
		saveContent := gfile.GetContents(item.file)
		saveContent = gstr.Replace(saveContent, "#VAR_NAME#", varName)
		saveContent = gstr.Replace(saveContent, "#MOUDLE_NAME#", moduleName)
		saveContent = gstr.Replace(saveContent, "#ROUTE_NAME#", fileName)
		//创建文件
		err := gfile.PutContents(saveFile, saveContent)
		if err != nil {
			fmt.Printf("文件：%s创建失败\n", saveFile)
		} else {
			fmt.Printf("文件：%s创建成功\n", saveFile)
		}
	}
}

// 单词转驼峰字符串
func wordToCamel(word string) string {
	return gstr.TrimAll(gstr.UcWords(gstr.Replace(word, "_", " ")))
}

// 文件和生成路径常量
const (
	// 生成文件路径
	apiV1Path      = "../../api/v1/"
	controllerPath = "../../internal/controller/"
	modelPath      = "../../internal/model/"
	servicePath    = "../../internal/service/"

	// 模板文件路径
	apiV1File      = "internal/api.v1.step"
	controllerFile = "internal/controller.step"
	modelFile      = "internal/model.step"
	serviceFile    = "internal/service.step"
)

// fileAndPathList 模板文件和生成路径
var fileAndPathList = []stepItem{
	{file: apiV1File, path: apiV1Path},
	{file: controllerFile, path: controllerPath},
	{file: modelFile, path: modelPath},
	{file: serviceFile, path: servicePath},
}

// stepItem 生成每一项
type stepItem struct {
	file string
	path string
}

// ============================ 结束 ==【不要在意上面这些细节】============================
