package database

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/spidernest-go/migrate"
)

func Synchronize() error {
	versions := *new([]uint8)
	times := *new([]time.Time)
	buffers := *new([]io.Reader)
	names := *new([]string)
	box := packr.NewBox("./migrations")

	box.Walk(func(n string, f packr.File) error {
		vals := strings.Split(n, "_")

		// Assign Times
		epoch, err := strconv.ParseInt(vals[0], 10, 64)
		if err != nil {
			return err
		}
		t := time.Unix(epoch, 0)

		times = append(times, t)

		// Assign Versioning
		ver, err := strconv.Atoi(vals[1])
		if err != nil {
			return err
		}
		versions = append(versions, uint8(ver))

		// Assign Readers
		data, err := box.Find(n)
		if err != nil {
			return err
		}
		buf := bytes.NewBuffer(data)
		buffers = append(buffers, buf)

		// Assign Names
		names = append(names, n)

		return nil
	})

	if err := migrate.UpTo(versions, names, times, buffers, DB); err != nil {
		return err
	}

	return nil
}
