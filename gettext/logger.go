package gettext

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type msgidLogger struct {
	msgids *sync.Map
}

var logger *msgidLogger = &msgidLogger{msgids: &sync.Map{}}

func (s *msgidLogger) Log(msgid string) {
	if _, ok := s.msgids.Load(msgid); !ok {
		s.msgids.Store(msgid, true)
	}
}

//SaveLog save the keywords not translate yet to messages.log.
//保存还没有翻译的关键字到 messages.log。
func SaveLog() {
	fp, err := os.OpenFile("messages.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	fp.WriteString(fmt.Sprintf("\n# %s\n", time.Now().Format("2006-01-02 15:04:05")))
	logger.msgids.Range(func(k1 interface{}, v1 interface{}) bool {
		k, ok := k1.(string)
		if !ok {
			return false
		}
		if strings.Contains(k, "\n") == false {
			fp.WriteString(fmt.Sprintf("\nmsgid \"%s\"\nmsgstr \"\"\n", k))
		} else {
			fp.WriteString("\nmsgid \"\"\n")
			strArray := strings.Split(k, "\n")
			strLen := len(strArray)
			for i := 0; i < strLen-1; i++ {
				fp.WriteString(fmt.Sprintf("\"%s\\n\"\n", strArray[i]))
			}
			fp.WriteString(fmt.Sprintf("\"%s\"\n", strArray[strLen-1]))
			fp.WriteString("msgstr \"\"\n")
		}
		return true
	})
}
