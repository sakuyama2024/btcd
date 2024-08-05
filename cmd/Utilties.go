package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"log"
	"math/big"
	"reflect"
	"strings"
)

// Convert a C++ hash (big-endian hex string) to BTCD format (little-endian chainhash.Hash)
func ConvertToBTCDFormat(cppHash string) (chainhash.Hash, error) {
	// Remove the "0x" prefix if it exists
	cppHash = strings.TrimPrefix(cppHash, "0x")

	// Decode the hex string to a byte slice
	bytes, err := hex.DecodeString(cppHash)
	if err != nil {
		return chainhash.Hash{}, err
	}

	// Reverse the byte slice to convert to little-endian
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	// Convert the byte slice to a chainhash.Hash
	var hash chainhash.Hash
	copy(hash[:], bytes)

	return hash, nil
}

// Format the chainhash.Hash into the Go source code format
func FormatHashAsGoSource(hash chainhash.Hash) string {
	formatted := "chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.\n"
	for i, b := range hash {
		if i%8 == 0 {
			formatted += "\t"
		}
		formatted += fmt.Sprintf("0x%02x", b)
		if i < len(hash)-1 {
			formatted += ", "
		}
		if (i+1)%8 == 0 {
			formatted += "\n"
		}
	}
	formatted += "})"
	return formatted
}

func hashConversion() {
	//genesis block hash
	//cppHash := "0x0000000cd159482c9663a50e6a23a63155f9477384843473b784449b897569bf"

	//merkle root hash
	cppHash := "0xc61f9003735f01c77c4a8b3554b86b8bda7ce1f3854f1e657abfad6f49462614"

	// Convert the C++ hash to BTCD format
	btcdHash, err := ConvertToBTCDFormat(cppHash)
	if err != nil {
		log.Fatalf("Failed to convert hash: %v", err)
	}

	// Format the BTCD hash as Go source code
	formattedHash := FormatHashAsGoSource(btcdHash)

	// Print the formatted Go source code
	fmt.Println("var alphaMainGenesisHash =", formattedHash)
}

// Function to sum numbers from 1 to 10 and print the result
func main() {

	hashConversion()
	return
	// Create a new big.Int with an initial value of 12345
	bigInt1 := big.NewInt(12345)
	fmt.Println("bigInt1:", bigInt1)

	var nBits uint32 = 0x1d0fffff
	fmt.Println("nbits:", nBits)

	target := blockchain.CompactToBig(nBits)

	// Print the result
	fmt.Println("Difficulty bits:", nBits)
	fmt.Println("Target as big.Int:", target)

	// Print the target as a 64-character hex string with leading zeros
	targetHex := fmt.Sprintf("%064x", target)
	fmt.Println("Target as 64-character hex:", targetHex)

	var nBits2 uint32 = 0x1d00ffff
	fmt.Println("nbits2:", nBits2)

	target = blockchain.CompactToBig(nBits2)

	// Print the result
	fmt.Println("Difficulty bits:", nBits2)
	fmt.Println("Target as big.Int:", target)

	// Print the target as a 64-character hex string with leading zeros
	targetHex = fmt.Sprintf("%064x", target)
	fmt.Println("Target as 64-character hex:", targetHex)

	var t2 *big.Int = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 228), big.NewInt(1))
	// Print t2 as a 64-character hex string with leading zeros
	fmt.Printf("t2 as 64-character hex: %064x\n", t2)

	PowLimitHex, _ := hex.DecodeString(
		"00000377ae000000000000000000000000000000000000000000000000000000",
	)

	//print as bytes
	fmt.Println("PowLimitHex: ", PowLimitHex)

	// Encode back to a hex string and print
	hexString := hex.EncodeToString(PowLimitHex)
	fmt.Println("PowLimitHex (Hex String):", hexString)

	// Print the type of PowLimitHex
	fmt.Println("Type of PowLimitHex:", reflect.TypeOf(PowLimitHex))
}
