package httplayer

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MyHandler struct {
	Router *mux.Router
	Http   *http.Server
}

    func Serverprovider() *MyHandler{
	    mH := &MyHandler{}
		mH.Router=mux.NewRouter()
		mH.mapRoutes()
		mH.Http= &http.Server{

			Addr: "127.0.0.1:2021",
			Handler: mH.Router,

		}
		fmt.Println("Trying to create the service , thanks")
		return mH
	}


	//creating a receiver function
	func (mH *MyHandler) mapRoutes(){

     mH.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w , "This is a response from the server")

	 })

	}

	//we need to create the listener but it may return an error so you need to capture it

	func (mH *MyHandler) Myserver() error {

		if err := mH.Http.ListenAndServe(); err != nil {

			return err
		}
		return nil
	


		}
		
	


