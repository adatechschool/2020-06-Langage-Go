package main

import (
  "fmt"
  "os"
  rd "encoding/csv"
)

func word_chooser() string {
      file, err := os.Open("words_csv/dico.csv")
      error_catch(err)

      var short, medium, long []string
      fmt.Println(short, medium, long)

      reader := rd.NewReader(file)
      words, err := reader.ReadAll()
      error_catch(err)

      for i := 0; i < len(words); i++ {
        length := len(words[i][0])
        switch {
        case length <= 5:
          short = append(short, words[i][0])
        case length < 8 && length > 5:
          medium = append(medium, words[i][0])
        case length >= 8:
          long = append(long, words[i][0])
        }
      }

      // fmt.Printf(short, medium, long)


  return "placeholder"
}

func error_catch(err error) {
  if err != nil {
    panic(err)
    fmt.Println("An error has occured: ", err)
  }
}
