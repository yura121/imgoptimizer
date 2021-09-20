package main

import (
	"fmt"
	"imgoptimizer/config"
	"imgoptimizer/convert"
	"imgoptimizer/database"
	"os"
	"path/filepath"
)

func main() {
	cfg, cfgErr := config.New("config.yaml")
	if cfgErr != nil {
		fmt.Println("Не удалось прочитать конфигурационный файл:", cfgErr)
		return
	}

	args := os.Args[1:]
	src := args[0]
	dst := args[1]
	cnv := convert.New(cfg)
	cnvErr := cnv.ToWebp(src, dst)
	if cnvErr != nil {
		fmt.Println("Не удалось выполнить конвертацию:", cnvErr)
	} else {
		fmt.Println("Конвертация выполнена успешно:", src, "->", dst)
	}

	db, dbErr := database.New(cfg)
	if dbErr != nil {
		fmt.Println("Результат не был записан в БД:", dbErr)
	} else {
		_, fileName := filepath.Split(src)
		lastId, _ := db.Add(fileName)
		fmt.Println("Добавлена запись под номером", lastId)
	}
}
