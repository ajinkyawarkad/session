// go test -run TestSendJSON -race -cpu 16

// Sample test to show how to test the execution of an internal endpoint.
package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"startup-code/day6/tests/example4/handlers"
)

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	handlers.Routes()
}

// TestSendJSON testing the sendjson internal endpoint.
func TestSendJSON(t *testing.T) {
	url := "/sendjson"
	statusCode := 200

	t.Log("Given the need to test the SendJSON endpoint.")
	{
		r := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)

		t.Logf("\tTest 0:\tWhen checking %q for status code %d", url, statusCode)
		{
			if w.Code != 200 {
				t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d].", failed, statusCode, w.Code)
			}
			t.Logf("\t%s\tShould receive a status code of %d for the response.", succeed, statusCode)

			var u struct {
				Name  string
				Email string
			}

			if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
				t.Fatalf("\t%s\tShould be able to decode the response.", failed)
			}
			t.Logf("\t%s\tShould be able to decode the response.", succeed)

			if u.Name == "Bill" {
				t.Logf("\t%s\tShould have \"Bill\" for Name in the response.", succeed)
			} else {
				t.Errorf("\t%s\tShould have \"Bill\" for Name in the response : %q", failed, u.Name)
			}

			if u.Email == "bill@goinggo.com" {
				t.Logf("\t%s\tShould have \"bill@goinggo.com\" for Email in the response.", succeed)
			} else {
				t.Errorf("\t%s\tShould have \"bill@goinggo.com\" for Email in the response : %q", failed, u.Email)
			}
		}
	}
}
