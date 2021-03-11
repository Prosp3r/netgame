package main

// Use this code snippet in your app.
// If you need more information about configurations or implementing the sample code, visit the AWS docs:
// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/setting-up.html

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"time"
)

//DatabaseAuth -Hold the structure for all data returned from aws secrets manager
type DatabaseAuth struct {
	Host                 string `json:"host"`
	UserName             string `json:"username"`
	Password             string `json:"password"`
	Port                 string `json:"port"`
	DbInstanceIdentifier string `json:"dbInstanceIdentifier"`
	DBEngine             string `json:"engine"`
	//os.Getenv("DB_NAME"))
}

//{"username":"pgadmin","password":"akomeno123,","engine":"postgres","host":"ngtest002.cxrap41ptykp.us-east-1.rds.amazonaws.com","port":5432,"dbInstanceIdentifier":"ngtest002"}
/*/DatabaseAuth -Hold the structure for all data returned from aws secrets manager
type DatabaseAuth struct{}*/

func main() {
	cnt := 0
	for {
		cnt++
		fmt.Printf("%v Attempting retrieval \n", cnt)
		dba, err := getSecret()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Access Host: %v\n Access Port: %v\n Access Username: %v\n Access password: %v\n ", dba.Host, dba.Port, dba.UserName, dba.Password)
		time.Sleep(time.Second * time.Duration(3))
	}
}

//getSecret - Fetches secret db access param from aws Secret manager
func getSecret() (DatabaseAuth, error) {
	secretName := "ng/masterpg"
	region := "us-east-1"

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html

	result, err := svc.GetSecretValue(input)
	var dbAuth = DatabaseAuth{}
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		//return dbAuth, error.Error()
	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString, decodedBinarySecret string
	if result.SecretString != nil {
		secretString = *result.SecretString
		fmt.Println(secretString)
		json.Unmarshal([]byte(secretString), &dbAuth)

	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			fmt.Println("Base64 Decode Error:", err)
			//return dbAuth, err
		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:len])
		//fmt.Printf("Secretes: %v\n", decodedBinarySecret)
		json.Unmarshal([]byte(decodedBinarySecret), &dbAuth)
	}

	// Your code goes here.
	return dbAuth, nil
}
