package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/reconmap/cli/internal/configuration"
	"github.com/reconmap/cli/internal/httputils"
)

func GetCommandById(id int) (*Command, error) {
	config, err := configuration.ReadConfig()
	if err != nil {
		return nil, err
	}
	var apiUrl string = config.ApiUrl + "/commands/" + strconv.Itoa(id)

	client := &http.Client{}
	req, err := httputils.NewRmapRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httputils.AddBearerToken(req)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("error from server: " + string(response.Status))
	}

	if err != nil {
		return nil, errors.New("Unable to read response from server")
	}

	var command *Command = &Command{}

	json.Unmarshal([]byte(body), command)

	return command, nil
}

func GetCommandsByKeywords(keywords string) (*Commands, error) {
	config, err := configuration.ReadConfig()
	if err != nil {
		return nil, err
	}
	var apiUrl string = config.ApiUrl + "/commands?keywords=" + keywords

	client := &http.Client{}
	req, err := httputils.NewRmapRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httputils.AddBearerToken(req)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("error from server: " + string(response.Status))
	}

	if err != nil {
		return nil, errors.New("Unable to read response from server")
	}

	var commands *Commands = &Commands{}

	json.Unmarshal(body, commands)

	return commands, nil
}
