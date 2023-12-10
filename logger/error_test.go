package logger

import (
	"testing"
)

func TestError(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			Error(table.args...)
		})
	}
}

func TestErrorf(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			Errorf(table.format, table.args...)
		})
	}
}

func TestErrorOpts(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOpts(*getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCaller(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCaller(table.skipCaller, table.args...)
		})
	}
}

func TestErrorH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorH(table.args...)
		})
	}
}

func TestErrorMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorMS(table.args...)
		})
	}
}

func TestErrorME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorME(table.args...)
		})
	}
}

func TestErrorfH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorfH(table.format, table.args...)
		})
	}
}

func TestErrorfMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorfMS(table.format, table.args...)
		})
	}
}

func TestErrorfME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorfME(table.format, table.args...)
		})
	}
}

func TestErrorOptsH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOptsH(*getOptionsTest(), table.args...)
		})
	}
}

func TestErrorOptsMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOptsMS(*getOptionsTest(), table.args...)
		})
	}
}

func TestErrorOptsME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOptsME(*getOptionsTest(), table.args...)
		})
	}
}

func TestErrorOptsf(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOptsf(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorOptsfH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOptsfH(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorOptsfMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOptsfMS(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorOptsfME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOptsfME(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCallerf(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerf(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestErrorSkipCallerH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerH(table.skipCaller, table.args...)
		})
	}
}

func TestErrorSkipCallerMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerMS(table.skipCaller, table.args...)
		})
	}
}

func TestErrorSkipCallerME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerME(table.skipCaller, table.args...)
		})
	}
}

func TestErrorSkipCallerfH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerfH(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestErrorSkipCallerfMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerfMS(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestErrorSkipCallerfME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerfME(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestErrorSkipCallerOpts(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOpts(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCallerOptsf(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOptsf(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCallerOptsH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOptsH(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCallerOptsMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOptsMS(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCallerOptsME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOptsME(table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCallerOptsfH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOptsfH(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCallerOptsfMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOptsfMS(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCallerOptsfME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOptsfME(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}
