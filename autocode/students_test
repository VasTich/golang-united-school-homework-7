package coverage

import (
	"os"
	"testing"
	"time"
	"strconv"
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

///////////////////////////////////////////////////////////////////

func generateMatrix() (*Matrix, error) {
	var rowCount int = 2
	var colCount int = 4
	var data []int;
	var counter int = 0
	var input string
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			data = append(data, counter)
			
			num := strconv.Itoa(counter)			
			input = input + num + " "
			counter = counter + 1
		}
		
		if i != rowCount -1 {
			input = input + "\n"
		}
	}
	
	mat, err := New(input)
	return mat, err
}

func TestNewMatrix(t *testing.T) {	
	mat, err := generateMatrix()
	if err != nil {
		t.Errorf("Tes New Matrix: not correct string transformation %s", err.Error())
	}
	
	if mat == nil {
		t.Errorf("Tes New Matrix: not correct string transformation matrix is nil")
	}
}

func TestNewMatrixDiffColumns(t *testing.T) {
	var rowCount int = 2
	var colCount int = 4
	var data []int;
	var counter int = 0
	var input string
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			data = append(data, counter)
			
			num := strconv.Itoa(counter)			
			input = input + num + " "
			counter = counter + 1
		}
		
		input = input + "\n"
	}
	
	mat, err := New(input)
	
	if err == nil || mat != nil {
		t.Errorf("Tes New Matrix Diff Columns: string not correct but not error")
	}
}

func TestNewMatrixNotInt(t *testing.T) {
	var input string = "5 100 hello"
	
	mat, err := New(input)
	
	if err == nil || mat != nil {
		t.Errorf("Tes New Matrix Diff Columns: string not correct but not error")
	}
}

func TestSet(t *testing.T) {
    mat, err := generateMatrix()
	if err != nil {
		t.Errorf("Test Set: not correct string transformation %s", err.Error())
	}
	
	if mat == nil {
		t.Errorf("Test Set: not correct string transformation matrix is nil")
	}
	
	var expected int = 20
	if !mat.Set(0, 0, expected) && mat.data[0] != expected {
		t.Errorf("Test Set: expected value %d get %d", expected, mat.data[0])
	}
}

func TestSetUncorrectRow(t *testing.T) {
    mat, err := generateMatrix()
	if err != nil {
		t.Errorf("Test Set Uncorrect row: not correct string transformation %s", err.Error())
	}
	
	if mat == nil {
		t.Errorf("Test Set Uncorrect row: not correct string transformation matrix is nil")
	}
	
	var expected int = 20
	if mat.Set(100, 0, expected) {
		t.Errorf("Test set uncorrect row: set uncorrect row")
	}
	
	if mat.Set(-100, 0, expected) {
		t.Errorf("Test set uncorrect row: set uncorrect row")
	}
}

func TestSetUncorrectCol(t *testing.T) {
    mat, err := generateMatrix()
	if err != nil {
		t.Errorf("Test Set Uncorrect col: not correct string transformation %s", err.Error())
	}
	
	if mat == nil {
		t.Errorf("Test Set Uncorrect col: not correct string transformation matrix is nil")
	}
	
	var expected int = 20
	if mat.Set(0, 100, expected) {
		t.Errorf("Test set uncorrect col: set uncorrect col")
	}
	
	if mat.Set(0, -100, expected) {
		t.Errorf("Test set uncorrect col: set uncorrect col")
	}
}

func TestRows(t *testing.T) {
	mat, err := generateMatrix()
	if err != nil {
		t.Errorf("Test Rows: not correct string transformation %s", err.Error())
	}
	
	if mat == nil {
		t.Errorf("Test Rows: not correct string transformation matrix is nil")
	}
	
	rows := mat.Rows()
	if len(rows) != 2 {
		t.Errorf("Test Rows: expected row %d get %d", 2, len(rows))
	}
}

func TestCols(t *testing.T) {
	mat, err := generateMatrix()
	if err != nil {
		t.Errorf("Test Cols: not correct string transformation %s", err.Error())
	}
	
	if mat == nil {
		t.Errorf("Test Cols: not correct string transformation matrix is nil")
	}
	
	cols := mat.Cols()
	if len(cols) != 4 {
		t.Errorf("Test Cols: expected cols %d get %d", 4, len(cols))
	}
}
