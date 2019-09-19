package sportmonks

import "strconv"

const squadUri = "/api/v2.0/squad/season/"

func (c Client) SquadBySeasonAndTeam(seasonId, teamId int, includes []string, retries int) (*SquadResponse, error) {
	url := c.BaseURL + squadUri + strconv.Itoa(seasonId) + "/team/" + strconv.Itoa(teamId) + "?api_token=" + c.ApiKey

	req, err := buildRequest("GET", url, nil, 0, includes)

	if err != nil {
		return nil, err
	}

	r := new(SquadResponse)

	err = c.sendRequest(req, &r, retries)

	return r, err
}
