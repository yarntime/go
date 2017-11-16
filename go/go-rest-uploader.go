package main

import (
	"flag"
	"fmt"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"

	"io"
	"io/ioutil"
	"os"
)

var http_port int
var root_path string

type Resource struct {
}

func init() {
	flag.IntVar(&http_port, "p,port", 8008, "http server port")
	flag.StringVar(&root_path, "r,root", "./data", "root path")
}

var html_before_tmpl string = `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>Simple HTTP</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/semantic-ui/1.11.8/semantic.min.css"/>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.3/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/semantic-ui/1.11.8/semantic.min.js"></script>
</head>
<body>

<form class="ui form" action="/users/upload/" enctype="multipart/form-data" method="post">
  <div class="field">
    <label>File</label>
    <input type="file" name="filename">
  </div>
  <button class="ui button" type="submit">upload</button>
</form>
<div id="filelist">
  <table class="ui celled striped table">
  <thead>
    <tr><th colspan="3">
      Files
    </th>
  </tr></thead>
  <tbody>
`
var html_after_tmpl string = `
  </tbody>
</table>
</div>

</body>
</html>
`

func (u Resource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON, "multipart/form-data").
		Produces(restful.MIME_JSON) // you can specify this per route as well

	ws.Route(ws.GET("/file/{file}").To(u.downLoad))
	ws.Route(ws.GET("/ui").To(u.showui))
	ws.Route(ws.POST("/upload").To(u.uploadfile))

	container.Add(ws)
}

func (u Resource) downLoad(request *restful.Request, response *restful.Response) {
	file := request.PathParameter("file")
	http.ServeFile(response.ResponseWriter, request.Request, string(http.Dir(root_path+"/"+file)))
}

func (u Resource) showui(request *restful.Request, response *restful.Response) {
	inner := func(filename, filesize, lasttime string) string {
		tr := "<tr><td class='collapsing'><a href='/users/file/" + filename + "'><i class='file icon'></i>" + filename + "</a></td>"
		tr = tr + "<td>" + filesize + "MB</td>"
		tr = tr + "<td class='right aligned collapsing'>" + lasttime + "</td>"
		tr = tr + "</td>"
		return tr
	}
	fileinfos, err := ioutil.ReadDir(root_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	trs := ""
	for _, fi := range fileinfos {
		trs = trs + inner(fi.Name(), fmt.Sprintf("%v", fi.Size()/1024/1024), fi.ModTime().String())
	}

	response.AddHeader("Content-Type", "text/html")
	response.Write([]byte(html_before_tmpl + trs + html_after_tmpl))

}

func (u *Resource) uploadfile(request *restful.Request, response *restful.Response) {
	fmt.Printf("request: %v\n", request.Request)
	file, header, err := request.Request.FormFile("filename")
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := header.Filename
	out, err := os.Create(root_path + "/" + filename)
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println(err)
	}
	response.AddHeader("Content-Type", "text/plain")
	response.Write([]byte("OK"))

}

func main() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	u := Resource{}
	u.Register(wsContainer)

	log.Print("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
