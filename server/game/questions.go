package game

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

// OpenTDBToken : The data object returned from the OpenTDB API on token request
type OpenTDBToken struct {
	ResponseCode    int8   `json:"response_code"`
	ResponseMessage string `json:"response_message"`
	Token           string `json:"token"`
}

// OpenTDBTokenJSON : The converted json structure of OpenTDBToken type
type OpenTDBTokenJSON struct {
	ResponseCode    int8   `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Token           string `json:"token"`
}

// GameCard : All game card metadata, including the question and list of options, including which answer is correct
type GameCard struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

// Category : The object structure of individual category items returned from OpenTDB API on category list request
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// QuestionAPIResult : The data structure returned from calling the get questions OPenTDB API
type QuestionAPIResult struct {
	ResponseCode int8       `json:"response_code"`
	Data         []GameCard `json:"results"`
}

// RequestToken : Generate a new token for OpenTDB API
func RequestToken() OpenTDBTokenJSON {
	var token OpenTDBToken

	response, err := http.Get(os.Getenv("OPENTDB_TOKEN_URL"))
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bodyByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(bodyByte, &token)

	return OpenTDBTokenJSON(token)
}

// GetQuestions : Return a list of GameCard objects to be used for a game
func GetQuestions(
	gameToken string,
	amount string,
	category string,
	difficulty string,
) []GameCard {
	var apiResult QuestionAPIResult

	client := &http.Client{}
	targetUrl, _ := url.Parse(os.Getenv("OPENTDB_HOST_URL"))
	params, _ := url.ParseQuery(targetUrl.RawQuery)
	params.Add("amount", amount)
	params.Add("category", category)
	params.Add("difficulty", difficulty)
	params.Add("token", gameToken)
	targetUrl.RawQuery = params.Encode()

	request, err := http.NewRequest("GET", targetUrl.String(), strings.NewReader(params.Encode()))
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bodyByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(bodyByte, &apiResult)

	return apiResult.Data
}
