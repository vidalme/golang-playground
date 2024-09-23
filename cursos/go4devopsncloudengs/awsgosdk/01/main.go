package main

import (
	"fmt"
	"os"
)

func main() {

	var (
		instanceID string
		err        error
	)

	if instanceID, err = createEC2("us-east-1"); err != nil {
		fmt.Printf("Error criando instancia de EC2; %s\n", instanceID)
		os.Exit(1)
	}

}

func createEC2(region string) (string, error) {
	return "", nil
}
