package esepunittests

type GradeCalculator struct {
    grades []Grade
	passCheck bool
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator(passCheck bool) *GradeCalculator {
    return &GradeCalculator{
        grades: make([]Grade, 0),
		passCheck: passCheck,
    }
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if gc.passCheck {
        if numericalGrade >= 60 {
            return "Pass"
        }
        return "Fail"
    }

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
    gc.grades = append(gc.grades, Grade{
        Name:  name,
        Grade: grade,
        Type:  gradeType,
    })
} 
func (gc *GradeCalculator) calculateNumericalGrade() int {
    // Calculate assignment average
    assignmentSum := 0
    assignmentCount := 0
    for _, grade := range gc.grades {
        if grade.Type == Assignment {
            assignmentSum += grade.Grade
            assignmentCount++
        }
    }
    assignment_average := assignmentSum / assignmentCount

    // Calculate exam average  
    examSum := 0
    examCount := 0
    for _, grade := range gc.grades {
        if grade.Type == Exam {
            examSum += grade.Grade
            examCount++
        }
    }
    exam_average := examSum / examCount

    // Calculate essay average
    essaySum := 0
    essayCount := 0
    for _, grade := range gc.grades {
        if grade.Type == Essay {
            essaySum += grade.Grade
            essayCount++
        }
    }
    essay_average := essaySum / essayCount

    // Keep the same weighted calculation
    weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

    return int(weighted_grade)
}

func computeAverage(grades []Grade) int {
	sum := 0

	for i := range grades {
		sum += grades[i].Grade
	}

	return sum / len(grades)
}
