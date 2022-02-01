package strategy

import "testing"

func Test_demo(t *testing.T) {
	normalStrategy := "file"
	storage, _ := NewStorageStrategy(normalStrategy)
	storage.Save("普通保存", []byte{})

	encryptStrategy := "encrypt_file"
	storage, _ = NewStorageStrategy(encryptStrategy)
	storage.Save("加密保存", []byte{})
}
