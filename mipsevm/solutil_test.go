package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSourcemap(t *testing.T) {
	contract, err := LoadContract("Hello")
	require.NoError(t, err)
	srcMap, err := contract.SourceMap([]string{"../contracts/src/Hello.sol"})
	require.NoError(t, err)
	for i := 0; i < len(contract.DeployedBytecode.Object); i++ {
		fmt.Println(srcMap.FormattedInfo(uint64(i)) + ": test")
	}
}
