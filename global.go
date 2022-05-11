package event

var (
    _mgr Manager
)
type gHandler struct {
    cb func(args F)
}
func (g gHandler) OnEvent(args F) {
    g.cb(args)
}
func init() {
    storage := NewMemoryStorage()
    dispatcher := NewDispatcher()
    _mgr = NewManager(storage, dispatcher, nil)
}
func On(name string, cb func(F)) Listener {
    var handler = &gHandler{
        cb: cb,
    }
    _mgr.Add(name, handler)
    return handler
}
func Off(name string, listener Listener) {
    _mgr.Remove(name, listener)
}
func Emit(name string, f F) {
    _mgr.Dispatch(name, f)
}
