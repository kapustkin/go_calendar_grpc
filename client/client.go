package main

import (
	"context"
	"log"
	"time"

	"github.com/kapustkin/go_calendar_grpc/calendarpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:5900", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	calendar := calendarpb.NewCalendarEventsClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	event, err := calendar.Get(ctx, &calendarpb.GetRequest{User: "test", Uuid: "1111"})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
			} else {
				log.Printf("undexpected error %s\n", statusErr.Message())
			}
		} else {
			log.Printf("Error while calling RPC CheckHomework: %v", err)
		}
	} else {
		log.Printf("%v\n", event.Event)
	}

	events, err := calendar.GetAll(ctx, &calendarpb.GetAllRequest{User: "test"})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
			} else {
				log.Printf("undexpected error %s\n", statusErr.Message())
			}
		} else {
			log.Printf("Error while calling RPC CheckHomework: %v", err)
		}
	} else {
		log.Printf("%v\n", events.Events)
	}

	editEvent, err := calendar.Edit(ctx, &calendarpb.EditRequest{User: "test", Event: &calendarpb.Event{Uuid: "123", Message: "111"}})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
			} else {
				log.Printf("undexpected error %s\n", statusErr.Message())
			}
		} else {
			log.Printf("Error while calling RPC CheckHomework: %v", err)
		}
	} else {
		log.Printf("Edit sucess: %v\n", editEvent.Sucess)
	}
}
