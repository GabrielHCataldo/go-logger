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

func TestWarningF(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningF(table.format, table.args...)
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

func TestWarningFH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningFH(table.format, table.args...)
		})
	}
}

func TestWarningFMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningFMS(table.format, table.args...)
		})
	}
}

func TestWarningFME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningFME(table.format, table.args...)
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

func TestWarningOptsF(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOptsF(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsFH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOptsFH(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsFMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOptsFMS(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningOptsFME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningOptsFME(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerF(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerF(table.format, table.skipCaller, table.args...)
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

func TestWarningSkipCallerFH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerFH(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerFMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerFMS(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestWarningSkipCallerFME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerFME(table.format, table.skipCaller, table.args...)
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

func TestWarningSkipCallerOptsF(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOptsF(table.format, table.skipCaller, *getOptionsTest(), table.args...)
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

func TestWarningSkipCallerOptsFH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOptsFH(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsFMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOptsFMS(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestWarningSkipCallerOptsFME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			WarningSkipCallerOptsFME(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}
