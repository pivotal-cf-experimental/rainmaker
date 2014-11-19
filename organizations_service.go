package rainmaker

type OrganizationsService struct {
	config Config
}

func NewOrganizationsService(config Config) *OrganizationsService {
	return &OrganizationsService{
		config: config,
	}
}

func (service OrganizationsService) Get(guid, token string) (Organization, error) {
	return FetchOrganization(service.config, "/v2/organizations/"+guid, token)
}

func (service OrganizationsService) ListUsers(guid, token string) (UsersList, error) {
	return FetchUsersList(service.config, "/v2/organizations/"+guid+"/users", token)
}

func (service OrganizationsService) ListBillingManagers(guid, token string) (UsersList, error) {
	return FetchUsersList(service.config, "/v2/organizations/"+guid+"/billing_managers", token)
}

func (service OrganizationsService) ListAuditors(guid, token string) (UsersList, error) {
	return FetchUsersList(service.config, "/v2/organizations/"+guid+"/auditors", token)
}

func (service OrganizationsService) ListManagers(guid, token string) (UsersList, error) {
	return FetchUsersList(service.config, "/v2/organizations/"+guid+"/managers", token)
}
