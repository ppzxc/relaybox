package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"relaybox/internal/application/port/input"
	"relaybox/internal/domain"
)

type Handler struct {
	receiveUC input.ReceiveMessageUseCase
	getUC     input.GetMessageUseCase
	listUC    input.ListMessagesUseCase
	requeueUC input.RequeueMessageUseCase
	configUC  input.ConfigQueryUseCase
}

func NewHandler(
	receiveUC input.ReceiveMessageUseCase,
	getUC input.GetMessageUseCase,
	listUC input.ListMessagesUseCase,
	requeueUC input.RequeueMessageUseCase,
	configUC input.ConfigQueryUseCase,
) *Handler {
	return &Handler{
		receiveUC: receiveUC,
		getUC:     getUC,
		listUC:    listUC,
		requeueUC: requeueUC,
		configUC:  configUC,
	}
}

func (h *Handler) PostMessage(w http.ResponseWriter, r *http.Request) {
	inputID := chi.URLParam(r, "inputId")
	resolvedInputID := inputIDFromContext(r.Context())

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1MB
	body, err := io.ReadAll(r.Body)
	if err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			writeError(w, r, http.StatusRequestEntityTooLarge, "Payload Too Large", "request body exceeds 1MB limit")
		} else {
			writeError(w, r, http.StatusBadRequest, "Bad Request", "failed to read body")
		}
		return
	}

	messageID, err := h.receiveUC.Receive(r.Context(), resolvedInputID, r.Header.Get("Content-Type"), body)
	if err != nil {
		mapError(w, r, err)
		return
	}

	resp := map[string]any{
		"id":        messageID,
		"inputId":   inputID,
		"status":    string(domain.MessageStatusPending),
		"createdAt": time.Now().UTC().Format(time.RFC3339),
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprintf("/inputs/%s/messages/%s", inputID, messageID))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetMessage(w http.ResponseWriter, r *http.Request) {
	messageID := chi.URLParam(r, "messageId")
	msg, err := h.getUC.GetByID(r.Context(), messageID)
	if err != nil {
		mapError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) ListMessages(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusNotImplemented, "Not Implemented", "not implemented")
}

func (h *Handler) PatchMessage(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusNotImplemented, "Not Implemented", "not implemented")
}

func (h *Handler) ListInputs(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusNotImplemented, "Not Implemented", "not implemented")
}

func (h *Handler) GetInput(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusNotImplemented, "Not Implemented", "not implemented")
}

func (h *Handler) ListOutputs(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusNotImplemented, "Not Implemented", "not implemented")
}

func (h *Handler) GetOutput(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusNotImplemented, "Not Implemented", "not implemented")
}
