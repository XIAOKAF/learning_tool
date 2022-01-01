package service

import (
	"learning-tool/dao"
	"learning-tool/modles"
)

func CreateReportCard(grades modles.Grades) error {
	err := dao.CreateReportCard(grades)
	if err != nil {
		return err
	}
	return nil
}

func InsertGrades(grades modles.Grades) error {
	err := dao.InsertGrades(grades)
	if err != nil {
		return err
	}
	return nil
}

func AverageScores(grades modles.Grades) (float32, error) {
	averageScores, err := dao.AverageScore(grades)
	if err != nil {
		return 0, err
	}
	return averageScores, nil
}

func HighestScore(grades modles.Grades) (float32, string, error) {
	maxScore, student, err := dao.HighestScore(grades)
	if err != nil {
		return maxScore, student, err
	}
	return maxScore, student, nil
}

func SpecifiedStudent(grades modles.Grades) (modles.Grades,error) {
	grades,err := dao.SpecifiedStudent(grades)
	if err != nil {
		return grades,err
	}
	return grades,nil
}
