// Code generated by rice embed-go; DO NOT EDIT.
package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "app.js",
		FileModTime: time.Unix(1587200577, 0),

		Content: string("'use strict'\r\n\r\n\r\nlet today = new Date();\r\nlet formatDate = today.toDateString();\r\nlet selectElement = document.getElementById('date');\r\nselectElement.innerHTML = formatDate;\r\nconsole.log('Here\\'s a hidden message');\r\nfunction initMap() {\r\n    // The location of Uluru\r\n    var uluru = {lat: -25.344, lng: 131.036};\r\n    // The map, centered at Uluru\r\n    var map = new google.maps.Map(\r\n        document.getElementById('map'), {zoom: 4, center: uluru});\r\n    // The marker, positioned at Uluru\r\n    var marker = new google.maps.Marker({position: uluru, map: map});\r\n  }"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1587200577, 0),

		Content: string("<!DOCTYPE html>\r\n<html lang=\"en\">\r\n<head>\r\n    <meta charset=\"UTF-8\">\r\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n    <title>Task Timeline</title>\r\n    <link rel=\"stylesheet\" href=\"main.css\">\r\n</head>\r\n<body>\r\n    <h1>Task Timeline</h1>\r\n    <p id=\"date\"></p>\r\n    <ul>\r\n        <li class=\"list\">Edit the head</li>\r\n        <li class=\"list\">Edit the body</li>\r\n        <li>Link to JavaScript</li>\r\n  </ul>\r\n  <h3>My Google Maps Demo</h3>\r\n  <div id=\"map\"></div>\r\n  <script src=\"app.js\"></script>\r\n  <script async defer\r\n   src=\"https://maps.googleapis.com/maps/api/js?key=AIzaSyCo2JSgqGWWDTEEUl6gv1Ys2Kj2FyuS630&callback=initMap\"></script>\r\n</body>\r\n</html>"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "main.css",
		FileModTime: time.Unix(1587200577, 0),

		Content: string("body {\r\n    font-family: helvetica;\r\n  }\r\n  \r\n  ul {\r\n    font-family: monospace;\r\n  }\r\n  li {\r\n    list-style: circle;\r\n  }\r\n  \r\n  .list {\r\n    list-style: square;\r\n  }\r\n\r\n  #date {\r\n    font-size: 12px;\r\n    font-style: italic;\r\n    font-weight: bold;\r\n  }\r\n\r\n  #map {\r\n    height: 400px;  /* The height is 400 pixels */\r\n    width: 50%;  /* The width is the width of the web page */\r\n   }"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1587201181, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "app.js"
			file3, // "index.html"
			file4, // "main.css"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`website`, &embedded.EmbeddedBox{
		Name: `website`,
		Time: time.Unix(1587201181, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.js":     file2,
			"index.html": file3,
			"main.css":   file4,
		},
	})
}
