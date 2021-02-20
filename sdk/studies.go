package sdk

import (
	"net/http"
	"path"
)

// StudyResp is a Rapid7 OpenData API Study JSON response
// Ref: https://opendata.rapid7.com/apihelp/
type StudyResp struct {
	Uniqid              string   `json:"uniqid"`
	Name                string   `json:"name"`
	ShortDesc           string   `json:"short_desc"`
	LongDesc            string   `json:"long_desc"`
	StudyURL            string   `json:"study_url"`
	StudyName           string   `json:"study_name"`
	StudyVenue          string   `json:"study_venue"`
	StudyBibtext        string   `json:"study_bibtext"`
	ContactName         string   `json:"contact_name"`
	ContactEmail        string   `json:"contact_email"`
	OrganizationName    string   `json:"organization_name"`
	OrganizationWebsite string   `json:"organization_website"`
	CreatedAt           string   `json:"created_at"`
	UpdatedAt           string   `json:"updated_at"`
	SonarfileSet        []string `json:"sonarfile_set"`
}

// ListStudies returns an array of available studies
func (c *Client) ListStudies() ([]StudyResp, error) {
	c.baseURL.RawPath = path.Join(c.baseURL.Path, "studies")
	req, err := http.NewRequest("GET", c.baseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var studies []StudyResp
	err = c.sendRequest(req, studies)

	return studies, err
}
