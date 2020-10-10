package gen

import (
  "fmt"
  "io/ioutil"
  "log"
  "regexp"
  "strings"

  "github.com/leonkozlowski/yaig/cmd/yaig/conf"
)

// Entry entry for index generation
type Entry struct {
  document, index interface{}
}

func generateInvertedIndex() {
  raw, err := ioutil.ReadFile("road.txt")
  if err != nil {
    log.Fatal(err)
    return
  }

  reg, err := regexp.Compile("[^a-zA-Z0-9]+")
  if err != nil {
    log.Fatal(err)
	  return
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
}
