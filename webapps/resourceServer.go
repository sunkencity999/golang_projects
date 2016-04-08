package main

/* In a restful architecture, URL's are calls to methods or functions on the server. This is a performance boost.
the resource server needs to listen on a tcp report and handle requests, so that it routes a url to a file or a function.
sometimes the server is refered to as a serveMux, http request router, or multiplexer. These are all the same thing, basically.
*/
import "net/http"

func main() {

	//http.Handle("/", new(MyHandler))

	http.ListenAndServe(":8080", http.FileServer(http.Dir("html"))) //nil just means we're using the default serveMux. This function must be called at the end.

}

/* The below is an example of a custom-built fileserver, used instead of Go's built-in version
MyHandler is a struct built to receive http requests.
type MyHandler struct {
	http.Handler
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "html/" + req.URL.Path //this sets the default route to the html folder in the hierarchy
	//data, err := ioutil.ReadFile(string(path))
	f, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(f)
		var contentType string

		if strings.HasSuffix(path, ".css") { //this allows us to tell the browser what type of file it is serving.
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content Type", contentType) //creates a header and writes the type to it
		//w.Write(data)                               //writes the information pulled from the file at the destination of the URL path..the html page
		bufferedReader.WriteTo(w) //writes buffered information to the w variable to be opened
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}

//func servedFunc(w http.ResponseWriter, req *http.Request) {
//	w.Write([]byte("Hello Universe!"))
//func servedFuncTwo(w http.ResponseWriter, req *http.Request) {
//	w.Write([]byte("Hello John in Particular"))
//}
*/
