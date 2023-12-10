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

func TestErrorF(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorF(table.format, table.args...)
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

func TestErrorFH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorFH(table.format, table.args...)
		})
	}
}

func TestErrorFMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorFMS(table.format, table.args...)
		})
	}
}

func TestErrorFME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorFME(table.format, table.args...)
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

func TestErrorOptsF(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOptsF(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorOptsFH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOptsFH(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorOptsFMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOptsFMS(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorOptsFME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorOptsFME(table.format, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCallerF(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerF(table.format, table.skipCaller, table.args...)
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

func TestErrorSkipCallerFH(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerFH(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestErrorSkipCallerFMS(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerFMS(table.format, table.skipCaller, table.args...)
		})
	}
}

func TestErrorSkipCallerFME(t *testing.T) {
	for _, table := range initTables() {
		initOptionsTest()
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerFME(table.format, table.skipCaller, table.args...)
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

func TestErrorSkipCallerOptsF(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOptsF(table.format, table.skipCaller, *getOptionsTest(), table.args...)
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

func TestErrorSkipCallerOptsFH(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOptsFH(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCallerOptsFMS(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOptsFMS(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}

func TestErrorSkipCallerOptsFME(t *testing.T) {
	for _, table := range initTables() {
		t.Run(table.name, func(t *testing.T) {
			ErrorSkipCallerOptsFME(table.format, table.skipCaller, *getOptionsTest(), table.args...)
		})
	}
}
