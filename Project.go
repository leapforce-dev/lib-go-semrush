package semrush

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Project struct {
	ProjectId     int    `json:"project_id"`
	ProjectName   string `json:"project_name"`
	Url           string `json:"url"`
	DomainUnicode string `json:"domain_unicode"`
	Tools         []struct {
		Tool string `json:"tool"`
	} `json:"tools"`
	OwnerId    int      `json:"owner_id"`
	Permission []string `json:"permission"`
}

// GetProjects returns all projects
func (service *Service) GetProjects() (*[]Project, *errortools.Error) {
	var projects []Project

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.urlManagement("projects"),
		ResponseModel: &projects,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &projects, nil
}
