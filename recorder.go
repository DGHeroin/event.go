package event

type (
    Recorder interface {
        TakeSnapShot(event string, args F)
    }

    nilRecorder struct{}
)

func (r *nilRecorder) TakeSnapShot(event string, args F) {

}
