package http

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	"github.com/baishancloud/octopux-gateway/g"
)

type Dto struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	httpserv *http.Server
	ln       *net.TCPListener
)

func Stop() {
	if ln != nil {
		log.Println("set http listen close!")
		ln.Close()
	}
}

func Start() {
	go startHttpServer()
}

func startHttpServer() {
	if !g.Config().Http.Enabled {
		return
	}

	addr := g.Config().Http.Listen
	if addr == "" {
		return
	}

	configCommonRoutes()
	configProcHttpRoutes()
	configApiHttpRoutes()

	httpserv = &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 30,
	}

	log.Println("http.startHttpServer ok, listening", addr)
	if addr == "" {
		addr = ":http"
	}
	hln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Start listen http error :", err)
		return
	}
	ln = hln.(*net.TCPListener)
	log.Println(httpserv.Serve(ln))
}

func RenderJson(w http.ResponseWriter, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bs)
}

func RenderDataJson(w http.ResponseWriter, data interface{}) {
	RenderJson(w, Dto{Msg: "success", Data: data})
}

func RenderMsgJson(w http.ResponseWriter, msg string) {
	RenderJson(w, map[string]string{"msg": msg})
}

func AutoRender(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		RenderMsgJson(w, err.Error())
		return
	}
	RenderDataJson(w, data)
}
