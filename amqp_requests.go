package setdata_questionnaire

import (
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	setdata_questionnaire_store "github.com/kirigaikabuto/setdata-questionnaire-store"
)

type AmqpRequests struct {
	clt amqp.Client
}

func NewAmqpRequests(clt amqp.Client) AmqpRequests {
	return AmqpRequests{clt: clt}
}

func (r *AmqpRequests) CreateQuestion(cmd *CreateQuestionCommand) (*setdata_questionnaire_store.Question, error) {
	response := &setdata_questionnaire_store.Question{}
	err := setdata_common.AmqpCall(r.clt, "questions.create", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) UpdateQuestion(cmd *UpdateQuestionsCommand) (*setdata_questionnaire_store.Question, error) {
	response := &setdata_questionnaire_store.Question{}
	err := setdata_common.AmqpCall(r.clt, "questions.update", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) GetQuestion(cmd *GetQuestionCommand) (*setdata_questionnaire_store.Question, error) {
	response := &setdata_questionnaire_store.Question{}
	err := setdata_common.AmqpCall(r.clt, "questions.get", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) DeleteQuestion(cmd *DeleteQuestionCommand) error {
	err := setdata_common.AmqpCall(r.clt, "questions.delete", cmd, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *AmqpRequests) ListQuestions(cmd *ListQuestionsCommand) ([]setdata_questionnaire_store.Question, error) {
	response := []setdata_questionnaire_store.Question{}
	err := setdata_common.AmqpCall(r.clt, "questions.list", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) CreateQuestionnaire(cmd *CreateQuestionnaireCommand) (*setdata_questionnaire_store.Questionnaire, error) {
	response := &setdata_questionnaire_store.Questionnaire{}
	err := setdata_common.AmqpCall(r.clt, "questionnaire.create", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) UpdateQuestionnaire(cmd *UpdateQuestionnaireCommand) (*setdata_questionnaire_store.Questionnaire, error) {
	response := &setdata_questionnaire_store.Questionnaire{}
	err := setdata_common.AmqpCall(r.clt, "questionnaire.update", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) ListQuestionnaire(cmd *ListQuestionnaireCommand) ([]setdata_questionnaire_store.Questionnaire, error) {
	response := []setdata_questionnaire_store.Questionnaire{}
	err := setdata_common.AmqpCall(r.clt, "questionnaire.list", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) GetQuestionnaireById(cmd *GetQuestionnaireByIdCommand) (*setdata_questionnaire_store.Questionnaire, error) {
	response := &setdata_questionnaire_store.Questionnaire{}
	err := setdata_common.AmqpCall(r.clt, "questionnaire.getById", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) GetQuestionnaireByName(cmd *GetQuestionnaireByNameCommand) (*setdata_questionnaire_store.Questionnaire, error) {
	response := &setdata_questionnaire_store.Questionnaire{}
	err := setdata_common.AmqpCall(r.clt, "questionnaire.getByName", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) DeleteQuestionnaire(cmd *DeleteQuestionnaireByIdCommand) error {
	err := setdata_common.AmqpCall(r.clt, "questionnaire.delete", cmd, nil)
	if err != nil {
		return err
	}
	return nil
}