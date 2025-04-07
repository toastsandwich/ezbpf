package generator

// TODO try using a visitor
type CGenerator struct{}

func NewCGenerator() *CGenerator {
	return &CGenerator{}
}

func (cg *CGenerator) GenerateType(type_ string) string {
	ctype := ""
	switch type_ {
	case "ethhdr":
		ctype = "struct ethhdr *"
	case "iphdr":
		ctype = "struct iphdr *"
	case "tcphdr":
		ctype = "struct tcphdr *"
	case "udphdr":
		ctype = "struct udphdr *"
	default:
		ctype = type_
	}
	return ctype
}
