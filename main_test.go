/*
 * File         : main_test.go
 * Project      : NetMon
 * Created Date : Tuesday, Oct 20th 2020, 7:27:59 PM
 * Author       : Pramod Devireddy
 *
 * Last Modified: Tuesday, 20th October 2020 10:07:26 pm
 * Modified By  : Pramod Devireddy
 *
 * Copyright (c)2020 - Pramod Devireddy
 * ************************* Description *************************
 *
 * ***************************************************************
 */

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/rx_log_data", serveRxLogData)
	router.HandleFunc("/tx_log_data", serveTxLogData)
	return router
}

func TestServeRxLogData(t *testing.T) {

	request, _ := http.NewRequest("GET", "/rx_log_data", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestServeTxLogData(t *testing.T) {

	request, _ := http.NewRequest("GET", "/tx_log_data", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
