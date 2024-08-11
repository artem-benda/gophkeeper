package service

import (
	"context"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"google.golang.org/protobuf/proto"

	"golang.org/x/crypto/pbkdf2"

	"github.com/artem-benda/gophkeeper/client/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/client/internal/domain/entity"
)

// Реализация интерфейса contract.SecretService
type secret struct {
	r          contract.SecretRepository
	PassPhrase string
}

// NewSecretService - создать экземпляр с интерфейсом contract.SecretService
func NewSecretService(r contract.SecretRepository, passPhrase string) contract.SecretService {
	return &secret{r: r, PassPhrase: passPhrase}
}

// AddLoginPassword добавить логин-пароль
func (s *secret) AddLoginPassword(ctx context.Context, name string, login string, password string, metadata string) (string, error) {
	model := &entity.SecretLoginPasswordPayload{Login: login, Password: password}
	payload := &entity.SecretPayload{
		Metadata: metadata,
		Secret:   &entity.SecretPayload_LoginPassword{LoginPassword: model},
	}
	binary, err := proto.Marshal(payload)
	if err != nil {
		return "", err
	}
	encPayload := encrypt(s.PassPhrase, binary)
	return s.r.Add(ctx, name, encPayload)
}

// AddText добавить текст
func (s *secret) AddText(ctx context.Context, name string, text string, metadata string) (string, error) {
	model := &entity.SecretTextPayload{Text: text}
	payload := &entity.SecretPayload{
		Metadata: metadata,
		Secret:   &entity.SecretPayload_Text{Text: model},
	}
	binary, err := proto.Marshal(payload)
	if err != nil {
		return "", err
	}
	encPayload := encrypt(s.PassPhrase, binary)
	return s.r.Add(ctx, name, encPayload)
}

// AddBinary добавить бинарную информацию
func (s *secret) AddBinary(ctx context.Context, name string, data []byte, metadata string) (string, error) {
	model := &entity.SecretBinaryPayload{Binary: data}
	payload := &entity.SecretPayload{
		Metadata: metadata,
		Secret:   &entity.SecretPayload_Binary{Binary: model},
	}
	binary, err := proto.Marshal(payload)
	if err != nil {
		return "", err
	}
	encPayload := encrypt(s.PassPhrase, binary)
	return s.r.Add(ctx, name, encPayload)
}

// AddBankingCard добавить банковскую карту
func (s *secret) AddBankingCard(ctx context.Context, name string, number string, owner string, dueTo string, cvv string, metadata string) (string, error) {
	model := &entity.SecretBankingCardPayload{CardNumber: number, OwnerName: owner, ValidThru: dueTo, CVV: cvv}
	payload := &entity.SecretPayload{
		Metadata: metadata,
		Secret:   &entity.SecretPayload_BankingCard{BankingCard: model},
	}
	binary, err := proto.Marshal(payload)
	if err != nil {
		return "", err
	}
	encPayload := encrypt(s.PassPhrase, binary)
	return s.r.Add(ctx, name, encPayload)
}

// ReplaceWithLoginPassword заменить на логин-пароль
func (s *secret) ReplaceWithLoginPassword(ctx context.Context, guid string, name string, login string, password string, metadata string) error {
	model := &entity.SecretLoginPasswordPayload{Login: login, Password: password}
	payload := &entity.SecretPayload{
		Metadata: metadata,
		Secret:   &entity.SecretPayload_LoginPassword{LoginPassword: model},
	}
	binary, err := proto.Marshal(payload)
	if err != nil {
		return err
	}
	encPayload := encrypt(s.PassPhrase, binary)
	return s.r.Edit(ctx, guid, name, encPayload)
}

// ReplaceWithText заменить на текст
func (s *secret) ReplaceWithText(ctx context.Context, guid string, name string, text string, metadata string) error {
	model := &entity.SecretTextPayload{Text: text}
	payload := &entity.SecretPayload{
		Metadata: metadata,
		Secret:   &entity.SecretPayload_Text{Text: model},
	}
	binary, err := proto.Marshal(payload)
	if err != nil {
		return err
	}
	encPayload := encrypt(s.PassPhrase, binary)
	return s.r.Edit(ctx, guid, name, encPayload)
}

// ReplaceWithBinary заменить на бинарную информацию
func (s *secret) ReplaceWithBinary(ctx context.Context, guid string, name string, data []byte, metadata string) error {
	model := &entity.SecretBinaryPayload{Binary: data}
	payload := &entity.SecretPayload{
		Metadata: metadata,
		Secret:   &entity.SecretPayload_Binary{Binary: model},
	}
	binary, err := proto.Marshal(payload)
	if err != nil {
		return err
	}
	encPayload := encrypt(s.PassPhrase, binary)
	return s.r.Edit(ctx, guid, name, encPayload)
}

// ReplaceWithBankingCard заменить на банковскую карту
func (s *secret) ReplaceWithBankingCard(ctx context.Context, guid string, name string, number string, owner string, dueTo string, cvv string, metadata string) error {
	model := &entity.SecretBankingCardPayload{CardNumber: number, OwnerName: owner, ValidThru: dueTo, CVV: cvv}
	payload := &entity.SecretPayload{
		Metadata: metadata,
		Secret:   &entity.SecretPayload_BankingCard{BankingCard: model},
	}
	binary, err := proto.Marshal(payload)
	if err != nil {
		return err
	}
	encPayload := encrypt(s.PassPhrase, binary)
	return s.r.Edit(ctx, guid, name, encPayload)
}

// Remove удалить секрет
func (s *secret) Remove(ctx context.Context, guid string) error {
	return s.r.Remove(ctx, guid)
}

// GetAll получить все секреты
func (s *secret) GetAll(ctx context.Context) ([]entity.Secret, error) {
	secretsEncrypted, err := s.r.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	secrets := make([]entity.Secret, len(secretsEncrypted))
	for i, e := range secretsEncrypted {
		payloadBinary := decrypt(s.PassPhrase, e.EncPayload)
		payload := new(entity.SecretPayload)
		err := proto.Unmarshal(payloadBinary, payload)
		if err != nil {
			return nil, err
		}
		secrets[i] = entity.Secret{GUID: e.GUID, Name: e.Name, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt, Payload: payload}
	}
	return secrets, nil
}

// GetByGUID получить секрет по GUID
func (s *secret) GetByGUID(ctx context.Context, guid string) (*entity.Secret, error) {
	e, err := s.r.GetByGUID(ctx, guid)
	if err != nil {
		return nil, err
	}
	payloadBinary := decrypt(s.PassPhrase, e.EncPayload)
	payload := new(entity.SecretPayload)
	err = proto.Unmarshal(payloadBinary, payload)
	if err != nil {
		return nil, err
	}
	secret := entity.Secret{GUID: e.GUID, Name: e.Name, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt, Payload: payload}
	return &secret, nil
}

// deriveKey - получить ключ на основе passphrase с использованием алгоритма pbkdf2
func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}

// encrypt - зашифровать данные
func encrypt(passphrase string, binary []byte) []byte {
	key, salt := deriveKey(passphrase, nil)
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, []byte(binary), nil)
	encString := hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data)
	return []byte(encString)
}

// decrypt - расшифровать данные
func decrypt(passPhrase string, encBinary []byte) []byte {
	encString := string(encBinary)
	arr := strings.Split(encString, "-")
	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	data, _ := hex.DecodeString(arr[2])
	key, _ := deriveKey(passPhrase, salt)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data, _ = aesgcm.Open(nil, iv, data, nil)
	return data
}
