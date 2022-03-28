package name

type Name struct {
	MaleFirstnames   []string
	MaleLastnames    []string
	FemaleFirstnames []string
	FemaleLastnames  []string
}

func New() Name {
	return Name{
		MaleFirstnames,
		MaleLastnames,
		FemaleFirstnames,
		FemaleLastnames,
	}
}
