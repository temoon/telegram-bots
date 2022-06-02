package helpers

import (
	"fmt"
	"testing"
	"time"

	"github.com/undefinedlabs/go-mpatch"
)

// region Bool

var testCasesBool = []bool{true, false}

func TestBool(t *testing.T) {
	for _, value := range testCasesBool {
		if ref := Bool(value); ref == nil || *ref != value {
			t.Error("Value:", value, "Got:", *ref)
		}
	}
}

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bool(testCasesBool[i%len(testCasesBool)])
	}
}

func ExampleBool() {
	fmt.Println(*Bool(true))
	// Output:
	// true
}

// endregion

// region Int

var testCasesInt = []int{-42, 0, 42}

func TestInt(t *testing.T) {
	for _, value := range testCasesInt {
		if ref := Int(value); ref == nil || *ref != value {
			t.Error("Value:", value, "Got:", *ref)
		}
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(testCasesInt[i%len(testCasesInt)])
	}
}

func ExampleInt() {
	fmt.Println(*Int(42))
	// Output:
	// 42
}

// endregion

// region Int32

var testCasesInt32 = []int32{-42, 0, 42}

func TestInt32(t *testing.T) {
	for _, value := range testCasesInt32 {
		if ref := Int32(value); ref == nil || *ref != value {
			t.Error("Value:", value, "Got:", *ref)
		}
	}
}

func BenchmarkInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int32(testCasesInt32[i%len(testCasesInt32)])
	}
}

func ExampleInt32() {
	fmt.Println(*Int32(42))
	// Output:
	// 42
}

// endregion

// region Int64

var testCasesInt64 = []int64{-42, 0, 42}

func TestInt64(t *testing.T) {
	for _, value := range testCasesInt64 {
		if ref := Int64(value); ref == nil || *ref != value {
			t.Error("Value:", value, "Got:", *ref)
		}
	}
}

func BenchmarkInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int64(testCasesInt64[i%len(testCasesInt64)])
	}
}

func ExampleInt64() {
	fmt.Println(*Int64(42))
	// Output:
	// 42
}

// endregion

// region String

var testCasesString = []string{"", " ", "foo"}

func TestString(t *testing.T) {
	for _, value := range testCasesString {
		if ref := String(value); ref == nil || *ref != value {
			t.Error("Value:", value, "Got:", *ref)
		}
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(testCasesString[i%len(testCasesString)])
	}
}

func ExampleString() {
	fmt.Println(*String("foo"))
	// Output:
	// foo
}

// endregion

// region BoolOrFalse

type CaseBoolOrFalse struct {
	Ref    *bool
	Return bool
}

var testCasesBoolOrFalse = []CaseBoolOrFalse{
	{Ref: nil, Return: false},
	{Ref: Bool(true), Return: true},
	{Ref: Bool(false), Return: false},
}

func TestBoolOrFalse(t *testing.T) {
	for _, testCase := range testCasesBoolOrFalse {
		if value := BoolOrFalse(testCase.Ref); value != testCase.Return {
			t.Error("Ref:", testCase.Ref, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkBoolOrFalse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BoolOrFalse(testCasesBoolOrFalse[i%len(testCasesBoolOrFalse)].Ref)
	}
}

func ExampleBoolOrFalse() {
	fmt.Println(BoolOrFalse(Bool(true)))
	fmt.Println(BoolOrFalse(nil))
	// Output:
	// true
	// false
}

// endregion

// region IntOrZero

type CaseIntOrZero struct {
	Ref    *int
	Return int
}

var testCasesIntOrZero = []CaseIntOrZero{
	{Ref: nil, Return: 0},
	{Ref: Int(-42), Return: -42},
	{Ref: Int(0), Return: 0},
	{Ref: Int(42), Return: 42},
}

func TestIntOrZero(t *testing.T) {
	for _, testCase := range testCasesIntOrZero {
		if value := IntOrZero(testCase.Ref); value != testCase.Return {
			t.Error("Ref:", testCase.Ref, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkIntOrZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntOrZero(testCasesIntOrZero[i%len(testCasesIntOrZero)].Ref)
	}
}

func ExampleIntOrZero() {
	fmt.Println(IntOrZero(Int(42)))
	fmt.Println(IntOrZero(nil))
	// Output:
	// 42
	// 0
}

// endregion

// region Int32OrZero

type CaseInt32OrZero struct {
	Ref    *int32
	Return int32
}

var testCasesInt32OrZero = []CaseInt32OrZero{
	{Ref: nil, Return: 0},
	{Ref: Int32(-42), Return: -42},
	{Ref: Int32(0), Return: 0},
	{Ref: Int32(42), Return: 42},
}

func TestInt32OrZero(t *testing.T) {
	for _, testCase := range testCasesInt32OrZero {
		if value := Int32OrZero(testCase.Ref); value != testCase.Return {
			t.Error("Ref:", testCase.Ref, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkInt32OrZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int32OrZero(testCasesInt32OrZero[i%len(testCasesInt32OrZero)].Ref)
	}
}

func ExampleInt32OrZero() {
	fmt.Println(Int32OrZero(Int32(42)))
	fmt.Println(Int32OrZero(nil))
	// Output:
	// 42
	// 0
}

// endregion

// region Int64OrZero

type CaseInt64OrZero struct {
	Ref    *int64
	Return int64
}

var testCasesInt64OrZero = []CaseInt64OrZero{
	{Ref: nil, Return: 0},
	{Ref: Int64(-42), Return: -42},
	{Ref: Int64(0), Return: 0},
	{Ref: Int64(42), Return: 42},
}

func TestInt64OrZero(t *testing.T) {
	for _, testCase := range testCasesInt64OrZero {
		if value := Int64OrZero(testCase.Ref); value != testCase.Return {
			t.Error("Ref:", testCase.Ref, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkInt64OrZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int64OrZero(testCasesInt64OrZero[i%len(testCasesInt64OrZero)].Ref)
	}
}

func ExampleInt64OrZero() {
	fmt.Println(Int64OrZero(Int64(42)))
	fmt.Println(Int64OrZero(nil))
	// Output:
	// 42
	// 0
}

// endregion

// region StringOrEmpty

type CaseStringOrEmpty struct {
	Ref    *string
	Return string
}

var testCasesStringOrEmpty = []CaseStringOrEmpty{
	{Ref: nil, Return: ""},
	{Ref: String(""), Return: ""},
	{Ref: String(" "), Return: " "},
	{Ref: String("foo"), Return: "foo"},
}

func TestStringOrEmpty(t *testing.T) {
	for _, testCase := range testCasesStringOrEmpty {
		if value := StringOrEmpty(testCase.Ref); value != testCase.Return {
			t.Error("Ref:", testCase.Ref, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkStringOrEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringOrEmpty(testCasesStringOrEmpty[i%len(testCasesStringOrEmpty)].Ref)
	}
}

func ExampleStringOrEmpty() {
	fmt.Println(StringOrEmpty(String("foo")))
	fmt.Println(StringOrEmpty(nil))
	// Output:
	// foo
}

// endregion

// region TimeOrNow

type CaseTimeOrNow struct {
	Ref    *time.Time
	Return time.Time
}

var timeNow = time.Unix(1234567890, 0)
var timeRef = time.Now().Add(time.Minute)

func init() {
	_, _ = mpatch.PatchMethod(time.Now, func() time.Time {
		return timeNow
	})
}

var testCasesTimeOrNow = []CaseTimeOrNow{
	{Ref: nil, Return: timeNow},
	{Ref: Time(timeRef), Return: timeRef},
}

func TestTimeOrNow(t *testing.T) {
	for _, testCase := range testCasesTimeOrNow {
		if value := TimeOrNow(testCase.Ref); value != testCase.Return {
			t.Error("Ref:", testCase.Ref, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkTimeOrNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeOrNow(testCasesTimeOrNow[i%len(testCasesTimeOrNow)].Ref)
	}
}

func ExampleTimeOrNow() {
	fmt.Println(TimeOrNow(Time(time.Now().In(time.UTC))))
	fmt.Println(TimeOrNow(nil).Unix())
	// Output:
	// 2009-02-13 23:31:30 +0000 UTC
	// 1234567890
}

// endregion

// region Weekday

type CaseWeekday struct {
	Datetime time.Time
	Return   string
}

var testCasesWeekday = []CaseWeekday{
	{Datetime: time.Unix(1561939200, 0), Return: DayMonday},
	{Datetime: time.Unix(1561939200, 0).Add(time.Hour * 24), Return: DayTuesday},
	{Datetime: time.Unix(1561939200, 0).Add(time.Hour * 24 * 2), Return: DayWednesday},
	{Datetime: time.Unix(1561939200, 0).Add(time.Hour * 24 * 3), Return: DayThursday},
	{Datetime: time.Unix(1561939200, 0).Add(time.Hour * 24 * 4), Return: DayFriday},
	{Datetime: time.Unix(1561939200, 0).Add(time.Hour * 24 * 5), Return: DaySaturday},
	{Datetime: time.Unix(1561939200, 0).Add(time.Hour * 24 * 6), Return: DaySunday},
}

func TestWeekday(t *testing.T) {
	for _, testCase := range testCasesWeekday {
		if value := Weekday(testCase.Datetime); value != testCase.Return {
			t.Error("Datetime:", testCase.Datetime, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkWeekday(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Weekday(testCasesWeekday[i%len(testCasesWeekday)].Datetime)
	}
}

func ExampleWeekday() {
	fmt.Println(Weekday(time.Unix(1561939200, 0)))
	// Output:
	// Пн
}

// endregion

// region Url

type CaseUrl struct {
	Caption string
	Url     string
	Return  string
}

var testCasesUrl = []CaseUrl{
	{Caption: "", Url: "", Return: "[]()"},
	{Caption: "foo", Url: "", Return: "[foo]()"},
	{Caption: "", Url: "https://example.com", Return: "[](https://example.com)"},
	{Caption: "foo", Url: "https://example.com", Return: "[foo](https://example.com)"},
	{Caption: "()", Url: "[]", Return: "[()]([])"},
	{Caption: "[]", Url: "()", Return: "[[]](())"},
}

func TestUrl(t *testing.T) {
	for _, testCase := range testCasesUrl {
		if value := Url(testCase.Caption, testCase.Url); value != testCase.Return {
			t.Error("Caption:", testCase.Caption, "Url:", testCase.Url, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkUrl(b *testing.B) {
	var testCase CaseUrl
	for i := 0; i < b.N; i++ {
		testCase = testCasesUrl[i%len(testCasesUrl)]
		Url(testCase.Caption, testCase.Url)
	}
}

func ExampleUrl() {
	fmt.Println(Url("foo", "https://example.com"))
	// Output:
	// [foo](https://example.com)
}

// endregion

// region UserUrl

type CaseUserUrl struct {
	UserId int64
	Return string
}

var testCasesUserUrl = []CaseUserUrl{
	{UserId: -42, Return: "tg://user?id=-42"},
	{UserId: 0, Return: "tg://user?id=0"},
	{UserId: 42, Return: "tg://user?id=42"},
}

func TestUserUrl(t *testing.T) {
	for _, testCase := range testCasesUserUrl {
		if value := UserUrl(testCase.UserId); value != testCase.Return {
			t.Error("User ID:", testCase.UserId, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkUserUrl(b *testing.B) {
	var testCase CaseUserUrl
	for i := 0; i < b.N; i++ {
		testCase = testCasesUserUrl[i%len(testCasesUserUrl)]
		UserUrl(testCase.UserId)
	}
}

func ExampleUserUrl() {
	fmt.Println(UserUrl(42))
	// Output:
	// tg://user?id=42
}

// endregion
