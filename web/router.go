package web

type Router interface {
    HandleFunc(string, func)
}