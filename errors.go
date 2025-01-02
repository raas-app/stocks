package raas

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
)

type (
	// LogicError описывает ошибки в бизнес логике
	LogicError struct {
		Meta *Meta
	}
)

type Meta struct {
	LocaleMessageID string
	Code            int       `json:"code"`
	Message         string    `json:"message"`
	Reason          ErrReason `json:"reason,omitempty"`
	Action          *Action
}

type Action struct {
	ViewType ViewType
	SlugID   string
	Button   Button
}

type Button struct {
	Text string
	URL  string
}

type ViewType string

const (
	ViewTypeDialog = "dialog"
	ViewTypeBDUI   = "bdui"
	ViewTypeToast  = "toast"
)

type ErrReason string

const (
	ReasonErrorCode = 467
)

var (
	ErrRouteBadRequest = NewBadRequestErrorf("route is not valid")
)

func NewReasonedError(reason ErrReason, messageID string, defaultMessage string) *LogicError {
	return NewReasonedErrorWithAction(reason, nil, messageID, defaultMessage)
}

func NewReasonedErrorWithAction(reason ErrReason, action *Action, messageID string, defaultMessage string) *LogicError {
	meta := &Meta{
		LocaleMessageID: messageID,
		Message:         defaultMessage,
		Code:            ReasonErrorCode,
		Reason:          reason,
		Action:          action,
	}
	return &LogicError{Meta: meta}
}

func (r ErrReason) String() string {
	return string(r)
}

func (e LogicError) Error() string {
	if e.Meta.Message == "" {
		return "UNKNOWN ERROR"
	}
	return e.Meta.Message
}

var ErrQueueRetryable = fmt.Errorf("ErrQueueRetryable")

func NewRetryableQueueError(err error) error {
	if err == nil {
		return nil
	}
	return errors.Join(ErrQueueRetryable, err)
}

func NewNotFoundErrorf(format string, args ...interface{}) error {
	return newError(http.StatusNotFound, format, args...)
}
func NewBadRequestErrorf(format string, args ...interface{}) error {
	return newError(http.StatusBadRequest, format, args...)
}
func NewConflictErrorf(format string, args ...interface{}) error {
	return newError(http.StatusConflict, format, args...)
}

func NewInternalErrorf(format string, args ...interface{}) error {
	return newError(http.StatusInternalServerError, format, args...)
}

func NewForbiddenErrorf(format string, args ...interface{}) error {
	return newError(http.StatusForbidden, format, args...)
}

func NewLocaledError(code int, messageID string, defaultMessage string) error {
	meta := &Meta{
		LocaleMessageID: messageID,
		Code:            code,
		Message:         defaultMessage,
	}

	return &LogicError{
		Meta: meta,
	}
}

func NewExternalMsgMarshallNilError(modelName string) error {
	return newError(0, "domain model: `%s` - nil reference while marshalling model", modelName)
}

func newError(code int, format string, args ...interface{}) error {
	meta := &Meta{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}

	return &LogicError{
		Meta: meta,
	}
}

// Validator is implemented by request structures.
type Validator interface {
	Validate() error
}

// *** MySQL ***

const (
	MySQLErrDuplicateCode = 1062
)

func IsMysqlDuplicateError(err error) bool {
	var mysqlError *mysql.MySQLError
	if errors.As(err, &mysqlError) {
		if mysqlError.Number == MySQLErrDuplicateCode {
			return true
		}
	}
	return false
}

// *** Redis ***
func IsRedisNilError(err error) bool {
	return errors.Is(err, redis.Nil)
}
