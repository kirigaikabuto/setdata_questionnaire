package setdata_questionnaire

import (
	setdata_common "github.com/kirigaikabuto/setdata-common"
	setdata_questionnaire_store "github.com/kirigaikabuto/setdata-questionnaire-store"
)

type CreateQuestionCommand struct {
	Name   string                 `json:"name"`
	Order  *int                   `json:"order"`
	Fields []setdata_common.Field `json:"fields"`
}

func (cmd *CreateQuestionCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).CreateQuestion(cmd)
}

type UpdateQuestionsCommand struct {
	*setdata_questionnaire_store.QuestionUpdate
}

func (cmd *UpdateQuestionsCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).UpdateQuestion(cmd)
}

type GetQuestionCommand struct {
	Id string `json:"id"`
}

func (cmd *GetQuestionCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).GetQuestion(cmd)
}

type ListQuestionsCommand struct {
}

func (cmd *ListQuestionsCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).ListQuestions(cmd)
}

type DeleteQuestionCommand struct {
	Id string `json:"id"`
}

func (cmd *DeleteQuestionCommand) Exec(svc interface{}) (interface{}, error) {
	return nil, svc.(QuestionsService).DeleteQuestion(cmd)
}
