package atomicity

type AtomicSection[T any] func(execAtomic ExecuteAtomicTask[T]) ([]T, *AtomicityError[T])

type ExecuteAtomicTask[T any] func(runnable Runnable[T])

type Runnable[T any] func(do Do[T])

type Do[T any] func() (T, error)