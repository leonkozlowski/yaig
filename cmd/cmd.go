package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
  "github.com/leonkozlowski/yaig/version"
)

// RootCmd CLI Command
var RootCmd = &cobra.Command{
  Use: "yaig",
  Short: "yaig is a CLI-based inveted index generator",
  Version: version.Version,
}

func must(err error) {
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
}
