package rainmaker_test

import (
	"time"

	"github.com/pivotal-golang/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Space", func() {
	var config rainmaker.Config

	BeforeEach(func() {
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}
	})

	Describe("FetchSpace", func() {
		It("retrieves the space", func() {
			space, err := rainmaker.FetchSpace(config, "/v2/spaces/space-001", "token-asd")
			Expect(err).NotTo(HaveOccurred())
			createdAt, err := time.Parse(time.RFC3339, "2014-10-09T22:02:26+00:00")
			if err != nil {
				panic(err)
			}

			expectedSpace := rainmaker.NewSpace(config)
			expectedSpace.GUID = "space-001"
			expectedSpace.URL = "/v2/spaces/space-001"
			expectedSpace.CreatedAt = createdAt
			expectedSpace.UpdatedAt = time.Time{}
			expectedSpace.Name = "development"
			expectedSpace.OrganizationGUID = "org-001"
			expectedSpace.SpaceQuotaDefinitionGUID = ""
			expectedSpace.OrganizationURL = "/v2/organizations/org-001"
			expectedSpace.DevelopersURL = "/v2/spaces/space-001/developers"
			expectedSpace.ManagersURL = "/v2/spaces/space-001/managers"
			expectedSpace.AuditorsURL = "/v2/spaces/space-001/auditors"
			expectedSpace.AppsURL = "/v2/spaces/space-001/apps"
			expectedSpace.RoutesURL = "/v2/spaces/space-001/routes"
			expectedSpace.DomainsURL = "/v2/spaces/space-001/domains"
			expectedSpace.ServiceInstancesURL = "/v2/spaces/space-001/service_instances"
			expectedSpace.AppEventsURL = "/v2/spaces/space-001/app_events"
			expectedSpace.EventsURL = "/v2/spaces/space-001/events"
			expectedSpace.SecurityGroupsURL = "/v2/spaces/space-001/security_groups"

			Expect(space).To(Equal(expectedSpace))
		})
	})
})
