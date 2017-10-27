package auth_backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"golang.org/x/oauth2"
)

// Backend is the main part of auth_backend.
// It stores oauth2.Config, app and edx url configuration
type Backend struct {
	conf    *oauth2.Config
	app_url string
	edx_url string
}

// Init to initialize a new backend authentication
func Init(client_id, client_secret, app_url, edx_url string) *Backend {
	return &Backend{
		conf: &oauth2.Config{
			ClientID:     client_id,
			ClientSecret: client_secret,
			RedirectURL:  app_url + URL_REDIRECT,
			Scopes:       Scopes(),
			Endpoint: oauth2.Endpoint{
				AuthURL:  edx_url + URL_AUTHORIZATION,
				TokenURL: edx_url + URL_ACCESS_TOKEN,
			},
		},
		app_url: app_url,
		edx_url: edx_url,
	}
}

// GetAuthorizationURL return authentication url based on given state
func (backend *Backend) GetAuthorizationURL(state string) string {
	return backend.conf.AuthCodeURL(state)
}

// Authenticate is callback handler when receiving redirection from edx site after successfully login
func (backend *Backend) Authenticate(receivedState string, values url.Values) (UserInfo, error) {
	var user UserInfo

	queryState := values.Get("state")
	if receivedState != queryState {
		return user, fmt.Errorf("Invalid session state: retrieved: %s; Param: %s", receivedState, queryState)
	}

	code := values.Get("code")
	token, err := backend.conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return user, err
	}

	client := backend.conf.Client(oauth2.NoContext, token)
	userinfo, err := client.Get(backend.edx_url + URL_USER_INFO)
	if err != nil {
		return user, err
	}
	defer userinfo.Body.Close()

	data, _ := ioutil.ReadAll(userinfo.Body)
	if err := json.Unmarshal(data, &user); err != nil {
		return user, err
	}

	return user, nil
}
