package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/mephux/gill/lib"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
)

func main() {
	app := cli.NewApp()

	app.Name = "gill"
	app.Usage = "Git, Clone, Cleanliness."
	app.Version = gill.VERSION
	app.Commands = []cli.Command{
		gill.AddCommand(),
		gill.RemoveCommand(),
		gill.ListCommand(),
		gill.ConfigCommand(),
	}

	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	config_path := path.Join(usr.HomeDir, gill.CONFIG_FILENAME)

	var default_json string

	if _, err := os.Stat(config_path); err != nil {
		// setup
		default_json = fmt.Sprintf(`{ "source": "%s", "repos": [] }`, path.Join(usr.HomeDir))
		fmt.Println(default_json)

		err = ioutil.WriteFile(config_path, []byte(default_json), 0644)
		if err != nil {
			panic(err)
		}
	}

	app.Run(os.Args)
}
