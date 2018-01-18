package go_sportradar_football


import (
	"fmt"
	"net/http"
	"errors"
	"encoding/json"
	"io/ioutil"

)


const (
	baseUrl = "https://api.sportradar.us/soccer-t3/"
	urlTournaments = "/tournaments.json"
)



type Client struct {
	token string
	lang string
}



func (c *Client) Init(token string, lang string) {
	c.token = token
	c.lang = lang
}



func (c *Client) GetTournaments() (TournamentsResult, error) {
	var (
		url string 
		res TournamentsResult
		err error
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
		return nil, errors.New( fmt.Sprintf("API refused with code %v", resp.Status) )
	}

	return body, nil
}


