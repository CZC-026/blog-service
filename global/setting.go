package global

import "github.com/go-programming-tour-book/blog-service/pkg/setting"

//在读取了文件的配置信息后，还是不够的，因为我们需要将配置信息和应用程序关联起来，我们才能够去使用
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)
