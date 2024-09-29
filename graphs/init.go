package graphs

type DAG[T comparable] map[T][]T

var (
	SampleGraph DAG[rune]
)

func init() {
	SampleGraph = DAG[rune]{
		'f': {'g', 'i'},
		'g': {'h'},
		'h': {},
		'i': {'g', 'k'},
		'j': {'i'},
		'k': {},
	}
}
