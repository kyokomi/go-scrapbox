package scrapbox

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// PageService pageList service
type PageService struct {
	*Client
}

// List get page list
func (s *PageService) List(ctx context.Context, projectName string, offset, limit uint) (PageListResponse, error) {
	u, err := s.Client.createURL("pages", projectName)
	if err != nil {
		return PageListResponse{}, err
	}

	val := url.Values{}
	val.Add("skip", strconv.FormatUint(uint64(offset), 10))
	val.Add("limit", strconv.FormatUint(uint64(limit), 10))
	u.RawQuery = val.Encode()

	req, err := s.Client.newRequest(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return PageListResponse{}, err
	}

	var res PageListResponse
	if err := s.Client.doAndJSONDecode(req, &res); err != nil {
		return PageListResponse{}, err
	}

	return res, nil
}
