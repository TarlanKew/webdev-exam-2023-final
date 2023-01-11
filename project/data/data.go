package data

import (
	"math/rand"
	"time"
)

/*type Route struct {
	Id          int64
	Name        string
	Description string
	MainObject  string
	Coords      [2][2]float64
}*/

type Routes struct {
	Arr [][]string
}

type Location struct {
	Start string
	Stop  string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (r *Routes) GetAll() {
	rand.Seed(time.Now().UnixNano())
	r.Arr = [][]string{}
	r.Arr = append(r.Arr, []string{
		"1",
		"Зоопарк-Арбат",
		"Прогулка от Московского зоопарка до нового Арбата",
		"Пешая прогулка от Московского зоопарка до нового Арбата",
		"Б. Грузинская ул., 1, Москва, 123242",
		"Новый Арбат ул., 7с1, Москва, 119019",
	})
	r.Arr = append(r.Arr, []string{
		"2",
		"Сити-Кремль",
		"Прогулка от Сити до Кремля",
		"Пешая прогулка от Сити до Кремля",
		"Москва, Пресненская набережная, 6с2, 123112",
		"Красная площадь, 7, Москва",
	})
	/*for i := 0; i < 10; i++ {
		tmp := []string{
			fmt.Sprintf("%v", time.Now().UnixNano()), // id
			RandStringBytes(10),                      // Name
			RandStringBytes(10),                      // Description
			RandStringBytes(10),                      // MainObject
			"ул. Прянишникова, 2А, Москва, 125008", // start
			"1-Y Baltiyskiy Pereulok",              // stop
		}
		r.Arr = append(r.Arr, tmp)
	}*/
}
