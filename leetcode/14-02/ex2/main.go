package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

type SyncGroupError struct {
	s []error
}

// https://pkg.go.dev/errors#Join
func (sge *SyncGroupError) String() string {
	return errors.Join(sge.s...).Error()
}

// SyncGroup представляет примитив синхронизации горутин
type SyncGroup struct {
	err SyncGroupError
	wg  sync.WaitGroup
	ch  chan int
	inc atomic.Int32
}

// Go запускает асинхронную выполнение функции и ловит паники
func (sg *SyncGroup) Go(f func() error) {
	sg.wg.Add(1)
	go func() {
		defer func() {
			sg.wg.Done()
			if r := recover(); r != nil {
				sg.err.s = append(sg.err.s, fmt.Errorf("panic: %v", r))
			}
		}()

		err := f()
		if err != nil {
			sg.err.s = append(sg.err.s, err)
		}
	}()
}

// Wait ожидает завершения всех горутин и возвращает все ошибки, которые произошли во время обработки задач горутинами
func (sg *SyncGroup) Wait() error {
	sg.wg.Wait()
	if len(sg.err.s) != 0 {
		return errors.Join(sg.err.s...)
	}

	return nil
}

func main() {
	// Пример использования SyncGroup
	sg := &SyncGroup{}

	// Запуск горутин
	sg.Go(func() error {
		fmt.Println("Горутина 1")
		return nil
	})

	sg.Go(func() error {
		fmt.Println("Горутина 2")
		return fmt.Errorf("Ошибка в горутине 2")
	})

	sg.Go(func() error {
		fmt.Println("Горутина 3")
		panic("Паника в горутине 3")
	})

	// Ожидание завершения всех горутин и получение ошибки
	if err := sg.Wait(); err != nil {
		fmt.Println("Произошла ошибка:", err)
	} else {
		fmt.Println("Все горутины завершились без ошибок")
	}
}
