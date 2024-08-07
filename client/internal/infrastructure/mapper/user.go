package mapper

import (
	"github.com/artem-benda/gophkeeper/client/internal/domain/contract"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MapUserError converts error to domain error
func MapUserError(grpcError error) error {
	if grpcError == nil {
		return nil
	}
	if e, ok := status.FromError(grpcError); ok {
		switch e.Code() {
		case codes.AlreadyExists:
			return contract.ErrUserAlreadyExists
		case codes.InvalidArgument:
			return contract.ErrUserInvalidCredentials
		default:
			return grpcError
		}
	} else {
		return grpcError
	}
}
