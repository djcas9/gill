package gill

import (
	"fmt"
	//     "log"
	"net/url"
	//     "path"
	"strings"
)

func FetchRepo(config_path string, u *url.URL, d Config) {
	fmt.Println(u.Scheme, u.Opaque, u.User, u.Host, u.Path, u.RawQuery, u.Fragment)

	var username string

	if len(u.Scheme) > 0 {

		switch u.Scheme {
		case "git":
			var items []string = strings.Split(u.Path, "/")
			u.Path = items[len(items)-1]
			username = strings.Join(items[:len(items)-1], "/")
		case "http", "https":
			if len(u.Host) > 0 {
				if strings.Contains(u.Host, ":") {
					var items []string = strings.Split(u.Host, ":")
					u.Host = items[0]
					username = items[1]
				} else {
					var items []string = strings.Split(u.Path, "/")
					fmt.Println("BEFORE:", items)

					u.Path = items[len(items)-1]

					for i := 0; i < len(items); i++ {
						if len(items[i]) <= 0 {
							items = append(items[:i], items[i+1:]...)
							i--
						}
					}

					username = items[0]
					fmt.Println("AFTER", items)

				}
			}
		}

	} else {

		if len(u.Path) > 0 && strings.Contains(u.Path, ":") {
			var split []string = strings.Split(u.Path, ":")
			var items []string = strings.Split(split[0], "@")
			u.Host = items[1]

			var userPath []string = strings.Split(split[1], "/")
			username = userPath[0]
			u.Path = userPath[1]
		}

	}

	if u.Host == "github.com" {
		// fetch json repo information
		fmt.Println("fetch github information")
	}

	fmt.Println(u.Scheme, u.Opaque, u.User, username, u.Host, u.Path, u.RawQuery, u.Fragment)
}

func RemoveRepo(config_path string, d Config) {
	fmt.Println(d)
	fmt.Println("REMOVE")
}
