package fakes

import (
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/organizations"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/service_instances"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/spaces"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/users"
)

type CloudController struct {
	server           *httptest.Server
	Organizations    *domain.Organizations
	Spaces           *domain.Spaces
	Users            *domain.Users
	ServiceInstances *domain.ServiceInstances
}

func NewCloudController() *CloudController {
	fake := &CloudController{
		Organizations:    domain.NewOrganizations(),
		Spaces:           domain.NewSpaces(),
		Users:            domain.NewUsers(),
		ServiceInstances: domain.NewServiceInstances(),
	}

	router := mux.NewRouter()
	router.Handle("/v2/organizations{anything:.*}", organizations.NewRouter(fake.Organizations, fake.Users))
	router.Handle("/v2/spaces{anything:.*}", spaces.NewRouter(fake.Spaces, fake.Users))
	router.Handle("/v2/users{anything:.*}", users.NewRouter(fake.Users, fake.Spaces))
	router.Handle("/v2/service_instances{anything:.*}", service_instances.NewRouter(fake.ServiceInstances))

	handler := fake.RequireToken(router)
	fake.server = httptest.NewUnstartedServer(handler)

	return fake
}

func (fake *CloudController) Start() {
	fake.server.Start()
}

func (fake *CloudController) Close() {
	fake.server.Close()
}

func (fake *CloudController) URL() string {
	return fake.server.URL
}

func (fake *CloudController) Reset() {
	fake.Organizations.Clear()
	fake.Spaces.Clear()
	fake.Users.Clear()
	fake.ServiceInstances.Clear()
}

func (fake *CloudController) RequireToken(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ok, err := regexp.MatchString(`Bearer .+`, req.Header.Get("Authorization"))
		if err != nil {
			panic(err)
		}

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 Not Authorized"))
			return
		}

		handler.ServeHTTP(w, req)
	})
}
