package main

func main() {
	app := App{}
	app.Initialize(
		"localhost",
		"godb",
		"movies",
	)

	app.Run("27017")
}
