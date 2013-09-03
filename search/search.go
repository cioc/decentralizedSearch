package search

import (
  "github.com/cioc/decentralizedSearch/providers"
  "github.com/cioc/decentralizedSearch/searchResult"

  "fmt"
)

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
