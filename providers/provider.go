package providers

import (
  "github.com/decentralizedSearch/searchResult"
)

type Provider interface {
  Search(string) ([]searchResult.SearchResult, error)
}
