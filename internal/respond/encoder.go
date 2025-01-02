package respond

import (
	"io"
	"net/http"
	"strconv"
	"sync/atomic"

	jsoniter "github.com/json-iterator/go"
)

var (
	// https://github.com/uber-go/guide/blob/master/style.md#verify-interface-compliance
	_    Encoder = (*jsonEncoder)(nil)
	JSON Encoder = &jsonEncoder{}
	json         = jsoniter.ConfigCompatibleWithStandardLibrary
)

type Encoder interface {
	Encode(w http.ResponseWriter, v interface{}) error
	ContentType() string
}

type jsonEncoder struct{}

type bytesCounterWriter struct {
	count atomic.Int32
}

func (r *bytesCounterWriter) Write(p []byte) (int, error) {
	r.count.Add(int32(len(p)))
	return len(p), nil
}

func (*jsonEncoder) Encode(w http.ResponseWriter, v interface{}) error {
	bytesCounter := bytesCounterWriter{}
	err := json.NewEncoder(io.MultiWriter(w, &bytesCounter)).Encode(v)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Length", strconv.Itoa(int(bytesCounter.count.Load())))

	return nil
}

func (*jsonEncoder) ContentType() string {
	return "application/json; charset=utf-8"
}
