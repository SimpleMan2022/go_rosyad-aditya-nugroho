package controllers

import (
	"net/http"
	"slices"
	"soal_priortias_2/entities"
	"soal_priortias_2/errorHandler"
	"strings"

	"github.com/labstack/echo/v4"
)

var words []entities.Word

func isPalindrome(word string) bool {
	lowerWord := strings.ToLower(word)
	arr := strings.Split(lowerWord, "")

	slices.Reverse(arr)
	reverseArr := strings.Join(arr, "")

	if reverseArr == lowerWord {
		return true
	}
	return false
}

func countVocals(name string) int {
	lowerName := strings.ToLower(name)
	arr := strings.Split(lowerName, "")

	count := 0
	for _, char := range arr {
		if char == "a" || char == "e" || char == "i" || char == "u" || char == "o" {
			count++
		}
	}
	return count
}

func AddNewWord(c echo.Context) error {
	var request entities.WordRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusInternalServerError, errorHandler.ErrorHandler{
			Message: err.Error(),
		})
	}

	if request.Word == "" {
		return c.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
			Message: "failed to add word data",
		})

	}

	isPalindrome := isPalindrome(request.Word)
	numOfVocals := countVocals(request.Word)

	word := entities.Word{
		Word:        request.Word,
		Length:      len(request.Word),
		NumOfVocals: numOfVocals,
		Palindrome:  isPalindrome,
	}

	words = append(words, word)

	response := entities.WordResponse{
		Message: "word added",
		Data:    word,
	}

	return c.JSON(http.StatusOK, response)

}

func GetAllWords(c echo.Context) error {
	if len(words) == 0 {
		return c.JSON(http.StatusNotFound, errorHandler.ErrorHandler{
			Message: "words is empty",
		})
	}
	response := entities.WordResponse{
		Message: "all words",
		Data:    words,
	}

	return c.JSON(http.StatusOK, response)
}
