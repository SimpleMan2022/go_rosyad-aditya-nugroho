package entities

type Word struct {
	Word        string `json:"word"`
	Length      int    `json:"length"`
	NumOfVocals int    `json:"num_of_vocals"`
	Palindrome  bool   `json:"palindrome"`
}

type WordRequest struct {
	Word string `json:"word"`
}

type WordResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
