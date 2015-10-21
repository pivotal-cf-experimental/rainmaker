package documents

type ApplicationSummaryResponse struct {
	GUID      string
	Name      string
	SpaceGUID string `json:"space_guid"`
}
