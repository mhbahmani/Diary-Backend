package routing

import(
	"github.com/gorilla/mux"
	"net/http"

	"/src/view"

)


func HandleServices(r *mux.Router) {
	routeUserService(r.PathPrefix("/user").Subrouter())
}


func routeUserService(subrouter *mux.Router){
	subrouter.HandleFunc("/get-all", view.GetAllUsers).Methods(http.MethodGet)
}