package rainmaker

import "github.com/pivotal-cf-experimental/rainmaker/internal/documents"

type Application struct {
	config    Config
	GUID      string
	Name      string
	SpaceGUID string
}

func NewApplication(config Config, guid string) Application {
	return Application{
		config: config,
		GUID:   guid,
	}
}

func newApplicationFromCreateResponse(config Config, response documents.ApplicationCreateResponse) Application {
	app := NewApplication(config, response.Metadata.GUID)

	app.Name = response.Entity.Name
	app.SpaceGUID = response.Entity.SpaceGUID

	return app
}

func newApplicationFromSummaryResponse(config Config, response documents.ApplicationSummaryResponse) Application {
	app := NewApplication(config, response.GUID)

	app.Name = response.Name
	app.SpaceGUID = response.SpaceGUID

	return app
}
