package steps_multi

import (
	"curltools/utility/steps_multi/internal"
	"github.com/gogf/gf/v2/text/gstr"
	"testing"
)

// TestStep 测试生成模板到指定目录，可修改 *.step文件，进行自定义生成文件规范
func TestStep(t *testing.T) {
	//根据数据表构建，多应用CURD
	fileNames := gstr.Split("article", ",")
	for _, fileName := range fileNames {
		//admin应用  curltools替换自己项目的模块名，isDelStep为true，则删除已生成的文件
		internal.StartStep(fileName, "curltools","admin", false)
		//api应用
		internal.StartStep(fileName, "curltools","api", false)
	}

}
