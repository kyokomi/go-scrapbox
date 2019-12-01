package scrapbox

// PageListResponse /api/pages/:projectName のresponse https://scrapbox.io/help-jp/API#58e676b197c29100006748f7
type PageListResponse struct {
	ProjectName string  `json:"projectName"`
	Skip        uint    `json:"skip"`
	Limit       uint    `json:"limit"`
	Count       uint    `json:"count"`
	Pages       []*Page `json:"pages"`
}

// Page page
type Page struct {
	ID                      string   `json:"id"`
	Title                   string   `json:"title"`
	Image                   string   `json:"image"`
	Descriptions            []string `json:"descriptions"`
	User                    User     `json:"user"`
	Pin                     uint64   `json:"pin"`
	Views                   uint     `json:"views"`
	Linked                  uint     `json:"linked"`
	CommitID                string   `json:"commitId"`
	CreatedUnixTime         uint64   `json:"created"`
	UpdatedUnixTime         uint64   `json:"updated"`
	AccessedUnixTime        uint64   `json:"accessed"`
	SnapshotCreatedUnixTime uint64   `json:"snapshotCreated"`
}

// User user
type User struct {
	ID string `json:"id"`
}

// ErrorResponse api error response
type ErrorResponse struct {
	Name       string `json:"name"`
	Message    string `json:"message"`
	StatusCode uint   `json:"statusCode"`
}
