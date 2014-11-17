package rainmaker

import (
	"time"

	"github.com/pivotal-golang/rainmaker/internal/documents"
)

type Space struct {
	GUID                     string
	URL                      string
	CreatedAt                time.Time
	UpdatedAt                time.Time
	Name                     string
	OrganizationGUID         string
	SpaceQuotaDefinitionGUID string
	OrganizationURL          string
	DevelopersURL            string
	ManagersURL              string
	AuditorsURL              string
	AppsURL                  string
	RoutesURL                string
	DomainsURL               string
	ServiceInstancesURL      string
	AppEventsURL             string
	EventsURL                string
	SecurityGroupsURL        string
}

func NewSpaceFromResponse(response documents.SpaceResponse) Space {
	if response.Metadata.CreatedAt == nil {
		response.Metadata.CreatedAt = &time.Time{}
	}

	if response.Metadata.UpdatedAt == nil {
		response.Metadata.UpdatedAt = &time.Time{}
	}

	return Space{
		GUID:                     response.Metadata.GUID,
		URL:                      response.Metadata.URL,
		CreatedAt:                *response.Metadata.CreatedAt,
		UpdatedAt:                *response.Metadata.UpdatedAt,
		Name:                     response.Entity.Name,
		OrganizationGUID:         response.Entity.OrganizationGUID,
		SpaceQuotaDefinitionGUID: response.Entity.SpaceQuotaDefinitionGUID,
		OrganizationURL:          response.Entity.OrganizationURL,
		DevelopersURL:            response.Entity.DevelopersURL,
		ManagersURL:              response.Entity.ManagersURL,
		AuditorsURL:              response.Entity.AuditorsURL,
		AppsURL:                  response.Entity.AppsURL,
		RoutesURL:                response.Entity.RoutesURL,
		DomainsURL:               response.Entity.DomainsURL,
		ServiceInstancesURL:      response.Entity.ServiceInstancesURL,
		AppEventsURL:             response.Entity.AppEventsURL,
		EventsURL:                response.Entity.EventsURL,
		SecurityGroupsURL:        response.Entity.SecurityGroupsURL,
	}
}
