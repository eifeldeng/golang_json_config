package json_config

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"os"
)

type JsonConfig struct {
	Path        string
	ConfigList  map[string]interface{}
	FielContent string
	configSize  int
}

func (jsonC *JsonConfig) GetInt(configName string, node string) (ret int) {

	jsonRet := jsonC.getValue(configName, node, "int")
	ret = jsonRet.(int)
	return ret

}

func (jsonC *JsonConfig) GetString(configName string, node string) (ret string) {
	jsonRet := jsonC.getValue(configName, node, "string")
	ret = jsonRet.(string)
	return ret
}

func (jsonC *JsonConfig) GetMap(configName string, node string) (ret map[string]interface{}) {
	jsonRet := jsonC.getValue(configName, node, "map")
	ret = jsonRet.(map[string]interface{})
	return ret
}

func (jsonC *JsonConfig) GetArray(configName string, node string) (ret []interface{}) {
	jsonRet := jsonC.getValue(configName, node, "array")
	ret = jsonRet.([]interface{})
	return ret

}

func (jsonC *JsonConfig) getValue(configName string, node string, nodeType string) (ret interface{}) {
	var value *simplejson.Json
	var ok bool
	value, ok = jsonC.ConfigList[configName].(*simplejson.Json)
	if ok == false {
		content := jsonC.readFile(configName)
		body := []byte(content)
		resultJson, _ := simplejson.NewJson(body)

		value = resultJson
		if jsonC.configSize == 0 {
			var newConfigList = make(map[string]interface{}, 10)
			newConfigList[configName] = value
			jsonC.ConfigList = newConfigList
		} else {
			jsonC.ConfigList[configName] = value
		}
		jsonC.configSize = jsonC.configSize + 1

	}

	switch nodeType {
	case "int":
		intRet := value.Get(node).MustInt()
		return intRet
	case "string":
		stringRet := value.Get(node).MustString()
		return stringRet
	case "array":
		arrayRet := value.Get(node).MustArray()
		return arrayRet
	case "map":
		mapRet := value.Get(node).MustMap()
		return mapRet
	default:
		stringRet := value.Get(node).MustString()
		return stringRet
	}

}

func (jsonC *JsonConfig) Init(path string) {
	jsonC.Path = path
}

func (jsonC *JsonConfig) readFile(configName string) (content string) {

	var filePath string = jsonC.Path + "/" + configName + ".json"

	fi, err := os.Open(filePath)
	defer fi.Close()
	if err != nil {
		content = "{}"
	} else {
		fd, _ := ioutil.ReadAll(fi)
		content = string(fd)
	}
	return content
}

func main() {
	path := "/Users/eifel/Downloads/"
	var serverConfig JsonConfig
	serverConfig.Init(path)
	port := serverConfig.GetInt("test", "port")
	serverIp := serverConfig.GetString("test", "server_ip")

	config := serverConfig.GetMap("test", "config")
	level := serverConfig.GetArray("test", "level")
	fmt.Println(port)
	fmt.Println(serverIp)
	fmt.Println(config["name"])
	fmt.Println(level[0])
	nodata :=serverConfig.GetInt("test", "man")
	fmt.Println(nodata)

	nofile :=serverConfig.GetInt("test1", "man")
	fmt.Println(nofile)

}
