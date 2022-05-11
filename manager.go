package event

type (
    Manager interface {
        Add(name string, listener Listener)
        Remove(name string, listener Listener)
        Dispatch(name string, args F)
    }
    manager struct {
        Storage    Storage
        Dispatcher Dispatcher
        Recorder   Recorder
    }
)

func NewManager(storage Storage, dispatcher Dispatcher, recorder Recorder) Manager {
    if recorder == nil {
        recorder = &nilRecorder{}
    }
    m := &manager{
        Storage:    storage,
        Dispatcher: dispatcher,
        Recorder:   recorder,
    }
    return m
}

func (m *manager) Add(name string, listener Listener) {
    m.Storage.Add(name, listener)
}

func (m *manager) Remove(name string, listener Listener) {
    m.Storage.Remove(name, listener)
}
func (m *manager) Dispatch(name string, args F) {
    // record
    m.Recorder.TakeSnapShot(name, args)
    // dispatch
    ss := m.Storage.Listeners(name)
    m.Dispatcher.Dispatch(args, ss...)
}
