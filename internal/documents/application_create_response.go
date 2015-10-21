package documents

type ApplicationCreateResponse struct {
	Metadata struct {
		GUID string
	}
	Entity struct {
		Name      string
		SpaceGUID string `json:"space_guid"`
		Diego     bool
	}
}
