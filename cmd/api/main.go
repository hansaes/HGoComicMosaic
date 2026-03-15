package main

import (
	"HGoComicMosaic/internal/bootstrap"
	"context"
	"log"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {

	app, err := bootstrap.NewApp()
	if err != nil {
		log.Fatalln(err)
	}
	if err := app.Run(context.Background()); err != nil {
		log.Fatalf("application stopped with error : %v", err)
	}

}
