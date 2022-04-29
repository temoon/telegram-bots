package bots

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/temoon/go-telegram-bots-api"
)

type YandexRequest struct {
	Body string `json:"body"`
}

type YandexResponse struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

//goland:noinspection GoUnusedExportedFunction
func YandexHandler(ctx context.Context, req *YandexRequest, router Router) (res *YandexResponse, err error) {
	var update telegram.Update
	if err = json.Unmarshal([]byte(req.Body), &update); err != nil {
		return &YandexResponse{StatusCode: http.StatusBadRequest}, err
	}

	h := Handler{}

	if err = h.onUpdate(ctx, &update, router); err != nil {
		log.WithError(err).Error("on update")

		var errWithStatusCode *ErrorWithStatusCode
		if errors.As(err, &errWithStatusCode) {
			res = &YandexResponse{StatusCode: errWithStatusCode.StatusCode}
		} else {
			res = &YandexResponse{StatusCode: http.StatusInternalServerError}
		}

		return
	}

	return &YandexResponse{StatusCode: http.StatusOK, Body: "OK"}, nil
}
