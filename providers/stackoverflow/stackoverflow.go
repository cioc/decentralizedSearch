package stackoverflow

import (
  "github.com/cioc/decentralizedSearch/searchResult"

  "net/http"
  "net/url"

  "io/ioutil"

  "fmt"

  "encoding/json"
)

type StackOverflow struct {
  endpoint string
}

func New() (*StackOverflow) {
  endpoint := "https://api.stackexchange.com/2.1/search/advanced?order=desc&sort=votes&q=%v&site=stackoverflow"
  return &StackOverflow{endpoint}
}

type soResult struct {
  Score float32
  Title string
  Tags []string
  Link string
}

type soResultArray struct {
  Items []soResult
}

func (s *StackOverflow) Search(q string) ([]searchResult.SearchResult, error) {
  queryStr := url.QueryEscape(q)
  path := fmt.Sprintf(s.endpoint, queryStr)
  resp, err := http.Get(path)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }
  var results soResultArray
  err = json.Unmarshal(body, &results)
  if err != nil {
    return nil, err
  }
  o := make([]searchResult.SearchResult, 0, len(results.Items))
  for _, r := range results.Items {
    o = append(o, searchResult.SearchResult{"so", "Stack Overflow", r.Title, r.Link, "", r.Tags, r.Score})
  }
  return o, nil
}
