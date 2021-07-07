package coutdown

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("测试Countdown打印输出内容", func(t *testing.T) {
		buffer := bytes.Buffer{}
		sleeper := &SpySleeper{0}
		Countdown(&buffer, sleeper)
		got := buffer.String()
		want := `3
2
1
GO!`
		if got != want {
			t.Errorf("get %s, want %s", got, want)
		}
		if sleeper.Calls != 4 {
			t.Errorf("睡眠次数不够：%d", sleeper.Calls)
		}
	})

	t.Run("测试Countdown睡眠和打印顺序", func(t *testing.T) {
		spy := SpyCountdownWriter{}
		Countdown(&spy, &spy)
		got := spy.record
		want := `sleep
write
sleep
write
sleep
write
sleep
write
`
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s,want %s", got, want)
		}
	})

}
