/*
 * File         : main.go
 * Project      : NetMon
 * Created Date : Sunday, Oct 4th 2020, 10:57:11 PM
 * Author       : Pramod Devireddy
 *
 * Last Modified: Tuesday, 20th October 2020 9:53:07 pm
 * Modified By  : Pramod Devireddy
 *
 * Copyright (c)2020 - Pramod Devireddy
 * ************************* Description *************************
 *
 * ***************************************************************
 */

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/handlers"
)

type traffic struct {
	Time  []string `json:"time"`
	Bytes []int    `json:"bytes"`
}

var rxLogFile, txLogFile *string

func init() {
	rxLogFile = flag.String("rx", "", "Receive Bytes Log File")
	txLogFile = flag.String("tx", "", "Transmit Bytes Log File")
}

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/rx_log_data", serveRxLogData)
	mux.HandleFunc("/tx_log_data", serveTxLogData)
	err := http.ListenAndServe(":9989", handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}

func serveRxLogData(w http.ResponseWriter, r *http.Request) {
	logData := extractTrafficLogData(rxLogFile)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(logData)
}

func serveTxLogData(w http.ResponseWriter, r *http.Request) {
	logData := extractTrafficLogData(txLogFile)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(logData)
}

func extractTrafficLogData(logfile *string) traffic {

	var logData traffic

	file, err := os.Open(*logfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)

	var initTime time.Time

	for scanner.Scan() {
		line := scanner.Text()

		if i == 0 {
			initTime, _ = time.Parse("Mon Jan 2 15:04:05 MST 2006", line)
		} else {
			logData.Time = append(logData.Time, initTime.Add(time.Second*time.Duration(i)).Format("2006-01-02 15:04:05"))
			rxKbitsString := strings.Fields(strings.TrimSpace(line))[0]
			rxKbits, _ := strconv.Atoi(rxKbitsString)
			logData.Bytes = append(logData.Bytes, rxKbits)
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return logData
}

func serving() {

	logTimestamp := time.Now().Format("02-01-2006_15:04:05")
	logFileName := os.Getenv("HOME") + "/logs/BW_" + logTimestamp + ".log"
	logFile, err := os.OpenFile(logFileName, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// out, _ := exec.Command("ping", "192.168.1.100", "-c 1").Output()
	// if strings.Contains(string(out), "Destination Host Unreachable") {
	// 	fmt.Println("IT'S DOWN")
	// } else {
	// 	fmt.Println("IT'S ALIVEEE")
	// }

	timestampOp, _ := exec.Command("date").Output()
	timestampString := strings.TrimSpace(string(timestampOp))
	timestamp, _ := time.Parse("Mon Jan 2 15:04:05 MST 2006", timestampString)
	timestampFormattedString := timestamp.Format("2006-01-02 15:04:05")

	fmt.Println(timestampFormattedString)

	oldRx, _ := exec.Command("cat", "/sys/class/net/enp0s3/statistics/rx_bytes").Output()
	totalOldRxBytes, _ := strconv.Atoi(strings.TrimSpace(string(oldRx)))

	oldTx, _ := exec.Command("cat", "/sys/class/net/enp0s3/statistics/tx_bytes").Output()
	totalOldTxBytes, _ := strconv.Atoi(strings.TrimSpace(string(oldTx)))

	for {
		time.Sleep(1 * time.Second)
		nowRx, _ := exec.Command("cat", "/sys/class/net/enp0s3/statistics/rx_bytes").Output()
		totalNowRxBytes, _ := strconv.Atoi(strings.TrimSpace(string(nowRx)))
		rxKBits := (totalNowRxBytes - totalOldRxBytes) * 8 / 1000
		totalOldRxBytes = totalNowRxBytes
		fmt.Println("Rx: ", rxKBits, "kBits/s")

		nowTx, _ := exec.Command("cat", "/sys/class/net/enp0s3/statistics/tx_bytes").Output()
		totalNowTxBytes, _ := strconv.Atoi(strings.TrimSpace(string(nowTx)))
		txKBits := (totalNowTxBytes - totalOldTxBytes) * 8 / 1000
		totalOldTxBytes = totalNowTxBytes
		fmt.Println("Tx: ", txKBits, "kBits/s")
	}
}
