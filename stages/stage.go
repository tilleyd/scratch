package stages

type Stage interface {
	Setup()
	Draw(delta float32)
	End()
}
