package setdata_questionnaire

import (
	"encoding/json"
	"github.com/gorilla/mux"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"io/ioutil"
	"net/http"
)

type HttpEndpoints interface {
	MakeCreateQuestionEndpoint() func(w http.ResponseWriter, r *http.Request)
	MakeListQuestionEndpoint() func(w http.ResponseWriter, r *http.Request)
	MakeUpdateQuestionEndpoint() func(w http.ResponseWriter, r *http.Request)
	MakeAddFieldToQuestionEndpoint() func(w http.ResponseWriter, r *http.Request)

	MakeCreateQuestionnaireEndpoint() func(w http.ResponseWriter, r *http.Request)
	MakeListQuestionnaireEndpoint() func(w http.ResponseWriter, r *http.Request)
	MakeAddQuestionToQuestionnaireEndpoint() func(w http.ResponseWriter, r *http.Request)
	MakeDeleteQuestionFromQuestionnaireEndpoint() func(w http.ResponseWriter, r *http.Request)
	MakeGetQuestionnaireByName(paramName string) func(w http.ResponseWriter, r *http.Request)
	MakeGetQuestionsByQuestionnaireName(paramName string) func(w http.ResponseWriter, r *http.Request)

	MakeCreateOrderEndpoint() func(w http.ResponseWriter, r *http.Request)
	MakeListOrderEndpoint(paramName string) func(w http.ResponseWriter, r *http.Request)
	MakeSendOrderForConsultation() func(w http.ResponseWriter, r *http.Request)
	MakeSendOrderEmail() func(w http.ResponseWriter, r *http.Request)
}

type httpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewHttpEndpoints(ch setdata_common.CommandHandler) HttpEndpoints {
	return &httpEndpoints{ch: ch}
}

func (h *httpEndpoints) MakeCreateQuestionEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		cmd := &CreateQuestionCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusCreated, response)
	}
}

func (h *httpEndpoints) MakeListQuestionEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		cmd := &ListQuestionsCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusCreated, response)
	}
}

func (h *httpEndpoints) MakeUpdateQuestionEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		cmd := &UpdateQuestionsCommand{}
		id := r.URL.Query().Get("id")
		cmd.Id = id
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeAddFieldToQuestionEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		cmd := &AddFieldToQuestionCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeCreateQuestionnaireEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		cmd := &CreateQuestionnaireCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeListQuestionnaireEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		cmd := &ListQuestionnaireCommand{}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeAddQuestionToQuestionnaireEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		cmd := &AddQuestionToQuestionnaireCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeDeleteQuestionFromQuestionnaireEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		cmd := &RemoveQuestionFromQuestionnaireCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeGetQuestionnaireByName(paramName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		params := mux.Vars(r)
		cmd := &GetQuestionnaireByNameCommand{}
		cmd.Name = params[paramName]
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeGetQuestionsByQuestionnaireName(paramName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		params := mux.Vars(r)
		cmd := &GetQuestionsByQuestionnaireNameCommand{}
		cmd.Name = params[paramName]
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeCreateOrderEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		cmd := &CreateOrderCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeListOrderEndpoint(paramName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		cmd := &ListOrderCommand{}
		cmd.QuestionnaireName = params[paramName]
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeSendOrderForConsultation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		cmd := &SendOrderForConsultationCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusCreated, response)
	}
}

func (h *httpEndpoints) MakeSendOrderEmail() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		cmd := &SendOrderEmailCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin")
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
