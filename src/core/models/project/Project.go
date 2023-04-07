package project

import "github.com/google/uuid"

type Project struct {
	id          uuid.UUID
	name        string
	description string
}

func (project Project) ID() uuid.UUID {
	return project.id
}

func (project Project) Name() string {
	return project.name
}

func (project Project) Description() string {
	return project.description
}

func New(id uuid.UUID, name, description string) *Project {
	return &Project{
		id:          id,
		name:        name,
		description: description,
	}
}
