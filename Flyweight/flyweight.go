package main

import "fmt"

type User struct {
	name string
}

func (u *User) Name() string {
	return u.name
}

type WebSite interface {
	Use(User)
}

type ConcreteWebsite struct {
	name string
}

func (c *ConcreteWebsite) Use(user User) {
	fmt.Println("网站分类", c.name, "用户：", user.Name())
}

type WebsiteFactory struct {
	flyweights map[string]WebSite
}

func (w *WebsiteFactory) Init() {
	w.flyweights = map[string]WebSite{}
}

func (w *WebsiteFactory) GetWebsiteCategory(key string) WebSite {
	if _, ok := w.flyweights[key]; !ok {
		w.flyweights[key] = &ConcreteWebsite{key}
	}
	return w.flyweights[key]
}
