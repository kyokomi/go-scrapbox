package scrapbox

// PageListResponse /api/pages/:projectName のresponse https://scrapbox.io/help-jp/API#58e676b197c29100006748f7
type PageListResponse struct {
	ProjectName string `json:"projectName"`
	Skip        uint   `json:"skip"`
	Limit       uint   `json:"limit"`
	Count       uint   `json:"count"`
	Pages       []*struct {
		ID           string   `json:"id"`
		Title        string   `json:"title"`
		Image        string   `json:"image"`
		Descriptions []string `json:"descriptions"`
		User         struct {
			ID string `json:"id"`
		} `json:"user"`
		Pin                     uint64 `json:"pin"`
		Views                   uint   `json:"views"`
		Linked                  uint   `json:"linked"`
		CommitID                string `json:"commitId"`
		CreatedUnixTime         uint64 `json:"created"`
		UpdatedUnixTime         uint64 `json:"updated"`
		AccessedUnixTime        uint64 `json:"accessed"`
		SnapshotCreatedUnixTime uint64 `json:"snapshotCreated"`
	} `json:"pages"`
}

// PageResponse page response https://scrapbox.io/help-jp/API#58e6775897c29100006748fe
type PageResponse struct {
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

	Persistent bool `json:"persistent"`
	Lines      []struct {
		ID      string `json:"id"`
		Text    string `json:"text"`
		UserID  string `json:"userId"`
		Created int    `json:"created"`
		Updated int    `json:"updated"`
	} `json:"lines"`
	Links        []string `json:"links"`
	Icons        struct{} `json:"icons"` // TODO: ?
	RelatedPages struct {
		Links1Hop []*LinkHop    `json:"links1hop"`
		Links2Hop []*LinkHop    `json:"links2hop"`
		Icons1Hop []interface{} `json:"icons1hop"` // TODO: ?
	} `json:"relatedPages"`
	Collaborators        []interface{} `json:"collaborators"` // TODO: ?
	LastAccessedUnixTime uint64        `json:"lastAccessed"`
}

// User user
type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Photo       string `json:"photo"`
}

// LinkHop page link hop
type LinkHop struct {
	ID               string   `json:"id"`
	Title            string   `json:"title"`
	TitleLc          string   `json:"titleLc"`
	Image            string   `json:"image"`
	Descriptions     []string `json:"descriptions"`
	LinksLc          []string `json:"linksLc"`
	UpdatedUnixTime  uint64   `json:"updated"`
	AccessedUnixTime uint64   `json:"accessed"`
}

// ErrorResponse api error response
type ErrorResponse struct {
	Name       string `json:"name"`
	Message    string `json:"message"`
	StatusCode uint   `json:"statusCode"`
}
