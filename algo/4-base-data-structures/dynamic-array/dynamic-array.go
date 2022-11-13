package dynamicArray

type DynamicArray[T any] interface {
	Push(item T) error
	Pop() (T, error)
	Add(item T, index int) error
	Remove(index int) (T, error)
	Get(int int) (T, error)
	Length() int
	Capacity() int
}
