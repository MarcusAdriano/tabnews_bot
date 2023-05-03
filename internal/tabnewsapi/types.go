package tabnewsapi

import (
	"fmt"
	"time"
)

const (
	StrategyNew      = "new"
	StrategyOld      = "old"
	StrategyRelevant = "relevant"
)

type ContentsConfig struct {
	Page     int
	PerPage  int
	Strategy string
}

type Content struct {
	Id                string    `json:"id"`
	OwnerId           string    `json:"owner_id"`
	ParentId          string    `json:"parent_id,omitempty"`
	Slug              string    `json:"slug"`
	Title             string    `json:"title"`
	Status            string    `json:"status"`
	SourceUrl         string    `json:"source_url,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
	PublishedAt       time.Time `json:"published_at"`
	DeletedAt         time.Time `json:"deleted_at,omitempty"`
	Tabcoins          int       `json:"tabcoins"`
	OwnerUsername     string    `json:"owner_username"`
	ChildrenDeepCount int       `json:"children_deep_count"`
}

func (c *Content) Link(baseUrl string) string {
	return fmt.Sprintf("%s/%s/%s", baseUrl, c.OwnerUsername, c.Slug)
}
