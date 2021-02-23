package model

type QuestionType int

const (
	QuestionTypeUnknown QuestionType = iota
	QuestionTypeToggle
	QuestionTypeSelectList
	QuestionTypeNumeric
	QuestionTypeNumericMeasurement
	QuestionTypeText
	QuestionTypeDates
	QuestionTypeUploadFile
	QuestionTypeUploadLink
	QuestionTypeTable
)
