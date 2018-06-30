//go:generate gqlgen -schema ../schema.graphql
package graph

import (
	"github.com/ChristianVestre/Spidey/account"
	"github.com/ChristianVestre/Spidey/catalog"
	"github.com/ChristianVestre/Spidey/order"
)

type GraphQLServer struct {
	accountClient *account.Client
	catalogClient *catalog.Client
	orderClient   *order.Client
}

func NewGraphQLServer(accountURL, catalogURL, orderURL string) (*GraphQLServer, error) {
	// Connect to account service
	accountClient, err := account.NewClient(accountURL)
	if err != nil {
		return nil, err
	}

	// connect to product service
	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		accountClient.Close()
		return nil, err
	}

	// connect to order service
	orderClient, err := order.NewClient(orderURL)
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return nil, err
	}

	return &GraphQLServer{
		accountClient,
		catalogClient,
		orderClient,
	}, nil
}
