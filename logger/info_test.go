package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			Info(table.args...)
		})
	}
}

func TestInfof(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			Infof(table.format, table.args...)
		})
	}
}

func TestInfoOpts(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOpts(*getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCaller(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCaller(table.skipCaller, table.args...)
		})
	}
}

func TestInfoH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoH(table.args...)
		})
	}
}

func TestInfoMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoMS(table.args...)
		})
	}
}

func TestInfoME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoME(table.args...)
		})
	}
}

func TestInfofH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfofH(table.format, table.args...)
		})
	}
}

func TestInfofMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfofMS(table.format, table.args...)
		})
	}
}

func TestInfofME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfofME(table.format, table.args...)
		})
	}
}

func TestInfoOptsH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOptsH(*getOptionsTest(), table.args...)
		})
	}
}

func TestInfoOptsMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOptsMS(*getOptionsTest(), table.args...)
		})
	}
}

func TestInfoOptsME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOptsME(*getOptionsTest(), table.args...)
		})
	}
}

func TestInfoOptsf(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOptsf(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoOptsfH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOptsfH(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoOptsfMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOptsfMS(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoOptsfME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOptsfME(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCallerf(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerf(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestInfoSkipCallerH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerH(table.skipCaller, table.args...)
		})
	}
}

func TestInfoSkipCallerMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerMS(table.skipCaller, table.args...)
		})
	}
}

func TestInfoSkipCallerME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerME(table.skipCaller, table.args...)
		})
	}
}

func TestInfoSkipCallerfH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerfH(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestInfoSkipCallerfMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerfMS(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestInfoSkipCallerfME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerfME(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestInfoSkipCallerOpts(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOpts(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCallerOptsf(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOptsf(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCallerOptsH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOptsH(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCallerOptsMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOptsMS(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCallerOptsME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOptsME(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCallerOptsfH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOptsfH(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCallerOptsfMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOptsfMS(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCallerOptsfME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOptsfME(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}
