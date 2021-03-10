package main

func main() {
	f := new(WebsiteFactory)
	f.Init()

	xiaoWang := f.GetWebsiteCategory("产品展示")
	xiaoWang.Use(User{"小王"})

	xiaoHuang := f.GetWebsiteCategory("产品展示")
	xiaoHuang.Use(User{"小黄"})

	xiaoChu := f.GetWebsiteCategory("博客")
	xiaoChu.Use(User{"小褚"})

	shiShu := f.GetWebsiteCategory("博客")
	shiShu.Use(User{"师叔"})
}
