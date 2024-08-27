// iter.go
package iter

// Seq[T] is a sequence type that allows iteration over elements of type T.
type Seq[T any] func(yield func(T) bool)

// Channel converts a Seq[T] to a channel of T, allowing iteration using range.
func (s Seq[T]) Channel() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		s(func(item T) bool {
			ch <- item
			return true
		})
	}()
	return ch
}

// Iterator provides a channel-based iterator for Seq[T].
func (s Seq[T]) Iterator() <-chan T {
	return s.Channel()
}
