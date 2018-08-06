package main

import "fmt"
import "os"
//import "io"
import "bufio"
import "strings"

// Our global map of config variables
var config map[string] string

// Global debug variables
var DBG_CONFIG, DBG_DNS, DBG_HTTP, DBG_STDIO bool

// Simple check func 
func check(e error){
	if e != nil {
		panic(e)
	}
}

// Simple debug func
func dbg(d bool, str string){
	if d{
		fmt.Println(str)		
	}
}

// Read configuration from passed config filename
func readConfig(fName string){
	// Init global defaults
	config = make(map[string]string)
	config["Port"] = "5000"
	config["SpoolDirectory"] = "/var/spool/derps"
	config["Debug"] = "" // Value set here will never be used, set defaults at next line
	DBG_CONFIG, DBG_DNS, DBG_HTTP, DBG_STDIO = false, false, false, false

	fd, err := os.Open(fName)
	check(err)
	defer fd.Close()

	confScanner := bufio.NewScanner(fd)
	for confScanner.Scan(){
		l := strings.TrimSpace(confScanner.Text()) 
		
		if len(l) == 0 {
			continue
		}
		if strings.Index(l, "#") == 0 {
			continue
		}

		kv := strings.Split(l, "=")
		if len(kv) == 1 {
			continue
		}else{
			k := strings.TrimSpace(kv[0])
			v := strings.TrimSpace(kv[1])
			
			if _, exists := config[k]; exists {
				switch {
				case k == "Debug":
					v = strings.ToUpper(v)
					config[k] = v
					if strings.Contains(v, "CONFIG") {
						DBG_CONFIG = true
					}
					if strings.Contains(v, "DNS") {
						DBG_DNS = true
					}
					if strings.Contains(v, "HTTP") {
						DBG_HTTP = true
					}
					if strings.Contains(v, "STDIO") {
						DBG_STDIO = true
					}

				default: // Handle all single value strings here
					config[k] = v
				}
			}
		}
	}
	
	for k, v := range config {
		dbg(DBG_CONFIG, "Key:" + k + " Value:" + v)
	}
}

