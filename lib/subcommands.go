package gill

import (
  "github.com/codegangsta/cli"
)

func AddCommand() cli.Command {
  return cli.Command{
    Name:   "add",
    Usage:  "Add a git repository to your gill store.",
    Action: AddAction,
  }
}

func RemoveCommand() cli.Command {
  return cli.Command{
    Name:   "remove",
    Usage:  "Remove a git repository from your gill store.",
    Action: RemoveAction,
  }
}

func ListCommand() cli.Command {
  return cli.Command{
    Name:   "list",
    Usage:  "List all git repositories in your gill store.",
    Action: ListAction,
  }
}

func ConfigCommand() cli.Command {
  return cli.Command{
    Name:   "config",
    Usage:  "Set gill configuration options. Example: gill config source ~/Source",
    Action: ConfigAction,
  }
}
