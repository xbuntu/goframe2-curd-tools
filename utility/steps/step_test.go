package steps

import (
	"curdtools/utility/steps/internal"
	"testing"
)

// TestStep 测试生成模板到指定目录，可修改 *.step文件，进行自定义生成文件规范
func TestStep(t *testing.T) {
	//循环构建
	fileNames := []string{
		"article",
	}
	for _, fileName := range fileNames {
		//curdtools替换自己项目的模块名，isDelStep为true，则删除已生成的文件
		internal.StartStep(fileName, "curdtools", false)
	}

}
