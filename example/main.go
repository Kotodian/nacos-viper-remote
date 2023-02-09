package main

import (
	"fmt"

	"github.com/spf13/viper"

	remote "github.com/yoyofxteam/nacos-viper-remote"
)

func main() {

	err := remote.SetOptions(&remote.Option{
		Url:         "10.43.0.12",
		Port:        8848,
		NamespaceId: "77b74411-f615-479e-a24e-6e848bc19121",
		GroupName:   "JXEU-CSMS",
		Config:      remote.Config{DataId: "jxeu-acos.properties"},
		Auth:        &remote.Auth{Enable: true, User: "csms", Password: "FsUGjIdnXcNflJdk"},
	})
	if err != nil {
		panic(err)
	}
	remoteViper := viper.New()
	err = remoteViper.AddRemoteProvider("nacos", "http://10.43.0.12:8848", "")
	if err != nil {
		panic(err)
	}
	//localSetting := runtime_viper.AllSettings()
	remoteViper.SetConfigType("properties")
	err = remoteViper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
	i := remoteViper.AllKeys()
	fmt.Println(i)

}
