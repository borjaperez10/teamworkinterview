package customerimporter

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadDomains(t *testing.T) {

	var in = [][]string{
		{"Borja", "Perez", "bperez@circontrol.com", "Male", "192.168.1.18"},
		{"Andres", "Fernandez", "afernandez@github.com", "Male", "127.44.9.19"},
		{"Alba", "Sanz", "asanz@123.com", "Female", "99.98.6.2"},
		{"Angel", "Campos", "acampos@circontrol.com", "Male", "1.56.21.231"},
		{"Andrea", "Rodriguez", "arodriguez@youtube.com", "Female", "192.92.21.1"},
		{"Bea", "Garcia", "bgarcia@123.com", "Female", "123.0.98.21"},
		{"Carlos", "Gonzalez", "cgonzalez@example.com", "Male", "98.12.21.34"},
		{"Sergio", "Luengo", "sluengo@circontrol.com", "Male", "84.129.98.34"},
		{"Silvia", "Gutierrez", "sgutierrez@456.com", "Female", "172.16.21.2"},
		{"Kevin", "Perez", "kperez@youtube.com", "Male", "200.21.2.15"},
	}

	var expectedResult = []string{
		"circontrol.com",
		"github.com",
		"123.com",
		"circontrol.com",
		"youtube.com",
		"123.com",
		"example.com",
		"circontrol.com",
		"456.com",
		"youtube.com",
	}

	result := ReadDomains(in)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Error in %s. Expected %s but got %s", t.Name(), expectedResult, result)
	}
}
func TestGetDomain(t *testing.T) {

	in := "bperez@circontrol.com"
	expectedResult := "circontrol.com"
	result, err := GetDomain(in)
	if err != nil {
		t.Errorf("Error in %s. Error %s", t.Name(), err)
	}
	if result != expectedResult {
		t.Errorf("Error in %s. Expected %s but got %s", t.Name(), expectedResult, result)
	}

}

func TestGetDomainListInfo(t *testing.T) {
	var in = []string{
		"circontrol.com",
		"github.com",
		"123.com",
		"circontrol.com",
		"youtube.com",
		"123.com",
		"example.com",
		"circontrol.com",
		"456.com",
		"youtube.com",
	}
	var expectedResult = map[string]int{
		"circontrol.com": 3,
		"github.com":     1,
		"123.com":        2,
		"youtube.com":    2,
		"example.com":    1,
		"456.com":        1,
	}
	result := GetDomainListInfo(in)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Error in %s. Expected %v but got %v", t.Name(), expectedResult, result)
	}
}
func TestReadCSV(t *testing.T) {
	var in bytes.Buffer
	in.WriteString("data1,data2,data3,data4")

	var expectedResult = [][]string{
		{"data1", "data2", "data3", "data4"},
	}

	result, err := ReadCSV(&in)
	if err != nil {
		t.Errorf("Error in %s. Error %s", t.Name(), err)
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Error in %s. Expected %v but got %v", t.Name(), expectedResult, result)
	}
}

func TestSortDomainList(t *testing.T) {
	var in = map[string]int{
		"circontrol.com": 3,
		"github.com":     1,
		"123.com":        2,
	}
	var expectedResult1 = []string{
		"123.com",
		"circontrol.com",
		"github.com",
	}
	var expectedResult2 = []string{
		"github.com",
		"circontrol.com",
		"123.com",
	}
	var expectedResult3 = []string{
		"github.com",
		"123.com",
		"circontrol.com",
	}
	var expectedResult4 = []string{
		"circontrol.com",
		"123.com",
		"github.com",
	}
	result1 := SortDomainList(in, 1)
	result2 := SortDomainList(in, 2)
	result3 := SortDomainList(in, 3)
	result4 := SortDomainList(in, 4)

	if !reflect.DeepEqual(result1, expectedResult1) {
		t.Errorf("Error in %s. Expected %v but got %v", t.Name(), expectedResult1, result1)
	}
	if !reflect.DeepEqual(result2, expectedResult2) {
		t.Errorf("Error in %s. Expected %v but got %v", t.Name(), expectedResult2, result2)
	}
	if !reflect.DeepEqual(result3, expectedResult3) {
		t.Errorf("Error in %s. Expected %v but got %v", t.Name(), expectedResult3, result3)
	}
	if !reflect.DeepEqual(result4, expectedResult4) {
		t.Errorf("Error in %s. Expected %v but got %v", t.Name(), expectedResult4, result4)
	}

}
