package apierrors

import (
	"github.com/taidevops/grafana/pkg/api/response"
	"github.com/taidevops/grafana/pkg/models"
)

// ToFolderErrorResponse returns a different response status according to the folder error type
func ToFolderErrorResponse(err error) response.Response {
	var dashboardErr models.Das
}
