package getstream

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Error represents an API error
type Error struct {
	Code            int               `json:"code"`
	Message         string            `json:"message"`
	ExceptionFields map[string]string `json:"exception_fields,omitempty"`
	StatusCode      int               `json:"StatusCode"`
	Duration        string            `json:"duration"`
	MoreInfo        string            `json:"more_info"`
	RateLimit       *RateLimitInfo    `json:"-"`
}

func (e Error) Error() string {
	return e.Message
}

// Response is the base response returned to the client
// type Response struct {
// 	RateLimitInfo *RateLimitInfo `json:"ratelimit"`
// }

// RateLimitInfo represents rate limit information (implementation omitted for brevity)
type RateLimitInfo struct {
	// Fields for rate limit info
}

// NewRateLimitFromHeaders creates a RateLimitInfo from HTTP headers (implementation omitted for brevity)
func NewRateLimitFromHeaders(headers http.Header) *RateLimitInfo {
	return &RateLimitInfo{}
}

// parseResponse parses the HTTP response into the provided result
func parseResponse[U any](c *Client, resp *http.Response, result *U) error {
	if resp.Body == nil {
		return errors.New("http body is nil")
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read HTTP response: %w", err)
	}

	if resp.StatusCode >= 399 {
		var apiErr Error
		err := json.Unmarshal(b, &apiErr)
		if err != nil {
			apiErr.Message = string(b)
			apiErr.StatusCode = resp.StatusCode
			return apiErr
		}
		apiErr.RateLimit = NewRateLimitFromHeaders(resp.Header)
		return apiErr
	}

	if _, ok := any(result).(*Response); !ok {
		err = json.Unmarshal(b, result)
		if err != nil {
			return fmt.Errorf("cannot unmarshal body: %w", err)
		}
	}

	return addRateLimitInfo(c, resp.Header, result)
}

// requestURL constructs the full request URL
func (c *Client) requestURL(path string, values url.Values, pathParams ...string) (string, error) {
	for i, param := range pathParams {
		placeholder := fmt.Sprintf("{param%d}", i+1)
		path = strings.ReplaceAll(path, placeholder, url.PathEscape(param))
	}

	u, err := url.Parse(c.BaseURL + "/" + path)
	if err != nil {
		return "", fmt.Errorf("url.Parse: %w", err)
	}

	if values == nil {
		values = make(url.Values)
	}

	values.Add("api_key", c.apiKey)
	u.RawQuery = values.Encode()

	return u.String(), nil
}

// newRequest creates a new HTTP request
func newRequest[T any](c *Client, ctx context.Context, method, path string, params url.Values, data T, pathParams ...string) (*http.Request, error) {
	u, err := c.requestURL(path, params, pathParams...)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequestWithContext(ctx, method, u, http.NoBody)
	if err != nil {
		return nil, err
	}

	c.setHeaders(r)

	switch t := any(data).(type) {
	case nil:
		r.Body = nil
	case io.ReadCloser:
		r.Body = t
	case io.Reader:
		r.Body = io.NopCloser(t)
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		r.Body = io.NopCloser(bytes.NewReader(b))
	}

	return r, nil
}

// setHeaders sets necessary headers for the request
func (c *Client) setHeaders(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Stream-Client", versionHeader())
	r.Header.Set("Authorization", c.authToken)
	r.Header.Set("Stream-Auth-Type", "jwt")
}

// makeRequest makes a generic HTTP request
func MakeRequest[T any, U any](c *Client, ctx context.Context, method, path string, params url.Values, data T, result *U, pathParams ...string) error {
	r, err := newRequest(c, ctx, method, path, params, data, pathParams...)
	if err != nil {
		return err
	}

	resp, err := c.HTTP.Do(r)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		return err
	}

	return parseResponse(c, resp, result)
}

// TODO: revisit this
// addRateLimitInfo adds rate limit information to the result
func addRateLimitInfo[U any](c *Client, headers http.Header, result *U) error {
	rl := map[string]interface{}{
		"ratelimit": NewRateLimitFromHeaders(headers),
	}

	b, err := json.Marshal(rl)
	if err != nil {
		return fmt.Errorf("cannot marshal rate limit info: %w", err)
	}

	err = json.Unmarshal(b, result)
	if err != nil {
		return fmt.Errorf("cannot unmarshal rate limit info: %w", err)
	}
	return nil
}

// versionHeader returns the version header (implementation omitted for brevity)
func (c *Client) version() string {
	return versionHeader()
}