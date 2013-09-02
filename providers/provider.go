package providers

import (
  "github.com/cioc/decentralizedSearch/searchResult"
)

type Provider interface {
  Search(string) ([]searchResult.SearchResult, error)
}
