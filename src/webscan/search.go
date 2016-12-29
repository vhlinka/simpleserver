package webscan

import (
	"context"
	"fmt"
	"google.golang.org/appengine/search"
	"io/ioutil"
	"net/http"
	"time"
)

//
// sample googe search - using Google API (and google's sample code)
//
// GET https://www.googleapis.com/customsearch/v1?key=INSERT_YOUR_API_KEY&cx=017576662512468239146:omuauf_lfve&q=lectures
//
// <script>
//  (function() {
//   var cx = '011864276067333636885:rgce82binry';
//    var gcse = document.createElement('script');
//    gcse.type = 'text/javascript';
//    gcse.async = true;
//    gcse.src = 'https://cse.google.com/cse.js?cx=' + cx;
//    var s = document.getElementsByTagName('script')[0];
//    s.parentNode.insertBefore(gcse, s);
//  })();
//</script>
//<gcse:search></gcse:search>
//

func SearchGoogle(w http.ResponseWriter, req *http.Request) {
	//	url := "https://www.googleapis.com/customsearch/v1?key=AIzaSyCb6vGXygPRtFCePOoXpu221gl01ggBZqA&cx=011864276067333636885:rgce82binry&q=smart+mesh+network+technology&tbs=qdr:y,sbd:1"
	apipart := "https://www.googleapis.com/customsearch/v1?"
	keypart := "key=AIzaSyCb6vGXygPRtFCePOoXpu221gl01ggBZqA"
	cxpart := "&cx=017576662512468239146:omuauf_lfve"
	//	cxpart := "011864276067333636885:rgce82binry"
	//cxpart := "017576662512468239146:omuauf_lfve&q=lectures"
	querypart := "&q=mesh+network+technology&tbs=qdr:y,sbd:1"
	url := apipart + keypart + cxpart + querypart

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w, "fetch: %v\n", err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(w, "fetch: reading %s: %v\n", url, err)
		return
	}

	fmt.Fprintf(w, "%s\n", b)
}

func oldSearchGoogle(w http.ResponseWriter, req *http.Request) {
	//	url := "https://www.google.com/#q=mesh+sensor+network"
	// ctx is the Context for this handler. Calling cancel closes the
	// ctx.Done channel, which is the cancellation signal for requests
	// started by this handler.
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		// The request has a timeout, so create a context that is
		// canceled automatically when the timeout expires.
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel() // Cancel ctx as soon as handleSearch returns

	// Check the search query.
	query := req.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}

	//	------------ added sample from google.golang.org/appengine/search example
	type Doc struct {
		Author   string
		Comment  string
		Creation time.Time
	}

	index, err := search.Open("comments")

	if err != nil {
		return
	}

	/*	newID, err := index.Put(ctx, "", &Doc{
			Author:   "gopher",
			Comment:  "the truth of the matter",
			Creation: time.Now(),
		})
		if err != nil {
			return
		}
	*/
	//	var doc Doc
	//	err = index.Get(ctx, id, &doc)
	//	if err != nil {
	//		return
	//	}

	for t := index.Search(ctx, "Comment:truth", nil); ; {
		var doc Doc
		id, err := t.Next(&doc)
		if err == search.Done {
			break
		}
		if err != nil {
			return
		}
		fmt.Fprintf(w, "%s -> %#v\n", id, doc)
	}

	// Store the user IP in ctx for use by code in other packages.
	/*	userIP, err := userip.FromRequest(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx = userip.NewContext(ctx, userIP)

		// Run the Google search and print the results.
		start := time.Now()
		results, err := google.Search(ctx, query)
		elapsed := time.Since(start)

		if err := resultsTemplate.Execute(w, struct {
			Results          google.Results
			Timeout, Elapsed time.Duration
		}{
			Results: results,
			Timeout: timeout,
			Elapsed: elapsed,
		}); err != nil {
			log.Print(err)
			return
		}
	*/
}
