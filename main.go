package main

import (
	"net/http"
	"os"
)

type fileServeHandler struct {
	path string
}

func (fsh *fileServeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fsh.path)
}

func FileServeHandler(path string) http.Handler {
	return &fileServeHandler{path}
}

func main() {
	mux := http.NewServeMux()

	// static files
	mux.Handle("/about/", FileServeHandler("static/about.html"))

	// redirects, now hosted at Tumblr
	mux.Handle("/feeds/feedburner.atom", http.RedirectHandler("http://tumblr.intranation.com/rss", http.StatusMovedPermanently))
	mux.Handle("/feeds/latest.atom", http.RedirectHandler("http://tumblr.intranation.com/rss", http.StatusMovedPermanently))
	mux.Handle("/entries/2010/03/why-i-dont-use-virtualbox/", http.RedirectHandler("http://tumblr.intranation.com/post/766290952/why-i-dont-use-virtualbox permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2010/02/using-dropbox-git-repository/", http.RedirectHandler("http://tumblr.intranation.com/post/766290743/using-dropbox-git-repository permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2010/01/how-set-up-your-own-private-git-server-linux/", http.RedirectHandler("http://tumblr.intranation.com/post/766290565/how-set-up-your-own-private-git-server-linux permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2009/12/python-virtualenv-quickstart-django/", http.RedirectHandler("http://tumblr.intranation.com/post/766290325/python-virtualenv-quickstart-django permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2009/06/week-using-emacs/", http.RedirectHandler("http://tumblr.intranation.com/post/766289934/week-using-emacs permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2009/03/development-virtual-machines-os-x-using-vmware-and/", http.RedirectHandler("http://tumblr.intranation.com/post/766289691/development-virtual-machines-os-x-using-vmware-ubuntu permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2009/02/ubuntu-intrepid-ibex-and-vmware-fusion-tools/", http.RedirectHandler("http://tumblr.intranation.com/post/766289439/ubuntu-intrepid-ibex-and-vmware-fusion-tools permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2009/01/resolute/", http.RedirectHandler("http://tumblr.intranation.com/post/766289166/resolute permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2008/11/django-static-asset-management/", http.RedirectHandler("http://tumblr.intranation.com/post/766288989/django-static-asset-management permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2008/10/synchronising-things-using-dropbox/", http.RedirectHandler("http://tumblr.intranation.com/post/766288807/synchronising-things-using-dropbox permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2008/10/installing-django-osx/", http.RedirectHandler("http://tumblr.intranation.com/post/766288554/installing-django-osx permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2008/09/using-nginx-reverse-proxy/", http.RedirectHandler("http://tumblr.intranation.com/post/766288369/using-nginx-reverse-proxy permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2008/09/hiatus/", http.RedirectHandler("http://tumblr.intranation.com/post/766288163/hiatus permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2008/08/unacceptable-marketing-practices-carsonified/", http.RedirectHandler("http://tumblr.intranation.com/post/766287422/unacceptable-marketing-practices-carsonified permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2008/08/ie-and-buttons/", http.RedirectHandler("http://tumblr.intranation.com/post/766287087/ie-and-buttons permanent", http.StatusMovedPermanently))
	mux.Handle("/entries/2008/07/first-post/", http.RedirectHandler("http://tumblr.intranation.com/post/766283182/first-post permanent", http.StatusMovedPermanently))

	err := http.ListenAndServe(":"+os.Getenv("PORT"), mux)

	if err != nil {
		panic(err)
	}
}
