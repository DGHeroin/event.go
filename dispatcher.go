package event

type (
    Dispatcher interface {
        Dispatch(name string, state interface{}, listeners ...Listener)
    }
    dispatcher int
)

func (d *dispatcher) Dispatch(name string, state interface{}, listeners ...Listener) {
    for _, s := range listeners {
        s.OnEvent(state)
    }
}

func NewDispatcher() Dispatcher {
    return new(dispatcher)
}
