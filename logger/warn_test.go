package logger

import (
	"testing"
)

func TestWarning(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			Warn(table.args...)
		})
	}
}

func TestWarningf(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			Warnf(table.format, table.args...)
		})
	}
}

func TestWarningOpts(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnOpts(*getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCaller(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCaller(table.skipCaller, table.args...)
		})
	}
}

func TestWarningH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnH(table.args...)
		})
	}
}

func TestWarningMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnMS(table.args...)
		})
	}
}

func TestWarningME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnME(table.args...)
		})
	}
}

func TestWarningfH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnfH(table.format, table.args...)
		})
	}
}

func TestWarningfMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnfMS(table.format, table.args...)
		})
	}
}

func TestWarningfME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnfME(table.format, table.args...)
		})
	}
}

func TestWarningOptsH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnOptsH(*getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnOptsMS(*getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnOptsME(*getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsf(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnOptsf(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsfH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnOptsfH(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsfMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnOptsfMS(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsfME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnOptsfME(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerf(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerf(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerH(table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerMS(table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerME(table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerfH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerfH(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerfMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerfMS(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerfME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerfME(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerOpts(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerOpts(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsf(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerOptsf(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerOptsH(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerOptsMS(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerOptsME(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsfH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerOptsfH(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsfMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerOptsfMS(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsfME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarnSkipCallerOptsfME(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}
