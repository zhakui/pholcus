package web

import (
	"net/http"
	"text/template"

	"github.com/henrylee2cn/pholcus/config"
	"github.com/henrylee2cn/pholcus/logs"
	"github.com/henrylee2cn/pholcus/runtime/status"
)

// 处理web页面请求
func pholcus(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("web/views/index.html") //解析模板文件
	if err != nil {
		logs.Log.Error("%v", err)
	}
	//获取pholcus信息
	data := map[string]interface{}{
		"title":   config.APP_NAME,
		"logo":    config.ICON_PNG,
		"version": config.APP_VERSION,
		"author":  config.APP_AUTHOR,
		"mode": map[string]int{
			"offline": status.OFFLINE,
			"server":  status.SERVER,
			"client":  status.CLIENT,
			"unset":   status.UNSET,
			"curr":    logicApp.GetAppConf("mode").(int),
		},
		"status": map[string]int{
			"unknow": status.UNKNOW,
			"stop":   status.STOP,
			"run":    status.RUN,
			"pause":  status.PAUSE,
		},
		"port": logicApp.GetAppConf("port").(int),
		"ip":   logicApp.GetAppConf("master").(string),
	}
	t.Execute(rw, data) //执行模板的merger操作
}
