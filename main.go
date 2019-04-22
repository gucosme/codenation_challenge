package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gucosme/codenation_challenge/data"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func julioDecypher(encrypted string, cases int) string {
	decyphered := make([]byte, len(encrypted))

	for i, e := range []byte(encrypted) {
		if e >= 'a' && e <= 'z' {
			if int(e)-cases < 'a' {
				casesLeft := cases - int(e-'a'+1)
				decyphered[i] = byte('z' - casesLeft)
			} else {
				decyphered[i] = byte(int(e) - cases)
			}
		} else {
			decyphered[i] = e
		}
	}

	return string(decyphered)
}

func main() {
	log.Println(">> APP STARTED")

	const answerFile string = "./answer.json"

	url := os.Getenv("URL")
	token := os.Getenv("TOKEN")
	getURL := fmt.Sprintf("%s/generate-data?token=%s", url, token)
	postURL := fmt.Sprintf("%s/submit-solution?token=%s", url, token)

	log.Printf(">> REQUESTING DATA TO %s\n", getURL)
	res, err := http.Get(getURL)
	check(err)

	log.Println(">> READING RESPONSE BODY")
	body, err := ioutil.ReadAll(res.Body)
	check(err)
	defer res.Body.Close()

	d := data.Data{}
	json.Unmarshal(body, &d)

	os.Remove(answerFile)
	file, err := os.Create(answerFile)
	check(err)
	defer file.Close()

	log.Println(">> DECYPHERING MESSAGE")
	d.Decripted = julioDecypher(d.Encrypted, d.NumberOfCases)

	log.Println(">> HASHING RESULT")
	h := sha1.New()
	io.WriteString(h, d.Decripted)
	d.EncryptedResume = fmt.Sprintf("%x", h.Sum(nil))

	log.Println(">> ENCODING DATA TO JSON")
	j, err := d.JSON()
	check(err)

	log.Println(">> UPDATING FILE")
	err = data.UpdateFile(file, j)
	check(err)

	log.Printf(">> SENDING FILE TO %s\n", postURL)
	content, err := data.SendData(postURL, file)
	check(err)
	log.Printf(">> RECEIVED: %s\n", string(content))

	log.Println(">> APP ENDED")
}
