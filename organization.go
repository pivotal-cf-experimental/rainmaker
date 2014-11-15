package rainmaker

import (
	"time"

	"github.com/pivotal-golang/rainmaker/internal/documents"
)

type Organization struct {
	GUID                     string
	Name                     string
	URL                      string
	BillingEnabled           bool
	Status                   string
	QuotaDefinitionGUID      string
	QuotaDefinitionURL       string
	SpacesURL                string
	DomainsURL               string
	PrivateDomainsURL        string
	UsersURL                 string
	ManagersURL              string
	BillingManagersURL       string
	AuditorsURL              string
	AppEventsURL             string
	SpaceQuotaDefinitionsURL string
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

func NewOrganizationFromResponse(response documents.OrganizationResponse) Organization {
	return Organization{
		GUID:                     response.Metadata.GUID,
		URL:                      response.Metadata.URL,
		CreatedAt:                response.Metadata.CreatedAt,
		UpdatedAt:                response.Metadata.UpdatedAt,
		Name:                     response.Entity.Name,
		BillingEnabled:           response.Entity.BillingEnabled,
		Status:                   response.Entity.Status,
		QuotaDefinitionGUID:      response.Entity.QuotaDefinitionGUID,
		QuotaDefinitionURL:       response.Entity.QuotaDefinitionURL,
		SpacesURL:                response.Entity.SpacesURL,
		DomainsURL:               response.Entity.DomainsURL,
		PrivateDomainsURL:        response.Entity.PrivateDomainsURL,
		UsersURL:                 response.Entity.UsersURL,
		ManagersURL:              response.Entity.ManagersURL,
		BillingManagersURL:       response.Entity.BillingManagersURL,
		AuditorsURL:              response.Entity.AuditorsURL,
		AppEventsURL:             response.Entity.AppEventsURL,
		SpaceQuotaDefinitionsURL: response.Entity.SpaceQuotaDefinitionsURL,
	}
}
