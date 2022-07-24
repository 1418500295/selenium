package cases

import (
	"flag"
	"fmt"
	"github.com/tebeka/selenium"
	"os"
	base2 "selenium/base"
	"testing"
)

var (
	driver selenium.WebDriver
	page   base2.Base
	path   = "C:\\Users\\14185\\go\\src\\selenium\\chromedriver.exe"
	port   = 8889
)

func TestMain(m *testing.M) {
	var err error
	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr), // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	_, err = selenium.NewChromeDriverService(path, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	//defer service.Stop()
	// set browser as chrome
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})

	// remote to selenium server
	if driver, err = selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port)); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}

	page = base2.Base{Driver: driver}
	err = driver.Get("https://github.com/login")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	flag.Parse()
	m.Run()

	//err1 := driver.Quit()
	//if err1 != nil {
	//	return
	//}

	//os.Exit(r)
}
