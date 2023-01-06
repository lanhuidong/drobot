package drobot

import (
	"fmt"
	"testing"
)

const webHook = "your dingding webhook"
const secret = "your dingding secret"

func TestSendMarkdown(t *testing.T) {
	robot := DingTalkRobot{
		Webhook: webHook,
		Secret:  secret,
	}
	if err := robot.SendMarkdown("测试", "D888888, <font color=#FF0000>1235</font>"); err != nil {
		t.Logf("%v", err)
	}
}

// 发送消息到钉钉群，注意：钉钉不支持完整的Markdown语法
func ExampleDingTalkRobot_SendMarkdown() {
	robot := DingTalkRobot{
		Webhook: webHook,
		Secret:  secret,
	}
	if err := robot.SendMarkdown("测试", "D888888, <font color=#FF0000>1235</font>"); err != nil {
		fmt.Printf("%v", err)
	}
}
