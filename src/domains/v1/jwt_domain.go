package domains

type Jwt struct {
	UID int `json:"uid"`
	Exp int64 `json:"exp"`
}
