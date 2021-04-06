package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func init() {
	//log.SetFormatter(&log.TextFormatter{FullTimestamp: true, ForceQuote: true})
	log.SetFormatter(&Formatter{})
	log.SetOutput(os.Stdout)
}

type Formatter struct{}

func (s *Formatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format(time.RFC3339)

	//fmt.Println(entry.Data)
	//s
	msg := fmt.Sprintf("%s [%s] %s\n", timestamp, strings.ToUpper(entry.Level.String())[:4], entry.Message)
	return []byte(msg), nil
}
