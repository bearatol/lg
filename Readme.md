# Simple logger for golang

### A simple extension of the standard library go "log" for minimises when and ease of use.

![lg](https://user-images.githubusercontent.com/54537638/130920979-8300b657-b953-4671-a567-584fac97f747.png)

<p>
The logger uses the standard go "log" package. Allows you to install and use expanded information (prefix, date and time, file and line) without cluttering up with unnecessary code. It also colors text on Linux and OS X for better reading of logs.
</p>

<p>
You can turn off the highlighting of logs. <br>
For example: <br>

```go
func init() {
    lg.Lg.OffColor()
}
```
</p>

## Install

```bash
go get -u github.com/bearatol/lg
```

## Exampels

```go
testArray := [...]int{1, 2, 3, 4, 5}

//Info
lg.Info("test text", testArray)
lg.Infoln("test text", testArray)
lg.Infof("some values: %s, %v", "test text", testArray)

//Debug
lg.Debug("test text", testArray)
lg.Debugln("test text", testArray)
lg.Debugf("some values: %s, %v", "test text", testArray)

//Trace
lg.Trace("test text", testArray)
lg.Traceln("test text", testArray)
lg.Tracef("some values: %s, %v", "test text", testArray)

//Warning
lg.Warn("test text", testArray)
lg.Warnln("test text", testArray)
lg.Warnf("some values: %s, %v", "test text", testArray)

//Error
lg.Error("test text", testArray)
lg.Errorln("test text", testArray)
lg.Errorf("some values: %s, %v", "test text", testArray)

//Fatal
lg.Fatal("test text", testArray)
lg.Fatalln("test text", testArray)
lg.Fatalf("some values: %s, %v", "test text", testArray)

//Panic
lg.Panic("test text", testArray)
lg.Panicln("test text", testArray)
lg.Panicf("some values: %s, %v", "test text", testArray)
```
