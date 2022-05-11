package event

type (
    Dispatcher interface {
        Dispatch(state F, listeners ...Listener)
    }
    dispatcher int
)

func (d *dispatcher) Dispatch(state F, listeners ...Listener) {
    for _, s := range listeners {
        s.OnEvent(state)
    }
}

func NewDispatcher() Dispatcher {
    return new(dispatcher)
}
