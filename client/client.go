package client

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/techrail/bark/models"
)

// The client defines 7 levels of errors:
// 1. Panic - The message you emit right before the program crashes
// 2. Alert - The message needs to be sent as an alert to someone who must resolve it ASAP
// 3. Error - The message indicating that there was an error and should be checked whenever possible
// 4. Warning - The message indicating that something wrong could have happened but was handled. Can be overlooked in some cases.
// 5. Notice - Something worth noticing, though it is fine to be ignored.
// 6. Info - Just a log of some data - does not indicate any error
// 7. Debug - used for debugging. It can represent any level of information but is only supposed to indicate a message emitted during a debug session

type Config struct {
	BaseUrl     string
	ErrorLevel  string
	ServiceName string
	SessionName string
}

func (c *Config) Panic(message string) {
	c.log(message, Panic)
}
func (c *Config) Alert(message string) {
	c.log(message, Alert)
}
func (c *Config) Error(message string) {
	c.log(message, Error)
}
func (c *Config) Warn(message string) {
	c.log(message, Warning)
}
func (c *Config) Notice(message string) {
	c.log(message, Notice)
}
func (c *Config) Info(message string) {
	c.log(message, Info)
}
func (c *Config) Debug(message string) {
	c.log(message, Debug)
}
func (c *Config) Println(message string) {
	c.log(message+"\n", Info)
}

func (c *Config) Panicf(message string, format ...any) {
	c.log(fmt.Sprintf(message, format...), Panic)
}
func (c *Config) Alertf(message string, format ...any) {
	c.log(fmt.Sprintf(message, format...), Alert)
}
func (c *Config) Errorf(message string, format ...any) {
	c.log(fmt.Sprintf(message, format...), Error)
}
func (c *Config) Warnf(message string, format ...any) {
	c.log(fmt.Sprintf(message, format...), Warning)
}
func (c *Config) Noticef(message string, format ...any) {
	c.log(fmt.Sprintf(message, format...), Notice)
}
func (c *Config) Infof(message string, format ...any) {
	c.log(fmt.Sprintf(message, format...), Info)
}
func (c *Config) Debugf(message string, format ...any) {
	c.log(fmt.Sprintf(message, format...), Debug)
}

func (c *Config) log(message, logLevel string) {
	// Todo: We have to parse the error message
	log := models.BarkLog{
		Message:     message,
		LogLevel:    logLevel,
		SessionName: c.SessionName,
		ServiceName: c.ServiceName,
		Code:        getLMIDForLogLevel(&logLevel),
	}

	go func() {
		_, err := PostLog(c.BaseUrl+"/insertSingle", log)
		if err.Severity == 1 {
			fmt.Println(err.Error())
			return
		}
	}()

	fmt.Printf("%s - %s\n", log.Code, message)
	// Todo: Add uber zap to avoid printing with PrintF (We don't want to handle log printing)
}

func getLMIDForLogLevel(logLevel *string) string {
	currTime := time.Now().Unix() - 1600000000
	convertToBase36 := strconv.FormatInt(currTime, 36)
	firstLetterOfLogLevel := string((*logLevel)[0])
	return firstLetterOfLogLevel + "#" + strings.ToUpper(convertToBase36)
}

func NewClient(url, errLevel, svcName, sessName string) *Config {
	return &Config{
		BaseUrl:     url,
		ErrorLevel:  errLevel,
		ServiceName: svcName,
		SessionName: sessName,
	}
}
