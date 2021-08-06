package setdata_questionnaire

import (
	"fmt"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	setdata_questionnaire_store "github.com/kirigaikabuto/setdata-questionnaire-store"
	"strings"
)

type QuestionsService interface {
	CreateQuestion(cmd *CreateQuestionCommand) (*setdata_questionnaire_store.Question, error)
	UpdateQuestion(cmd *UpdateQuestionsCommand) (*setdata_questionnaire_store.Question, error)
	GetQuestion(cmd *GetQuestionCommand) (*setdata_questionnaire_store.Question, error)
	DeleteQuestion(cmd *DeleteQuestionCommand) error
	ListQuestions(cmd *ListQuestionsCommand) ([]setdata_questionnaire_store.Question, error)
	AddFieldToQuestion(cmd *AddFieldToQuestionCommand) (*setdata_questionnaire_store.Question, error)

	CreateQuestionnaire(cmd *CreateQuestionnaireCommand) (*setdata_questionnaire_store.Questionnaire, error)
	UpdateQuestionnaire(cmd *UpdateQuestionnaireCommand) (*setdata_questionnaire_store.Questionnaire, error)
	ListQuestionnaire(cmd *ListQuestionnaireCommand) ([]setdata_questionnaire_store.Questionnaire, error)
	GetQuestionnaireById(cmd *GetQuestionnaireByIdCommand) (*QuestionnaireDetail, error)
	GetQuestionnaireByName(cmd *GetQuestionnaireByNameCommand) (*QuestionnaireDetail, error)
	DeleteQuestionnaire(cmd *DeleteQuestionnaireByIdCommand) error
	AddQuestionToQuestionnaire(cmd *AddQuestionToQuestionnaireCommand) (*setdata_questionnaire_store.Questionnaire, error)
	RemoveQuestionFromQuestionnaire(cmd *RemoveQuestionFromQuestionnaireCommand) (*setdata_questionnaire_store.Questionnaire, error)
	GetQuestionsByQuestionnaireName(cmd *GetQuestionsByQuestionnaireNameCommand) ([]setdata_questionnaire_store.Question, error)

	CreateOrder(cmd *CreateOrderCommand) (*setdata_questionnaire_store.Order, error)
	ListOrder(cmd *ListOrderCommand) ([]setdata_questionnaire_store.Order, error)
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

func (q *questionsService) AddFieldToQuestion(cmd *AddFieldToQuestionCommand) (*setdata_questionnaire_store.Question, error) {
	question, err := q.GetQuestion(&GetQuestionCommand{Id: cmd.QuestionId})
	if err != nil {
		return nil, err
	}
	question.Fields = append(question.Fields, setdata_common.Field{
		Name:        cmd.Name,
		Type:        cmd.Type,
		Placeholder: cmd.Placeholder,
	})
	fmt.Println("cmd", cmd)
	fmt.Println(question)
	cmdUpdate := &UpdateQuestionsCommand{}
	cmdUpdate.Id = question.Id
	cmdUpdate.Fields = &question.Fields
	return q.UpdateQuestion(cmdUpdate)
}

func (q *questionsService) CreateQuestionnaire(cmd *CreateQuestionnaireCommand) (*setdata_questionnaire_store.Questionnaire, error) {
	for _, v := range cmd.Questions {
		_, err := q.amqpRequest.GetQuestion(&GetQuestionCommand{Id: v})
		if err != nil {
			return nil, err
		}
	}
	return q.amqpRequest.CreateQuestionnaire(cmd)
}

func (q *questionsService) UpdateQuestionnaire(cmd *UpdateQuestionnaireCommand) (*setdata_questionnaire_store.Questionnaire, error) {
	return q.amqpRequest.UpdateQuestionnaire(cmd)
}

func (q *questionsService) ListQuestionnaire(cmd *ListQuestionnaireCommand) ([]setdata_questionnaire_store.Questionnaire, error) {
	return q.amqpRequest.ListQuestionnaire(cmd)
}

func (q *questionsService) GetQuestionnaireById(cmd *GetQuestionnaireByIdCommand) (*QuestionnaireDetail, error) {
	questionnaire, err := q.amqpRequest.GetQuestionnaireById(cmd)
	if err != nil {
		return nil, err
	}
	response := &QuestionnaireDetail{
		Id:        questionnaire.Id,
		Name:      questionnaire.Name,
		Questions: nil,
	}
	questions := []setdata_questionnaire_store.Question{}
	for _, v := range questionnaire.Questions {
		q, err := q.amqpRequest.GetQuestion(&GetQuestionCommand{Id: v})
		if err != nil {
			return nil, err
		}
		questions = append(questions, *q)
	}
	response.Questions = questions
	return response, nil
}

func (q *questionsService) GetQuestionnaireByName(cmd *GetQuestionnaireByNameCommand) (*QuestionnaireDetail, error) {
	questionnaire, err := q.amqpRequest.GetQuestionnaireByName(cmd)
	if err != nil {
		return nil, err
	}
	response := &QuestionnaireDetail{
		Id:        questionnaire.Id,
		Name:      questionnaire.Name,
		Questions: nil,
	}
	questions := []setdata_questionnaire_store.Question{}
	for _, v := range questionnaire.Questions {
		q, err := q.amqpRequest.GetQuestion(&GetQuestionCommand{Id: v})
		if err != nil {
			return nil, err
		}
		questions = append(questions, *q)
	}
	response.Questions = questions
	return response, nil
}

func (q *questionsService) DeleteQuestionnaire(cmd *DeleteQuestionnaireByIdCommand) error {
	return q.amqpRequest.DeleteQuestionnaire(cmd)
}

func (q *questionsService) AddQuestionToQuestionnaire(cmd *AddQuestionToQuestionnaireCommand) (*setdata_questionnaire_store.Questionnaire, error) {
	questionnaire, err := q.amqpRequest.GetQuestionnaireById(&GetQuestionnaireByIdCommand{cmd.Id})
	if err != nil {
		return nil, err
	}
	questions := questionnaire.Questions
	_, err = q.GetQuestion(&GetQuestionCommand{Id: cmd.QuestionId})
	if err != nil {
		return nil, err
	}
	questions = append(questions, cmd.QuestionId)
	return q.UpdateQuestionnaire(&UpdateQuestionnaireCommand{
		Id:        cmd.Id,
		Name:      nil,
		Questions: &questions,
	})
}

func (q *questionsService) RemoveQuestionFromQuestionnaire(cmd *RemoveQuestionFromQuestionnaireCommand) (*setdata_questionnaire_store.Questionnaire, error) {
	questionnaire, err := q.amqpRequest.GetQuestionnaireById(&GetQuestionnaireByIdCommand{cmd.Id})
	if err != nil {
		return nil, err
	}
	questions := []string{}
	for _, v := range questionnaire.Questions {
		if v != cmd.QuestionId {
			questions = append(questions, v)
		}
	}
	return q.UpdateQuestionnaire(&UpdateQuestionnaireCommand{
		Id:        cmd.Id,
		Name:      nil,
		Questions: &questions,
	})
}

func (q *questionsService) GetQuestionsByQuestionnaireName(cmd *GetQuestionsByQuestionnaireNameCommand) ([]setdata_questionnaire_store.Question, error) {
	questionnaire, err := q.amqpRequest.GetQuestionnaireByName(&GetQuestionnaireByNameCommand{Name: cmd.Name})
	if err != nil {
		return nil, err
	}
	questions := []setdata_questionnaire_store.Question{}
	for _, v := range questionnaire.Questions {
		q, err := q.amqpRequest.GetQuestion(&GetQuestionCommand{Id: v})
		if err != nil {
			return nil, err
		}
		questions = append(questions, *q)
	}
	return questions, nil
}

func (q *questionsService) CreateOrder(cmd *CreateOrderCommand) (*setdata_questionnaire_store.Order, error) {
	order, err := q.amqpRequest.CreateOrder(cmd)
	if err != nil {
		return nil, err
	}
	text := ""
	for _, v := range cmd.QuestionnaireAnswers {
		text += fmt.Sprintf("<pre>%s<b>%s</b></pre> \n", v.QuestionName, strings.Join(v.Answers, ","))
	}
	message := fmt.Sprintf(`<strong>%s</strong>
	%s
	`, cmd.QuestionnaireName, text)
	err = q.amqpRequest.SendTelegramMessage(&setdata_questionnaire_store.SendMessageCommand{
		Message:   message,
		ParseMode: "HTML",
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (q *questionsService) ListOrder(cmd *ListOrderCommand) ([]setdata_questionnaire_store.Order, error) {
	return q.amqpRequest.ListOrder(cmd)
}
