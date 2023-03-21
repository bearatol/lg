package lg

import (
	"log"
	"testing"
)

func TestLogger_TurnColor(t *testing.T) {
	logger = newLogEnd()
	defer resetVariables()
	t.Run("color on", func(t *testing.T) {
		colorReset := logger.infoLogger.c.Reset == "\033[0;0;0m"
		if !colorReset {
			t.Errorf("color Reset isn't valid: %s", logger.infoLogger.c.Reset)
		}
		colorRed := logger.infoLogger.c.Red == "\033[0;0;31m"
		if !colorRed {
			t.Errorf("color ReRedset isn't valid: %s", logger.infoLogger.c.Red)
		}
		colorYellow := logger.infoLogger.c.Yellow == "\033[0;0;33m"
		if !colorYellow {
			t.Errorf("color Yellow isn't valid: %s", logger.infoLogger.c.Yellow)
		}
		colorPurple := logger.infoLogger.c.Purple == "\033[0;0;35m"
		if !colorPurple {
			t.Errorf("color Purple isn't valid: %s", logger.infoLogger.c.Purple)
		}
		colorPurpleUnderline := logger.infoLogger.c.PurpleUnderline == "\033[0;4;35m"
		if !colorPurpleUnderline {
			t.Errorf("color PurpleUnderline isn't valid: %s", logger.infoLogger.c.PurpleUnderline)
		}
		colorCyan := logger.infoLogger.c.Cyan == "\033[0;0;36m"
		if !colorCyan {
			t.Errorf("color Cyan isn't valid: %s", logger.infoLogger.c.Cyan)
		}
		colorGray := logger.infoLogger.c.Gray == "\033[0;0;37m"
		if !colorGray {
			t.Errorf("color Gray isn't valid: %s", logger.infoLogger.c.Gray)
		}
		colorGrayUnderline := logger.infoLogger.c.GrayUnderline == "\033[0;4;37m"
		if !colorGrayUnderline {
			t.Errorf("color GrayUnderline isn't valid: %s", logger.infoLogger.c.GrayUnderline)
		}
	})

	t.Run("color off", func(t *testing.T) {
		logger.OffColor()
		emptyReset := logger.infoLogger.c.Reset == ""
		if !emptyReset {
			t.Errorf("color Reset isn't empty: %s", logger.infoLogger.c.Reset)
		}
		emptyRed := logger.infoLogger.c.Red == ""
		if !emptyRed {
			t.Errorf("color Red isn't empty: %s", logger.infoLogger.c.Red)
		}
		emptyYellow := logger.infoLogger.c.Yellow == ""
		if !emptyYellow {
			t.Errorf("color Yellow isn't empty: %s", logger.infoLogger.c.Yellow)
		}
		emptyPurple := logger.infoLogger.c.Purple == ""
		if !emptyPurple {
			t.Errorf("color Purple isn't empty: %s", logger.infoLogger.c.Purple)
		}
		emptyPurpleUnderline := logger.infoLogger.c.PurpleUnderline == ""
		if !emptyPurpleUnderline {
			t.Errorf("color PurpleUnderline isn't empty: %s", logger.infoLogger.c.PurpleUnderline)
		}
		emptyCyan := logger.infoLogger.c.Cyan == ""
		if !emptyCyan {
			t.Errorf("color Cyan isn't empty: %s", logger.infoLogger.c.Cyan)
		}
		emptyGray := logger.infoLogger.c.Gray == ""
		if !emptyGray {
			t.Errorf("color Gray isn't empty: %s", logger.infoLogger.c.Gray)
		}
		emptyGrayUnderline := logger.infoLogger.c.GrayUnderline == ""
		if !emptyGrayUnderline {
			t.Errorf("color GrayUnderline isn't empty: %s", logger.infoLogger.c.GrayUnderline)
		}
	})
}

func TestLogger_TurnPrefixOn(t *testing.T) {
	logger := newLogEnd()
	defer resetVariables()
	p := newPrefix()
	c := newColor()

	t.Run("prefix info on", func(t *testing.T) {
		prefix := logger.infoLogger.log.Prefix()
		if prefix != c.Cyan+p.Info+"\t" {
			t.Errorf("prefix info isn't valid: %s", prefix)
		}
	})
	t.Run("prefix debug on", func(t *testing.T) {
		prefix := logger.debugLogger.log.Prefix()
		if prefix != c.GrayUnderline+p.Debug+"\t" {
			t.Errorf("prefix debug isn't valid: %s", prefix)
		}
	})
	t.Run("prefix trace on", func(t *testing.T) {
		prefix := logger.traceLogger.log.Prefix()
		if prefix != c.Gray+p.Trace+"\t" {
			t.Errorf("prefix trace isn't valid: %s", prefix)
		}
	})
	t.Run("prefix warn on", func(t *testing.T) {
		prefix := logger.warnLogger.log.Prefix()
		if prefix != c.Yellow+p.Warn+"\t" {
			t.Errorf("prefix warn isn't valid: %s", prefix)
		}
	})
	t.Run("prefix error on", func(t *testing.T) {
		prefix := logger.errorLogger.log.Prefix()
		if prefix != c.Red+p.Error+"\t" {
			t.Errorf("prefix error isn't valid: %s", prefix)
		}
	})
	t.Run("prefix fatal on", func(t *testing.T) {
		prefix := logger.fatalLogger.log.Prefix()
		if prefix != c.Purple+p.Fatal+"\t" {
			t.Errorf("prefix fatal isn't valid: %s", prefix)
		}
	})
	t.Run("prefix panic on", func(t *testing.T) {
		prefix := logger.panicLogger.log.Prefix()
		if prefix != c.PurpleUnderline+p.Panic+"\t" {
			t.Errorf("prefix panic isn't valid: %s", prefix)
		}
	})
}

func TestLogger_TurnPrefixOff(t *testing.T) {
	logger := newLogEnd()
	defer resetVariables()
	c := newColor()
	logger.OffPrefix()
	t.Run("prefix info off", func(t *testing.T) {
		prefix := logger.infoLogger.log.Prefix()
		if prefix != c.Cyan {
			t.Errorf("prefix info isn't valid: %s", prefix)
		}
	})
	t.Run("prefix debug off", func(t *testing.T) {
		prefix := logger.debugLogger.log.Prefix()
		if prefix != c.GrayUnderline {
			t.Errorf("prefix debug isn't valid: %s", prefix)
		}
	})
	t.Run("prefix trace off", func(t *testing.T) {
		prefix := logger.traceLogger.log.Prefix()
		if prefix != c.Gray {
			t.Errorf("prefix trace isn't valid: %s", prefix)
		}
	})
	t.Run("prefix warn off", func(t *testing.T) {
		prefix := logger.warnLogger.log.Prefix()
		if prefix != c.Yellow {
			t.Errorf("prefix warn isn't valid: %s", prefix)
		}
	})
	t.Run("prefix error off", func(t *testing.T) {
		prefix := logger.errorLogger.log.Prefix()
		if prefix != c.Red {
			t.Errorf("prefix error isn't valid: %s", prefix)
		}
	})
	t.Run("prefix fatal off", func(t *testing.T) {
		prefix := logger.fatalLogger.log.Prefix()
		if prefix != c.Purple {
			t.Errorf("prefix fatal isn't valid: %s", prefix)
		}
	})
	t.Run("prefix panic off", func(t *testing.T) {
		prefix := logger.panicLogger.log.Prefix()
		if prefix != c.PurpleUnderline {
			t.Errorf("prefix panic isn't valid: %s", prefix)
		}
	})
}

func TestLogger_TurnDateOff(t *testing.T) {
	logger := newLogEnd()
	defer resetVariables()
	logger.OffDate()
	logFlagsTest := log.Ltime | log.Lshortfile
	t.Run("date info", func(t *testing.T) {
		flag := logger.infoLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("date info isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("date debug", func(t *testing.T) {
		flag := logger.debugLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("date debug isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("date trace", func(t *testing.T) {
		flag := logger.traceLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("date trace isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("date warn", func(t *testing.T) {
		flag := logger.warnLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("date warn isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("date error", func(t *testing.T) {
		flag := logger.errorLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("date error isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("date fatal", func(t *testing.T) {
		flag := logger.fatalLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("date fatal isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("date panic", func(t *testing.T) {
		flag := logger.panicLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("date panic isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
}

func TestLogger_TurnTimeOff(t *testing.T) {
	logger := newLogEnd()
	defer resetVariables()
	logger.OffTime()
	logFlagsTest := log.Ldate | log.Lshortfile
	t.Run("time info", func(t *testing.T) {
		flag := logger.infoLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("time info isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("time debug", func(t *testing.T) {
		flag := logger.debugLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("time debug isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("time trace", func(t *testing.T) {
		flag := logger.traceLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("time trace isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("time warn", func(t *testing.T) {
		flag := logger.warnLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("time warn isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("data error", func(t *testing.T) {
		flag := logger.errorLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("data error isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("time fatal", func(t *testing.T) {
		flag := logger.fatalLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("time fatal isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("time panic", func(t *testing.T) {
		flag := logger.panicLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("time panic isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
}

func TestLogger_TurnShortFileOff(t *testing.T) {
	logger := newLogEnd()
	defer resetVariables()
	logger.OffShortFile()
	logFlagsTest := log.Ldate | log.Ltime
	t.Run("shortfile info", func(t *testing.T) {
		flag := logger.infoLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("shortfile info isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("shortfile debug", func(t *testing.T) {
		flag := logger.debugLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("shortfile debug isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("shortfile trace", func(t *testing.T) {
		flag := logger.traceLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("shortfile trace isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("shortfile warn", func(t *testing.T) {
		flag := logger.warnLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("shortfile warn isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("shortfile error", func(t *testing.T) {
		flag := logger.errorLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("shortfile error isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("shortfile fatal", func(t *testing.T) {
		flag := logger.fatalLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("shortfile fatal isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
	t.Run("shortfile panic", func(t *testing.T) {
		flag := logger.panicLogger.log.Flags()
		if flag != logFlagsTest {
			t.Errorf("shortfile panic isn't valid: %d, valid value: %d", flag, logFlagsTest)
		}
	})
}

func TestLogger_Info(t *testing.T) {
	t.Run("Info", func(t *testing.T) {
		if err := Info("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
	t.Run("Infoln", func(t *testing.T) {
		if err := Infoln("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
	t.Run("Infof", func(t *testing.T) {
		if err := Infof("%v,%v,%v,%v,%v,%v", "test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
}

func TestLogger_Debug(t *testing.T) {
	t.Run("Debug", func(t *testing.T) {
		if err := Debug("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
	t.Run("Debugln", func(t *testing.T) {
		if err := Debugln("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
	t.Run("Debugf", func(t *testing.T) {
		if err := Debugf("%v,%v,%v,%v,%v,%v", "test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
}

func TestLogger_Trace(t *testing.T) {
	t.Run("Trace", func(t *testing.T) {
		if err := Trace("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
	t.Run("Traceln", func(t *testing.T) {
		if err := Traceln("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
	t.Run("Tracef", func(t *testing.T) {
		if err := Tracef("%v,%v,%v,%v,%v,%v", "test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
}

func TestLogger_Warn(t *testing.T) {
	t.Run("Warn", func(t *testing.T) {
		if err := Warn("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
	t.Run("Warnln", func(t *testing.T) {
		if err := Warnln("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
	t.Run("Warnf", func(t *testing.T) {
		if err := Warnf("%v,%v,%v,%v,%v,%v", "test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
}

func TestLogger_Error(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		if err := Error("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
	t.Run("Errorln", func(t *testing.T) {
		if err := Errorln("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
	t.Run("Errorf", func(t *testing.T) {
		if err := Errorf("%v,%v,%v,%v,%v,%v", "test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
}
