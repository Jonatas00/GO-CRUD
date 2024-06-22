package model

type User struct {
	ID           uint   `json:"id"`
	NomeCompleto string `json:"nome_completo"`
	Email        string `json:"email"`
	CPF          string `json:"cpf"`
	Senha        string `json:"senha"`
}
