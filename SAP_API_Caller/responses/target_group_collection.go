package responses

type TargetGroupCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID            string `json:"ObjectID"`
			ETag                string `json:"ETag"`
			ID                  string `json:"ID"`
			LifeCycleStatusCode string `json:"LifeCycleStatusCode"`
			Description         string `json:"Description"`
			MemberNumberValue   int    `json:"MemberNumberValue"`
			SalesOrganizationID string `json:"SalesOrganizationID"`
			EntityLastChangedOn string `json:"EntityLastChangedOn"`
			TargetGroupMember   struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"TargetGroupMember"`
			TargetGroupNotes struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"TargetGroupNotes"`
		} `json:"results"`
	} `json:"d"`
}
