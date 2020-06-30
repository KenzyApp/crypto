package des

import (
	"crypto/des"
	"encoding/hex"
	"strings"
	"testing"
)

func TestDESEncryption(t *testing.T) {
	keyBytes, _ := hex.DecodeString("0123456789ABCDEF")
	keyBlock, _ := des.NewCipher(keyBytes)
	cipher := DESCipher{keyBlock, keyBytes}

	testData := map[string]string{
		"":                                 "",
		"4123456789012345":                 "B76ADDCE71CCC6BE",
		"41234567890123454123456789012345": "B76ADDCE71CCC6BEB76ADDCE71CCC6BE",
	}

	for plainText, expectedCipherText := range testData {
		cipherBytes, err := cipher.EncryptHex(plainText)
		if err != nil {
			t.Errorf("Did not expect an error but got %q", err)
		}
		cipherText := hex.EncodeToString(cipherBytes)
		if !strings.EqualFold(expectedCipherText, cipherText) {
			t.Errorf("Expected value %s but got %s instead", expectedCipherText, cipherText)
		}
	}

	if _, err := cipher.EncryptHex("1234"); err == nil {
		t.Error("should be an error if the input text is not valid")
	}
}

func TestTripleDESEncryption(t *testing.T) {
	keyBytes, _ := hex.DecodeString("A1FA4BF45ECDA0C1198CF971365C148CA1FA4BF45ECDA0C1")
	keyBlock, _ := des.NewTripleDESCipher(keyBytes)
	cipher := DESCipher{keyBlock, keyBytes}

	testData := map[string]string{
		"":                                 "",
		"4123456789012345":                 "26ECB8D84DDFF9E3",
		"41234567890123454123456789012345": "26ECB8D84DDFF9E326ECB8D84DDFF9E3",
	}

	for plainText, expectedCipherText := range testData {
		cipherBytes, err := cipher.EncryptHex(plainText)
		if err != nil {
			t.Errorf("Did not expect an error but got %q", err)
		}
		cipherText := hex.EncodeToString(cipherBytes)
		if !strings.EqualFold(expectedCipherText, cipherText) {
			t.Errorf("Expected value %s but got %s instead", expectedCipherText, cipherText)
		}
	}

	if _, err := cipher.EncryptHex("1234"); err == nil {
		t.Error("should be an error if the input text is not valid")
	}
}

func TestDESDecryption(t *testing.T) {
	keyBytes, _ := hex.DecodeString("0123456789ABCDEF")
	keyBlock, _ := des.NewCipher(keyBytes)
	cipher := DESCipher{keyBlock, keyBytes}

	testData := map[string]string{
		"":                                 "",
		"B76ADDCE71CCC6BE":                 "4123456789012345",
		"B76ADDCE71CCC6BEB76ADDCE71CCC6BE": "41234567890123454123456789012345",
	}

	for cipherText, expectedPlainText := range testData {
		plainBytes, err := cipher.DecryptHex(cipherText)
		if err != nil {
			t.Errorf("Did not expect an error but got %q", err)
		}
		plainText := hex.EncodeToString(plainBytes)
		if !strings.EqualFold(expectedPlainText, plainText) {
			t.Errorf("Expected value %s but got %s instead", expectedPlainText, cipherText)
		}
	}

	if _, err := cipher.EncryptHex("1234"); err == nil {
		t.Error("should be an error if the input text is not valid")
	}
}

func TestTripleDESDecryption(t *testing.T) {
	keyBytes, _ := hex.DecodeString("A1FA4BF45ECDA0C1198CF971365C148CA1FA4BF45ECDA0C1")
	keyBlock, _ := des.NewTripleDESCipher(keyBytes)
	cipher := DESCipher{keyBlock, keyBytes}

	testData := map[string]string{
		"":                                 "",
		"26ECB8D84DDFF9E3":                 "4123456789012345",
		"26ECB8D84DDFF9E326ECB8D84DDFF9E3": "41234567890123454123456789012345",
	}

	for cipherText, expectedPlainText := range testData {
		plainBytes, err := cipher.DecryptHex(cipherText)
		if err != nil {
			t.Errorf("Did not expect an error but got %q", err)
		}
		plainText := hex.EncodeToString(plainBytes)
		if !strings.EqualFold(expectedPlainText, plainText) {
			t.Errorf("Expected value %s but got %s instead", expectedPlainText, cipherText)
		}
	}

	if _, err := cipher.EncryptHex("1234"); err == nil {
		t.Error("should be an error if the input text is not valid")
	}
}

func TestDESCheckValueVerification(t *testing.T) {
	keyBytes, _ := hex.DecodeString("0123456789ABCDEF")
	keyBlock, _ := des.NewCipher(keyBytes)
	cipher := DESCipher{keyBlock, keyBytes}

	if !cipher.VerifyCheckValue("D5D44F") {
		t.Error("expect checkValue to be valid")
	}

	if cipher.VerifyCheckValue("D5D44E") {
		t.Error("expect checkValue to be invalid")
	}
}

func TestTripleDESCheckValueVerification(t *testing.T) {
	keyBytes, _ := hex.DecodeString("F94AC55104B0E5532D0A61D2D2C6C655F94AC55104B0E553")
	keyBlock, _ := des.NewTripleDESCipher(keyBytes)
	cipher := DESCipher{keyBlock, keyBytes}

	if !cipher.VerifyCheckValue("6FAAD3") {
		t.Error("expect checkValue to be valid")
	}

	if cipher.VerifyCheckValue("6FAAD4") {
		t.Error("expect checkValue to be invalid")
	}
}
