package model

import "crawler/engine"

//SearchResult SearchResult返回的结果集
type SearchResult struct {
	Hits int
	Start int
	Items []engine.Items
}