package dynamicArray

type DynamicArray[T any] interface {
	Push(value T) error
	Pop() (T, error)
	Add(value T, index int) error
	Remove(index int) (T, error)
	Get(int int) (T, error)
	Set(value T, int int) error
	Length() int
	Capacity() int
}
