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
	Id     string                  `json:"id"`
	Name   *string                 `json:"name"`
	Order  *int                    `json:"order"`
	Fields *[]setdata_common.Field `json:"fields"`
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

type AddFieldToQuestionCommand struct {
	QuestionId  string                   `json:"question_id"`
	Name        string                   `json:"name"`
	Type        setdata_common.FieldType `json:"type"`
	Placeholder string                   `json:"placeholder"`
}

func (cmd *AddFieldToQuestionCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).AddFieldToQuestion(cmd)
}

type CreateQuestionnaireCommand struct {
	Name      string   `json:"name"`
	Questions []string `json:"questions"`
}

func (cmd *CreateQuestionnaireCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).CreateQuestionnaire(cmd)
}

type GetQuestionnaireByIdCommand struct {
	Id string `json:"id"`
}

func (cmd *GetQuestionnaireByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).GetQuestionnaireById(cmd)
}

type GetQuestionnaireByNameCommand struct {
	Name string `json:"name"`
}

func (cmd *GetQuestionnaireByNameCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).GetQuestionnaireByName(cmd)
}

type UpdateQuestionnaireCommand struct {
	Id        string    `json:"id"`
	Name      *string   `json:"name"`
	Questions *[]string `json:"questions"`
}

func (cmd *UpdateQuestionnaireCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).UpdateQuestionnaire(cmd)
}

type DeleteQuestionnaireByIdCommand struct {
	Id string `json:"id"`
}

func (cmd *DeleteQuestionnaireByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return nil, svc.(QuestionsService).DeleteQuestionnaire(cmd)
}

type ListQuestionnaireCommand struct {
}

func (cmd *ListQuestionnaireCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).ListQuestionnaire(cmd)
}

type AddQuestionToQuestionnaireCommand struct {
	Id         string `json:"id"`
	QuestionId string `json:"question_id"`
}

func (cmd *AddQuestionToQuestionnaireCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).AddQuestionToQuestionnaire(cmd)
}

type RemoveQuestionFromQuestionnaireCommand struct {
	Id         string `json:"id"`
	QuestionId string `json:"question_id"`
}

func (cmd *RemoveQuestionFromQuestionnaireCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionsService).RemoveQuestionFromQuestionnaire(cmd)
}

type QuestionnaireDetail struct {
	Id        string                                 `json:"id"`
	Name      string                                 `json:"name"`
	Questions []setdata_questionnaire_store.Question `json:"questions"`
}
