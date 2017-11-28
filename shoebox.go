package shoebox

type Shoebox struct {

}

func NewShoebox () (*Shoebox, error) {

     sb := Shoebox{

     }

     return &sb, nil
}

func (sb *Shoebox) Archive(root string) error {


}

func (sb *Shoebox) GetItems() error {

	pages := -1
	page := 1
	per_page := 100

	method := "cooperhewitt.shoebox.items.getList"
	params := url.Values{}
	params.Set("per_page", string(per_page))

	for pages == -1 || pages >= page {

		params.Set("page", strconv.Itoa(page))
		rsp, err := sb.client.ExecuteMethod(method, &params)

		if err != nil {
		   	return err
		}

		body := rsp.Body()

		if pages == -1 {
			pages_float := body.Path("pages").Data().(float64)
			pages = int(pages_float)
		}

		items, _ := body.S("items").Children()

		wg := new(sync.WaitGroup)

		for _, item := range items {

			// don't put this bit in a goroutine because we want to preserve
			// the ordering in the "stuff" array (20160313/thisisaaronland)

			action := item.Path("action").Data().(string)

			if action != "collect" {
				fmt.Println("not a collect")
				continue
			}

			isa := item.Path("refers_to_a").Data().(string)

			if isa != "object" {
				fmt.Println("not an object")
				continue
			}

			item_id := item.Path("id").Data().(string)
			refers_to := item.Path("refers_to_uid").Data().(string)

			// fmt.Printf("%s refers to %s (%s)\n", item_id, refers_to, action)

			id, _ := strconv.Atoi(item_id)

			parent := Id2Path(id)
			root := filepath.Join(*shoebox, parent)

			_, err = os.Stat(root)

			if os.IsNotExist(err) {
				os.MkdirAll(root, 0755)
			}

			data := filepath.Join(root, "index.json")
			ioutil.WriteFile(data, []byte(item.String()), 0644) // todo - check errors

			stuff = append(stuff, data)

			wg.Add(1)

			// sudo make me a not-anonymous function

			go func() {

				<-ch
				defer func() { ch <- true }()

				defer wg.Done()

				method := "cooperhewitt.objects.getInfo"

				params := url.Values{}
				params.Set("object_id", refers_to)

				rsp, err := client.ExecuteMethod(method, &params)

				if err != nil {
					return
				}

				body := rsp.Body()

				data := filepath.Join(root, refers_to+".json")
				ioutil.WriteFile(data, []byte(body.String()), 0644) // todo - check errors

				images, _ := body.Path("object.images").Children()
				sizes := []string{"b", "n", "d", "sq", "z"}

				for _, img := range images {
					for _, sz := range sizes {

						p := fmt.Sprintf("%s.url", sz)
						url := img.Path(p).Data().(string)

						local := filepath.Join(root, filepath.Base(url))
						_, err = os.Stat(local)

						if !os.IsNotExist(err) {
							continue
						}

						wg.Add(1)

						// sudo make me a not-anonymous function

						go func() {

							<-ch
							defer func() { ch <- true }()

							defer wg.Done()

							err := GetStore(url, local)

							if err != nil {
								fmt.Printf("failed to fetch %s because %v", url, err)
							}
						}()
					}
				}

			}()

		}

		wg.Wait()

		page += 1
	}

	// todo - move this in to a separate function/file ?
	// todo - use html/template

	count := len(stuff)

	offset := 0
	idx := 0

	page = 1
	per_page = 3

	fl_count := float64(count)
	fl_per_page := float64(per_page)

	fl_pages := math.Ceil(fl_count / fl_per_page)
	pages = int(fl_pages)

	for page <= pages {

		var index_html string

		// todo - add next/prev rel links (20160314/thisisaaronland)

		index_html = fmt.Sprintf(`<!DOCTYPE html>
<html>
  <head>
    <title>Your shoebox</title>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta name="referrer" content="origin" />
    <meta http-equiv="X-UA-Compatible" content="IE=9" />
    <link rel="stylesheet" type="text/css" href="css/shoebox.css" />
    <script type="text/javascript" src="javascript/shoebox.common.js"></script>
    <script type="text/javascript" src="javascript/shoebox.index.js"></script>
    <script type="text/javascript">
    window.onload = function(e){
    		  shoebox.common.init();
    		  shoebox.index.init("item");
    }
    </script>
  </head>
  <body>
  <ul class="items">`)

		start := offset
		end := start + per_page

		if end > count {
			end = count
		}

		items := stuff[start:end]

		for _, path := range items {

			body, err := ioutil.ReadFile(path)

			if err != nil {
				continue
			}

			item, err := gabs.ParseJSON(body)

			if err != nil {
				continue
			}

			root := filepath.Dir(path)

			refers_to := item.Path("refers_to_uid").Data().(string)

			ref_path := filepath.Join(root, refers_to+".json")
			ref_body, err := ioutil.ReadFile(ref_path)

			if err != nil {
				continue
			}

			ref, err := gabs.ParseJSON(ref_body)

			if err != nil {
				continue
			}

			images, _ := ref.Path("object.images").Children()

			var local_sq string
			var local_b string
			var local_n string

			for _, image := range images {
				is_primary := image.Path("b.is_primary").Data().(string)

				if is_primary == "1" {

					remote_b := image.Path("b.url").Data().(string)
					remote_sq := image.Path("sq.url").Data().(string)
					remote_n := image.Path("n.url").Data().(string)

					local_b = filepath.Base(remote_b)
					local_sq = filepath.Base(remote_sq)
					local_n = filepath.Base(remote_n)

					break
				}
			}

			item_id := item.Path("id").Data().(string)
			id, _ := strconv.Atoi(item_id)

			item_root := Id2Path(id)

			parts := strings.Split(item_root, "/")

			for i, _ := range parts {
				parts[i] = ".."
			}

			find_root := strings.Join(parts, "/")

			// created := item.Path("created").Data().(string)
			// title := item.Path("title").Data().(string)
			// desc := item.Path("description").Data().(string)

			ref_title := item.Path("refers_to.title").Data().(string)
			ref_acc := item.Path("refers_to.accession_number").Data().(string)
			ref_url := item.Path("refers_to.url").Data().(string)
			//ref_text := item.Path("refers_to.gallery_text").Data().(string)

			item_href := filepath.Join(item_root, "index.html")
			item_sq := filepath.Join(item_root, local_sq)
			item_n := filepath.Join(item_root, local_n)
			item_b := local_b

			index_html += fmt.Sprintf(`<li class="item"><div class="thumb-container"><a href="%s"><img src="%s" class="shoebox-thumb" data-alt-src="%s" title="%s"/></a></div></li>`, item_href, item_sq, item_n, ref_title)

			var item_html string

			// todo - add next/prev rel links (20160314/thisisaaronland)

			item_html = fmt.Sprintf(`<!DOCTYPE html>
<html>
  <head>
    <title>%s</title>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta name="referrer" content="origin">
    <meta http-equiv="X-UA-Compatible" content="IE=9" />
    <link rel="stylesheet" type="text/css" href="%s/css/shoebox.css" />
    <link rel="stylesheet" type="text/css" href="%s/css/mapzen.whosonfirst.yesnofix.css" />    
    <script type="text/javascript" src="%s/javascript/shoebox.common.js"></script>
    <script type="text/javascript" src="%s/javascript/shoebox.item.js"></script>
    <script type="text/javascript" src="%s/javascript/mapzen.whosonfirst.yesnofix.js"></script>    
    <script type="text/javascript">
    window.onload = function(e){
    		  shoebox.common.init();
    		  shoebox.item.init();
    }
    </script>
  </head>
  <body>
  <div class="item-big">
  <div class="item-img">
  <img src="%s" id="item-image" />
  </div>
  <h2>%s <small><a href="%s">%s</a></small></h2></div>`, ref_title, find_root, find_root, find_root, find_root, find_root, item_b, ref_title, ref_url, ref_acc)


  	 	   	item_html += fmt.Sprintf(`<div id="details">`)
  	 	   	item_html += fmt.Sprintf(`<div id="details-item"></div>`)
  	 	   	item_html += fmt.Sprintf(`<div id="details-refers-to"></div>`)
  	 	   	item_html += fmt.Sprintf(`</div>`)
			
			item_html += fmt.Sprintf(`<ul class="pagination">`)

			if idx == 0 {
				item_html += fmt.Sprintf(`<li class="next prev">previous</li>`)
			} else {

				prev := idx - 1
				prev_path := stuff[prev]

				prev_path = strings.Replace(prev_path, *shoebox, "", -1)
				prev_path = strings.Replace(prev_path, "index.json", "index.html", -1)

				prev_path = filepath.Join(find_root, prev_path)

				item_html += fmt.Sprintf(`<li class="prev"><a href="%s" id="rel-prev">previous</a></li>`, prev_path)
			}

			// todo - figure out what page this item is on and link to that directly

			item_html += fmt.Sprintf(`<li class="top"><a href="%s/index.html">index</a></li>`, find_root)

			if idx+1 == count {
				item_html += fmt.Sprintf(`<li class="next last">last</li>`)
			} else {

				next := idx + 1
				next_path := stuff[next]

				next_path = strings.Replace(next_path, *shoebox, "", -1)
				next_path = strings.Replace(next_path, "index.json", "index.html", -1)

				next_path = filepath.Join(find_root, next_path)

				item_html += fmt.Sprintf(`<li class="next"><a href="%s" id="rel-next">next</a></li>`, next_path)
			}

			item_html += fmt.Sprintf(`</ul>`)
			item_html += fmt.Sprintf(`</body></html>`)

			item_parent := filepath.Join(*shoebox, item_root)
			item_path := filepath.Join(item_parent, "index.html")

			ioutil.WriteFile(item_path, []byte(item_html), 0644) // todo - check errors

			idx += 1
		}

		index_html += fmt.Sprintf(`</ul>`)
		index_html += fmt.Sprintf(`<ul class="pagination">`)

		if page == 1 {
			index_html += fmt.Sprintf(`<li class="prev first">this is the first page</li>`)
		} else {
			prev := page - 1
			prev_html := fmt.Sprintf("page%d.html", prev)
			index_html += fmt.Sprintf(`<li class="prev"><a href="%s" id="rel-prev">previous</a></li>`, prev_html)
		}

		if page == pages {
			index_html += fmt.Sprintf(`<li class="next last">this is the last page</li>`)
		} else {
			next := page + 1
			next_html := fmt.Sprintf("page%d.html", next)
			index_html += fmt.Sprintf(`<li class="next"><a href="%s" id="rel-next">next</a></li>`, next_html)
		}

		index_html += fmt.Sprintf("</ul>")
		index_html += fmt.Sprintf("</body></html>")

		page_html := fmt.Sprintf("page%d.html", page)
		index_path := filepath.Join(*shoebox, page_html)

		ioutil.WriteFile(index_path, []byte(index_html), 0644) // todo - check errors

		offset += per_page
		page += 1
	}

}
