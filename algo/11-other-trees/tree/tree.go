package tree

type Tree interface {
	Insert(val int)
	Search(val int) bool
	Remove(val int)
	DumpValuesInDetails()
}
