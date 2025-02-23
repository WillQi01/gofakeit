package gofakeit

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

type strTyp string

func (t strTyp) Fake(faker *Faker) interface{} {
	return faker.FirstName()
}

type strTypPtr string

func (t *strTypPtr) Fake(faker *Faker) interface{} {
	return strTypPtr("hello test ptr")
}

type testStruct1 struct {
	B string `fake:"{firstname}"`
}

type testStruct2 struct {
	B strTyp
}

func TestIsFakeable(t *testing.T) {
	var t1 testStruct2
	var t2 *testStruct2
	var t3 strTyp
	var t4 *strTyp
	var t5 strTypPtr
	var t6 *strTypPtr

	if isFakeable(reflect.ValueOf(t1).Type()) {
		t.Errorf("expected testStruct2 not to be fakeable")
	}

	if isFakeable(reflect.ValueOf(t2).Type()) {
		t.Errorf("expected *testStruct2 not to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t3).Type()) {
		t.Errorf("expected strTyp to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t4).Type()) {
		t.Errorf("expected *strTyp to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t5).Type()) {
		t.Errorf("expected strTypPtr to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t6).Type()) {
		t.Errorf("expected *strTypPtr to be fakeable")
	}
}

func ExampleFakeable() {
	var t1 testStruct1
	var t2 testStruct1
	var t3 testStruct2
	var t4 testStruct2
	New(314).Struct(&t1)
	New(314).Struct(&t2)
	New(314).Struct(&t3)
	New(314).Struct(&t4)

	fmt.Printf("%#v\n", t1)
	fmt.Printf("%#v\n", t2)
	fmt.Printf("%#v\n", t3)
	fmt.Printf("%#v\n", t4)
	// Expected Output:
	// gofakeit.testStruct1{B:"Margarette"}
	// gofakeit.testStruct1{B:"Margarette"}
	// gofakeit.testStruct2{B:"Margarette"}
	// gofakeit.testStruct2{B:"Margarette"}
}

type gammaFloat64 float64

func (gammaFloat64) Fake(faker *Faker) interface{} {
	alpha := 2.0

	// Generate a random value from the Gamma distribution
	var r float64
	for r == 0 {
		u := faker.Float64Range(0, 1)
		v := faker.Float64Range(0, 1)
		w := u * (1 - u)
		y := math.Sqrt(-2 * math.Log(w) / w)
		x := alpha * (y*v + u - 0.5)
		if x > 0 {
			r = x
		}
	}
	return gammaFloat64(r)
}

func ExampleGammaFloat64() {
	f1 := New(100)

	// Fakes random values from the Gamma distribution
	var A1 gammaFloat64
	var A2 gammaFloat64
	var A3 gammaFloat64
	f1.Struct(&A1)
	f1.Struct(&A2)
	f1.Struct(&A3)

	fmt.Println(A1)
	fmt.Println(A2)
	fmt.Println(A3)
	// Output:
	// 10.300651760129734
	// 5.391434877284098
	// 2.0575989252140676
}

type poissonInt64 int64

func (poissonInt64) Fake(faker *Faker) interface{} {
	lambda := 15.0

	// Generate a random value from the Poisson distribution
	var k int64
	var p float64 = 1.0
	var L float64 = math.Exp(-lambda)
	for p > L {
		u := faker.Float64Range(0, 1)
		p *= u
		k++
	}
	return poissonInt64(k - 1)
}

type customerSupportEmployee struct {
	Name             string `fake:"{firstname} {lastname}"`
	CallCountPerHour poissonInt64
}

func ExamplecustomerSupportEmployee() {
	f1 := New(100)

	// Fakes random values from the Gamma distribution
	var A1 customerSupportEmployee
	var A2 customerSupportEmployee
	var A3 customerSupportEmployee
	f1.Struct(&A1)
	f1.Struct(&A2)
	f1.Struct(&A3)

	fmt.Printf("%#v\n", A1)
	fmt.Printf("%#v\n", A2)
	fmt.Printf("%#v\n", A3)
	// Output:
	// gofakeit.customerSupportEmployee{Name:"Pearline Rippin", CallCountPerHour:12}
	// gofakeit.customerSupportEmployee{Name:"Sammie Renner", CallCountPerHour:23}
	// gofakeit.customerSupportEmployee{Name:"Katlyn Runte", CallCountPerHour:8}
}
