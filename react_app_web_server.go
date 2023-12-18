package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Data struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func setCustomPath() {
	// Specify your custom PATH here
	customPath := "/usr/local/bin:/usr/bin:/bin"

	// Get the current environment PATH
	currentPath := os.Getenv("PATH")

	// Append the custom path to the current path, separating them with a colon (for Unix-like systems)
	newPath := currentPath + ":" + customPath

	// Set the new environment PATH
	os.Setenv("PATH", newPath)
}

func getWho() string {
	cmd := exec.Command("who")
	output, err := cmd.Output()
	if err != nil {
		return err.Error()
	}

	// Convert the output byte slice to a string
	outputStr := string(output)

	return outputStr

	// Q8 uses busybox who, which doesn't support the -h option and ip address is in the last field.

	// // Split the output into lines
	// lines := strings.Split(outputStr, "\n")

	// // Extract the third field from each line and store them in a slice
	// var thirdFields []string
	// for _, line := range lines {
	// 	fields := strings.Fields(line)
	// 	if len(fields) >= 3 {
	// 		thirdFields = append(thirdFields, fields[2])
	// 	}
	// }

	// // Join the third fields into a single string
	// thirdFieldString := strings.Join(thirdFields, " ")

	// return thirdFieldString
	// // return strings.TrimSpace(string(output))
}
func runCommand(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		return err.Error()
	}

	return strings.TrimSpace(string(output))
}

func sensorInfoHandler(w http.ResponseWriter, r *http.Request) {
	input := runCommand("sensors")
	data := parseTemperatureData(input)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func _sensorInfoHandler(w http.ResponseWriter, r *http.Request) {
	// data := map[string]map[string]interface{}{
	// 	"ps_temp": {
	// 		"value": 99,
	// 		"unit":  "Celsius",
	// 	},
	// 	"pl_temp": {
	// 		"value": 12,
	// 		"unit":  "Celsius",
	// 	},
	// 	"remote_temp": {
	// 		"value": 13,
	// 		"unit":  "Celsius",
	// 	},
	// }

	data := get_temperatures()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "1234")
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	// Redirect to the same host and path with HTTPS scheme
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
}

func main() {

	setCustomPath()

	// Set the path to the React build directory
	reactBuildDir := os.Getenv("REACT") // "./client/build"
	if reactBuildDir == "" {
		reactBuildDir = "./client/build"
	}
	// Create a file server handler for the React static files
	fs := http.FileServer(http.Dir(reactBuildDir))

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if filepath.Ext(r.URL.Path) == ".css" {
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
			fmt.Println("CSS")

		}
		fmt.Println("URL: ", r.URL.Path)

		fs.ServeHTTP(w, r)
	}

	// Set the handler function for all routes
	http.HandleFunc("/", handler)

	// set the handler for system info
	http.HandleFunc("/deviceSystemInfo", func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Entries: []Entry{
				{Name: "webserver version", Value: "0.11"},
				{Name: "device name", Value: runCommand("uname", "-a")},
				{Name: "system time", Value: runCommand("date")},
				{Name: "uptime", Value: runCommand("uptime")},
				{Name: "who", Value: getWho()},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	// set the handler for system info
	http.HandleFunc("/deviceSensorInfo", sensorInfoHandler)

	http.HandleFunc("/test", _sensorInfoHandler)

	// Get the port from the environment variable or use the default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the web server
	if port == "80" {
		log.Printf("Server started on http://localhost:%s", port)
		log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(redirectToHTTPS)))
	}
	// Start HTTPS Web Server
	// Use the paths to your certificate and key
	if port == "443" {
		log.Printf("Starting HTTPS Server")
		err := http.ListenAndServeTLS(":"+port, "/etc/letsencrypt/live/wanxuaneducation.com/fullchain.pem",
			"/etc/letsencrypt/live/wanxuaneducation.com/privkey.pem", nil)
		if err != nil {
			log.Fatal("ListenAndServeTLS: ", err)
		}
	}
	log.Fatal("Unsupported web server port: ", port)

}
