package core

import (
	"fmt"

	"github.com/alikWu/go-tomcat/servlet"
)

type ApplicationFilterChain struct {
	stage   int
	filters []servlet.Filter
	servlet servlet.Servlet
}

func NewApplicationFilterChain() *ApplicationFilterChain {
	return &ApplicationFilterChain{}
}

func (afc *ApplicationFilterChain) DoFilter(request servlet.ServletRequest, response servlet.ServletResponse) error {
	fmt.Println("ApplicationFilterChain DoFilter")
	return afc.internalDoFilter(request, response)
}

func (afc *ApplicationFilterChain) internalDoFilter(request servlet.ServletRequest, response servlet.ServletResponse) error {
	curStage := afc.stage
	afc.stage++
	if curStage < len(afc.filters) {
		fmt.Println("ApplicationFilterChain internalDoFilter stage=", curStage)
		//调用filter的过滤逻辑，根据规范，filter中要再次调用filterChain.doFilter
		//这样又会回到internalDoFilter()方法，就会再拿到下一个filter，如此实现一个一个往下传
		return afc.filters[curStage].DoFilter(request, response, afc)
	}

	return afc.servlet.Service(NewServletRequestFacade(request), NewServletResponseFacade(response))
}

func (afc *ApplicationFilterChain) AddFilter(filter servlet.Filter) {
	afc.filters = append(afc.filters, filter)
}

func (afc *ApplicationFilterChain) SetServlet(servlet servlet.Servlet) {
	afc.servlet = servlet
}

func (afc *ApplicationFilterChain) Release() {
	afc.servlet = nil
	afc.filters = nil
	afc.stage = 0
}
