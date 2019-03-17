package templates

import (
	"bufio"
	"gitops/backend/utils"
	"os"
	"strings"
)

// GetSecrets return the list of decrypted secrets
func GetSecrets(basePath string) (secrets map[string]interface{}, err error) {
	repositoryPath := utils.ComposeStrings(basePath , "/secrets.yaml")

	secrets = map[string]interface{}{}
	file, err := os.Open(repositoryPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ": ")
		secrets[values[0]] = values[1]
	}

	return
}
