package setting

import "github.com/spf13/viper" //适用于go的完整配置解决方案

//把配置的基础属性安排好，初始化，配置信息能启动，但是配置信息和应用程序还没有关联起来
type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) { //初始化本项目配置基本属性
	//设置项目配置文件名称为config，配置类型为yaml，设置配置路径为相对路径configs/
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
