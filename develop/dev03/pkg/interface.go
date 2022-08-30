package pkg

// Интерфейс стратегии сортировки
type Operation interface {
	Exec(elems *[]string)
}
