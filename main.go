package main

import (
  "fmt"

  "github.com/cioc/decentralizedSearch/providers"
  "github.com/cioc/decentralizedSearch/searchResult"
//  "github.com/cioc/decentralizedSearch/providers/stackoverflow"
  "github.com/cioc/decentralizedSearch/providers/wikipedia"
)

//github
//http://developer.github.com/v3/search/

func Search(query string, ps []providers.Provider) ([][]searchResult.SearchResult) {
  results := make([][]searchResult.SearchResult, 0, len(ps))
  for _,p :=  range ps {
    r, err := p.Search(query)
    if err != nil {
      fmt.Println(err)
      continue
    }
    results = append(results, r)
  }
  return results
}

func main() {
//  so := stackoverflow.New()
  wikipedia := wikipedia.New()
  providers := []providers.Provider{wikipedia}
  results := Search("mergesort", providers)
  for i := range results {
    for _, r := range results[i] {
      fmt.Println(r)
    }
  }
}
