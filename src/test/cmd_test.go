package cmd_test

import (
	be "beanstalkd"
	"fmt"
	"testing"
)

var c *be.Conn
var job_id int64

func init() {
	_ = fmt.Println
	c = be.Dial("tcp", "127.0.0.1:11300")
}

func Test_use(t *testing.T) {
	tube := "default"
	res, err := c.Do("use", be.Argv{Tube: tube})
	if err != nil {
		t.Error(err)
		return
	}
	if tube != res.(string) {
		t.Error("tube not match")
		return
	}
}

func Test_put(t *testing.T) {
	res, err := c.Do("put", be.Argv{Pri: 1024, Delay: 1, Ttr: 1, Message: "aabbcc"})
	if err != nil {
		t.Error(err)
		return
	}
	if res.(int64) <= 0 {
		t.Error("put failed")
		return
	}
}

func Test_reserve(t *testing.T) {
	res, err := c.Do("reserve")
	if err != nil {
		t.Error(err)
		return
	}
	items := res.(map[string]interface{})
	if _, ok := items["id"]; !ok {
		t.Error(ok)
		return
	}
	if items["id"].(int64) <= 0 {
		t.Error("id error")
		return
	}
	if _, ok := items["data"]; !ok {
		t.Error(ok)
		return
	}

	job_id = items["id"].(int64)
}

func Test_release(t *testing.T) {
	res, err := c.Do("release", be.Argv{Id: job_id, Pri: 10, Delay: 10})
	if err != nil {
		t.Error(err)
		return
	}
	if res != true {
		t.Error("error")
		return
	}
}

func Test_bury(t *testing.T) {
	res, err := c.Do("bury", be.Argv{Id: job_id, Pri: 10})
	if err != nil {
		t.Error(err)
		return
	}
	if res != true {
		t.Error("error")
		return
	}

}

func Test_touch(t *testing.T) {
	var id int64
	id = 12
	res, err := c.Do("touch", be.Argv{Id: id})
	if err != nil {
		t.Error(err)
		return
	}
	if res != true {
		t.Error("not true")
		return
	}
}

func Test_watch(t *testing.T) {
	res, err := c.Do("watch", be.Argv{Tube: "default_test"})
	if err != nil {
		t.Error(err)
		return
	}
	if res.(int64) <= 0 {
		t.Error("watch error")
		return
	}
}

func Test_ignore(t *testing.T) {
	res, err := c.Do("ignore", be.Argv{Tube: "default_test"})
	if err != nil {
		t.Error(err)
		return
	}
	if res.(int64) <= 0 {
		t.Error("ignore error")
		return
	}
}

func Test_peek(t *testing.T) {
	res, err := c.Do("peek", be.Argv{Id: 9})
	if err != nil {
		t.Error(err)
		return
	}
	items := res.(map[string]interface{})
	if _, ok := items["id"]; !ok {
		t.Error(ok)
		return
	}
	if items["id"].(int64) <= 0 {
		t.Error("id error")
		return
	}
	if _, ok := items["data"]; !ok {
		t.Error(ok)
		return
	}
}

func Test_peek_ready(t *testing.T) {
	res, err := c.Do("peek-ready")
	if err != nil {
		t.Error(err)
		return
	}
	items := res.(map[string]interface{})
	if _, ok := items["id"]; !ok {
		t.Error(ok)
		return
	}
	if items["id"].(int64) <= 0 {
		t.Error("id error")
		return
	}
	if _, ok := items["data"]; !ok {
		t.Error(ok)
		return
	}
}

func Test_peek_delayed(t *testing.T) {
	res, err := c.Do("peek-delayed")
	if err != nil {
		t.Error(err)
		return
	}
	items := res.(map[string]interface{})
	if _, ok := items["id"]; !ok {
		t.Error(ok)
		return
	}
	if items["id"].(int64) <= 0 {
		t.Error("id error")
		return
	}
	if _, ok := items["data"]; !ok {
		t.Error(ok)
		return
	}
}

func Test_peek_buried(t *testing.T) {
	res, err := c.Do("peek-buried")
	if err != nil {
		t.Error(err)
		return
	}
	items := res.(map[string]interface{})
	if _, ok := items["id"]; !ok {
		t.Error(ok)
		return
	}
	if items["id"].(int64) <= 0 {
		t.Error("id error")
		return
	}
	if _, ok := items["data"]; !ok {
		t.Error(ok)
		return
	}
}

func Test_kick(t *testing.T) {
	res, err := c.Do("kick", be.Argv{Bound: 100})
	if err != nil {
		t.Error(err)
		return
	}
	if res.(int64) < 0 {
		t.Error("error")
		return
	}
}

func Test_stats_job(t *testing.T) {
	res, err := c.Do("stats-job", be.Argv{Id: job_id})
	if err != nil {
		t.Error(err)
		return
	}
	items := res.(map[string]interface{})
	if _, ok := items["id"]; !ok {
		t.Error("id not exits")
		return
	}
}

func Test_stats_tube(t *testing.T) {
	res, err := c.Do("stats-tube", be.Argv{Tube: "default"})
	if err != nil {
		t.Error(err)
		return
	}
	items := res.(map[string]interface{})
	if _, ok := items["name"]; !ok {
		t.Error("name not exits")
		return
	}
}

func Test_stats(t *testing.T) {
	res, err := c.Do("stats")
	if err != nil {
		t.Error(err)
		return
	}
	items := res.(map[string]interface{})
	if _, ok := items["total-jobs"]; !ok {
		t.Error("total-jobs not exits")
		return
	}
}

func Test_list_tubes(t *testing.T) {
	res, err := c.Do("list-tubes")
	if err != nil {
		t.Error(err)
		return
	}
	if len(res.([]string)) <= 0 {
		t.Error("length  <= 0")
		return
	}
}

func Test_list_tube_used(t *testing.T) {
	res, err := c.Do("list-tube-used")
	if err != nil {
		t.Error(err)
		return
	}
	if res.(string) != "default" {
		t.Error("length  <= 0")
		return
	}
}

func Test_list_tubes_watched(t *testing.T) {
	res, err := c.Do("list-tubes-watched")
	if err != nil {
		t.Error(err)
		return
	}
	if len(res.([]string)) <= 0 {
		t.Error("length  <= 0")
		return
	}
}

func Test_pause_tube(t *testing.T) {
	res, err := c.Do("pause-tube", be.Argv{Tube: "default", Delay: 10})
	if err != nil {
		t.Error(err)
		return
	}
	if res != true {
		t.Error("not true")
		return
	}
}

func Test_delete(t *testing.T) {
	res, err := c.Do("delete", be.Argv{Id: job_id})
	if err != nil {
		t.Error(err)
		return
	}
	if res != true {
		t.Error("delete job error")
		return
	}
}
