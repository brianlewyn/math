package arithmetic

import c "github.com/brianlewyn/math/tools/component"

func Multiply(x, gx *string) error {

	var opn, cls string = "(", ")"
	c.Signs = "-+ .^" + opn + cls

	err := c.FullFields(*x, *gx)
	if err != nil {
		return err
	}

	err = c.CheckSyntax(*x, *gx)
	if err != nil {
		return err
	}

	err = c.CheckParentheses(*gx, opn, cls)
	if err != nil {
		return err
	}

	c.RmUnnecessarySpacesSigns(gx)
	setPolynomial := c.SplitByParentheses(*gx, opn, cls)

	c.SetFullPolynomial(*x, &setPolynomial)
	groupK, groupN := c.GetGroupsKN(*x, setPolynomial)
	setK, setN := c.SimplifyGroupKN_Multiply(groupK, groupN)

	setKN := c.SingleSet(setK, setN)
	c.RebuildFunc(x, gx, setKN)

	return nil
}
