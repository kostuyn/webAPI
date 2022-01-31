package apperror

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(handler appHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var appErr *AppError
		err := handler(writer, request)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					writer.WriteHeader(http.StatusNotFound)
					writer.Write(ErrNotFound.Marshal())
					return
				}

				err = err.(*AppError)
				writer.WriteHeader(http.StatusBadRequest)
				writer.Write(appErr.Marshal())
				return
			}

			writer.WriteHeader(http.StatusTeapot)
			writer.Write(systemError(err).Marshal())
		}
	}
}
