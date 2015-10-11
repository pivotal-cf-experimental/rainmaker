package rainmaker

import (
	"io"

	"github.com/pivotal-cf-experimental/rainmaker/internal/network"
)

type Config struct {
	SkipVerifySSL bool
	Host          string
	TraceWriter   io.Writer
}

type Client struct {
	config           Config
	Organizations    *OrganizationsService
	Spaces           *SpacesService
	Users            *UsersService
	ServiceInstances *ServiceInstancesService
}

func NewClient(config Config) Client {
	return Client{
		config:           config,
		Organizations:    NewOrganizationsService(config),
		Spaces:           NewSpacesService(config),
		Users:            NewUsersService(config),
		ServiceInstances: NewServiceInstancesService(config),
	}
}

func newNetworkClient(config Config) network.Client {
	return network.NewClient(network.Config{
		Host:          config.Host,
		SkipVerifySSL: config.SkipVerifySSL,
		TraceWriter:   config.TraceWriter,
	})
}
