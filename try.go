package main


type CatchHandler interface {
    Catch(handler func(interface{}))
}

type catchHandler struct {
    err interface{}
}

func (c *catchHandler) Catch(handler func(interface{}))  {
    if c.err != nil {
        handler(c.err)
    }
}

func Try(block func()) CatchHandler {
    c := &catchHandler{}
    defer func() {
        defer func() {
            if v := recover(); v != nil {
                c.err = c
            }
        }()
        block()
    }()
    return c
}

func ThrowIfError(e error)  {
    if e != nil {
        panic(e)
    }
}
