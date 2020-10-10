package cmd

import (
  "fmt"
  "io/ioutil"
  "log"
  "regexp"
  "strings"

  "github.com/spf13/cobra"

  "github.com/leonkozlowski/yaig/cmd/yaig/conf"
)

var filePath string
var indexCmd = &cobra.Command{
  Use:   "index [filename]",
  Short: "Generate an inverted index for .txt file input",
  Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    fp := args[0]
    generateInvertedIndex(fp)
  },
}

func init() {
  indexCmd.Flags().StringVarP(&filePath, "filepath", "f", "./", "Input filepath")
  RootCmd.AddCommand(indexCmd)
}

// Entry entry for index generation
type Entry struct {
  document, index interface{}
}

func generateInvertedIndex(filepath string) map[string][]Entry {
  raw, err := ioutil.ReadFile(filepath)
  if err != nil {
    log.Fatal(err)
  }

  reg, err := regexp.Compile("[^a-zA-Z0-9]+")
  if err != nil {
    log.Fatal(err)
  }

  cleaned := reg.ReplaceAllString(string(raw), " ")
  stripped := strings.TrimRight(cleaned, " ")
  tokens := strings.Split(string(stripped), " ")

  stopWords := conf.StopWords()
  master := map[string][]Entry{}
  for index, element := range tokens {
    isin := false
    lower := strings.ToLower(element)
    if stopWords[lower] {
      isin = true
    }
    if isin == false {
      lower := strings.ToLower(element)
      value, ok := master[lower]
      if ok {
        existing := value
        updated := append(existing, Entry{1, index})
        master[lower] = updated
      } else {
        master[lower] = append(master[lower], Entry{1, index})
      }
    }
  }
  fmt.Println(master)
  return master
}
