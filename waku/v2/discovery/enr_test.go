package discovery

import (
	"testing"

	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/stretchr/testify/require"
)

func TestEnodeToMultiAddr(t *testing.T) {
	enr := "enr:-IS4QAmC_o1PMi5DbR4Bh4oHVyQunZblg4bTaottPtBodAhJZvxVlWW-4rXITPNg4mwJ8cW__D9FBDc9N4mdhyMqB-EBgmlkgnY0gmlwhIbRi9KJc2VjcDI1NmsxoQOevTdO6jvv3fRruxguKR-3Ge4bcFsLeAIWEDjrfaigNoN0Y3CCdl8"
	parsedNode := enode.MustParse(enr)
	expectedMultiAddr := "/ip4/134.209.139.210/tcp/30303/p2p/16Uiu2HAmPLe7Mzm8TsYUubgCAW1aJoeFScxrLj8ppHFivPo97bUZ"
	actualMultiAddr, err := enodeToMultiAddr(parsedNode)

	require.NoError(t, err)
	require.Equal(t, expectedMultiAddr, actualMultiAddr)
}

// TestRetrieveNodes uses a live connection, so it could be
// flaky, it should though pay for itself and should be fairly stable
func TestRetrieveNodes(t *testing.T) {
	url := "enrtree://AOFTICU2XWDULNLZGRMQS4RIZPAZEHYMV4FYHAPW563HNRAOERP7C@test.nodes.vac.dev"

	nodes, err := RetrieveNodes(url)
	require.NoError(t, err)
	require.NotEmpty(t, nodes)
}
