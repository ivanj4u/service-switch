/*
 * Copyright (c) 2018.
 */

package main

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ivanj4u/service-switch/framework"
)

func main() {
	log.Println("Starting Application")

	// Initialize Application
	framework.Init()
}


