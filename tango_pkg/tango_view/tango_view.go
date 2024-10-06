package tango_view

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type TangoView struct {
	Component templ.Component
	Data      ViewData
}

func New() TangoView {
	return TangoView{
		Data: NewViewData(),
	}
}

func (me *TangoView) SetComponent(component templ.Component) {
	me.Component = component
}

func (me *TangoView) Render(ctx echo.Context, statusCode int) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := me.Component.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

// func (me *TangoView) Render() error {
// 	return me.Component.Render(context.Background(), os.Stdout)
// 	// return me.Component
// }
