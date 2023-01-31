package lg

import "testing"

func TestLogger_TurnColor(t *testing.T) {
	t.Run("color off", func(t *testing.T) {
		lg := newLogEnd()
		lg.OffColor()

		if lg.infoLogger.c.Reset != "" ||
			lg.infoLogger.c.Red != "" ||
			lg.infoLogger.c.Yellow != "" ||
			lg.infoLogger.c.Purple != "" ||
			lg.infoLogger.c.PurpleUnderline != "" ||
			lg.infoLogger.c.Cyan != "" ||
			lg.infoLogger.c.Gray != "" ||
			lg.infoLogger.c.GrayUnderline != "" {
			t.Logf("%+v", lg.infoLogger.c)
			t.Error("color isn't empty")
		}
	})
	t.Run("color on", func(t *testing.T) {
		lg := newLogEnd()

		if lg.infoLogger.c.Reset != "\033[0;0;0m" ||
			lg.infoLogger.c.Red != "\033[0;0;31m" ||
			lg.infoLogger.c.Yellow != "\033[0;0;33m" ||
			lg.infoLogger.c.Purple != "\033[0;0;35m" ||
			lg.infoLogger.c.PurpleUnderline != "\033[0;4;35m" ||
			lg.infoLogger.c.Cyan != "\033[0;0;36m" ||
			lg.infoLogger.c.Gray != "\033[0;0;37m" ||
			lg.infoLogger.c.GrayUnderline != "\033[0;4;37m" {
			t.Logf("%+v", lg.infoLogger.c)
			t.Error("color isn't valid")
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
