package basic

import "github.com/brianlewyn/math/tool"

func Add(x, gx *string) error {

	err := tool.FullFields(*x, *gx)
	if err != nil {
		return err
	}

	err = tool.CheckSyntax(*x, *gx)
	if err != nil {
		return err
	}

	tool.RmUnnecessarySpacesSigns(gx)
	polynomial := tool.SplitBySpaces(*gx)

	tool.FullPolynomial(*x, &polynomial)
	setN, setKN := tool.StoreSetsNandKN(*x, polynomial)

	tool.RmDuplicateValues(&setN)
	tool.FromHighToLow(&setN)

	tool.SimplifyKN(setN, &setKN)
	tool.RebuildFunc(x, gx, setKN)

	return nil
}
