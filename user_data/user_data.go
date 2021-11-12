package user_data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/kirsle/configdir"
)

var ConfigPath = configdir.LocalConfig("umw-class-monitor", "data")

func SetupConfigPath() {
	configdir.MakePath(ConfigPath)
}

func ReadConfigData() *UserData {
	SetupConfigPath()

	var userData UserData

	file, err := os.Open(path.Join(ConfigPath, "config.json"))
	defer file.Close()
	if err != nil {
		return &userData
	}

	fileBytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(fileBytes, &userData)

	return &userData
}

func UpdateConfigData(userData *UserData) error{
	file, err := os.Create(path.Join(ConfigPath, "config.json"))
	defer file.Close()
	if err != nil {
		return err
	}

	fileBytes, err := json.Marshal(userData)
	file.Write(fileBytes)
	
	return nil
}