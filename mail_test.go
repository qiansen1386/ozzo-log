// Copyright 2015 Qiang Xue. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	"github.com/go-ozzo/ozzo-log"
	"testing"
)

func TestNewMailTarget(t *testing.T) {
	target := log.NewMailTarget()
	if target.MaxLevel != log.LevelDebug {
		t.Errorf("NewMailTarget.MaxLevel = %v, expected %v", target.MaxLevel, log.LevelDebug)
	}
}
