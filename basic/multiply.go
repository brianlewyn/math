package basic

import "github.com/brianlewyn/math/tool"

func Multiply(x, gx *string) error {

	var opn, cls string = "(", ")"
	tool.E_SIGNS = "-+ .^" + opn + cls

	err := tool.FullFields(*x, *gx)
	if err != nil {
		return err
	}

	err = tool.CheckSyntax(*x, *gx)
	if err != nil {
		return err
	}

	err = tool.CheckParentheses(*gx, opn, cls)
	if err != nil {
		return err
	}

	tool.RmUnnecessarySpacesSigns(gx)
	setPolynomial := tool.SplitByParentheses(*gx, opn, cls)

	tool.SetFullPolynomial(*x, &setPolynomial)
	groupK, groupN := tool.GetGroupsKN(*x, setPolynomial)
	setK, setN := tool.SimplifyGroupKN(groupK, groupN)

	setKN := tool.SingleSet(setK, setN)
	tool.RebuildFunc(x, gx, setKN)

	return nil
}
