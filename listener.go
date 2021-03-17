package event

type (
    Listener interface {
        OnEvent(args interface{})
    }
)
