package gill

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"net/url"
	//     "os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
)

type Repo struct {
	Name     string   `json:"name"`
	Path     string   `json:"path"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

type Config struct {
	Repos  []Repo `json:"repos"`
	Source string `json:"source"`
}

// hasExpectedArgs checks whether the number of args are as expected.
func hasArgs(args []string, expected int) bool {
	switch expected {
	case -1:
		if len(args) > 0 {
			return true
		} else {
			return false
		}
	default:
		if len(args) == expected {
			return true
		} else {
			return false
		}
	}
}

func getConfigPath() string {
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	return path.Join(usr.HomeDir, CONFIG_FILENAME)
}

func getConfig(config_path string) (Config, error) {
	file, err := ioutil.ReadFile(config_path)

	if err != nil {
		panic(err)
	}

	var config Config
	errr := json.Unmarshal(file, &config)

	return config, errr
}

// Edit Configuration
func ConfigAction(c *cli.Context) {
	if !hasArgs(c.Args(), 2) {
		log.Fatal("Incorrect number of arguments to 'config' command")
	}

	config_path := getConfigPath()
	d, _ := getConfig(config_path)

	fmt.Println("WORD:", config_path, d)

	key := c.Args()[0]
	value := c.Args()[1]

	// TODO - Switch for different config key types
	if !filepath.IsAbs(value) {
		abs_path, err := filepath.Abs(value)
		if err != nil {
			log.Fatal("Incorrect config value.")
		}
		value = abs_path
	}

	fmt.Println(key, value)

	//     d.AddOption("Paths", key, value)
	//     d.WriteFile(config_path, 0644, "Gill - Git, Clone, Cleanliness.")
}

// Reove
func RemoveAction(c *cli.Context) {
	config_path := getConfigPath()
	d, _ := getConfig(config_path)

	RemoveRepo(config_path, d)
}

// List
func ListAction(c *cli.Context) {
	config_path := getConfigPath()
	d, _ := getConfig(config_path)

	fmt.Println(d)

	//     if d.HasSection("Repos") {
	//         repos, _ := d.Options("Repos")

	//         for i := 0; i < len(repos); i++ {
	//             fmt.Println(repos[i])
	//         }

	//     } else {
	//         d.AddSection("Repos")
	//         fmt.Println("No repositories have been added to your gill store.")
	//         os.Exit(1)
	//     }

}

// Add
func AddAction(c *cli.Context) {
	if !hasArgs(c.Args(), 1) {
		log.Fatal("Incorrect number of arguments to 'add' command")
	}

	gill_url, err := url.Parse(c.Args()[0])

	if err != nil {
		log.Fatal("Unable to parse argument for 'add' command")
	}

	if gill_url.Scheme == "http" || gill_url.Scheme == "https" || gill_url.Scheme == "git" {

		match, _ := regexp.MatchString(".git", gill_url.Path)
		fmt.Println(match)

		if !match {
			log.Fatal("Unable to parse argument for 'add' command")
		}

		fmt.Println(gill_url.Host, gill_url.Path, err)

		config_path := getConfigPath()
		d, _ := getConfig(config_path)

		FetchRepo(config_path, gill_url, d)

	} else {
		log.Fatalf("Unable to parse argument for 'add' command - Scheme '%s' not supported.", gill_url.Scheme)
	}

}
