package derivatives

import c "github.com/brianlewyn/math/tools/component"

func Add(x, gx *string) error {

	err := c.FullFields(*x, *gx)
	if err != nil {
		return err
	}

	err = c.CheckSyntax(*x, *gx)
	if err != nil {
		return err
	}

	c.RmUnnecessarySpacesSigns(gx)
	polynomial := c.SplitBySpaces(*gx)

	c.FullPolynomial(*x, &polynomial)
	_, setKN := c.StoreSetsNandKN(*x, polynomial)

	c.KxnRuleSetKN(*x, &setKN)
	c.RebuildFunc(x, gx, setKN)

	return nil
}
