# gophkeeper

## Второй дипломный проект курса "Продвинутый Go-разработчик" от Яндекс Практикум

GophKeeper представляет собой клиент-серверную систему, позволяющую пользователю надёжно и безопасно хранить логины, пароли, бинарные данные и прочую приватную информацию.

### Сервер реализует следующую бизнес-логику:

- регистрация, аутентификация и авторизация пользователей;
- хранение приватных данных;
- синхронизация данных между несколькими авторизованными клиентами одного владельца;
- передача приватных данных владельцу по запросу.

### Клиент реализует следующую бизнес-логику:

- аутентификация и авторизация пользователей на удалённом сервере;
- доступ к приватным данным по запросу.

### Особенности реализации

- Сервер и клиент - CLI приложения
- Используется GRPC в качестве протокола взаимодействия клиента и сервера
- Секретные данные шифруются на клиенте с помощью алгоритмов pbkdf2 на основе секретного пароля, а так же AESGCM

### Для генерации  моделей proto и GRPC-кода можно использовать следующие команды

protoc --go_out=server/ --go_opt=paths=import --go-grpc_out=server/ --go-grpc_opt=paths=import proto/gophkeeper.proto

protoc --go_out=client/ --go_opt=paths=import --go-grpc_out=client/ --go-grpc_opt=paths=import proto/gophkeeper.proto

protoc --go_out=client/ --go_opt=paths=import --go-grpc_out=client/ --go-grpc_opt=paths=import proto/secret_payload.proto

### Использование клиента

#### Просмотреть список доступных команд

`client -h`

```
Usage:
  client [OPTIONS] <command>

Application Options:
  -e, --endpoint=    address and port of server to connect to (default:
                     localhost:8080)
  -k, --pass-phrase= pass phrase used to encrypt secret information (default:
                     supersecretkey123)

Help Options:
  -h, --help         Show this help message

Available commands:
  add-secret-banking-card        Add secret banking card (aliases: abc)
  add-secret-binary              Add secret binary from file (aliases: ab)
  add-secret-login-password      Add secret login password (aliases: alp)
  add-secret-text                Add secret text (aliases: at)
  get-all-secrets                Get all secrets (aliases: ga)
  get-secret                     Get secret (aliases: g)
  register                       Register user (aliases: r)
  remove-secret                  Remove secret (aliases: rm)
  replace-secret-banking-card    Replace secret with banking card (aliases: rbc)
  replace-secret-binary          Replace secret with binary from file (aliases: rb)
  replace-secret-login-password  Replace secret with login password (aliases: rlp)
  replace-secret-text            Replace secret with text (aliases: rt)
```

#### Просмотреть список параметров конкретной команды

`client [command name] -h`

