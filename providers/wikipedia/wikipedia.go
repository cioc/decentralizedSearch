package wikipedia
//wikipedia
//http://en.wikipedia.org/w/api.php?action=query&list=search&format=json&srsearch=matrix
//go to http://en.wikipedia.org/wiki/<term>
import (
  "github.com/decentralizedSearch/searchResult"

  "net/http"
  "net/url"

  "io/ioutil"

  "fmt"

  "strings"

  "encoding/json"
)

type Wikipedia struct {
  endpoint string
  linkFormat string
}

func New() (*Wikipedia) {
  endpoint := "http://en.wikipedia.org/w/api.php?action=query&list=search&format=json&srsearch=%v"
  linkFormat := "http://en.wikipedia.org/wiki/%v"
  return &Wikipedia{endpoint, linkFormat}
}

type wikipediaResult struct {
  Query struct {
    Search []struct {
      Title string
      Snippet string
    }
  }
}

func (s *Wikipedia) Search(q string) ([]searchResult.SearchResult, error) {
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
  var results wikipediaResult
  err = json.Unmarshal(body, &results)
  if err != nil {
    return nil, err
  }
  o := make([]searchResult.SearchResult, 0, len(results.Query.Search))
  for _, v := range results.Query.Search {
    encodedTitle := strings.Replace(v.Title, " ", "_", -1)
    link := fmt.Sprintf(s.linkFormat, encodedTitle)
    o = append(o, searchResult.SearchResult{"wikipedia", "Wikipedia", v.Title, link, v.Snippet, []string{}, 0.0})
  }
  return o, nil
}
