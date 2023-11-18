package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

var FileContents []byte
var IsReadingFromStdin bool
var FileName string
var Arg string

func main() {
  fileInfo, err := os.Stdin.Stat(); if err != nil {
    fmt.Println("Error checking stdin")
    return
  }

  if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
    IsReadingFromStdin = true
    FileContents, _ = ioutil.ReadAll(os.Stdin)
  }

  if(IsReadingFromStdin) {
    if len(os.Args) == 2 {
      Arg = string(os.Args[1])
    }
  } else {
    FileName = os.Args[1]
    FileContents, err = readContentsOfFile(FileName); if err != nil {
      panic(err)
    }
    if len(os.Args) == 3 {
      Arg = string(os.Args[2])
    }
  }

  if Arg == "" {
    fmt.Printf("  %d  %d %d %s\n",  lineCount(FileContents), wordCount(FileContents), byteCount(FileContents), FileName)
    return
  }

  if Arg == "-c" {
    fmt.Println(byteCount(FileContents))
  } else if Arg == "-l" {
    fmt.Println(lineCount(FileContents))
  } else if Arg == "-w" {
    fmt.Println(wordCount(FileContents))
  } else if Arg == "-m" {
    fmt.Println(charCount(FileContents))
  }
}

func readContentsOfFile(f string) ([]byte, error) {
  return ioutil.ReadFile(f)
}

func charCount(file []byte) int {
  return utf8.RuneCount(file)
}

func wordCount(file []byte) int {
  words := strings.Fields(string(file))
  return len(words)
}

func lineCount(file []byte) int {
  return len(strings.Split(string(file), "\n")) - 1
}

func byteCount(file []byte) int {
  return len(file)
}
