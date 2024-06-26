package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"go-restful-fiber/model/dto"
	"go-restful-fiber/pkg"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	// Init Logger
	logger := pkg.NewLogger()

	// Return error
	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		logger.WithFields(logrus.Fields{
			"error":  err,
			"host":   request.Host,
			"method": request.Method,
			"uri":    request.RequestURI,
		}).Infof("Validation Error")
		return
	}

	logger.Error(err)
	internalServerError(writer, request, err)
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	apiResponse := dto.ApiResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	pkg.WriteToResponseBody(writer, apiResponse)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := dto.ApiResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		pkg.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := dto.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		pkg.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}
