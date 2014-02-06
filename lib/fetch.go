package gill

import (
	"fmt"
	//     "log"
	"net/url"
	//     "path"
	//     "regexp"
	//     "strings"
)

func FetchRepo(config_path string, u *url.URL, d Config) {
	fmt.Println(d)
	fmt.Println("ADD")
}

func RemoveRepo(config_path string, d Config) {
	fmt.Println(d)
	fmt.Println("REMOVE")
}
