package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"servergps/server"
	"time"
)

const MSG_LOGIN = 0x01
const MSG_POSITION = 0x12
const MSG_ALARM = 0x16
const MSG_HEARBET = 0x13
const MSG_RESPONSE = 0x15

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello GopherCon Israel 2019!")
}

func json(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This json data")
}

func main() {
	runtime.GOMAXPROCS(1)
	bytes, err := hex.DecodeString("78780D01035341353215036200022D060D0A")
	if err != nil {
		return
	}

	answerLogin(bytes)
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%02X", bytes[i])
	}
	fmt.Println("")
	var input string
	fmt.Scanln(&input)
}
func parseLogin(bytes []byte) {
	if bytes[0] == 0x78 && bytes[1] == 0x78 {
		var length = bytes[2]
		var protocol = bytes[3]
		var imei = fmt.Sprintf("%02X", bytes[4:12])
		fmt.Println("Length=>", length, ",Protocol=>", protocol)
		fmt.Println("Imei=>", imei)
	}
}
func answerLogin(bytes []byte) {
	var length uint8 = 5
	//length = 5
	//7878 0D 01 0353413532150362 0002 2D06 0D0A
	var response [10]uint8
	response[0] = 0x78
	response[1] = 0x78
	response[2] = length
	response[3] = MSG_LOGIN
	response[4] = bytes[12]
	response[5] = bytes[13]
	var calc = response[2:6]
	var crc = server.Crc16(calc)
	fmt.Printf("CRC %04X \r\n", crc)
	fmt.Printf("CRC %02X \r\n", (crc >> 8))
	fmt.Printf("CRC %02X \r\n", (crc & 0xFF))
	response[6] = uint8(crc >> 8)
	response[7] = uint8(crc & 0xFF)
	response[8] = 0X0D
	response[9] = 0X0A
	//7878 05 01 0002 B447 0D0A
	for i := 0; i < len(response); i++ {
		fmt.Printf("%02X", response[i])
	}

}
func cetak(t int) {
	var ctr = 0
	for ctr < 50 {
		fmt.Println("Thread ", t, ". Msg =>", ctr)

		time.Sleep(1 * time.Second)
		ctr++
	}
}
