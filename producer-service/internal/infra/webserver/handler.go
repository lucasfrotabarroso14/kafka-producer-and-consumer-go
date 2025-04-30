package webserver

import (
	"async_kafka_services/producer-service/internal/domain/port"
	"context"
	"encoding/json"
	"net/http"
)

type Request struct {
	Message_to_kafka string `json:"message_to_kafka"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ProducerHandler struct {
	ExampleUseCase port.ExampleUseCase
}

func NewProducerHandler(exampleUseCase port.ExampleUseCase) *ProducerHandler {
	return &ProducerHandler{
		ExampleUseCase: exampleUseCase,
	}
}

func (h *ProducerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	inputUseCase := port.InputDTO{
		Message: request.Message_to_kafka,
	}

	code, err := h.ExampleUseCase.Execute(context.Background(), inputUseCase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Code:    code,
		Message: "Mensagem enviada com sucesso",
	}
	if err = json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

}
