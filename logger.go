package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
	"time"
	"strconv"
	"encoding/csv"
)

var init_time = 0.0
var file *os.File = nil
var writer *csv.Writer = nil
var err error = nil

var bandwidth_file *os.File = nil
var bandwidth_writer *csv.Writer = nil
var bandwidth_err error = nil

func Logger(connection net.Conn, client string) {
	for {
		reader := bufio.NewReader(connection)
		bytes_read, err := reader.ReadString('\n')
		if err != nil {
			// fmt.Fprintf(os.Stderr, "Disconnect")
			// current_time := time.Now().Unix()
			current_time := time.Now()
			time_in_seconds := float64(current_time.Unix())
			time_in_nanoseconds := float64(current_time.Nanosecond()) / 1000000000.0
			total_time := strconv.FormatFloat((time_in_seconds + time_in_nanoseconds), 'f', 7, 64)
			fmt.Fprintf(os.Stdout, total_time + " - " + client + "disconnected\n")
			return
		}

		res := strings.Split(bytes_read, " ")
		timestamp, _ := strconv.ParseFloat(res[0], 64)

		current_time := time.Now()
		time_in_seconds := float64(current_time.Unix())
		time_in_nanoseconds := float64(current_time.Nanosecond()) / 1000000000.0
		total_time := float64(time_in_seconds) + time_in_nanoseconds

		diff_time := time_in_seconds - init_time // seconds since logger started
		delay := total_time - timestamp // delay since event generated

		csv_write := []string{strconv.FormatFloat(diff_time, 'f', 7, 64), strconv.FormatFloat(delay, 'f', 7, 64)}
		writer.Write(csv_write)
		writer.Flush()

		bandwidth := []string{strconv.FormatFloat(diff_time, 'f', 7, 64), strconv.FormatFloat(float64(len(bytes_read)*8), 'f', 7, 64)}
		bandwidth_writer.Write(bandwidth)
		bandwidth_writer.Flush()

		// fmt.Println(bytes_read)
		fmt.Fprintf(os.Stdout, bytes_read)

		// if diff_time >= 95 {
		// 	fmt.Fprintf(os.Stderr, "100+ seconds\n")
		// }
	}
	connection.Close()
}

func main() {
	args := os.Args
	if len(args) != 2 {
		return
	}

	file, err = os.Create("delay.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "File not created")
	}
	defer file.Close()

	writer = csv.NewWriter(file)
	defer writer.Flush()

	bandwidth_file, bandwidth_err = os.Create("bandwidth.csv")
	if bandwidth_err != nil {
		fmt.Fprintf(os.Stderr, "File not created")
	}
	defer bandwidth_file.Close()

	bandwidth_writer = csv.NewWriter(bandwidth_file)
	defer bandwidth_writer.Flush()

	port := args[1]
	server, err := net.Listen("tcp", ":" + port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen - Listen error")
		return
	}
	defer server.Close()

	init_time = float64(time.Now().Unix())

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Accept failure")
			return
		}

		reader := bufio.NewReader(connection)
		client, err := reader.ReadString(' ')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Incomplete connection - missing name")
			return
		}

		current_time := time.Now()
		time_in_seconds := float64(current_time.Unix())
		time_in_nanoseconds := float64(current_time.Nanosecond()) / 1000000000.0
		total_time := strconv.FormatFloat((time_in_seconds + time_in_nanoseconds), 'f', 7, 64)
		fmt.Fprintf(os.Stdout, total_time + " - " + client + "connected\n")

		go Logger(connection, client)
	}
}

