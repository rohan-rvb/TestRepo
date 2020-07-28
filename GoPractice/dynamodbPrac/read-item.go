package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rohan-rvb/TestRepo/dynamodbPrac/types"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region: aws.String("us-east-1"),
	}))

	db := dynamodb.New(sess)

	artist := "Sonu Nigam"
	songtitle := "Hello Again"

	params := &dynamodb.GetItemInput{
		TableName:aws.String("Music1"),
		Key:map[string]*dynamodb.AttributeValue{
			"Artist" :{
				S:aws.String(artist),
			},
			"SongTitle" :{
				S:aws.String(songtitle),
			},
		},
	}

	resp, err := db.GetItem(params)

	if err!= nil {

	}

	var music = types.Music{}
	erre := dynamodbattribute.UnmarshalMap(resp.Item, &music)

	if erre!= nil {

	}

	fmt.Println(music)

	//fmt.Println(resp)

}
