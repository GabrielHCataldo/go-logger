package logger

import (
	"testing"
)

func TestWarning(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			Warning(table.args...)
		})
	}
}

func TestWarningf(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			Warningf(table.format, table.args...)
		})
	}
}

func TestWarningOpts(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOpts(*getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCaller(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCaller(table.skipCaller, table.args...)
		})
	}
}

func TestWarningH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningH(table.args...)
		})
	}
}

func TestWarningMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningMS(table.args...)
		})
	}
}

func TestWarningME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningME(table.args...)
		})
	}
}

func TestWarningfH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningfH(table.format, table.args...)
		})
	}
}

func TestWarningfMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningfMS(table.format, table.args...)
		})
	}
}

func TestWarningfME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningfME(table.format, table.args...)
		})
	}
}

func TestWarningOptsH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOptsH(*getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOptsMS(*getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOptsME(*getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsf(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOptsf(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsfH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOptsfH(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsfMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOptsfMS(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsfME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOptsfME(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerf(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerf(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerH(table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerMS(table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerME(table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerfH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerfH(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerfMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerfMS(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerfME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerfME(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerOpts(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOpts(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsf(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOptsf(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOptsH(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOptsMS(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOptsME(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsfH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOptsfH(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsfMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOptsfMS(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsfME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOptsfME(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}
