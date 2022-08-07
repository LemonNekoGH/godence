package godence

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/onflow/flow-go-sdk"
	flowGrpc "github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Helper for unsupport type test.
type unsupportType string

var flowCli *flowGrpc.Client

func initFlowClient() {
	client, err := flowGrpc.NewClient(
		"localhost:3569",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	flowCli = client
	fmt.Println("Flow client init success.")
}

func TestMain(m *testing.M) {
	// start: Init a flow client.
	initFlowClient()
	m.Run()
}

// waitForTransactionSealed. Only for test
func waitForTransactionSealed(tx *flow.Transaction, a *assert.Assertions) *flow.TransactionResult {
	for {
		result, err := flowCli.GetTransactionResult(context.Background(), tx.ID())
		a.NoError(err)
		if result.Error != nil {
			return result
		}
		if result.Status == flow.TransactionStatusSealed {
			return result
		}
		// sleep for 200ms
		time.Sleep(200 * time.Millisecond)
	}
}

func buildSimpleTx(script []byte, a *assert.Assertions) *flow.Transaction {
	latestBlock, err := flowCli.GetLatestBlock(context.Background(), true)
	a.NoError(err)

	account, err := flowCli.GetAccount(context.Background(), flow.HexToAddress("0xf8d6e0586b0a20c7"))
	a.NoError(err)

	tx := flow.NewTransaction().
		SetScript([]byte(script)).
		SetReferenceBlockID(latestBlock.ID).
		SetProposalKey(account.Address, 0, account.Keys[0].SequenceNumber).
		SetPayer(account.Address).
		SetGasLimit(999)

	privateKey, err := crypto.DecodePrivateKeyHex(crypto.ECDSA_P256, "c47db93881bc34a6155192c2bec0d124731e08ff105672afdb09892e3dc9ccae")
	a.NoError(err)

	signer, err := crypto.NewInMemorySigner(privateKey, crypto.SHA3_256)
	a.NoError(err)

	tx.SignEnvelope(account.Address, 0, signer)
	return tx
}
