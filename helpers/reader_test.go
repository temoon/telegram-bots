package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
)

const ExampleFilename = "NOTICE"
const ExampleTemplateFilename = "example.tmpl"

const ExampleOkTemplate = `This is a template. There is foo = {{.foo}}.`
const ExampleOkTemplateText = `This is a template. There is foo = bar.`
const ExampleInvalidTemplate = `This is a template. There is foo = {{.foo`

var ExampleData []byte
var ExampleTemplateData = map[string]string{"foo": "bar"}

func init() {
	var err error
	if ExampleData, err = os.ReadFile(path.Join("..", ExampleFilename)); err != nil {
		panic(err)
	}
}

// region ReadTemplate

func TestReadTemplate(t *testing.T) {
	filenameOk := path.Join(t.TempDir(), ExampleTemplateFilename)
	filenameInvalid := path.Join(t.TempDir(), ExampleTemplateFilename)

	var text string
	var err error

	// region Prepare templates
	if err = createTemplate(filenameOk, ExampleOkTemplate); err != nil {
		t.Fatal("Template not created:", err)
	}

	if err = createTemplate(filenameInvalid, ExampleInvalidTemplate); err != nil {
		t.Fatal("Template not created:", err)
	}
	// endregion

	if text, err = ReadTemplate(ExampleTemplateData, filenameOk); text != ExampleOkTemplateText || err != nil {
		t.Error("Filename:", filenameOk, "Text:", text, "Error:", err)
	}

	if text, err = ReadTemplate(ExampleTemplateData, "..", "NOT_FOUND"); text != "" || !errors.Is(err, os.ErrNotExist) {
		t.Error("Filename:", filenameOk, "Text:", text, "Error:", err)
	}

	if text, err = ReadTemplate(ExampleTemplateData, filenameInvalid); text != "" || err == nil {
		t.Error("Filename:", filenameInvalid, "Text:", text, "Error:", err)
	}

	if text, err = ReadTemplate(ReadTemplate, filenameOk); text != "" || err == nil {
		t.Error("Filename:", filenameOk, "Text:", text, "Error:", err)
	}
}

func BenchmarkReadTemplate(b *testing.B) {
	filename := path.Join(b.TempDir(), ExampleTemplateFilename)
	if err := createTemplate(filename, ExampleOkTemplate); err != nil {
		b.Fatal("Template not created:", err)
	}

	for i := 0; i < b.N; i++ {
		_, _ = ReadTemplate(ExampleTemplateData, filename)
	}
}

func ExampleReadTemplate() {
	data, _ := ReadTemplate(ExampleTemplateData, "..", ExampleFilename)
	fmt.Println(len(data))
	// Output:
	// 107
}

func createTemplate(filename string, content string) (err error) {
	var file *os.File
	if file, err = os.Create(filename); err != nil {
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		return
	}

	return
}

// endregion

// region ReadStaticFile

type CaseReadStaticFile struct {
	Filenames   []string
	Text        string
	ExampleText bool
	Error       error
}

func (c *CaseReadStaticFile) GetText() string {
	if c.ExampleText {
		return string(ExampleData)
	}

	return c.Text
}

var testCasesReadStaticFile = []CaseReadStaticFile{
	{Filenames: []string{"..", ExampleFilename}, ExampleText: true, Error: nil},
	{Filenames: []string{"..", "NOT_FOUND"}, Text: "", Error: os.ErrNotExist},
	{Filenames: []string{path.Join("..", ExampleFilename)}, ExampleText: true, Error: nil},
	{Filenames: []string{os.DevNull}, Text: "", Error: nil},
}

func TestReadStaticFile(t *testing.T) {
	var text string
	var err error
	for _, testCase := range testCasesReadStaticFile {
		if text, err = ReadStaticFile(testCase.Filenames...); text != testCase.GetText() || !errors.Is(err, testCase.Error) {
			t.Error("Filenames:", strings.Join(testCase.Filenames, " / "), "Text:", text, "Error:", err)
		}
	}
}

func BenchmarkReadStaticFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ReadStaticFile(testCasesReadStaticFile[i%len(testCasesReadStaticFile)].Filenames...)
	}
}

func ExampleReadStaticFile() {
	data, _ := ReadStaticFile("..", ExampleFilename)
	fmt.Println(len(data))
	// Output:
	// 107
}

// endregion

// region ReadStaticFileData

type CaseReadStaticFileData struct {
	Filenames   []string
	Data        []byte
	ExampleData bool
	Error       error
}

func (c *CaseReadStaticFileData) GetData() []byte {
	if c.ExampleData {
		return ExampleData
	}

	return c.Data
}

var testCasesReadStaticFileData = []CaseReadStaticFileData{
	{Filenames: []string{"..", ExampleFilename}, ExampleData: true, Error: nil},
	{Filenames: []string{"..", "NOT_FOUND"}, Data: nil, Error: os.ErrNotExist},
	{Filenames: []string{path.Join("..", ExampleFilename)}, ExampleData: true, Error: nil},
	{Filenames: []string{os.DevNull}, Data: nil, Error: nil},
}

func TestReadStaticFileData(t *testing.T) {
	var data []byte
	var err error
	for _, testCase := range testCasesReadStaticFileData {
		if data, err = ReadStaticFileData(testCase.Filenames...); !bytes.Equal(data, testCase.GetData()) || !errors.Is(err, testCase.Error) {
			t.Error("Filenames:", strings.Join(testCase.Filenames, " / "), "Data:", data, "Error:", err)
		}
	}
}

func BenchmarkReadStaticFileData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ReadStaticFileData(testCasesReadStaticFileData[i%len(testCasesReadStaticFileData)].Filenames...)
	}
}

func ExampleReadStaticFileData() {
	data, _ := ReadStaticFileData("..", ExampleFilename)
	fmt.Println(len(data))
	// Output:
	// 107
}

// endregion
