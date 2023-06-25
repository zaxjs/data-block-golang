package data_block

type (
	BLOCK_TYPE  string
	TEXT_STATUS string
)

const (
	BT_KV    BLOCK_TYPE = "kv"
	BT_BLOCK BLOCK_TYPE = "block"
)

const (
	TS_PUBLISHED TEXT_STATUS = "PUBLISHED"
	TS_DELETED   TEXT_STATUS = "DELETED"
)
