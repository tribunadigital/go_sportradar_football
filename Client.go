package go_sportradar_football

import (
	"fmt"
	"net/http"
	"errors"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

const (
	baseUrl                   = "https://api.sportradar.us/soccer-p3/"
	urlTournaments            = "/tournaments.json"
	urlTournamentStanding     = "/tournaments/%s/standings.json"
	urlScheduleForTournaments = "/tournaments/%s/schedule.json"
	urlSeasons                = "/tournaments/%s/seasons.json"
	urlMatchSummary           = "/matches/%s/summary.json"
	urlMatchesLineup          = "/matches/%s/lineups.json"
	urlTeam                   = "/teams/%s/profile.json"
	urlDailySchedules         = "/schedules/%s/schedule.json"
	urlDailyResults           = "/schedules/%s/results.json"
	urlLiveResults            = "/schedules/live/results.json"
	urlPlayerProfile          = "/players/%s/profile.json"
)

type Client struct {
	token     string
	lang      string
	LastQuery HttpLog
}

func (c *Client) Init(token string, lang string) {
	c.token = token
	c.lang = lang
	c.LastQuery = HttpLog{}
}

func (c *Client) GetDailySchedule(date string) (DailySchedule, error) {
	var (
		res  DailySchedule
		url  string
		err  error
		body []byte
	)
	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang,
		fmt.Sprintf(urlDailySchedules, date), c.token)

	if body, err = c.getUrl(url); err != nil {
		return DailySchedule{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return DailySchedule{}, err
	}

	return res, nil
}

func (c *Client) GetLiveResult() (LiveResult, error) {
	var (
		res  LiveResult
		url  string
		err  error
		body []byte
	)
	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang,
		urlLiveResults, c.token)

	fmt.Println(url)

	if body, err = c.getUrl(url); err != nil {
		return LiveResult{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return LiveResult{}, err
	}

	return res, nil
}

func (c *Client) GetDailyResult(date string) (DailyResult, error) {
	var (
		res  DailyResult
		url  string
		err  error
		body []byte
	)
	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang,
		fmt.Sprintf(urlDailyResults, date), c.token)

	fmt.Println(url)

	if body, err = c.getUrl(url); err != nil {
		return DailyResult{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return DailyResult{}, err
	}

	return res, nil
}

func (c *Client) GetPlayerProfile(id string) (PlayerProfile, error) {
	var (
		res  PlayerProfile
		url  string
		err  error
		body []byte
	)
	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang,
		fmt.Sprintf(urlPlayerProfile, id), c.token)

	if body, err = c.getUrl(url); err != nil {
		return PlayerProfile{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return PlayerProfile{}, err
	}

	return res, nil
}

func (c *Client) GetTeamProfile(id string) (TeamProfile, error) {
	var (
		res  TeamProfile
		url  string
		err  error
		body []byte
	)

	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang,
		fmt.Sprintf(urlTeam, id), c.token)

	if body, err = c.getUrl(url); err != nil {
		return TeamProfile{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return TeamProfile{}, err
	}

	return res, nil
}

func (c *Client) GetMatchLineups(id string) (MatchLineups, error) {

	var (
		res  MatchLineups
		url  string
		err  error
		body []byte
	)

	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang,
		fmt.Sprintf(urlMatchesLineup, id), c.token)

	fmt.Println(url)
	if body, err = c.getUrl(url); err != nil {
		return MatchLineups{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return MatchLineups{}, err
	}

	return res, nil
}

func (c *Client) GetDetailMatch(id string) (MatchSummary, error) {
	var (
		res  MatchSummary
		url  string
		err  error
		body []byte
	)

	url = fmt.Sprintf("%s%s%s?api_key=%s", baseUrl, c.lang,
		fmt.Sprintf(urlMatchSummary, id), c.token)

	fmt.Println(url)
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

	fmt.Println(url)

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

	fmt.Println(url)

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

	c.LastQuery.Request = url

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.LastQuery.Response = string(body[:])
	c.LastQuery.CodeStatus = strconv.Itoa(resp.StatusCode)

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("API refused with code %v", resp.Status))
	}

	return body, nil
}
