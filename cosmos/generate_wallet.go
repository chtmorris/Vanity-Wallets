package main

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var attempts int

func genAddress(inputLength int) string {

	cstore := keys.NewInMemory()

	info, seed, err := cstore.CreateMnemonic("john", keys.English, "secretcpw", keys.Secp256k1)
	if err != nil {
		panic(err)
	}

	address := sdk.AccAddress(info.GetPubKey().Address()).String()

	fmt.Println(address, seed)

	return strings.ToLower(address[0:inputLength])

}

func runIteration(generatedPubKey, inputText string) {
	if generatedPubKey == inputText {
		fmt.Println("Successful key generated ;-)")
	} else {
		attempts++
		fmt.Println("Attempt number:", attempts)
		runIteration(genAddress(len(inputText)), inputText)
	}
}

func main() {
	fmt.Print("What would you like it say after cosmos1:")
	var input string
	fmt.Scanln(&input)
	input = "cosmos1" + input
	var inputLength = len(input)

	runIteration(genAddress(inputLength), input)
}
