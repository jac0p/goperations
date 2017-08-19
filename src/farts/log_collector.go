package main

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/cloudwatchlogs"

)

func main() {

  svc := cloudwatchlogs.New(session.New())

  params := &cloudwatchlogs.FilterLogEventsInput{
    LogGroupName: aws.String("/aws/lambda/NEWDEVVpcProxy"),
  }

  result, err := svc.FilterLogEvents(params)
  if err != nil {
    fmt.Println(err.Error())
  }

  fmt.Println(result)
}



