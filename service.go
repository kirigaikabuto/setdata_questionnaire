package setdata_questionnaire

import setdata_questionnaire_store "github.com/kirigaikabuto/setdata-questionnaire-store"

type QuestionsService interface {
	CreateQuestion(cmd *CreateQuestionCommand) (*setdata_questionnaire_store.Question, error)
	UpdateQuestion(cmd *UpdateQuestionsCommand) (*setdata_questionnaire_store.Question, error)
	GetQuestion(cmd *GetQuestionCommand) (*setdata_questionnaire_store.Question, error)
	DeleteQuestion(cmd *DeleteQuestionCommand) error
	ListQuestions(cmd *ListQuestionsCommand) ([]setdata_questionnaire_store.Question, error)
}

type questionsService struct {
	amqpRequest AmqpRequests
}

func NewQuestionsService(requests AmqpRequests) QuestionsService {
	return &questionsService{amqpRequest: requests}
}

func (q *questionsService) CreateQuestion(cmd *CreateQuestionCommand) (*setdata_questionnaire_store.Question, error) {
	return q.amqpRequest.CreateQuestion(cmd)
}

func (q *questionsService) UpdateQuestion(cmd *UpdateQuestionsCommand) (*setdata_questionnaire_store.Question, error) {
	return q.amqpRequest.UpdateQuestion(cmd)
}

func (q *questionsService) GetQuestion(cmd *GetQuestionCommand) (*setdata_questionnaire_store.Question, error) {
	return q.amqpRequest.GetQuestion(cmd)
}

func (q *questionsService) DeleteQuestion(cmd *DeleteQuestionCommand) error {
	return q.amqpRequest.DeleteQuestion(cmd)
}

func (q *questionsService) ListQuestions(cmd *ListQuestionsCommand) ([]setdata_questionnaire_store.Question, error) {
	return q.amqpRequest.ListQuestions(cmd)
}
