package rainmaker_test

import (
	"time"

	"github.com/pivotal-golang/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SpacesService", func() {
	var service *rainmaker.SpacesService

	BeforeEach(func() {
		client := rainmaker.NewClient(rainmaker.Config{
			Host: fakeCloudController.URL(),
		})
		service = rainmaker.NewSpacesService(client)
	})

	Describe("Get", func() {
		It("returns the space matching the given GUID", func() {
			space := service.Get("space-001")
			createdAt, err := time.Parse(time.RFC3339, "2014-10-09T22:02:26+00:00")
			if err != nil {
				panic(err)
			}

			Expect(space).To(Equal(rainmaker.Space{
				GUID:                     "space-001",
				URL:                      "/v2/spaces/space-001",
				CreatedAt:                createdAt,
				UpdatedAt:                time.Time{},
				Name:                     "development",
				OrganizationGUID:         "org-001",
				SpaceQuotaDefinitionGUID: "",
				OrganizationURL:          "/v2/organizations/org-001",
				DevelopersURL:            "/v2/spaces/space-001/developers",
				ManagersURL:              "/v2/spaces/space-001/managers",
				AuditorsURL:              "/v2/spaces/space-001/auditors",
				AppsURL:                  "/v2/spaces/space-001/apps",
				RoutesURL:                "/v2/spaces/space-001/routes",
				DomainsURL:               "/v2/spaces/space-001/domains",
				ServiceInstancesURL:      "/v2/spaces/space-001/service_instances",
				AppEventsURL:             "/v2/spaces/space-001/app_events",
				EventsURL:                "/v2/spaces/space-001/events",
				SecurityGroupsURL:        "/v2/spaces/space-001/security_groups",
			}))
		})
	})
})
