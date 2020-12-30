package app

import (	
	"TOMforMSSQL/app/model"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render = render.New()

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	//http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect)
	
	filePrefix, _ := filepath.Abs("./app/views/")

	//tmpl, err := template.ParseGlob("*.html")
	tmpl, err := template.New("index").ParseFiles(		
		filePrefix + "/sidebar.html",		
		filePrefix + "/index.html",
	)
	if err != nil {
		panic(err)
	}	
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func (a *AppHandler) renderGetServiceListHandler(w http.ResponseWriter, r *http.Request) {
	filePrefix, _ := filepath.Abs("./app/views/")

	tmpl, err := template.New("servicelist").ParseFiles(		
		filePrefix + "/header.html",
		filePrefix + "/footer.html",
		filePrefix + "/sidebar.html",		
		filePrefix + "/app/servicelist.html",
	)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "servicelist.html", nil)
}

func (a *AppHandler) getServiceListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetServiceList()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) getServiceByServiceNoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceno, _ := strconv.Atoi(vars["serviceno"])
	list := a.db.GetServiceByServiceNo(serviceno)
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) renderGetServiceByServiceNoHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// serviceno, _ := strconv.Atoi(vars["serviceno"])
	// list := a.db.GetServiceByServiceNo(serviceno)

	filePrefix, _ := filepath.Abs("./app/views/")
	tmpl, err := template.New("service").ParseFiles(		
		filePrefix + "/header.html",
		filePrefix + "/footer.html",
		filePrefix + "/sidebar.html",		
		filePrefix + "/app/service.html",
	)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "service.html", nil)

	// e, err := json.Marshal(list)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(e))

	// //fmt.Println(services[0].created_date)	
	
}

func (a *AppHandler) getHostListByServiceNoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceno, _ := strconv.Atoi(vars["serviceno"])
	list := a.db.GetHostListByServiceNo(serviceno)
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) Close() {
	a.db.Close()
}

func MakeHandler() *AppHandler {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)
	n.UseHandler(r)

	a := &AppHandler{
		Handler: n,
		db:      model.NewDBHandler(),
	}

	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/render_service", a.renderGetServiceByServiceNoHandler).Methods("GET")
	r.HandleFunc("/render_servicelist", a.renderGetServiceListHandler).Methods("GET")	
	r.HandleFunc("/servicelist", a.getServiceListHandler).Methods("GET")
	r.HandleFunc("/service/{serviceno:[0-9]+}", a.getServiceByServiceNoHandler).Methods("GET")
	
	r.HandleFunc("/host/{serviceno:[0-9]+}", a.getHostListByServiceNoHandler).Methods("GET")

	return a
}
