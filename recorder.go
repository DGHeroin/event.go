package event

type (
    Recorder interface {
        TakeSnapShot(event string, args interface{})
    }

    nilRecorder struct{}
)

func (r *nilRecorder) TakeSnapShot(event string, args interface{}) {

}
