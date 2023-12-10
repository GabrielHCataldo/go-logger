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

func TestInfoF(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoF(table.format, table.args...)
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

func TestInfoFH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoFH(table.format, table.args...)
		})
	}
}

func TestInfoFMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoFMS(table.format, table.args...)
		})
	}
}

func TestInfoFME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoFME(table.format, table.args...)
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

func TestInfoOptsF(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOptsF(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoOptsFH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOptsFH(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoOptsFMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOptsFMS(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoOptsFME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoOptsFME(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCallerF(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerF(table.format, table.skipCaller, table.args...)
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

func TestInfoSkipCallerFH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerFH(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestInfoSkipCallerFMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerFMS(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestInfoSkipCallerFME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerFME(table.format, table.skipCaller, table.args...)
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

func TestInfoSkipCallerOptsF(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOptsF(table.format, table.skipCaller, *getOptionsTest(), table.args...)
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

func TestInfoSkipCallerOptsFH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOptsFH(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCallerOptsFMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOptsFMS(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestInfoSkipCallerOptsFME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			InfoSkipCallerOptsFME(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}
