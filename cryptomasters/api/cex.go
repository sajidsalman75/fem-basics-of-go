package api

import (
	"encoding/json"
	"fem/go/cryptomasters/datatypes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var cryptoRate datatypes.Rate
		err = json.Unmarshal(bodyBytes, &cryptoRate)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: 20}
	return &rate, nil
}
