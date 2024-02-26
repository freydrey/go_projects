package main

import(
	"context"
	"fmt"
	"log"
	"os"
	"github.com/shomali11/slacker"
	"strconv"
)

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-30583237CX")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06JXLG5573961e8524b")

    bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my job is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		//Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil{
				fmt.Println("error")
			}

			age := 2021 - yob
			r := fmt.Sprintf("my yob is %d", age)
			response.Reply(r)

		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()  //called towards the end of the function

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}


}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}

}