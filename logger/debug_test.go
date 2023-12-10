package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			Debug(table.args...)
		})
	}
}

func TestDebugF(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugF(table.format, table.args...)
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

func TestDebugFH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugFH(table.format, table.args...)
		})
	}
}

func TestDebugFMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugFMS(table.format, table.args...)
		})
	}
}

func TestDebugFME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugFME(table.format, table.args...)
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

func TestDebugOptsF(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOptsF(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugOptsFH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOptsFH(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugOptsFMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOptsFMS(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugOptsFME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugOptsFME(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCallerF(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerF(table.format, table.skipCaller, table.args...)
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

func TestDebugSkipCallerFH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerFH(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestDebugSkipCallerFMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerFMS(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestDebugSkipCallerFME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerFME(table.format, table.skipCaller, table.args...)
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

func TestDebugSkipCallerOptsF(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOptsF(table.format, table.skipCaller, *getOptionsTest(), table.args...)
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

func TestDebugSkipCallerOptsFH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOptsFH(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCallerOptsFMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOptsFMS(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestDebugSkipCallerOptsFME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			DebugSkipCallerOptsFME(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}
