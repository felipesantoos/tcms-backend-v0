package handlers

import (
	"net/http"
	"tcms/src/api/handlers/dto/response"
	"tcms/src/core/errors"

	"github.com/labstack/echo/v4"
)

func bindError(err errors.Error) response.Error {
	dto := response.Error{}

	dto.Message = err.FriendlyMessage()

	switch v := err.(type) {
	case *errors.ConflictError:
		dto.StatusCode = http.StatusConflict
		dto.DuplicatedFields = v.ConflictingFields()
	case *errors.UnexpectedError:
		dto.StatusCode = http.StatusInternalServerError
	case *errors.MissingInformationError:
		dto.StatusCode = http.StatusBadRequest
	case *errors.UnavailableServiceError:
		dto.StatusCode = http.StatusServiceUnavailable
	case *errors.NotFoundError:
		dto.StatusCode = http.StatusNotFound
	case *errors.ValidationError:
		dto.StatusCode = http.StatusUnprocessableEntity
		invalidFields := v.InvalidFields()
		for _, invalidField := range invalidFields {
			dto.InvalidFields = append(dto.InvalidFields, response.InvalidField{
				FieldName:   invalidField.Name,
				Description: invalidField.Description,
			})
		}
	}

	return dto
}

func getHttpHandledErrorResponse(context echo.Context, err errors.Error) error {
	handledErr := bindError(err)
	return context.JSON(handledErr.StatusCode, handledErr)
}
