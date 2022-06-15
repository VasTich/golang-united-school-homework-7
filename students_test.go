package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLen(t *testing.T) {
	var persons People
	firstMan := Person{firstName: "Ivan", lastName : "Ivanov", birthDay : time.Date(1987, time.May, 30, 5, 0, 0, 0, time.UTC)}
	var expectedCount int = 5
	for i := 0; i < 5; i ++ {
		persons = append(persons, firstMan)
	}
	testLen := persons.Len()
	if testLen != expectedCount {
		t.Errorf("Test Len: expected len %d, got %d", expectedCount, testLen)
	}
}

func TestSwap(t *testing.T) {
	var persons People
	firstMan := Person{firstName: "Ivan", lastName : "Ivanov", birthDay : time.Date(1987, time.May, 30, 5, 0, 0, 0, time.UTC)}
	secMan := Person{firstName: "Petr", lastName : "Petrov", birthDay : time.Date(1997, time.June, 30, 5, 0, 0, 0, time.UTC)}
	
	persons = append(persons, firstMan)
	persons = append(persons, secMan)
	
	persons.Swap(0,1)
	
	if persons[0] != secMan || persons[1] != firstMan {
		t.Errorf("Test Swap: first person %+v second person %+v", persons[0], persons[1])
	}
}

func TestLessDifferentBirthday(t *testing.T) {
	var persons People
	firstMan := Person{firstName: "Ivan", lastName : "Ivanov", birthDay : time.Date(1987, time.May, 30, 5, 0, 0, 0, time.UTC)}
	secMan := Person{firstName: "Petr", lastName : "Petrov", birthDay : time.Date(1997, time.June, 30, 5, 0, 0, 0, time.UTC)}
	
	persons = append(persons, firstMan)
	persons = append(persons, secMan)
	
	if !persons.Less(1, 0) {
		t.Errorf("Test Less Different Birthday: not correct less operation")
	}
}

func TestLessEqualBirthday(t *testing.T) {
	var persons People
	firstMan := Person{firstName: "Ivan", lastName : "Ivanov", birthDay : time.Date(1997, time.June, 30, 5, 0, 0, 0, time.UTC)}
	secMan := Person{firstName: "Petr", lastName : "Petrov", birthDay : time.Date(1997, time.June, 30, 5, 0, 0, 0, time.UTC)}
	
	persons = append(persons, firstMan)
	persons = append(persons, secMan)
	
	if !persons.Less(0, 1) {
	t.Errorf("Test Less Equal Birthday: not correct less operation")
	}
}

func TestLessEqualBirthdayEqualFirstName(t *testing.T) {
	var persons People
	firstMan := Person{firstName: "Ivan", lastName : "Ivanov", birthDay : time.Date(1997, time.June, 30, 5, 0, 0, 0, time.UTC)}
	secMan := Person{firstName: "Ivan", lastName : "Petrov", birthDay : time.Date(1997, time.June, 30, 5, 0, 0, 0, time.UTC)}
	
	persons = append(persons, firstMan)
	persons = append(persons, secMan)
	
	if !persons.Less(0, 1) {
		t.Errorf("Test Less Equal Birthday Equal First Name: not correct less operation")
	}
}