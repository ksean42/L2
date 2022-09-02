Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Выведутся числа и за ними бесконечное количество нулей.
Так происходит потому что в блоке select мы не проверяем каналы на закрытость и постоянно получаем zero value и не закрываем канал c
Пример исправленной функции merge:

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
	Loop:
		for {
			select {
			case v, ok := <-a:
				if ok {
					c <- v

				} else {
					break Loop
				}
			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					break Loop
				}
			}
		}
		close(c)
	}()
	return c
}


```
