package sportmonks

//// Retrieve a team squad resource by season ID and team ID. Use the includes slice to enrich the response data.
//func (c *HTTPClient) TeamSquad(seasonId, teamId int, includes []string) ([]PlayerSeasonStats, *Meta, error) {
//	url := fmt.Sprintf("/api/v2.0/squad/season/%d/team/%d", seasonId, teamId)
//
//	response := new(TeamSquadResponse)
//
//	err := c.handleRequest(url, includes, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return response.Data, &response.Meta, err
//}
