package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
    "github.com/oschwald/geoip2-golang"
	"net"
	"log"
)

type input struct {
    IP string `json:"ip"`
}

type Response struct {
    Message string `json:"message"`
    Ok bool `json:"ok"`
}

func Handler(event input) (Response, error){
    db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ip := net.ParseIP(event.IP)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

    return Response {
        Message: fmt.Sprintf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"]),
        Ok: true,
    }, nil
}
func main() {
    lambda.Start(Handler)
}