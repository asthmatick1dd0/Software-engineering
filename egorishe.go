package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Egorishe struct {
	Name       string
	University string
	Faculty    string
	Cookies    int
	BassSkill  int
	Spaghetti  bool
	Energy     int
}

func (e *Egorishe) EatCookie() string {
	e.Cookies++
	e.Energy += 15
	if e.Energy > 100 {
		e.Energy = 100
	}
	return fmt.Sprintf("%s съел печенье #%d · Энергия %d", e.Name, e.Cookies, e.Energy)
}

func (e *Egorishe) Study() string {
	loss := rand.Intn(12) + 5
	e.Energy -= loss
	if e.Energy < 0 {
		e.Energy = 0
	}
	return fmt.Sprintf("%s учится · -%d энергии (%d)", e.Name, loss, e.Energy)
}

func (e *Egorishe) PlayBass() string {
	styles := []string{"мрачно", "весело", "медитативно", "бодро"}
	e.Energy -= 8
	if e.Energy < 0 {
		e.Energy = 0
	}
	return fmt.Sprintf("%s играет на басу (%s) · уровень %d · %d", e.Name, styles[rand.Intn(len(styles))], e.BassSkill, e.Energy)
}

func (e *Egorishe) WriteSpaghetti() string {
	t := []string{
		"много вложенных if'ов",
		"функция длиной в килобайт",
		"неожиданные глобальные переменные",
		"комментарии вместо архитектуры",
	}
	e.Energy -= 10
	if e.Energy < 0 {
		e.Energy = 0
	}
	return fmt.Sprintf("%s пишет спагетти: %s · %d", e.Name, t[rand.Intn(len(t))], e.Energy)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	e := &Egorishe{
		Name:       "Егор Медянцев",
		University: "ЮФУ",
		Faculty:    "Мехмат",
		Cookies:    0,
		BassSkill:  7,
		Spaghetti:  true,
		Energy:     70,
	}
	tasks := []string{"study", "cookie", "code", "bass", "cookie", "spaghetti", "study", "cookie"}
	for _, t := range tasks {
		now := time.Now().Format("15:04:05")
		var out string
		switch t {
		case "cookie":
			out = e.EatCookie()
		case "study":
			out = e.Study()
		case "bass":
			out = e.PlayBass()
		case "spaghetti", "code":
			out = e.WriteSpaghetti()
		default:
			out = fmt.Sprintf("%s делает %s", e.Name, t)
		}
		fmt.Println(now, out)
		if e.Energy < 40 && rand.Intn(3) == 0 {
			fmt.Println(time.Now().Format("15:04:05"), e.EatCookie())
		}
		time.Sleep(time.Duration(150+rand.Intn(250)) * time.Millisecond)
	}
	fmt.Println(time.Now().Format("15:04:05"), e.Name, "закончил · печенек:", e.Cookies)
}
