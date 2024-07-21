package pkg

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type PersonSessionWithContext interface {
	BeginTransaction(ctx context.Context, configurers ...func(*neo4j.TransactionConfig)) (neo4j.ExplicitTransaction, error)
	// ExecuteRead executes the given unit of work in a AccessModeRead transaction with
	// retry logic in place
	// Contexts terminating too early negatively affect connection pooling and degrade the driver performance.
	ExecuteRead(ctx context.Context, work neo4j.ManagedTransactionWork, configurers ...func(*neo4j.TransactionConfig)) (any, error)
	// ExecuteWrite executes the given unit of work in a AccessModeWrite transaction with
	// retry logic in place
	// Contexts terminating too early negatively affect connection pooling and degrade the driver performance.
	ExecuteWrite(ctx context.Context, work neo4j.ManagedTransactionWork, configurers ...func(*neo4j.TransactionConfig)) (any, error)
	// Run executes an auto-commit statement and returns a result
	// Contexts terminating too early negatively affect connection pooling and degrade the driver performance.
	Run(ctx context.Context, cypher string, params map[string]any, configurers ...func(*neo4j.TransactionConfig)) (neo4j.ResultWithContext, error)
	// Close closes any open resources and marks this session as unusable
	// Contexts terminating too early negatively affect connection pooling and degrade the driver performance.
	Close(ctx context.Context) error
}

type ManagedTransaction interface {
	Run(ctx context.Context, cypher string, params map[string]any) (neo4j.ResultWithContext, error)
	legacy() neo4j.Transaction
}
