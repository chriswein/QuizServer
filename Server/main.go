package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type QuestionItem struct {
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

type QuestionSet struct {
	Text    string         `json:"text"`
	Answers []QuestionItem `json:"answers"`
}

var (
	release         = "false"
	questions       = []QuestionSet{}
	currentQuestion = 0
	configfile      = "./config.json"
)

func getQuestion(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions[currentQuestion])
	currentQuestion = (currentQuestion + 1) % len(questions)
}

type Config struct {
	Port int `json:"port"`
}

func readPortFromJSONFile(filename string) (int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return 0, err
	}

	return config.Port, nil
}

func readQuestionsFromFile(file string) []QuestionSet {
	var result []QuestionSet
	data, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf(`Error reading file: %s`, err))
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling JSON: %s", err))
	}
	return result
}

func main() {
	if release == "true" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}
	port, err := readPortFromJSONFile(configfile)
	if err != nil {
		fmt.Println("Could not read config file containing port. Using 8080")
		port = 8080
	}

	router := gin.Default()
	router.GET("/question", getQuestion)
	router.Static("/s", "./dist")

	questions = readQuestionsFromFile("questions.json")
	router.Run(fmt.Sprintf("localhost:%d", port))
}
