package organizations

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type associateBillingManagerHandler struct {
	organizations *domain.Organizations
	users         *domain.Users
}

func (h associateBillingManagerHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := regexp.MustCompile(`^/v2/organizations/(.*)/billing_managers/(.*)$`)
	matches := r.FindStringSubmatch(req.URL.Path)

	org, ok := h.organizations.Get(matches[1])
	if !ok {
		common.NotFound(w)
		return
	}

	billingManager, ok := h.users.Get(matches[2])
	if !ok {
		common.NotFound(w)
		return
	}

	org.BillingManagers.Add(billingManager)

	response, err := json.Marshal(org)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
