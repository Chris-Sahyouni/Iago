package cli

/* ---------------------------- Integration Tests --------------------------- */

import (
	"encoding/binary"
	"github.com/Chris-Sahyouni/iago/global"
	"os"
	"testing"
)

func TestCommandSequence(t *testing.T) {

	t.Cleanup(cleanup)

	workingDir := "../../"

	loadCmd := Load{
		args: Args{
			"default": workingDir + "test_programs/bin/square32",
		},
	}
	setTargetCmd := SetTarget{
		args: Args{
			"default": workingDir + "test_programs/test_targets/square",
		},
	}
	ropCmd := Rop{
		args: Args{
			"-o": workingDir + "cli_test_rop_chain",
		},
	}

	var testGlobalState global.GlobalState
	var err error
	err = loadCmd.Execute(&testGlobalState)
	if err != nil {
		t.Errorf("error during load square32: %s", err)
	}
	err = setTargetCmd.Execute(&testGlobalState)
	if err != nil {
		t.Errorf("error during set-target square: %s", err)
	}
	err = ropCmd.Execute(&testGlobalState)
	if err != nil {
		t.Errorf("error during rop: %s", err)
	}

	actual, err := os.ReadFile(workingDir + "cli_test_rop_chain")
	if err != nil {
		t.Error("could not open rop's output file")
	}

	// note that this only tests to ensure the generated chain is the same
	// as the one written to the file, not that the chain itself is correct

	expected := testGlobalState.CurrentPayload.Chain

	if len(expected) != 4*len(actual) {
		t.Error("chain in file different length than chain generated")
	}

	for i := range len(expected) {
		if binary.LittleEndian.Uint32(actual[i*4:(i*5)]) != uint32(expected[i]) {
			t.Error("chain in file differs from chain generated")
		}
	}
}

func cleanup() {
	os.Remove("../../cli_test_rop_chain")
}
