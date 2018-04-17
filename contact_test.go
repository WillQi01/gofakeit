package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleContact() {
	Seed(11)
	contact := Contact()
	fmt.Println(contact.Phone)
	fmt.Println(contact.Email)
	// Output: 3287271570
	// santinostanton@carroll.biz
}

func BenchmarkContact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Contact()
	}
}

func ExamplePhone() {
	Seed(11)
	fmt.Println(Phone())
	// Output: 3287271570
}

func BenchmarkPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phone()
	}
}

func ExamplePhoneFormatted() {
	Seed(11)
	fmt.Println(PhoneFormatted())
	// Output: 287-271-5702
}

func BenchmarkPhoneFormatted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhoneFormatted()
	}
}

func ExampleEmail() {
	Seed(11)
	fmt.Println(Email())
	// Output: markusmoen@pagac.net
}

func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Email()
	}
}
