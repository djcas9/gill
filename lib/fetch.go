package gill

import (
  "fmt"
  "net/url"
  "github.com/sbinet/go-config/config"
  "regexp"
  "strings"
  "path"
  "log"
)

func FetchRepo(config_path string, u *url.URL, d *config.Config) {

  reg, err := regexp.Compile("(.git)$")
  split := strings.Split(u.Path, "/")

  if err != nil {
    log.Fatal(err)
  }

  key := reg.ReplaceAllString(split[len(split) - 1], "")

  if len(u.Fragment) > 0 {
    key = path.Join(u.Fragment, key)
  }

  source_path, _ := d.String("Paths", "source")

  fmt.Println("key:", key, source_path)

  fmt.Println("Word, In Fetch", key, d)
  
  if !d.HasSection("Repos") {
    d.AddSection("Repos")
  }

  d.AddOption("Repos", key, path.Join(source_path, key))
  d.WriteFile(config_path, 0644, "Gill - Git, Clone, Cleanliness.")
}

func RemoveRepo(config_path string, d *config.Config) {
  fmt.Println("REMOVE")
}
