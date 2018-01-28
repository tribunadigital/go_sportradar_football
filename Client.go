package go_sportradar_football

import (
	"fmt"
	"net/http"
	"errors"
	"encoding/json"
	"io/ioutil"
)

const (
	baseUrl                   = "https://api.sportradar.us/soccer-t3/"
	urlTournaments            = "/tournaments.json"
	urlTournamentStanding     = "/tournaments/%s/standings.json"
	urlScheduleForTournaments = "/tournaments/%s/schedule.json"
	urlSeasons                = "/tournaments/%s/seasons.json"
	urlMatchSummary           = "/matches/%s/summary.json"
)

type Client struct {
	token string
	lang  string
}

func (c *Client) Init(token string, lang string) {
	c.token = token
	c.lang = lang
}

func (c *Client) GetDetailMatch(id string) (MatchSummary,  error) {
	var (
		res MatchSummary
		url  string
		err  error
		body []byte
	)

	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang,
		fmt.Sprintf(urlMatchSummary, id), c.token)
	if body, err = c.getUrl(url); err != nil {
		return MatchSummary{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return MatchSummary{}, err
	}

	return res, nil

}

func (c *Client) GetSeasons(id string) (TournamentSeason, error) {
	var (
		res  TournamentSeason
		url  string
		err  error
		body []byte
	)

	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang,
		fmt.Sprintf(urlSeasons, id), c.token)

	if body, err = c.getUrl(url); err != nil {
		return TournamentSeason{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return TournamentSeason{}, err
	}

	return res, nil
}

func (c *Client) GetScheduleForTournament(id string) (Schedule, error) {

	var (
		res  Schedule
		url  string
		err  error
		body []byte
	)

	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang,
		fmt.Sprintf(urlScheduleForTournaments, id), c.token)

	fmt.Println(url)

	if body, err = c.getUrl(url); err != nil {
		return Schedule{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return Schedule{}, err
	}

	return res, nil
}

func (c *Client) GetTournamentStanding(id string) (TournamentStanding, error) {
	var (
		url  string
		res  TournamentStanding
		err  error
		body []byte
	)

	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang,
		fmt.Sprintf(urlTournamentStanding, id), c.token)

	if body, err = c.getUrl(url); err != nil {
		return TournamentStanding{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return TournamentStanding{}, err
	}

	return res, nil
}

func (c *Client) GetTournaments() (TournamentsResult, error) {
	var (
		url  string
		res  TournamentsResult
		err  error
		body []byte
	)

	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang, urlTournaments, c.token)

	if body, err = c.getUrl(url); err != nil {
		return TournamentsResult{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return TournamentsResult{}, err
	}

	return res, nil
}

func (c *Client) getUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("API refused with code %v", resp.Status))
	}

	return body, nil
}
