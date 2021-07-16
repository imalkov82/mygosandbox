package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

/*
The main drawback of the singleton pattern
that it's ofter breaks the dependency inversion principle. (fixed by introducing DataBase interface)
*/

type DataBase interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int { return db.capitals[name] }

// sync.Once or init() -- thread safety
// laziness (can be garantied by using sync.Once) --
var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, err := readData(".\\capitals.txt")
		db := singletonDatabase{caps}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func readData(path string) (map[string]int, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	file, err := os.Open(exPath + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

func GetTotalPopulation(db DataBase, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (db *DummyDatabase) GetPopulation(city string) int {
	if len(db.dummyData) == 0 {
		db.dummyData = map[string]int{
			"alpha": 0,
			"beta":  1,
			"gamma": 2}
	}
	return db.dummyData[city]
}

func main() {
	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulation(&DummyDatabase{}, names)
	fmt.Println(tp == 2)
}
