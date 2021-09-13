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

func(h *httpEndpoints) MakeSendOrderForConsultation() func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
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
		respondJSON(w, http.StatusOK, response)
	}
}

func(h *httpEndpoints) MakeSendOrderEmail() func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
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

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
