package pkg

// Operation Интерфейс стратегии сортировки
type Operation interface {
	Exec(elems *[]string)
}
