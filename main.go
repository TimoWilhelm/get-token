package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jessevdk/go-flags"

	"github.com/timowilhelm/get-token/models"
)

func main() {
	var err error
	var opts struct {
		TenantID               string `short:"t" long:"tenantID" description:"The directory ID" required:"true"`
		ClientID               string `short:"c" long:"clientID" description:"The application ID" required:"true"`
		Scope                  string `short:"s" long:"scope" description:"A space-separated list of scopes" required:"true"`
		AuthorizationServerURL string `short:"a" long:"authorizationServer" description:"The authorization Server URL" default:"https://login.microsoftonline.com"`
	}

	_, err = flags.ParseArgs(&opts, os.Args)

	if err != nil {
		panic(err)
	}

	if len(opts.TenantID) == 0 {
		panic("TenantID is empty")
	}

	if len(opts.ClientID) == 0 {
		panic("ClientID is empty")
	}

	deviceCodeUrl := fmt.Sprintf("%s/%s/oauth2/v2.0/devicecode", opts.AuthorizationServerURL, opts.TenantID)
	tokenUrl := fmt.Sprintf("%s/%s/oauth2/v2.0/token", opts.AuthorizationServerURL, opts.TenantID)

	rsp, err := http.PostForm(deviceCodeUrl, url.Values{
		"client_id": {opts.ClientID},
		"scope":     {opts.Scope},
	})

	if !isSuccessStatusCode(rsp.StatusCode) {
		if rsp.ContentLength != 0 {
			panic(errorFromResponseBody(rsp))
		} else {
			panic(rsp.Status)
		}
	}

	var body []byte
	body, err = ioutil.ReadAll(rsp.Body)

	if err != nil {
		panic(err)
	}

	var dtr models.DeviceTokenResponse
	err = json.Unmarshal(body, &dtr)

	if err != nil {
		panic(err)
	}

	expiresOn := time.Now().UTC().Unix() + dtr.ExpiresIn

	fmt.Println(dtr.Message)

	for t := time.Now().UTC().Unix(); expiresOn-t > 0; t = time.Now().UTC().Unix() {

		time.Sleep(time.Duration(dtr.Interval) * time.Second)

		rsp, err := http.PostForm(tokenUrl, url.Values{
			"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
			"client_id":   {opts.ClientID},
			"device_code": {dtr.DeviceCode},
		})

		body, err = ioutil.ReadAll(rsp.Body)

		if isSuccessStatusCode(rsp.StatusCode) {
			var tsr models.TokenSucessResponse
			err = json.Unmarshal(body, &tsr)

			if err != nil {
				panic(err)
			}

			fmt.Printf("%+v\n", tsr)
			return

		} else {
			var ter models.TokenErrorResponse
			err = json.Unmarshal(body, &ter)

			if err != nil {
				panic(err)
			}

			if ter.Error != "" && ter.Error != models.AuthorizationPending {
				panic(ter.Error)
			}
		}
	}

	panic("Timeout!")
}

func errorFromResponseBody(rsp *http.Response) error {
	var result map[string]interface{}
	json.NewDecoder(rsp.Body).Decode(&result)
	return errors.New(fmt.Sprint(result))
}

func isSuccessStatusCode(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}
