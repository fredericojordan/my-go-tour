package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    var word_count = make(map[string]int)

    for _, word := range strings.Fields(s) {
        if count, ok := word_count[word]; ok {
            word_count[word] = count+1
        } else {
            word_count[word] = 1
        }
    }

    return word_count
}

func main() {
    wc.Test(WordCount)
}
