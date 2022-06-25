package response

import (
	"bytes"
	"net/http"

	"github.com/taidevops/grafana/pkg/services/models"
)

type Response interface {
	WriteTo(ctx *models.ReqContext)

	Body() []byte

	Status() int
}

func CreateNormalResponse(header http.Header, body []byte, status int) *NormalResponse {

}

type NormalResponse struct {
	status     int
	body       *bytes.Buffer
	header     http.Header
	errMessage string
	err        error
}

func (r *NormalResponse) WriteTo(ctx *models.ReqContext) {
	if r.err != nil {
		v := map[string]interface{}{}

	}

	header := ctx.Resp.Header()
	for k, v := range r.header {
		header[k] = v
	}
	ctx.Resp.WriteHeader(r.status)
	if _, err := ctx.Resp.Write(r.body.Bytes()); err != nil {
		
	}
}
