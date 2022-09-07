package middleware

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Logger - логгер запросов.
type Logger struct {
	W io.Writer
}

// Log основной метод. По умолчанию выводит в STDOUT
func (l *Logger) Log(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer handler(w, r)
		body, err := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		data := strings.Builder{}
		data.WriteString(fmt.Sprintf("Request time: %s\nMethod: %s\n", time.Now(), "GET"))
		data.WriteString("\nHeader: \n")
		for k, v := range r.Header {
			data.WriteString(fmt.Sprintf("%s : ", k))
			for _, s := range v {
				data.WriteString(fmt.Sprintf("%s", s))
			}
			data.WriteString("\n")
		}
		data.WriteString("\nQuery params : \n")
		err = r.ParseForm()
		if err == nil {
			for k, v := range r.Form {
				data.WriteString(fmt.Sprintf("%s : ", k))
				for _, s := range v {
					data.WriteString(fmt.Sprintf("%s", s))
				}
				data.WriteString("\n")
			}
		} else {
			data.WriteString(err.Error() + "\n")
		}
		data.WriteString(fmt.Sprintf("Body:\n%s\n", string(body)))
		l.W.Write([]byte(data.String()))
	}

}
