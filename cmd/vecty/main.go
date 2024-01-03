package main

import (
	"fmt"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/hexops/vecty/event"

	//"github.com/patrickmn/go-cache"
	//"github.com/go-co-op/gocron/v2"
)

func main() {
	vecty.SetTitle("Vecty Form")
	//flowbite from cdn
	vecty.AddStylesheet("dist/css/flowbite.min.css")	
	
	// TODO.QueryClient Instance
	/*
	pg := &PageView{
		queryclient : QueryClient{
			endpoint :
			sched : gocron.NewScheduler()
			cache : cache.New(60*time.Second, 60*time.Second)
		}
	}
	*/
	fmt.Println("form start")
	vecty.RenderBody(&PageView{})
}

/*
type QueryClient {
	endpoint string
	cache cache.Cache
}
*/

// PageView is our main page component.
type PageView struct {
	vecty.Core
	se string //submit event
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		//Form Control
		elem.Form(
			vecty.Markup(
				vecty.Class("max-w-sm"),vecty.Class("mx-auto"),
				//event.Click(p.onSubmitHandle).PreventDefault(),
			),
			p.MakeLabelAndInput("Username",prop.TypeText,"kokusai taro"),	
			p.MakeLabelAndInput("Password",prop.TypePassword,""),	
			//p.MakeButton("Submit",nil),	
			p.MakeButton("Submit",p.onSubmitHandle),	
			vecty.Text(p.se),	
		),	
	)
}

func (p *PageView) onSubmitHandle(e *vecty.Event) {
	//test code
	p.se = "Form Submitted"	
	//fmt.Println(e.Target.Get("value").String())
	fmt.Println(e.Target.Length())
	fmt.Println(e.Value.Length())
	//fmt.Println(e.Target.Get("value").Get("Username").String())
	//p.se = e.Value.Get("Username").String()	
	//fmt.Println(p.se)
	vecty.Rerender(p)	
}

func (p *PageView) MakeButton(labelname string,callback func(*vecty.Event)) *vecty.HTML {
	return elem.Button(
		vecty.Markup(
			vecty.Class("text-white"),
			vecty.Class("bg-blue-700"),
			vecty.Class("hover:bg-blue-800"),
			vecty.Class("focus:ring-4"),
			vecty.Class("focus:outline-none"),
			vecty.Class("focus:ring-blue-300"),
			vecty.Class("font-medium"),
			vecty.Class("rounded-lg"),
			vecty.Class("text-sm"),
			vecty.Class("w-full"),
			vecty.Class("sm:w-auto"),
			vecty.Class("px-5"),
			vecty.Class("py-2.5"),
			vecty.Class("text-center"),
			vecty.Class("dark:bg-blue-600"),
			vecty.Class("dark:hover:bg-blue-700"),
			vecty.Class("dark:focus:ring-blue-800"),
			prop.Type(prop.TypeSubmit),
			//event.Click(p.onSubmitHandle).PreventDefault(),
			event.Click(callback).PreventDefault(),
		),
		vecty.Text(labelname),
	);	
}

func (p *PageView) MakeLabelAndInput(labelname string,ptype prop.InputType,plhd string) *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("mb-5"),
		),
		elem.Label(
			p.MarkupLabel(),
			vecty.Text(labelname),
		),
		elem.Input(
			p.MarkupInput(labelname,ptype,plhd),	
		),
	);
}

func (p *PageView) MarkupLabel() vecty.MarkupList {
	return vecty.Markup(
		vecty.Class("block"),
		vecty.Class("mb-2"),
		vecty.Class("text-sm"),
		vecty.Class("font-medium"),
		vecty.Class("text-gray-900"),
		vecty.Class("dark:text-white"),
	);
}

func (p *PageView) MarkupInput(id string,ptype prop.InputType,plhd string) vecty.MarkupList {
	return vecty.Markup(
		vecty.Class("bg-gray-50"),
		vecty.Class("border"),
		vecty.Class("border-gray-300"),
		vecty.Class("text-gray-900"),
		vecty.Class("test-sm"),
		vecty.Class("rounded-lg"),
		vecty.Class("focus:ring-blue-500"),
		vecty.Class("focus:border-blue-500"),
		vecty.Class("block"),
		vecty.Class("w-full"),
		vecty.Class("p-2.5"),
		vecty.Class("dark:bg-gray-700"),
		vecty.Class("dark:border-gray-600"),
		vecty.Class("dark:placeholder-gray-400"),
		vecty.Class("dark:text-white"),
		vecty.Class("dark:focus:ring-blue-500"),
		vecty.Class("dark:focus:border-blue-500"),
		prop.Type(ptype),
		prop.Name(id),
		prop.Placeholder(plhd),
	);
}

