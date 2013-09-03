package main

import (
  "fmt"
  "log"

  "github.com/cioc/decentralizedSearch/providers"
  "github.com/cioc/decentralizedSearch/search"
  "github.com/cioc/decentralizedSearch/providers/stackoverflow"
  "github.com/cioc/decentralizedSearch/providers/wikipedia"
  "github.com/cioc/decentralizedSearch/searchResult"

  "net/http"
  "net/url"

  "encoding/json"
)

const handleSearchLen = len("/api/search/")

var myProviders []providers.Provider

func handleIndex(w http.ResponseWriter, r *http.Request) {

}

func handleSearch(w http.ResponseWriter, r *http.Request) {
  query, err := url.QueryUnescape(r.URL.Path[handleSearchLen:])
  if err != nil {
    log.Fatal(err)
  }
  results := search.Search(query, myProviders)
  totalResults := 0
  for _, r := range results {
    totalResults +=  len(r)
  }
  o := make([]searchResult.SearchResult, 0, totalResults)
  for _,r := range results {
    o = append(o, r...)
  }
  b, err := json.Marshal(o)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Fprintf(w, string(b))
}

func main() {
  //setup all search providers
  so := stackoverflow.New()
  wikipedia := wikipedia.New()
  myProviders = []providers.Provider{so, wikipedia}

  //setup the web server
  http.Handle("/", http.FileServer(http.Dir("./content")))
  http.HandleFunc("/api/search/", handleSearch)
  log.Fatal(http.ListenAndServe(":8080", nil))
}


