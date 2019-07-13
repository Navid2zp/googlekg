package googlekg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (r *KGRequest) setParam(param string, paramValues []string) {
	for _, p := range paramValues {
		r.URL += "&" + param + "=" + p
	}
}

func Make(apiKey string) (*KGRequest, error) {
	if apiKey == "" {
		return nil, errors.New("no valid api key provided")
	}

	request := KGRequest{
		URL: "https://kgsearch.googleapis.com/v1/entities:search?key=" + apiKey,
	}
	return &request, nil
}

func (r *KGRequest) SetLimit(limit int) {
	r.setParam("limit", []string{string(limit)})
}

func (r *KGRequest) Query(query string) {
	r.setParam("query", []string{query})
}

func (r *KGRequest) SetIndent(indent bool) {
	indentString := "true"
	if !indent {
		indentString = "false"
	}
	r.setParam("indent", []string{indentString})
}

func (r *KGRequest) SetPreFix(prefix bool) {
	prefixString := "true"
	if !prefix {
		prefixString = "false"
	}
	r.setParam("prefix", []string{prefixString})
}

func (r *KGRequest) SetIDs(ids []string) {
	r.setParam("ids", ids)
}

func (r *KGRequest) SetLanguages(languages []string) {
	r.setParam("languages", languages)
}

func (r *KGRequest) SetTypes(types []string) {
	r.setParam("types", types)
}

func (r *KGRequest) Do() (*KG, error) {
	if r.URL == "" {
		return nil, errors.New("call Make() func first")
	}

	res, err := http.Get(r.URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("request failed with status code " + string(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	responseData := KG{}
	err = json.Unmarshal(body, &responseData)
	return &responseData, err
}
