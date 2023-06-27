package data_block

import (
	"time"
)

type Kv struct {
	KvId           string       `json:"kvId,omitempty"`
	K              string       `json:"k,omitempty"`
	V              interface{}  `json:"v,omitempty"`
	Description    string       `json:"description,omitempty"`
	OrganizationId string       `json:"organizationId,omitempty"`
	BlockStatus    *TEXT_STATUS `json:"blockStatus,omitempty"`
	SyncAt         *time.Time   `json:"syncAt,omitempty"`
	SysId          uint         `json:"sysId,omitempty"`
	CreatedBy      string       `json:"createdBy,omitempty"`
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	UpdatedBy      interface{}  `json:"updatedBy,omitempty"`
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	PublishedBy    interface{}  `json:"publishedBy,omitempty"`
	PublishedAt    interface{}  `json:"publishedAt,omitempty"`
	Organization   Organization `json:"organization,omitempty"`
}
type Organization struct {
	OrganizationName string `json:"organizationName,omitempty"`
}

type Block struct {
	BlockCode       *string                  `json:"blockCode,omitempty"`
	BlockData       []map[string]interface{} `json:"blockData,omitempty"`
	Slugs           string                   `json:"slugs,omitempty"`
	Stage           string                   `json:"stage,omitempty"`
	IsMultipleGroup *string                  `json:"isMultipleGroup,omitempty"`
	AtUsers         []string                 `json:"atUsers,omitempty"`
	SpaceId         string                   `json:"spaceId,omitempty"`
	SpaceName       string                   `json:"spaceName,omitempty"`
	SysId           uint                     `json:"sysId,omitempty"`
	BlockStatus     *TEXT_STATUS             `json:"blockStatus,omitempty"`
	SyncAt          *time.Time               `json:"syncAt,omitempty"`
	CreatedBy       string                   `json:"createdBy,omitempty"`
	CreatedAt       *time.Time               `json:"createdAt,omitempty"`
	UpdatedBy       interface{}              `json:"updatedBy,omitempty"`
	UpdatedAt       *time.Time               `json:"updatedAt,omitempty"`
	PublishedBy     interface{}              `json:"publishedBy,omitempty"`
	PublishedAt     interface{}              `json:"publishedAt,omitempty"`
	ModelCode       string                   `json:"modelCode,omitempty"`
}

type BlockData struct {
	ID           string                   `json:"id,omitempty"`
	Cid          string                   `json:"cid,omitempty"`
	Data         []map[string]interface{} `json:"data,omitempty"`
	GroupName    string                   `json:"groupName,omitempty"`
	GroupPercent string                   `json:"groupPercent,omitempty"`
}
