package server

func Init() {
	r := NewRouter()
	r.Run("localhost:2345")
}
