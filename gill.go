package main

import (
	"github.com/codegangsta/cli"
	"github.com/mephux/gill/lib"
	"github.com/sbinet/go-config/config"
	"log"
	"os"
	"os/user"
	"path"
)

func main() {
	app := cli.NewApp()

	app.Name = "gill"
	app.Usage = "Git, Clone, Cleanliness."
	app.Version = gill.GILL_VERSION
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

	if _, err := os.Stat(config_path); err != nil {
		c := config.NewDefault()
		c.AddSection("Paths")
		c.AddSection("Repos")
		c.AddOption("Paths", "source", path.Join(usr.HomeDir, "Source"))
		c.WriteFile(config_path, 0644, "Gill - Git, Clone, Cleanliness.")
	}

	app.Run(os.Args)
}
