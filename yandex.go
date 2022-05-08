package bots

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/temoon/telegram-bots-api"
)

type YandexRequest struct {
	Body string `json:"body"`
}

type YandexResponse struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

//goland:noinspection GoUnusedExportedFunction
func YandexHandler(ctx context.Context, req *YandexRequest, h BaseHandler) (res *YandexResponse, err error) {
	var update telegram.Update
	if err = json.Unmarshal([]byte(req.Body), &update); err != nil {
		return &YandexResponse{StatusCode: http.StatusBadRequest}, err
	}

	f := Frame{
		Handler: h,
	}

	if err = f.onUpdate(ctx, &update); err != nil {
		log.WithError(err).Error("on update")

		var errWithStatusCode *ErrorWithStatusCode
		if errors.As(err, &errWithStatusCode) {
			return &YandexResponse{StatusCode: errWithStatusCode.StatusCode}, nil
		} else {
			return &YandexResponse{StatusCode: http.StatusInternalServerError}, err
		}
	}

	return &YandexResponse{StatusCode: http.StatusOK, Body: "OK"}, nil
}
