package postgres

import (
	"context"
	"github.com/google/uuid"
	"strings"
	"tcms/src/core/domain/project"
	"tcms/src/core/errors"
	"tcms/src/core/interfaces/repository"
	"tcms/src/core/messages"
	"tcms/src/infra/postgres/connection"
	"tcms/src/infra/postgres/converters"
	repositoryMessages "tcms/src/infra/postgres/messages"
	"tcms/src/infra/postgres/queries/connector"
)

var _ repository.ProjectLoader = ProjectPostgresRepository{}

type ProjectPostgresRepository struct {
	connection.Manager
}

func (instance ProjectPostgresRepository) CreateProject(projectInstance project.Project) (*uuid.UUID, errors.Error) {
	connectionInstance, err := instance.GetConnection()
	if err != nil {
		return nil, errors.NewUnavailableServiceError(messages.UnavailableDatabaseErrorMessage, err)
	}
	defer instance.CloseConnection(connectionInstance)

	queries := connector.New(connectionInstance)
	params := connector.InsertIntoProjectParams{
		Name:        projectInstance.Name(),
		Description: converters.StringToNullString(projectInstance.Description()),
	}
	id, err := queries.InsertIntoProject(context.Background(), params)
	if err != nil {
		return nil, instance.handleError(err)
	}

	return &id, nil
}

func (ProjectPostgresRepository) handleError(err error) errors.Error {
	message := err.Error()

	if strings.Contains(message, repositoryMessages.SQLNoRowsInResultSetErrorMessage) {
		return errors.NewNotFoundError(messages.NoRecordsWereFoundInTheDatabase, err)
	} else if strings.Contains(message, repositoryMessages.PKProjectID) {
		return errors.NewConflictError(messages.DuplicatedProjectIDErrorMessage)
	} else if strings.Contains(message, repositoryMessages.UKProjectName) {
		return errors.NewConflictError(messages.DuplicatedProjectNameErrorMessage)
	}

	return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
}
