package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTestNil()
		t.Run(table.name, func(t *testing.T) {
			Debug(table.args...)
		})
	}
}

func TestDebugf(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			Debugf(table.format, table.args...)
		})
	}
}

func TestDebugOpts(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOpts(*getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCaller(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		SetOptions(&Options{Mode: ModeJson})
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCaller(table.skipCaller, table.args...)
		})
	}
}

func TestDebugH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugH(table.args...)
		})
	}
}

func TestDebugMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugMS(table.args...)
		})
	}
}

func TestDebugME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugME(table.args...)
		})
	}
}

func TestDebugfH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugfH(table.format, table.args...)
		})
	}
}

func TestDebugfMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugfMS(table.format, table.args...)
		})
	}
}

func TestDebugfME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugfME(table.format, table.args...)
		})
	}
}

func TestDebugOptsH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOptsH(*getOptionsTest(), table.args...)
		})
	}
}

func TestDebugOptsMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOptsMS(*getOptionsTest(), table.args...)
		})
	}
}

func TestDebugOptsME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOptsME(*getOptionsTest(), table.args...)
		})
	}
}

func TestDebugOptsf(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOptsf(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugOptsfH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOptsfH(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugOptsfMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOptsfMS(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugOptsfME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOptsfME(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCallerf(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerf(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestDebugSkipCallerH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerH(table.skipCaller, table.args...)
		})
	}
}

func TestDebugSkipCallerMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerMS(table.skipCaller, table.args...)
		})
	}
}

func TestDebugSkipCallerME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerME(table.skipCaller, table.args...)
		})
	}
}

func TestDebugSkipCallerfH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerfH(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestDebugSkipCallerfMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerfMS(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestDebugSkipCallerfME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerfME(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestDebugSkipCallerOpts(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOpts(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCallerOptsf(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOptsf(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCallerOptsH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOptsH(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCallerOptsMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOptsMS(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCallerOptsME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOptsME(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCallerOptsfH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOptsfH(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCallerOptsfMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOptsfMS(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCallerOptsfME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOptsfME(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}
