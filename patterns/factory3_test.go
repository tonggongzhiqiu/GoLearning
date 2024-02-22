package patterns

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewHttpClient() Doer {
	return &http.Client{}
}

type mockHttpClient struct{}

func (*mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	res := httptest.NewRecorder()
	return res.Result(), nil
}

func NewMockHttpClient() Doer {
	return &mockHttpClient{}
}

func QueryUser(doer Doer) error {
	req, err := http.NewRequest("Get", "http://iam.api.marmotedu.com:8080/v1/secrets", nil)
	if err != nil {
		return err
	}

	_, err = doer.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func TestQueryUser(t *testing.T) {
	doer := NewMockHttpClient()
	if err := QueryUser(doer); err != nil {
		t.Errorf("QueryUser failed, err = %v", err)
	}
}
