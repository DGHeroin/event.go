package event

type (
    Listener interface {
        OnEvent(args F)
    }
)
