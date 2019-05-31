package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/urfave/cli"
)

func Init() []cli.Command {
	return []cli.Command{
		Get(),
	}
}

func Get() cli.Command {
	return cli.Command{
		Name:        "get",
		ShortName:   "g",
		Description: "Peforms a GET request",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "url", Usage: "", Value: ""},
			cli.StringFlag{Name: "L", Usage: "--L", Value: "false"},
		},
		Action: func(c *cli.Context) error {

			// First, let's validate the URL. If it's invalid, throw an error.
			rawURL := c.String("url")
			targetURL, err := url.Parse(rawURL)
			if err != nil {
				fmt.Println(err)
				return err
			}
			// if no URL scheme detected, we will default to http
			if targetURL.Scheme == "" {
				targetURL.Scheme = "http"
			}
			// now that we have a valid URL, let's peform our HTTP Request
			client := http.DefaultClient

			// override the default behaivor of our client so that it DOESN'T follow redirects (just like curl)
			if c.String("L") != "true" {
				client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				}
			}

			res, err := client.Get(targetURL.String())
			if err != nil {
				fmt.Println(err)
				return err
			}

			defer res.Body.Close()

			// read the response body
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(body))
			return nil
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			fmt.Println("error: " + err.Error())
			return err
		},
	}
}
