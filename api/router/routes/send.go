package routes

import (
	"net/http"
	"strings"
)

func Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := r.URL.Query()
		code, _ := data["code"]
		text, _ := data["text"]
		key := code[0]
		digit := text[0]

		fakeKey := "131"

		if key == fakeKey && len(digit) == 0 {
			w.Write([]byte("CON 1. Buy MTN Packet 2. Buy Tigo Packet"))
			return
		}
		if key == fakeKey && len(digit) != 0 {
			s := strings.Split(digit, "*")
			if len(s) == 1 {
				switch s[0] {
				case "1":
					w.Write([]byte("CON 1. Mobile Money 2. Mtn Balance"))
					return
				case "2":
					w.Write([]byte("CON 1. Tigo Cash 2. Tigo Balance"))
					return
				}
				return
			}

			if len(s) == 2 {
				switch s[1] {
				case "1":
					w.Write([]byte("CON 1. Enter Pin 2. Go Back"))
					return
				case "2":
					w.Write([]byte("CON 1. Enter Pin 2. Go Back"))
					return
				}
				return
			}
			if len(s) == 3 {
				switch s[2] {
				case "1":
					w.Write([]byte("CON Enter Your Pin"))
					return
				case "2":
					w.Write([]byte("END Not Yet Integrated"))
					return
				}
				return
			}
			if len(s) == 4 && len(s[3]) == 4 {

				w.Write([]byte("END Wait For results for your request"))
				return
			} else {
				w.Write([]byte("END Invalid Pin"))
				return
			}

		} else {
			w.Write([]byte("END Invalid Input"))
			return
		}

	}
}
