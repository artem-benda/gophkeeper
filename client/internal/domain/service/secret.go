package service

import "github.com/artem-benda/gophkeeper/client/internal/domain/contract"

type secret struct {
	r contract.SecretRepository
}

func NewSecretService(r contract.SecretRepository) contract.SecretService {
	return &secret{r}
}
