package model

type QuestionStatus int

const (
	QuestionStatusUnknown QuestionStatus = iota
	QuestionStatusPublic
	QuestionStatusCustom
	QuestionStatusDraft
)
