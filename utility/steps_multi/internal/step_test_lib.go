package internal

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// ============================ 开始 ==【不要在意下面这些细节】============================

// StartStep 生成的文件名、项目模块名、是否删除创建
func StartStep(fileName string, moduleName string, appName string,isDelStep bool) {
	//根据文件名，获得模板变量名
	varName := wordToCamel(fileName)
	for stepDir, stepFile := range stepsMap {
		saveFile := fmt.Sprintf("../../internal/app/%s/%s/%s.go",appName,stepDir,fileName);
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
		saveContent := gfile.GetContents(stepFile)
		saveContent = gstr.Replace(saveContent, "#VAR_NAME#", varName)
		saveContent = gstr.Replace(saveContent, "#APP_NAME#", appName)
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

// fileAndPathList 模板文件和生成路径
var stepsMap = map[string]string{
	"router":"internal/router.step",
	"controller":"internal/controller.step",
	"model":"internal/model.step",
	"service":"internal/service.step",
}
// ============================ 结束 ==【不要在意上面这些细节】============================
