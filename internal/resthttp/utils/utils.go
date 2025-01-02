package utils

import (
	"context"
	"net/http"

	raas "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/respond"
	"github.com/raas-app/stocks/internal/resthttp/dto"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

const (
	AuthorizationHeader     = "Authorization"
	UserAgentHeader         = "User-Agent"
	ContentTypeHeader       = "Content-Type"
	XForwardedForHeaderName = "X-Forwarded-For"
	XFingerprintHeaderName  = "X-Fingerprint"
)

func GetForwardedIP(r *http.Request) string {
	return r.Header.Get(XForwardedForHeaderName)
}

func GetShieldID(r *http.Request) string {
	header := r.Header.Get(XFingerprintHeaderName)
	body := struct {
		ShieldSessionID string `json:"shield_session_id"`
	}{}
	_ = jsoniter.Unmarshal([]byte(header), &body)
	return body.ShieldSessionID
}

func SendLogicError(
	ctx context.Context,
	err *raas.LogicError,
	debugID string,
	logger *zap.Logger,
	responder *respond.Responder,
	w http.ResponseWriter,
) {
	var action *dto.MetaAction

	switch err.Meta.Code {
	case http.StatusBadRequest:
		SendBadRequest(ctx, err, debugID, logger, responder, w)
		return
	case http.StatusNotFound, raas.ReasonErrorCode, http.StatusForbidden:
		action = &dto.MetaAction{
			ViewType: dto.MetaActionViewTypeToast,
		}
		if err.Meta.Action != nil {
			action = dto.NewMetaAction(err.Meta.Action)
		}
	}

	httpErr := &dto.ResponseError{
		Meta: &dto.Meta{
			Code:    err.Meta.Code,
			Message: ctx.Err().Error(),
			DebugID: debugID,
			Action:  action,
			Reason:  err.Meta.Reason.String(),
		},
	}
	responder.WriteResponse(w, httpErr, httpErr.Meta.Code)
}

func SendBadRequest(ctx context.Context, err error, debugID string, logger *zap.Logger, responder *respond.Responder, w http.ResponseWriter) {
	meta := &dto.Meta{
		Code:    http.StatusBadRequest,
		Message: getErrorMessage(ctx, err),
		DebugID: debugID,
		Action:  dto.MetaActionBadRequest,
	}
	responseError := &dto.ResponseError{
		Meta: meta,
	}
	allFields := []zap.Field{zap.Error(responseError), zap.String("request_id", debugID)}
	if logger != nil {
		logger.Debug("bad request", allFields...)
	}
	responder.BadRequest(w, responseError)
}

func SendInternalServerError(ctx context.Context, err error, debugID string, logger *zap.Logger, responder *respond.Responder, w http.ResponseWriter, fields ...zap.Field) {
	meta := &dto.Meta{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
		DebugID: debugID,
	}
	responseError := &dto.ResponseError{
		Meta: meta,
	}
	allFields := []zap.Field{zap.Error(responseError), zap.String("request_id", debugID)}
	allFields = append(allFields, fields...)

	if !raas.IsContextCanceled(ctx) {
		if logger != nil {
			logger.Error("internal server error", allFields...)
		}
	}

	responder.InternalServerError(w, responseError)
}

func getErrorMessage(ctx context.Context, err error) string {
	var nErr = err.Error()
	if err == nil {
		nErr = ctx.Err().Error()
	}

	return nErr
}
