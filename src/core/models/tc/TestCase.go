package tc

import (
	"github.com/felipesantoos/tcms/src/core/models/tc/tcs"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type TestCase struct {
	id            uuid.UUID
	createdAt     time.Time
	updatedAt     time.Time
	deletedAt     time.Time
	code          uint64
	version       uint64
	title         string
	summary       string
	importance    string
	executionType string
	precondition  string
	isActive      bool
	steps         []tcs.TestCaseStep
}

func (instance *TestCase) ID() uuid.UUID {
	return instance.id
}

func (instance *TestCase) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *TestCase) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *TestCase) DeletedAt() time.Time {
	return instance.deletedAt
}

func (instance *TestCase) Code() uint64 {
	return instance.code
}

func (instance *TestCase) Version() uint64 {
	return instance.version
}

func (instance *TestCase) Title() string {
	return instance.title
}

func (instance *TestCase) Summary() string {
	return instance.summary
}

func (instance *TestCase) Importance() string {
	return instance.importance
}

func (instance *TestCase) ExecutionType() string {
	return instance.executionType
}

func (instance *TestCase) Precondition() string {
	return instance.precondition
}

func (instance *TestCase) IsActive() bool {
	return instance.isActive
}

func (instance *TestCase) Steps() []tcs.TestCaseStep {
	return instance.steps
}

func (instance *TestCase) IsZero() bool {
	return reflect.DeepEqual(instance, &TestCase{})
}
