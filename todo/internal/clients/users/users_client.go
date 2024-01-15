package users

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"todo/config"
	"todo/internal/models"
	"todo/pkg/grpc_stubs/users"
)

type UsersClient struct {
	client users.UserServiceClient
}

func NewUsersClient(cfg *config.Config, logger *zerolog.Logger) (*UsersClient, error) {
	appAddr := fmt.Sprintf("%s:%s", cfg.UsersClient.AppHost, cfg.UsersClient.AppGrpcPort)

	logger.Info().Msgf("[NewUsersClient] connecting via GRPC to usres at %s", appAddr)

	conn, err := grpc.Dial(
		appAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("[NewApp] connect to users GRPC service: %w", err)
	}

	c := users.NewUserServiceClient(conn)
	return &UsersClient{
		client: c,
	}, nil
}

func (c *UsersClient) GetUserByID(ctx context.Context, userID int) (*models.UserDTO, error) {
	user, err := c.client.GetUserByID(ctx, &users.UserID{
		Id: int32(userID),
	})
	if err != nil {
		return nil, err
	}

	return models.NewEmptyUserDTO().FromGRPC(user), nil
}
