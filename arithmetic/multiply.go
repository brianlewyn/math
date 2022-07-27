package arithmetic

import "github.com/brianlewyn/math/tools/component"

func Multiply(x, gx *string) error {

	var opn, cls string = "(", ")"
	component.Signs = "-+ .^" + opn + cls

	err := component.FullFields(*x, *gx)
	if err != nil {
		return err
	}

	err = component.CheckSyntax(*x, *gx)
	if err != nil {
		return err
	}

	err = component.CheckParentheses(*gx, opn, cls)
	if err != nil {
		return err
	}

	component.RmUnnecessarySpacesSigns(gx)
	setPolynomial := component.SplitByParentheses(*gx, opn, cls)

	component.SetFullPolynomial(*x, &setPolynomial)
	groupK, groupN := component.GetGroupsKN(*x, setPolynomial)
	setK, setN := component.SimplifyGroupKN_Multiply(groupK, groupN)

	setKN := component.SingleSet(setK, setN)
	component.RebuildFunc(x, gx, setKN)

	return nil
}
