package cmd

import (
	"encoding/json"
	"errors"
	"log"
	"path/filepath"
	"reflect"

	"github.com/spf13/cobra"

	"github.com/gtsteffaniak/filebrowser/auth"
	"github.com/gtsteffaniak/filebrowser/settings"
)

type settingsFile struct {
	Settings *settings.Settings `json:"settings"`
	Server   *settings.Server   `json:"server"`
	Auther   interface{}        `json:"auther"`
}

var configImportCmd = &cobra.Command{
	Use:   "import <path>",
	Short: "Import a configuration file",
	Long: `Import a configuration file. This will replace all the existing
configuration. Can be used with or without unexisting databases.

If used with a nonexisting database, a key will be generated
automatically. Otherwise the key will be kept the same as in the
database.

The path must be for a json or yaml file.`,
	Args: jsonYamlArg,
	Run: python(func(cmd *cobra.Command, args []string, d pythonData) {
		var key []byte
		if d.hadDB {
			settings, err := d.store.Settings.Get()
			checkErr(err)
			key = settings.Key
		} else {
			key = generateKey()
		}
		file := settingsFile{}
		err := unmarshal(args[0], &file)
		checkErr(err)
		log.Println(file.Settings)
		file.Settings.Key = key
		err = d.store.Settings.Save(file.Settings)
		checkErr(err)

		err = d.store.Settings.SaveServer(file.Server)
		checkErr(err)

		var rawAuther interface{}
		if filepath.Ext(args[0]) != ".json" { //nolint:goconst
			rawAuther = cleanUpInterfaceMap(file.Auther.(map[interface{}]interface{}))
		} else {
			rawAuther = file.Auther
		}
		log.Println("config_import", file.Settings.Auth)
		var auther auth.Auther
		switch file.Settings.Auth.Method {
		case "password":
			auther = getAuther(auth.JSONAuth{}, rawAuther).(*auth.JSONAuth)
		case "noauth":
			auther = getAuther(auth.NoAuth{}, rawAuther).(*auth.NoAuth)
		case "proxy":
			auther = getAuther(auth.ProxyAuth{}, rawAuther).(*auth.ProxyAuth)
		case "hook":
			auther = getAuther(&auth.HookAuth{}, rawAuther).(*auth.HookAuth)
		default:
			checkErr(errors.New("invalid auth method"))
		}

		err = d.store.Auth.Save(auther)
		checkErr(err)

	}, pythonConfig{allowNoDB: true}),
}

func getAuther(sample auth.Auther, data interface{}) interface{} {
	authType := reflect.TypeOf(sample)
	auther := reflect.New(authType).Interface()
	bytes, err := json.Marshal(data)
	checkErr(err)
	err = json.Unmarshal(bytes, &auther)
	checkErr(err)
	return auther
}
