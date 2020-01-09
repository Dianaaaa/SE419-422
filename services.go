package main

import (
	"fmt"
	"crypto/md5"
	"time"
	"math/rand"
	"strconv"
)

var tenToAny map[byte]string = map[byte]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
	16: "g",
	17: "h",
	18: "i",
	19: "j",
	20: "k",
	21: "l",
	22: "m",
	23: "n",
	24: "o",
	25: "p",
	26: "q",
	27: "r",
	28: "s",
	29: "t",
	30: "u",
	31: "v",
	32: "w",
	33: "x",
	34: "y",
	35: "z",
	36: "A",
	37: "B",
	38: "C",
	39: "D",
	40: "E",
	41: "F",
	42: "G",
	43: "H",
	44: "I",
	45: "J",
	46: "K",
	47: "L",
	48: "M",
	49: "N",
	50: "O",
	51: "P",
	52: "Q",
	53: "R",
	54: "S",
	55: "T",
	56: "U",
	57: "V",
	58: "W",
	59: "X",
	60: "Y",
	61: "Z",
	62: "_",
	63: "-"}

func get_md5_string(str string) []byte {
	h := md5.New()
    h.Write([]byte(str))
    return h.Sum(nil)
}

func to_short_url(ip string) string {
	const base_format = "2006-01-02 15:04:05"
	t := time.Now()
	rand.Seed(time.Now().Unix())
	r := rand.Int()
	str := ip + t.Format(base_format) + strconv.Itoa(r)
	fmt.Println("str ", str)
	md5_code := get_md5_string(str)
	fmt.Println("md5_str ", md5_code)
	code_len := len(md5_code)
	fmt.Println("code len ", code_len)
	url := ""
	for j := 0; j < code_len; j += 2 {
		index1 := md5_code[j]
		index2 := md5_code[j+1]

		index := (index1 + index2) % 64
		fmt.Println("index str", index)
		url+=tenToAny[index]
	} 
	fmt.Println("url ", url)
	return url
}

func Gen_short_url(ip string, path string) string {
	var shortlink string
	shortlink = to_short_url(ip)
	for !InsertUrl(shortlink, path) {
		shortlink = to_short_url(ip)
	}
	return shortlink
}

func Get_path(shortlink string) string {
	var path string
	path = GetPath(shortlink)
	fmt.Println("shortlink get_path: ", shortlink)
	return path
}

