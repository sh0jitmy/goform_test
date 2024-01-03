package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	//"github.com/hexops/vecty/style"
)

func main() {
	vecty.SetTitle("Vecty Form")
	//flowbite from cdn
	vecty.AddStylesheet("dist/css/flowbite.min.css")	
	vecty.RenderBody(&PageView{})
}

// PageView is our main page component.
type PageView struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		//Form Control
		elem.Form(
			vecty.Markup(
				vecty.Class("max-w-sm"),vecty.Class("mx-auto"),
			),
			p.MakeLabelAndInput("Username",prop.TypeText),	
			p.MakeLabelAndInput("Password",prop.TypePassword),	
			p.MakeButton("Submit"),	
		),	
	)
}

func (p *PageView) MakeButton(labelname string) *vecty.HTML {
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
		),
		vecty.Text(labelname),
	);	
}

func (p *PageView) MakeLabelAndInput(labelname string,ptype prop.InputType) *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("mb-5"),
		),
		elem.Label(
			p.MarkupLabel(),
			vecty.Text(labelname),
		),
		elem.Input(
			p.MarkupInput(ptype),	
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

func (p *PageView) MarkupInput(ptype prop.InputType) vecty.MarkupList {
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
	);
}

