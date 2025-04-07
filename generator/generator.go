package generator

type Generator interface {
	GenerateType(string) string
}
