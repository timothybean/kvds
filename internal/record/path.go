package record

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func leadingZeros(v int, size int) string {
	out := strconv.FormatInt(int64(v), 10)
	if len(out) >= size {
		return out
	}

	out2 := ""
	for i := 0; i < (size - len(out)); i++ {
		out2 += "0"
	}

	return out2 + out
}

type path time.Time

func (p *path) decode(format string, v string) bool {
	re := format

	for _, ch := range []uint8{
		'\\',
		'.',
		'+',
		'*',
		'?',
		'^',
		'$',
		'(',
		')',
		'[',
		']',
		'{',
		'}',
		'|',
	} {
		re = strings.ReplaceAll(re, string(ch), "\\"+string(ch))
	}

	re = strings.ReplaceAll(re, "%path", "(.*?)")
	re = strings.ReplaceAll(re, "%Y", "([0-9]{4})")
	re = strings.ReplaceAll(re, "%m", "([0-9]{2})")
	re = strings.ReplaceAll(re, "%d", "([0-9]{2})")
	re = strings.ReplaceAll(re, "%H", "([0-9]{2})")
	re = strings.ReplaceAll(re, "%M", "([0-9]{2})")
	re = strings.ReplaceAll(re, "%S", "([0-9]{2})")
	re = strings.ReplaceAll(re, "%f", "([0-9]{6})")
	re = strings.ReplaceAll(re, "%s", "([0-9]{10})")
	r := regexp.MustCompile(re)

	var groupMapping []string
	cur := format
	for {
		i := strings.Index(cur, "%")
		if i < 0 {
			break
		}

		cur = cur[i:]

		for _, va := range []string{
			"%path",
			"%Y",
			"%m",
			"%d",
			"%H",
			"%M",
			"%S",
			"%f",
			"%s",
		} {
			if strings.HasPrefix(cur, va) {
				groupMapping = append(groupMapping, va)
			}
		}

		cur = cur[1:]
	}

	matches := r.FindStringSubmatch(v)
	if matches == nil {
		return false
	}

	values := make(map[string]string)

	for i, match := range matches[1:] {
		values[groupMapping[i]] = match
	}

	var year int
	var month time.Month = 1
	day := 1
	var hour int
	var minute int
	var second int
	var micros int
	var unixSec int64 = -1

	for k, v := range values {
		switch k {
		case "%Y":
			tmp, _ := strconv.ParseInt(v, 10, 64)
			year = int(tmp)

		case "%m":
			tmp, _ := strconv.ParseInt(v, 10, 64)
			month = time.Month(int(tmp))

		case "%d":
			tmp, _ := strconv.ParseInt(v, 10, 64)
			day = int(tmp)

		case "%H":
			tmp, _ := strconv.ParseInt(v, 10, 64)
			hour = int(tmp)

		case "%M":
			tmp, _ := strconv.ParseInt(v, 10, 64)
			minute = int(tmp)

		case "%S":
			tmp, _ := strconv.ParseInt(v, 10, 64)
			second = int(tmp)

		case "%f":
			tmp, _ := strconv.ParseInt(v, 10, 64)
			micros = int(tmp)

		case "%s":
			unixSec, _ = strconv.ParseInt(v, 10, 64)
		}
	}

	if unixSec > 0 {
		*p = path(time.Unix(unixSec, 0))
	} else {
		*p = path(time.Date(year, month, day, hour, minute, second, micros*1000, time.Local))
	}

	return true
}

func (p path) encode(format string) string {
	format = strings.ReplaceAll(format, "%Y", strconv.FormatInt(int64(time.Time(p).Year()), 10))
	format = strings.ReplaceAll(format, "%m", leadingZeros(int(time.Time(p).Month()), 2))
	format = strings.ReplaceAll(format, "%d", leadingZeros(time.Time(p).Day(), 2))
	format = strings.ReplaceAll(format, "%H", leadingZeros(time.Time(p).Hour(), 2))
	format = strings.ReplaceAll(format, "%M", leadingZeros(time.Time(p).Minute(), 2))
	format = strings.ReplaceAll(format, "%S", leadingZeros(time.Time(p).Second(), 2))
	format = strings.ReplaceAll(format, "%f", leadingZeros(time.Time(p).Nanosecond()/1000, 6))
	format = strings.ReplaceAll(format, "%s", strconv.FormatInt(time.Time(p).Unix(), 10))
	return format
}
