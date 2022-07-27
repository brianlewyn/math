package arithmetic

import "github.com/brianlewyn/math/tools/component"

func Add(x, gx *string) error {

	err := component.FullFields(*x, *gx)
	if err != nil {
		return err
	}

	err = component.CheckSyntax(*x, *gx)
	if err != nil {
		return err
	}

	component.RmUnnecessarySpacesSigns(gx)
	polynomial := component.SplitBySpaces(*gx)

	component.FullPolynomial(*x, &polynomial)
	setN, setKN := component.StoreSetsNandKN(*x, polynomial)

	component.RmDuplicateValues(&setN)
	component.FromHighToLow(&setN)

	component.SimplifyKN(setN, &setKN)
	component.RebuildFunc(x, gx, setKN)

	return nil
}
