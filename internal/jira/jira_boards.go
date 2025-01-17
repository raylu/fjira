package jira

import (
	"encoding/json"
	"fmt"
)

const (
	FindAllBoardsUrl          = "/rest/agile/1.0/board"
	FindBoardConfigurationUrl = "/rest/agile/1.0/board/%d/configuration"
)

type findBoardsQueryParams struct {
	ProjectKeyOrId string `url:"projectKeyOrId"`
}

func (api *httpApi) FindBoards(projectKeyOrId string) ([]BoardItem, error) {
	resultBytes, err := api.jiraRequest("GET", FindAllBoardsUrl, &findBoardsQueryParams{ProjectKeyOrId: projectKeyOrId}, nil)
	if err != nil {
		return nil, err
	}
	var result BoardsResponse
	err = json.Unmarshal(resultBytes, &result)
	if err != nil {
		return nil, err
	}
	return result.Values, nil
}

func (api *httpApi) GetBoardConfiguration(boardId int) (*BoardConfiguration, error) {
	resultBytes, err := api.jiraRequest("GET", fmt.Sprintf(FindBoardConfigurationUrl, boardId), &nilParams{}, nil)
	if err != nil {
		return nil, err
	}
	var result BoardConfiguration
	err = json.Unmarshal(resultBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
