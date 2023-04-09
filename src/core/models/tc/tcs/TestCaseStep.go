package tcs

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type TestCaseStep struct {
	id             uuid.UUID
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      time.Time
	testCaseID     uuid.UUID
	description    string
	expectedResult string
}

func (instance *TestCaseStep) ID() uuid.UUID {
	return instance.id
}

func (instance *TestCaseStep) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *TestCaseStep) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *TestCaseStep) DeletedAt() time.Time {
	return instance.deletedAt
}

func (instance *TestCaseStep) TestCaseID() uuid.UUID {
	return instance.testCaseID
}

func (instance *TestCaseStep) Description() string {
	return instance.description
}

func (instance *TestCaseStep) ExpectedResult() string {
	return instance.expectedResult
}

func (instance *TestCaseStep) IsZero() bool {
	return reflect.DeepEqual(instance, &TestCaseStep{})
}
