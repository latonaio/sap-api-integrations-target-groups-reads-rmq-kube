package sap_api_output_formatter

type TargetGroup struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	APISchema       string `json:"api_schema"`
	TargetGroupCode string `json:"target_group_code"`
	Deleted         bool   `json:"deleted"`
}

type TargetGroupCollection struct {
	ObjectID            string `json:"ObjectID"`
	ETag                string `json:"ETag"`
	ID                  string `json:"ID"`
	LifeCycleStatusCode string `json:"LifeCycleStatusCode"`
	Description         string `json:"Description"`
	MemberNumberValue   int    `json:"MemberNumberValue"`
	SalesOrganizationID string `json:"SalesOrganizationID"`
	EntityLastChangedOn string `json:"EntityLastChangedOn"`
}
