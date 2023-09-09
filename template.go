package main

import "html/template"

var html = template.Must(template.New("index").Parse(`
<html> 
<head> 
    <title>{{.count}}</title>
</body>
	  <h1> 我抄襲你的抄襲 wahahahha !!! </h1>
    <h2> This Page already viewed {{.count}} times</h2>
</html>
`))
