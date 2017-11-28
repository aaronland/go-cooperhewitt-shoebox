package util

import (
       "io/ioutil"
       "log"
       "net/http"
       "strconv"
       "strings"
)

func GetStore(remote string, local string) error {

	log.Println("getstore", remote, local)

	rsp, err := http.Get(remote)

	if err != nil {
		return err
	}

	contents, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		log.Printf("failed to fetch %s because %v", remote, err)
		return err
	}

	rsp.Body.Close()
	err = ioutil.WriteFile(local, contents, 0644)

	if err != nil {
		return err
	}

	return nil
}

func Id2Path(id int) string {

	parts := []string{}
	input := strconv.Itoa(id)

	for len(input) > 3 {

		chunk := input[0:3]
		input = input[3:]
		parts = append(parts, chunk)
	}

	if len(input) > 0 {
		parts = append(parts, input)
	}

	path := strings.Join(parts, "/")
	return path
}
