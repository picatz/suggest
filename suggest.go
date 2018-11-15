package suggest

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// Suggestions is a slice of strings that is returned by Suggest
type Suggestions = []string

// Suggest queries google's unofficial suggestion API for Suggestions.
func Suggest(str string) (Suggestions, error) {
	query := url.QueryEscape(str)

	resp, err := http.Get(BaseURL + query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw := []json.RawMessage{}

	err = json.NewDecoder(resp.Body).Decode(&raw)
	if err != nil {
		return nil, err
	}

	if len(raw) < 2 {
		return nil, errors.New("no suggestion data found")
	}

	sgs := Suggestions{}

	err = json.Unmarshal(raw[1], &sgs)
	if err != nil {
		return nil, err
	}

	return sgs, nil
}
