package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("survey.csv")
	if err != nil {
		log.Fatal("Error file not found")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var recordsMentor []int
	var recordsLearning []int
	courseMap := make(map[string]int)

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		recordMentor, _ := strconv.Atoi(record[4])
		recordsMentor = append(recordsMentor, recordMentor)

		recordLearning, _ := strconv.Atoi(record[3])
		recordsLearning = append(recordsLearning, recordLearning)
		courseMap[record[2]] += recordMentor

	}
	fmt.Println("Mean Mentor Satisfaction Score:", meanMentorSatisfactionScore(recordsMentor))
	fmt.Println("Mean Learning Satisfaction Score:", meanLearningSatisfactionScore(recordsLearning))
	fmt.Println("Courses the highest mentor satisfaction: ", mostHighMeanMentorSatisfaction(courseMap, recordsMentor))

}

func meanMentorSatisfactionScore(records []int) int {
	var total int
	for _, record := range records {
		total += record
	}
	lenRecords := len(records)
	return total / lenRecords
}

func meanLearningSatisfactionScore(records []int) int {
	var total int
	for _, record := range records {
		total += record
	}
	lenRecords := len(records)
	return total / lenRecords
}

func mostHighMeanMentorSatisfaction(courseMap map[string]int, recordMentors []int) string {
	var highestMean float64
	var highestCourse string
	for course, mentorScore := range courseMap {
		mean := float64(mentorScore) / float64(len(recordMentors))
		if mean > highestMean {
			highestMean = mean
			highestCourse = course
		}
	}
	return highestCourse
}
