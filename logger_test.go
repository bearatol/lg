package lg

import "testing"

func TestLogger_TurnColor(t *testing.T) {
	t.Run("color off", func(t *testing.T) {
		lg := newLogEnd()
		lg.OffColor()

		emptyReset := lg.infoLogger.c.Reset == ""
		if !emptyReset {
			t.Errorf("color Reset isn't empty: %s", lg.infoLogger.c.Reset)
		}
		emptyRed := lg.infoLogger.c.Red == ""
		if !emptyRed {
			t.Errorf("color Red isn't empty: %s", lg.infoLogger.c.Red)
		}
		emptyYellow := lg.infoLogger.c.Yellow == ""
		if !emptyYellow {
			t.Errorf("color Yellow isn't empty: %s", lg.infoLogger.c.Yellow)
		}
		emptyPurple := lg.infoLogger.c.Purple == ""
		if !emptyPurple {
			t.Errorf("color Purple isn't empty: %s", lg.infoLogger.c.Purple)
		}
		emptyPurpleUnderline := lg.infoLogger.c.PurpleUnderline == ""
		if !emptyPurpleUnderline {
			t.Errorf("color PurpleUnderline isn't empty: %s", lg.infoLogger.c.PurpleUnderline)
		}
		emptyCyan := lg.infoLogger.c.Cyan == ""
		if !emptyCyan {
			t.Errorf("color Cyan isn't empty: %s", lg.infoLogger.c.Cyan)
		}
		emptyGray := lg.infoLogger.c.Gray == ""
		if !emptyGray {
			t.Errorf("color Gray isn't empty: %s", lg.infoLogger.c.Gray)
		}
		emptyGrayUnderline := lg.infoLogger.c.GrayUnderline == ""
		if !emptyGrayUnderline {
			t.Errorf("color GrayUnderline isn't empty: %s", lg.infoLogger.c.GrayUnderline)
		}
	})
	t.Run("color on", func(t *testing.T) {
		lg := newLogEnd()

		colorReset := lg.infoLogger.c.Reset == "\033[0;0;0m"
		if !colorReset {
			t.Errorf("color Reset isn't valid: %s", lg.infoLogger.c.Reset)
		}
		colorRed := lg.infoLogger.c.Red == "\033[0;0;31m"
		if !colorRed {
			t.Errorf("color ReRedset isn't valid: %s", lg.infoLogger.c.Red)
		}
		colorYellow := lg.infoLogger.c.Yellow == "\033[0;0;33m"
		if !colorYellow {
			t.Errorf("color Yellow isn't valid: %s", lg.infoLogger.c.Yellow)
		}
		colorPurple := lg.infoLogger.c.Purple == "\033[0;0;35m"
		if !colorPurple {
			t.Errorf("color Purple isn't valid: %s", lg.infoLogger.c.Purple)
		}
		colorPurpleUnderline := lg.infoLogger.c.PurpleUnderline == "\033[0;4;35m"
		if !colorPurpleUnderline {
			t.Errorf("color PurpleUnderline isn't valid: %s", lg.infoLogger.c.PurpleUnderline)
		}
		colorCyan := lg.infoLogger.c.Cyan == "\033[0;0;36m"
		if !colorCyan {
			t.Errorf("color Cyan isn't valid: %s", lg.infoLogger.c.Cyan)
		}
		colorGray := lg.infoLogger.c.Gray == "\033[0;0;37m"
		if !colorGray {
			t.Errorf("color Gray isn't valid: %s", lg.infoLogger.c.Gray)
		}
		colorGrayUnderline := lg.infoLogger.c.GrayUnderline == "\033[0;4;37m"
		if !colorGrayUnderline {
			t.Errorf("color GrayUnderline isn't valid: %s", lg.infoLogger.c.GrayUnderline)
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
