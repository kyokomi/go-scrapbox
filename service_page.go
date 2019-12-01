package scrapbox

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"golang.org/x/xerrors"
)

// PageService page service
type PageService struct {
	*Client
}

// List get page list
func (s *PageService) List(ctx context.Context, projectName string, offset, limit uint) (PageListResponse, error) {
	u, err := s.Client.createURL("pages", projectName)
	if err != nil {
		return PageListResponse{}, xerrors.Errorf("create url error: %w", err)
	}

	val := url.Values{}
	val.Add("skip", strconv.FormatUint(uint64(offset), 10))
	val.Add("limit", strconv.FormatUint(uint64(limit), 10))
	u.RawQuery = val.Encode()

	req, err := s.Client.newRequest(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return PageListResponse{}, xerrors.Errorf("new request error: %w", err)
	}

	var res PageListResponse
	if err := s.Client.doAndJSONDecode(req, &res); err != nil {
		return PageListResponse{}, xerrors.Errorf("doAndJSONDecode error: %w", err)
	}

	return res, nil
}

// Get get page
func (s *PageService) Get(ctx context.Context, projectName string, title string) (PageResponse, error) {
	u, err := s.Client.createURL("pages", path.Join(projectName, title))
	if err != nil {
		return PageResponse{}, xerrors.Errorf("create url error: %w", err)
	}

	req, err := s.Client.newRequest(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return PageResponse{}, xerrors.Errorf("new request error: %w", err)
	}

	var res PageResponse
	if err := s.Client.doAndJSONDecode(req, &res); err != nil {
		return PageResponse{}, xerrors.Errorf("doAndJSONDecode error: %w", err)
	}

	return res, nil
}

// Text get page text
func (s *PageService) Text(ctx context.Context, projectName string, title string) (string, error) {
	u, err := s.Client.createURL("pages", path.Join(projectName, title, "text"))
	if err != nil {
		return "", xerrors.Errorf("create url error: %w", err)
	}

	req, err := s.Client.newRequest(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return "", xerrors.Errorf("new request error: %w", err)
	}

	return s.Client.doAndGetText(req)
}

// IconURL get page icon url
func (s *PageService) IconURL(ctx context.Context, projectName string, title string) (bool, *url.URL, error) {
	u, err := s.Client.createURL("pages", path.Join(projectName, title, "icon"))
	if err != nil {
		return false, nil, xerrors.Errorf("create url error: %w", err)
	}

	req, err := s.Client.newRequest(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return false, nil, xerrors.Errorf("new request error: %w", err)
	}

	return s.Client.doAndGetRedirectURL(req)
}
