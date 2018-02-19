package main

/*

#include <string.h>
#include <stdlib.h>

typedef unsigned char Ripemd160[20];

typedef struct {
	unsigned char Version;
	Ripemd160 Key;
} Address;

typedef unsigned char Checksum[4];

typedef unsigned char PubKey[33];
typedef unsigned char SecKey[32];

typedef unsigned char Pubkey[33];

*/
import "C"

import (
	"unsafe"

	"github.com/skycoin/skycoin/src/cipher"
)

/**
 * Functions in github.com/skycoin/skycoin/src/cipher/address.go
 */

//export SKY_Cipher_DecodeBase58Address
func SKY_Cipher_DecodeBase58Address(_strAddr string, _retAddr *C.Address) C.int {
	addr, err := cipher.DecodeBase58Address(_strAddr)
	*_retAddr = *(*C.Address)(unsafe.Pointer(&addr))

	errCode := 1
	if err != nil {
		errCode = 0
	}
	return C.int(errCode)
}

// export SKY_Cipher_AddressFromPubKey
func SKY_Cipher_AddressFromPubKey(_pubKey C.PubKey, _retAddr *C.Address) {
	var pubKey cipher.PubKey
	pubKey = *(*cipher.PubKey)(unsafe.Pointer(&_pubKey))
	addr := cipher.AddressFromPubKey(pubKey)
	*_retAddr = *(*C.Address)(unsafe.Pointer(&addr))
}

// export SKY_Cipher_AddressFromSecKey
func SKY_Cipher_AddressFromSecKey(_secKey C.SecKey, _retAddr *C.Address) {
	var secKey cipher.SecKey
	secKey = *(*cipher.SecKey)(unsafe.Pointer(&_secKey))
	addr := cipher.AddressFromSecKey(secKey)
	*_retAddr = *(*C.Address)(unsafe.Pointer(&addr))
}

// export SKY_Cipher_BitcoinDecodeBase58Address
func SKY_Cipher_BitcoinDecodeBase58Address(_strAddr string, _retAddr *C.Address) C.int {
	addr, err := cipher.BitcoinDecodeBase58Address(_strAddr)
	*_retAddr = *(*C.Address)(unsafe.Pointer(&addr))

	errCode := 1
	if err != nil {
		errCode = 0
	}
	return C.int(errCode)
}

// export SKY_Cipher_Address_Bytes
func SKY_Cipher_Address_Bytes(_addr *C.Address, _ret *C.uchar) {
	addr := (*cipher.Address)(unsafe.Pointer(&_addr))
	bytes := addr.Bytes()
	C.memcpy(unsafe.Pointer(_ret), unsafe.Pointer(&bytes[0]), C.size_t(len(bytes)))
}

// export SKY_Cipher_BitcoinBytes
func SKY_Cipher_BitcoinBytes(_addr *C.Address, _ret *C.uchar) {
	addr := (*cipher.Address)(unsafe.Pointer(&_addr))
	bytes := addr.BitcoinBytes()
	C.memcpy(unsafe.Pointer(_ret), unsafe.Pointer(&bytes[0]), C.size_t(len(bytes)))
}

// export SKY_Cipher_Verify
func SKY_Cipher_Verify(_addr *C.Address, _key *C.PubKey) C.uint {
	addr := (*cipher.Address)(unsafe.Pointer(&_addr))
	key := (*cipher.PubKey)(unsafe.Pointer(&_key))
	err := addr.Verify(*key)
	if err != nil {
		return 1
	}
	return 0
}

/*
// String address as Base58 encoded string
// Returns address as printable
// version is first byte in binary format
// in printed address its key, version, checksum
func SKY_Cipher_(addr Address) String() string {
}

// BitcoinString convert bitcoin address to hex string
func SKY_Cipher_(addr Address) BitcoinString() string {
}

// Checksum returns Address Checksum which is the first 4 bytes of sha256(key+version)
func SKY_Cipher_(addr *Address) Checksum() Checksum {
}

// BitcoinChecksum bitcoin checksum
func SKY_Cipher_(addr *Address) BitcoinChecksum() Checksum {
}
*/

/*
Bitcoin Functions
*/

/*
// BitcoinAddressFromPubkey prints the bitcoin address for a seckey
func SKY_Cipher_BitcoinAddressFromPubkey(pubkey PubKey) string {
}

// BitcoinWalletImportFormatFromSeckey exports seckey in wallet import format
// key must be compressed
func SKY_Cipher_BitcoinWalletImportFormatFromSeckey(seckey SecKey) string {
}

// BitcoinAddressFromBytes Returns an address given an Address.Bytes()
func SKY_Cipher_BitcoinAddressFromBytes(b []byte) (Address, error) {
}

// SecKeyFromWalletImportFormat extracts a seckey from wallet import format
func SKY_Cipher_SecKeyFromWalletImportFormat(input string) (SecKey, error) {
}

// MustSecKeyFromWalletImportFormat SecKeyFromWalletImportFormat or panic
func SKY_Cipher_MustSecKeyFromWalletImportFormat(input string) SecKey {
}
*/
