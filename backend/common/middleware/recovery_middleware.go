package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/RustamSafiulin/3d_reconstruction_service/common/helpers"
)

func PanicRecoveryMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// log the error
				logrus.Error(fmt.Sprint(rec))

				// write the error response
				errorResponse := map[string]interface{}{
					"error": "Internal Error",
				}

				body, _ := json.Marshal(errorResponse)

				helpers.WriteJSONResponse(w, 500, body)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
