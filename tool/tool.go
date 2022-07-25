package tool

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianlewyn/math/msg"
)

// ! My constants.
const ALPHABET string = "abcdefghijklmnopqrstuvwxyz"
const NUMBERS string = "0123456789"

// ! My variables.
var E_SIGNS string = "-+ .^"

// ! The Add Func.

func FullFields(x, gx string) error {
	switch {
	case x == "" && gx == "":
		return msg.Error("Fill in both input fields.")
	case x == "":
		return msg.Error("Fill in the field 'x'.")
	case gx == "":
		return msg.Error("Fill in the field 'gx'.")
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
	set := strings.Split(keyword, "")
	ReplaceSet(w, "", set...)
}

func CorrectFieldSyntax(x, gx string) error {
	switch {
	case x != "" && gx != "":
		return msg.Error("Invalid fields.")
	case x != "":
		return msg.Errorf("'%s' is not valid.", x)
	case gx != "":
		return msg.Errorf("'%s' is not valid.", gx)
	default:
		return nil
	}
}

func HasIncorrectSyntax(x, gx string) error {

	w := x
	GetInconrrectSyntax(&x, ALPHABET)
	GetInconrrectSyntax(&gx, w+NUMBERS+E_SIGNS)

	err := CorrectFieldSyntax(x, gx)
	if err != nil {
		return err
	}

	return nil
}

func CheckSyntax(x, gx string) error {

	if len(x) > 1 {
		return msg.Error("Only one literal is accepted.")
	}

	err := HasIncorrectSyntax(x, gx)
	if err != nil {
		return err
	}

	return nil
}

func RmUnnecessarySpacesSigns(gx *string) {
	*gx = strings.Trim(*gx, " ")
	*gx = strings.ReplaceAll(*gx, "- ", "-")
	ReplaceSet(gx, "", "+ ", "+", "^")
}

func SplitBySpaces(gx string) []string {
	polynomial := strings.Split(gx, " ")
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
	var temporl float64
	for x := range *s {
		for y := range *s {
			if (*s)[x] > (*s)[y] {
				temporl = (*s)[x]
				(*s)[x] = (*s)[y]
				(*s)[y] = temporl
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

	var k float64
	for i := range monomials {
		k += monomials[i][0]
	}

	return []float64{k, n}
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
	return " " + kxn
}

func BuildKxn(x string, kFloat, nFloat float64) string {

	k := fmt.Sprint(kFloat)
	n := fmt.Sprint(nFloat)

	if !(k == "0") {
		return k + x + "^" + n
	}

	return ""
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

	if polynomial != "" {
		*gx = polynomial
	} else {
		*gx = "0"
	}
}

// ! The Mutiply Func.

func CheckParentheses(gx, opn, cls string) error {
	mathEquality := strings.Count(gx, opn) == strings.Count(gx, cls)
	if !mathEquality {
		return msg.Error("The number of parentheses is unequal.")
	}
	return nil
}

func SplitByParentheses(gx, opn, cls string) []string {
	ReplaceSet(&gx, "|", cls+opn, cls+" "+opn)
	ReplaceSet(&gx, "", cls, opn)
	return strings.Split(gx, "|")
}

func SetFullPolynomial(x string, setPolynomial *[]string) {
	for i := range *setPolynomial {
		polynomial := SplitBySpaces((*setPolynomial)[i])
		FullPolynomial(x, &polynomial)
		(*setPolynomial)[i] = strings.Join(polynomial, " ")
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

func SimplifyGroupKN(groupK, groupN [][]float64) ([]float64, []float64) {
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
