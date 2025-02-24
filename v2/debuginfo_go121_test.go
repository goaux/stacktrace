//go:build go1.21

package stacktrace_test

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/goaux/stacktrace/v2"
)

func TestDebugInfo_slog(t *testing.T) {
	info := stacktrace.DebugInfo{
		StackEntries: []string{"entry#1", "entry#2"},
		Detail:       "debuginfo-detail",
	}
	t.Run("text", func(t *testing.T) {
		buf := new(bytes.Buffer)
		log := slog.New(slog.NewTextHandler(buf, nil))
		log.Error("error", slog.Any("err", info))
		txt := buf.String()
		i := strings.Index(txt, ` err="{`)
		got := txt[i:]
		want := ` err="{StackEntries:[entry#1 entry#2] Detail:debuginfo-detail}"` + "\n"
		if got != want {
			t.Errorf("got=%q want=%q", got, want)
		}
	})
	t.Run("json", func(t *testing.T) {
		buf := new(bytes.Buffer)
		log := slog.New(slog.NewJSONHandler(buf, nil))
		log.Error("error", slog.Any("err", info))
		txt := buf.String()
		i := strings.Index(txt, `"err":{`)
		got := txt[i:]
		want := `"err":{"stack_entries":["entry#1","entry#2"],"detail":"debuginfo-detail"}}` + "\n"
		if got != want {
			t.Errorf("got=%q want=%q", got, want)
		}
	})
}
