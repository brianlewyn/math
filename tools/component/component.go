package component

import (
	"fmt"
	"strconv"
	"strings"

	m "github.com/brianlewyn/math/tools/message"
	u "github.com/brianlewyn/math/tools/utility"
)

const (
	alphabet   = u.Alphabet
	numbers    = u.Numbers
	blankSpace = u.BlankSpace
	emptySpace = u.EmptySpace
)

var Signs = u.Signs

// ! The Add Func.

func FullFields(x, gx string) error {
	switch {
	case x == emptySpace && gx == emptySpace:
		return m.Error("Fill in both input fields.")
	case x == emptySpace:
		return m.Error("Fill in the field 'x'.")
	case gx == emptySpace:
		return m.Error("Fill in the field 'gx'.")
	default:
		return nil
	}
}

func ReplaceSet(w *string, new string, set ...string) {
	for _, item := range set {
		*w = strings.ReplaceAll(*w, item, new)
	}
}

func GetInconrrectSyntax(w *string, keyword string) {
	set := strings.Split(keyword, emptySpace)
	ReplaceSet(w, emptySpace, set...)
}

func CorrectFieldSyntax(x, gx string) error {
	switch {
	case x != emptySpace && gx != emptySpace:
		return m.Error("Invalid fields.")
	case x != emptySpace:
		return m.Errorf("'%s' is not valid.", x)
	case gx != emptySpace:
		return m.Errorf("'%s' is not valid.", gx)
	default:
		return nil
	}
}

func HasIncorrectSyntax(x, gx string) error {
	w := x
	GetInconrrectSyntax(&x, alphabet)
	GetInconrrectSyntax(&gx, w+numbers+Signs)

	err := CorrectFieldSyntax(x, gx)
	if err != nil {
		return err
	}

	return nil
}

func CheckSyntax(x, gx string) error {
	if len(x) > 1 {
		return m.Error("Only one literal is accepted.")
	}

	err := HasIncorrectSyntax(x, gx)
	if err != nil {
		return err
	}

	return nil
}

func RmUnnecessarySpacesSigns(gx *string) {
	*gx = strings.Trim(*gx, blankSpace)
	*gx = strings.ReplaceAll(*gx, "- ", "-")
	ReplaceSet(gx, emptySpace, "+ ", "+", "^")
}

func SplitBySpaces(gx string) []string {
	polynomial := strings.Split(gx, blankSpace)
	return polynomial
}

func FullPolynomial(x string, polynomial *[]string) {
	for i, monomial := range *polynomial {
		FullMonomial(x, &monomial)
		(*polynomial)[i] = monomial
	}
}

func HasPrefixOne(x string, monomial *string) {
	positivePrefix := strings.HasPrefix(*monomial, x)
	negativePrefix := strings.HasPrefix(*monomial, "-"+x)

	if positivePrefix || negativePrefix {
		*monomial = strings.Replace(*monomial, x, "1"+x, 1)
	}
}

func HasSuffixOne(x string, monomial *string) {
	if strings.HasSuffix(*monomial, x) {
		*monomial = strings.Replace(*monomial, x, x+"1", 1)
	}
}

func FullMonomial(x string, monomial *string) {
	if strings.Contains(*monomial, x) {
		HasPrefixOne(x, monomial)
		HasSuffixOne(x, monomial)
	} else {
		*monomial += x + "0"
	}
}

func StoreSetsNandKN(x string, polynomial []string) ([]float64, [][]float64) {
	setN := make([]float64, len(polynomial))
	setKN := make([][]float64, len(polynomial))

	for i, monomial := range polynomial {
		kStr, nStr, _ := strings.Cut(monomial, x)
		kFloat, _ := strconv.ParseFloat(kStr, 64)
		nFloat, _ := strconv.ParseFloat(nStr, 64)
		setN[i], setKN[i] = nFloat, []float64{kFloat, nFloat}
	}

	return setN, setKN
}

func RmDuplicateValues(setN *[]float64) {
	s, dicc := []float64{}, make(map[float64]bool)
	for _, n := range *setN {
		if key := dicc[n]; !key {
			dicc[n] = true
			s = append(s, n)
		}
	}
	*setN = s
}

func FromHighToLow(s *[]float64) {
	var temp float64
	for j := range *s {
		for i := range *s {
			if (*s)[j] > (*s)[i] {
				temp = (*s)[j]
				(*s)[j] = (*s)[i]
				(*s)[i] = temp
			}
		}
	}
}

func EqualBaseMonomials(n2 float64, setKN [][]float64) [][]float64 {
	monomials := [][]float64{}
	for i := range setKN {
		k1, n1 := setKN[i][0], setKN[i][1]
		if n2 == n1 {
			monomials = append(monomials, []float64{k1, n1})
		}
	}
	return monomials
}

func SumMonomials(n float64, monomials [][]float64) []float64 {
	var sumK float64
	for i := range monomials {
		sumK += monomials[i][0]
	}
	return []float64{sumK, n}
}

func SimplifyKN(setN []float64, setKN *[][]float64) {
	t1, temp := 0, make([][]float64, len(setN))
	for _, n := range setN {
		monomials := EqualBaseMonomials(n, *setKN)
		temp[t1] = SumMonomials(n, monomials)
		t1++
	}
	*setKN = temp
}

func AddSpaceSigns(kxn string) string {
	if !strings.HasPrefix(kxn, "-") {
		kxn = "+" + kxn
	}
	return blankSpace + kxn
}

func BuildKxn(x string, kFloat, nFloat float64) string {
	k, n := fmt.Sprint(kFloat), fmt.Sprint(nFloat)
	if !(k == "0") {
		return k + x + "^" + n
	}

	return emptySpace
}

func BuildPolynomial(kxn string, i int, polynomial *string) {
	if i == 0 {
		*polynomial = kxn
	} else {
		*polynomial += AddSpaceSigns(kxn)
	}
}

func RebuildFunc(x, gx *string, setKN [][]float64) {
	var polynomial string
	for i, monomial := range setKN {
		kxn := BuildKxn(*x, monomial[0], monomial[1])
		BuildPolynomial(kxn, i, &polynomial)
	}

	if polynomial != emptySpace {
		*gx = polynomial
	} else {
		*gx = "0"
	}
}

// ! The Mutiply Func.

func CheckParentheses(gx, opn, cls string) error {
	mathEquality := strings.Count(gx, opn) == strings.Count(gx, cls)
	if !mathEquality {
		return m.Error("The number of parentheses is unequal.")
	}
	return nil
}

func SplitByParentheses(gx, opn, cls string) []string {
	ReplaceSet(&gx, "|", cls+opn, cls+blankSpace+opn)
	ReplaceSet(&gx, emptySpace, cls, opn)
	return strings.Split(gx, "|")
}

func SetFullPolynomial(x string, setPolynomial *[]string) {
	for i := range *setPolynomial {
		polynomial := SplitBySpaces((*setPolynomial)[i])
		FullPolynomial(x, &polynomial)
		(*setPolynomial)[i] = strings.Join(polynomial, blankSpace)
	}
}

func StoreSetsKN(x string, polynomial []string) ([]float64, []float64) {
	setK := make([]float64, len(polynomial))
	setN := make([]float64, len(polynomial))

	for i, monomial := range polynomial {
		k, n, _ := strings.Cut(monomial, x)
		kFloat, _ := strconv.ParseFloat(k, 64)
		nFloat, _ := strconv.ParseFloat(n, 64)
		setK[i], setN[i] = kFloat, nFloat
	}

	return setK, setN
}

func GetGroupsKN(x string, setPolynomial []string) ([][]float64, [][]float64) {
	groupSetK := make([][]float64, len(setPolynomial))
	groupSetN := make([][]float64, len(setPolynomial))

	for i := range setPolynomial {
		polynomial := SplitBySpaces(setPolynomial[i])
		setK, setN := StoreSetsKN(x, polynomial)
		groupSetK[i], groupSetN[i] = setK, setN
	}

	return groupSetK, groupSetN
}

func OperationsBetweenSets(symbol string, setA, setB []float64) (setC []float64) {
	for _, monomial_A := range setA {
		for _, monomial_B := range setB {
			if symbol == "+" {
				setC = append(setC, monomial_A+monomial_B)
			}
			if symbol == "*" {
				setC = append(setC, monomial_A*monomial_B)
			}
		}
	}
	return
}

func OperationsBetweenGroups(symbol string, group [][]float64) []float64 {
	n := len(group) - 1
	for i, j := 0, 1; i < n; i, j = i+1, j+1 {
		group[j] = OperationsBetweenSets(symbol, group[i], group[j])
	}
	return group[n]
}

func SimplifyGroupKN_Multiply(groupK, groupN [][]float64) ([]float64, []float64) {
	setK := OperationsBetweenGroups("*", groupK)
	setN := OperationsBetweenGroups("+", groupN)
	return setK, setN
}

func SingleSet(setK, setN []float64) [][]float64 {
	setKN := make([][]float64, len(setN))
	for i := range setN {
		setKN[i] = []float64{setK[i], setN[i]}
	}
	return setKN
}

// ! The Add Func [derivatives].

func KxnRuleSetKN(x string, polynomial *[][]float64) {
	t, n := 0, len(*polynomial)
	for _, monomial := range *polynomial {
		k := monomial[0] * monomial[1]
		n := monomial[1] - 1
		if !(k == 0) {
			(*polynomial)[t] = []float64{k, n}
			t++
		}
	}

	*polynomial = append((*polynomial)[:t], (*polynomial)[n:]...)
	temporal := make([][]float64, t)
	for i := range temporal {
		temporal[i] = (*polynomial)[i]
	}

	*polynomial = temporal
}
