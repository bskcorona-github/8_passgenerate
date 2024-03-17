package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
	symbols          = "!@#$%^&*()-_=+,.?/:;{}[]`~"
)

func main() {
	// パスワードの長さを指定します
	passwordLength := 12

	// パスワードを生成します
	password := generatePassword(passwordLength)

	// 生成されたパスワードを表示します
	fmt.Println("Generated Password:", password)
}

func generatePassword(length int) string {
	// 乱数のシードを設定します
	rand.Seed(time.Now().UnixNano())

	// 使用する文字セットを定義します
	charset := lowercaseLetters + uppercaseLetters + digits + symbols

	// パスワードを格納するためのバッファを作成します
	password := make([]byte, length)

	// パスワードを生成します
	for i := 0; i < length; i++ {
		// ランダムなインデックスを生成します
		randomIndex := rand.Intn(len(charset))
		// ランダムな文字を選択してパスワードに追加します
		password[i] = charset[randomIndex]
	}

	// 生成されたパスワードを文字列として返します
	return string(password)
}
