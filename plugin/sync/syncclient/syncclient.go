package syncclient
import (
	"github.com/wjpxxx/letgo/file"
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/plugin/sync/syncconfig"
	"fmt"
)

//config
var config syncconfig.ClientConfig
//init
func init() {
	clientFile:="config/sync_client.config"
	cfgFile:=file.GetContent(clientFile)
	if cfgFile==""{
		var paths []syncconfig.PathList
		paths=append(paths,syncconfig.PathList{
			LocationPath: "./",
			RemotePath: "./",
			Filter:[]string{},
		})
		clientConfig:=syncconfig.ClientConfig{
			Paths:paths,
			Server: syncconfig.SyncServer{
				Server: syncconfig.Server{
					IP: "127.0.0.1",
					Port: "5566",
				},
				Slave: []syncconfig.Server{},
			},
		}
		file.PutContent(clientFile,fmt.Sprintf("%v",clientConfig))
		panic("please setting sync client config in config/sync_client.config file")
	}
	if !lib.StringToObject(cfgFile, &config) {
		panic("config/sync_client.config file format error, Please check carefully")
	}
}