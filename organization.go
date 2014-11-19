package rainmaker

import (
	"encoding/json"
	"time"

	"github.com/pivotal-golang/rainmaker/internal/documents"
)

type Organization struct {
	config                   Config
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

func NewOrganization(config Config) Organization {
	return Organization{
		config: config,
	}
}

func NewOrganizationFromResponse(config Config, response documents.OrganizationResponse) Organization {
	if response.Metadata.CreatedAt == nil {
		response.Metadata.CreatedAt = &time.Time{}
	}

	if response.Metadata.UpdatedAt == nil {
		response.Metadata.UpdatedAt = &time.Time{}
	}

	organization := NewOrganization(config)
	organization.GUID = response.Metadata.GUID
	organization.URL = response.Metadata.URL
	organization.CreatedAt = *response.Metadata.CreatedAt
	organization.UpdatedAt = *response.Metadata.UpdatedAt
	organization.Name = response.Entity.Name
	organization.BillingEnabled = response.Entity.BillingEnabled
	organization.Status = response.Entity.Status
	organization.QuotaDefinitionGUID = response.Entity.QuotaDefinitionGUID
	organization.QuotaDefinitionURL = response.Entity.QuotaDefinitionURL
	organization.SpacesURL = response.Entity.SpacesURL
	organization.DomainsURL = response.Entity.DomainsURL
	organization.PrivateDomainsURL = response.Entity.PrivateDomainsURL
	organization.UsersURL = response.Entity.UsersURL
	organization.ManagersURL = response.Entity.ManagersURL
	organization.BillingManagersURL = response.Entity.BillingManagersURL
	organization.AuditorsURL = response.Entity.AuditorsURL
	organization.AppEventsURL = response.Entity.AppEventsURL
	organization.SpaceQuotaDefinitionsURL = response.Entity.SpaceQuotaDefinitionsURL

	return organization
}

func FetchOrganization(config Config, path, token string) (Organization, error) {
	_, body, err := NewClient(config).makeRequest(requestArguments{
		Method: "GET",
		Path:   path,
		Token:  token,
	})
	if err != nil {
		return Organization{}, err
	}

	var response documents.OrganizationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Organization{}, err
	}

	return NewOrganizationFromResponse(config, response), nil
}
