package setdata_questionnaire

import (
	"encoding/json"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"io/ioutil"
	"net/http"
)

type HttpEndpoints interface {
	MakeCreateQuestionEndpoint() func(w http.ResponseWriter, r *http.Request)
	MakeListQuestionEndpoint() func(w http.ResponseWriter, r *http.Request)
	MakeUpdateQuestionEndpoint() func(w http.ResponseWriter, r *http.Request)
}

type httpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewHttpEndpoints(ch setdata_common.CommandHandler) HttpEndpoints {
	return &httpEndpoints{ch: ch}
}

func (h *httpEndpoints) MakeCreateQuestionEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := &CreateQuestionCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToAmqpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToAmqpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToAmqpResponse(err))
			return
		}
		respondJSON(w, http.StatusCreated, response)
	}
}

func (h *httpEndpoints) MakeListQuestionEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := &ListQuestionsCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToAmqpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToAmqpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToAmqpResponse(err))
			return
		}
		respondJSON(w, http.StatusCreated, response)
	}
}

func (h *httpEndpoints) MakeUpdateQuestionEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := &UpdateQuestionsCommand{}
		cmd.Id = r.URL.Query().Get("id")
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToAmqpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToAmqpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToAmqpResponse(err))
			return
		}
		respondJSON(w, http.StatusCreated, response)
	}
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
