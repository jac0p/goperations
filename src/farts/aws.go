package main


import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
  svc := ec2.New(session.New())

  params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running"), aws.String("pending")},
			},
		},
	}

  result, err := svc.DescribeInstances(params)
  if err != nil {
   fmt.Println(err.Error())
  }

  fmt.Println(result)


}



