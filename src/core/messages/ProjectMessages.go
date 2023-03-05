package messages

const (
	Project            = "Projeto"
	ProjectID          = "ID do projeto"
	ProjectName        = "Nome do projeto"
	ProjectDescription = "Descrição do projeto"
	ProjectIsActive    = "O projeto está ativo?"
	ProjectIsDeleted   = "O projeto foi deletado?"
	ProjectCreatedAt   = "Data de criação do projeto"
	ProjectUpdatedAt   = "Data de atualização do projeto"

	InvalidProjectErrorMessage            = "O projeto informado é inválido!"
	InvalidProjectIDErrorMessage          = "O ID do projeto é inválido!"
	InvalidProjectNameErrorMessage        = "O nome do projeto é inválido!"
	InvalidProjectDescriptionErrorMessage = "A descrição do projeto é inválida!"
	InvalidProjectIsActiveErrorMessage    = "O valor do atributo \"projeto está ativo?\" é inválido!"
	InvalidProjectIsDeletedErrorMessage   = "O valor do atributo \"projeto foi deletado?\" é inválido!"
	InvalidProjectCreatedAtErrorMessage   = "A data de criação do projeto é inválida!"
	InvalidProjectUpdatedAtErrorMessage   = "A data de atualização do projeto é inválida!"

	DuplicatedProjectIDErrorMessage   = "Um projeto com esse mesmo identificador já está cadastrado no banco de dados!"
	DuplicatedProjectNameErrorMessage = "Um projeto com esse mesmo nome já está cadastrado no banco de dados!"
)
