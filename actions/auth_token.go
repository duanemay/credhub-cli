package actions

import (
	"encoding/json"
	"net/http"

	"github.com/pivotal-cf/credhub-cli/client"
	"github.com/pivotal-cf/credhub-cli/config"
	"github.com/pivotal-cf/credhub-cli/errors"
	"github.com/pivotal-cf/credhub-cli/models"
)

func NewAuthToken(httpClient client.HttpClient, config config.Config) ServerInfo {
	return ServerInfo{httpClient: httpClient, config: config}
}

func (serverInfo ServerInfo) GetAuthToken(user string, pass string) (models.Token, error) {
	request := client.NewAuthTokenRequest(serverInfo.config, user, pass)
	response, err := serverInfo.httpClient.Do(request)
	if err != nil {
		return models.Token{}, errors.NewNetworkError(err)
	}

	if response.StatusCode != http.StatusOK {
		return models.Token{}, errors.NewAuthorizationError()
	}

	token := new(models.Token)

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(token)

	if err != nil {
		return models.Token{}, err
	}

	return *token, nil
}