package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator(false)

	gradeCalculator.AddGrade("open source assignment", 90, Assignment)
	gradeCalculator.AddGrade("exam 1", 90, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 90, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator(false)

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator(false)

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator(false)

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator(false)

	gradeCalculator.AddGrade("open source assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 55, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeStringLabels(t *testing.T) {
    if Assignment.String() != "assignment" {
        t.Errorf("Expected assignment, got %q", Assignment.String())
    }
    if Exam.String() != "exam" {
        t.Errorf("Expected exam, got %q", Exam.String())
    }
    if Essay.String() != "essay" {
        t.Errorf("Expected essay, got %q", Essay.String())
    }
}

func TestGetGradePass(t *testing.T) {
    expected_value := "Pass"

    gradeCalculator := NewGradeCalculator(true)  // true = pass/fail mode

    gradeCalculator.AddGrade("open source assignment", 70, Assignment)
    gradeCalculator.AddGrade("exam 1", 75, Exam)
    gradeCalculator.AddGrade("essay on ai ethics", 68, Essay)

    actual_value := gradeCalculator.GetFinalGrade()

    if expected_value != actual_value {
        t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
    }
}

func TestGetGradeFail(t *testing.T) {
    expected_value := "Fail"

    gradeCalculator := NewGradeCalculator(true)

    gradeCalculator.AddGrade("open source assignment", 50, Assignment)
    gradeCalculator.AddGrade("exam 1", 55, Exam)
    gradeCalculator.AddGrade("essay on ai ethics", 45, Essay)

    actual_value := gradeCalculator.GetFinalGrade()

    if expected_value != actual_value {
        t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
    }
}