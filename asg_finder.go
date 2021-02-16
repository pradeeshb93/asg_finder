package main
import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/service/autoscaling"
)

func main(){
sess := session.Must(session.NewSessionWithOptions(session.Options{
    Profile: "xxxxx",
    Config: aws.Config{
        Region: aws.String("xxxxx"),
    },}))
svc := autoscaling.New(sess)
input := &autoscaling.DescribeAutoScalingGroupsInput{}
result, err := svc.DescribeAutoScalingGroups(input)
if err != nil {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        case autoscaling.ErrCodeInvalidNextToken:
            fmt.Println(autoscaling.ErrCodeInvalidNextToken, aerr.Error())
        case autoscaling.ErrCodeResourceContentionFault:
            fmt.Println(autoscaling.ErrCodeResourceContentionFault, aerr.Error())
        default:
            fmt.Println(aerr.Error())
        }
    } else {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
    }
    return
}
for _,out := range result.AutoScalingGroups {
fmt.Println(*out.AutoScalingGroupName)
}
}
