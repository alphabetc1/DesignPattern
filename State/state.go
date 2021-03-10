/*
state状态模式
定义：当一个对象的内在状态改变时允许改变其行为，这个对象看起来像是改变了其类。
主要解决的是当控制一个对象状态转换的条件表达式过于复杂时的情况，即“长方法(Long Method)”。
方法是把状态的判断逻辑转移到不同状态的一系列类当中。
作者：   alphabetc1
日期：   20210218
*/
package main

import "fmt"

func NewWork() *Work {
	state := new(ForenoonState)
	return &Work{current: state}
}

type Work struct {
	current  State
	hour     int
	finished bool
}

func (w *Work) Hour() int {
	if w == nil {
		return -1
	}
	return w.hour
}

func (w *Work) SetHour(hour int) {
	if w == nil {
		return
	}
	w.hour = hour
}

func (w *Work) SetState(s State) {
	if w == nil {
		return
	}
	w.current = s
}

func (w *Work) WriteProgram() {
	if w == nil {
		return
	}
	w.current.WriteProgram(w)
}

type State interface {
	WriteProgram(w *Work)
}

type ForenoonState struct {
}

func (f *ForenoonState) WriteProgram(w *Work) {
	if w.Hour() < 12 {
		fmt.Println("当前时间：{", w.Hour(), "}点 上午工作，精神百倍")
	} else {
		w.SetState(new(NoonState))
		w.WriteProgram()
	}
}

type NoonState struct {
}

func (f *NoonState) WriteProgram(w *Work) {
	if w.Hour() < 13 {
		fmt.Println("当前时间：{", w.Hour(), "}点 午饭时间，犯困")
	} else {
		w.SetState(new(AfternoonState))
		w.WriteProgram()
	}
}

type AfternoonState struct {
}

func (a *AfternoonState) WriteProgram(w *Work) {
	if w.Hour() < 19 {
		fmt.Println("当前时间：{", w.Hour(), "}点 下午时间，状态还行")
	} else {
		w.SetState(new(EveningState))
		w.WriteProgram()
	}
}

type EveningState struct {
}

func (e *EveningState) WriteProgram(w *Work) {
	if w.finished {
		w.SetState(new(RestState))
		w.WriteProgram()
	} else {
		if w.Hour() < 21 {
			fmt.Println("当前时间：{", w.Hour(), "}点 加班时间，有点累")
		} else {
			w.SetState(new(SleepingState))
			w.WriteProgram()
		}
	}
}

type RestState struct {
}

func (r *RestState) WriteProgram(w *Work) {
	fmt.Println("当前时间：{", w.Hour(), "}点 到点了，下班！")
}

type SleepingState struct {
}

func (s *SleepingState) WriteProgram(w *Work) {
	fmt.Println("当前时间：{", w.Hour(), "}点 顶不住了，Zzz")
}
