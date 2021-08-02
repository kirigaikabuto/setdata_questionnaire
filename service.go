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
}

func NewQuestionsService() QuestionsService {
	return &questionsService{}
}

func (q *questionsService) CreateQuestion(cmd *CreateQuestionCommand) (*setdata_questionnaire_store.Question, error) {
	return nil, nil
}

func (q *questionsService) UpdateQuestion(cmd *UpdateQuestionsCommand) (*setdata_questionnaire_store.Question, error) {
	return nil, nil
}

func (q *questionsService) GetQuestion(cmd *GetQuestionCommand) (*setdata_questionnaire_store.Question, error) {
	return nil, nil
}

func (q *questionsService) DeleteQuestion(cmd *DeleteQuestionCommand) error {
	return nil
}

func (q *questionsService) ListQuestions(cmd *ListQuestionsCommand) ([]setdata_questionnaire_store.Question, error) {
	return nil, nil
}
